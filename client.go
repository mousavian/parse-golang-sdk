package parse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const HOST_NAME string = "https://api.parse.com"

const API_VERSION string = "1"

type Client struct {
	AppId     string
	RestKey   string
	MasterKey string
}

func (client Client) getRequest(url string) []byte {
	return client.request("GET", url, nil)
}

func (client Client) postRequest(url string, params map[string]interface{}) []byte {
	return client.request("POST", url, params)
}

func (client Client) putRequest(url string, params map[string]interface{}) []byte {
	return client.request("PUT", url, params)
}

func (client Client) deleteRequest(url string) []byte {
	return client.request("DELETE", url, nil)
}

func (client Client) request(method string, url string, params map[string]interface{}) []byte {
	var response []byte
	httpClient := &http.Client{}

	jsonParams, _ := json.Marshal(params)
	paramsReader := bytes.NewReader(jsonParams)

	httpRequest, err := http.NewRequest(method, HOST_NAME+url, paramsReader)
	httpRequest.Header.Add("X-Parse-Application-Id", client.AppId)
	httpRequest.Header.Add("X-Parse-REST-API-Key", client.RestKey)
	// httpRequest.Header.Add("X-Parse-Master-Key", client.MasterKey)

	if currentSession.hasToken() {
		httpRequest.Header.Add("X-Parse-Session-Token", currentSession.token)
	}

	httpResponse, err := httpClient.Do(httpRequest)

	if err != nil {
		fmt.Printf("Err1: %s", err)
	} else {
		defer httpResponse.Body.Close()

		contents, err := ioutil.ReadAll(httpResponse.Body)

		if err != nil {
			fmt.Printf("Err2: %s", err)
		} else {
			response = contents
		}
	}

	return response
}
