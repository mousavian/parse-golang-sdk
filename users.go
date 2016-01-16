package parse

import (
	"encoding/json"
	"fmt"
)

type UserType struct {
	client   *Client
	email    string
	username string
	password string
	objectId string
}

func (client Client) User() UserType {
	return UserType{client: &client}
}

func (self *UserType) Username() string {
	return self.username
}

func (self *UserType) Password() string {
	return self.password
}

func (self *UserType) Email() string {
	return self.email
}

func (self *UserType) ObjectId() string {
	return self.objectId
}

func (self *UserType) SetUsername(username string) {
	self.username = username
}

func (self *UserType) SetPassword(password string) {
	self.password = password
}

func (self *UserType) SetEmail(email string) {
	self.email = email
}

func (self *UserType) setObjectId(objectId string) {
	self.objectId = objectId
}

func (self *UserType) Signup() map[string]interface{} {
	var url string = fmt.Sprintf("/%s/users", API_VERSION)

	var result map[string]interface{}

	user := map[string]interface{}{
		"email":    self.email,
		"username": self.username,
		"password": self.password,
	}

	response := self.client.postRequest(url, user)

	json.Unmarshal(response, &result)

	if sessionToken, exists := result["sessionToken"]; exists {
		self.client.sessionToken = sessionToken.(string)
	}

	if objectId, exists := result["objectId"]; exists {
		self.objectId = objectId.(string)
	}

	return result
}

func (self *UserType) Login(username string, password string) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/login?username=%s&password=%s", API_VERSION, username, password)

	var result map[string]interface{}

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	if sessionToken, exists := result["sessionToken"]; exists {
		self.client.sessionToken = sessionToken.(string)
	}

	return result
}

func (self *UserType) Logout() map[string]interface{} {
	var url string = fmt.Sprintf("/%s/logout", API_VERSION)

	var result map[string]interface{}

	response := self.client.postRequest(url, nil)

	json.Unmarshal(response, &result)

	self.client.sessionToken = ""
	self.objectId = ""

	return result
}

func (self *UserType) Get(objectId string) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/users/%s", API_VERSION, objectId)

	var result map[string]interface{}

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result
}

func (self *UserType) Me() map[string]interface{} {
	var url string = fmt.Sprintf("/%s/users/me", API_VERSION)

	var result map[string]interface{}

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result
}

func (self *UserType) Update(user map[string]interface{}) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/users/%s", API_VERSION, user["objectId"])

	var result map[string]interface{}

	response := self.client.putRequest(url, user)

	json.Unmarshal(response, &result)

	return result
}

func (self *UserType) GetAll() []map[string]interface{} {
	var url string = fmt.Sprintf("/%s/users", API_VERSION)

	var result map[string][]map[string]interface{}

	response := self.client.getRequest(url)

	json.Unmarshal(response, &result)

	return result["results"]
}

func (self *UserType) Delete(objectId string) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/users/%s", API_VERSION, objectId)

	var result map[string]interface{}

	response := self.client.deleteRequest(url)

	json.Unmarshal(response, &result)

	return result
}

func (self *UserType) ResetPassword(emailAddress string) map[string]interface{} {
	var url string = fmt.Sprintf("/%s/requestPasswordReset", API_VERSION)

	var result map[string]interface{}

	data := make(map[string]interface{})

	data["email"] = emailAddress

	response := self.client.postRequest(url, data)

	json.Unmarshal(response, &result)

	return result
}

func (self *UserType) IsAuthenticated() bool {
	if len(self.client.sessionToken) > 0 {
		return true
	}

	return false
}
