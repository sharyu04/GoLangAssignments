package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWebsiteList(t *testing.T) {

	var testData = map[string][]string{
		"websites": {"www.google.com", "www.facebook.com", "www.fakewebsite1.com"},
	}
	expectedStatusCode := 200

	requestObj, _ := json.Marshal(testData)

	req, err := http.NewRequest(http.MethodPost, "/input", bytes.NewBuffer(requestObj))
	if err != nil {
		t.Errorf("Error occured while making http request. Error: %v", err.Error())
	}

	server := httptest.NewServer(http.HandlerFunc(getWebsiteList))
	resp, err := http.Post(server.URL, "JSON", req.Body)

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("getWebsiteList api endpoint falied. Error: %v", err.Error())
	}

}

func TestCheckWebsiteStatusHandler(t *testing.T) {

	tests := []struct {
		name               string
		queryName          string
		websiteObj         httpChecker
		expectedStatusCode int
	}{
		{
			name:               "Success for particular website",
			expectedStatusCode: http.StatusOK,
			queryName:          "www.facebook.com",
			websiteObj: httpChecker{map[string]string{
				"www.facebook.com":     "UP",
				"www.fakewebsite1.com": "DOWN",
				"www.google.com":       "UP",
			}},
		},
		{
			name:               "Empty map",
			websiteObj:         httpChecker{map[string]string{}},
			expectedStatusCode: 400,
		},
		{
			name: "Key not found",
			websiteObj: httpChecker{map[string]string{
				"www.facebook.com":     "UP",
				"www.fakewebsite1.com": "DOWN",
				"www.google.com":       "UP",
			}},
			expectedStatusCode: 400,
			queryName:          "www.abc.com",
		},
		{
			name: "Status of all websites",
			websiteObj: httpChecker{map[string]string{
				"www.facebook.com":     "UP",
				"www.fakewebsite1.com": "DOWN",
				"www.google.com":       "UP",
			}},
			expectedStatusCode: 200,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", fmt.Sprintf("/check?name=%s", test.queryName), bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(test.websiteObj.checkWebsiteStatusHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
