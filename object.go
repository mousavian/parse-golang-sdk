package parse

import (
	"encoding/json"
	"fmt"
)

type ObjectType struct {
	client    *Client
	className string
}

func (client Client) Object(className string) ObjectType {
	return ObjectType{&client, className}
}

func (self ObjectType) GetAll() []map[string]interface{} {
	var url string = fmt.Sprintf("/%s/classes/%s/", API_VERSION, self.className)

	var result map[string][]map[string]interface{}

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result["results"]
}

func (self ObjectType) Get(objectId string) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/classes/%s/%s/", API_VERSION, self.className, objectId)

	var result map[string]interface{}

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result
}

func (self ObjectType) Create(object map[string]interface{}) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/classes/%s", API_VERSION, self.className)

	var result map[string]interface{}

	response := self.client.postRequest(url, object)

	json.Unmarshal(response, &result)

	return result
}

func (self ObjectType) Update(object map[string]interface{}) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/classes/%s/%s", API_VERSION, self.className, object["objectId"])

	var result map[string]interface{}

	response := self.client.putRequest(url, object)

	json.Unmarshal(response, &result)

	return result
}

func (self ObjectType) Delete(objectId string) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/classes/%s/%s", API_VERSION, self.className, objectId)

	var result map[string]interface{}

	response := self.client.deleteRequest(url)

	json.Unmarshal(response, &result)

	return result
}
