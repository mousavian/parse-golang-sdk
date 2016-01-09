package parse

import (
	"fmt"
	"encoding/json"
)

type UserType struct {
	client *Client
}

func (client Client) User() UserType {
	return UserType{&client}
}

func (self UserType) Signup(user map[string]interface{}) (map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/users", API_VERSION)

	var result map[string]interface{}

	response := self.client.postRequest(url, user)

	json.Unmarshal(response, &result)

	return result
}

func (self UserType) Login(username string, password string) (map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/login", API_VERSION)
	
	var result map[string]interface{}
	
	// TODO
	// 

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result
}

func (self UserType) Logout() (map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/logout", API_VERSION)

	var result map[string]interface{}

	// TODO:
	// Add header: "X-Parse-Session-Token"

	response := self.client.postRequest(url, nil)

	json.Unmarshal(response, &result)

	return result
}

func (self UserType) Get(objectId string) (map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/users/%s", API_VERSION, objectId)

	var result map[string]interface{}

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result
}

func (self UserType) Me() (map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/users/me", API_VERSION)
	
	var result map[string]interface{}
	
	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result
}

func (self UserType) Update(user map[string]interface{}) (map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/users/%s", API_VERSION, user["objectId"])

	var result map[string]interface{}

	response := self.client.putRequest(url, user)

	json.Unmarshal(response, &result)

	return result
}

func (self UserType) GetAll() ([]map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/users", API_VERSION)

	var result map[string][]map[string]interface{}

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result["results"]
}

func (self UserType) Delete(objectId string) (map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/users/%s", API_VERSION, objectId)
	
	var result map[string]interface{}
	
	response := self.client.deleteRequest(url)

	json.Unmarshal(response, &result)

	return result
}

func (self UserType) ResetPassword(emailAddress string) (map[string]interface{}) {
	var url string = fmt.Sprintf("/%s/requestPasswordReset", API_VERSION)

	var result map[string]interface{}

	data := make(map[string]interface{})

	data["email"] = emailAddress

	response := self.client.postRequest(url, data)

	json.Unmarshal(response, &result)

	return result
}
