package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Websites struct {
	List []string `json:"websites"`
}

var websiteMap map[string]string = make(map[string]string)
var list []string

func main() {
	fmt.Println("Starting server")

	r := mux.NewRouter()

	r.HandleFunc("/", genericHandler)

	r.HandleFunc("/input", getWebsiteList).Methods("POST")

	r.HandleFunc("/check", checkWebsiteStatusHandler).Methods("GET")

	go func() {
		for {
			if list != nil {
				checkWebsiteStatus(list)
				time.Sleep(time.Minute)
			}
		}
	}()

	http.ListenAndServe("localhost:8080", r)

}

func getWebsiteList(w http.ResponseWriter, r *http.Request) {
	var website Websites
	err := json.NewDecoder(r.Body).Decode(&website)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list = website.List

	w.WriteHeader(http.StatusOK)

}

func checkWebsiteStatus(web []string) {
	fmt.Println("running checkWebsiteStatus")
	for _, data := range web {

		go func(data string) {
			response, err := http.Get("https://" + data)
			if err == nil {
				responseStatus := response.StatusCode
				if responseStatus == 200 {
					websiteMap[data] = "UP"
				} else {
					websiteMap[data] = "DOWN"
				}

			} else {
				websiteMap[data] = "DOWN"
			}
		}(data)

	}
}

func checkWebsiteStatusHandler(w http.ResponseWriter, r *http.Request) {

	if len(websiteMap) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("The map is empty"))
		return
	}

	name := r.URL.Query().Get("name")

	if name != "" {

		_, ok := websiteMap[name]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("name not in the list"))
			return
		}

		respJson, err := json.Marshal(map[string]string{name: websiteMap[name]})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
		return
	}

	respJson, err := json.Marshal(websiteMap)
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
