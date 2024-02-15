package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type UserInput struct {
	List []string `json:"websites"`
}

type httpChecker struct {
	websiteMap map[string]string
}

// var websiteMap map[string]string = make(map[string]string)
var list []string

func main() {
	fmt.Println("Starting server")

	ctx := context.Background()

	r := mux.NewRouter()

	websiteObj := httpChecker{make(map[string]string)}

	r.HandleFunc("/", genericHandler)
	r.HandleFunc("/input", getWebsiteList).Methods("POST")
	r.HandleFunc("/check", websiteObj.checkWebsiteStatusHandler).Methods("GET")

	go func() {
		for {
			if list != nil {
				websiteObj.checkWebsiteStatus(ctx, list)
				time.Sleep(time.Minute)
			}
		}
	}()

	http.ListenAndServe("localhost:8080", r)

}

func getWebsiteList(w http.ResponseWriter, r *http.Request) {
	var website UserInput
	err := json.NewDecoder(r.Body).Decode(&website)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list = website.List

	w.WriteHeader(http.StatusOK)

}

func (websiteObj *httpChecker) checkWebsiteStatus(ctx context.Context, web []string) {
	fmt.Println("running checkWebsiteStatus")
	for _, data := range web {

		go func(data string) {
			response, err := http.Get("https://" + data)
			if err == nil {
				responseStatus := response.StatusCode
				if responseStatus == 200 {
					websiteObj.websiteMap[data] = "UP"
				} else {
					websiteObj.websiteMap[data] = "DOWN"
				}

			} else {
				websiteObj.websiteMap[data] = "DOWN"
			}
		}(data)

	}
}

func (websiteObj *httpChecker) checkWebsiteStatusHandler(w http.ResponseWriter, r *http.Request) {

	if len(websiteObj.websiteMap) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(map[string]string{"error": "The map is empty"})
		w.Write(res)
		return
	}

	name := r.URL.Query().Get("name")

	if name != "" {

		_, ok := websiteObj.websiteMap[name]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(map[string]string{"error": "Name not in list"})
			w.Write(res)
			return
		}

		respJson, err := json.Marshal(map[string]string{name: websiteObj.websiteMap[name]})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
		return
	}

	respJson, err := json.Marshal(websiteObj.websiteMap)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func genericHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the server, the page you are looking for does not exist!")
}
