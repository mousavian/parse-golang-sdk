package parse

import (
    "fmt"
)


func ExampleUsers() {
    client := Client{
        AppId: "",
        RestKey: "",
        MasterKey: "",
    }
    
    var username string = "admin"
    var password string = "123456"
    var email    string = "admin@localhost.dev"
    var response map[string]interface{}

    // Users
    users := client.User()

    // sign up
    signupUser := make(map[string]interface{})
    signupUser["username"] = username
    signupUser["password"] = password
    signupUser["email"] = email
    response = users.Signup(signupUser)
    objectId := response["objectId"].(string)
    signupUser["objectId"] = objectId
    fmt.Println("Signup: ", response)

    // login user
    response = users.Login(username, password)
    fmt.Println("Login: ", response)

    // logout user
    response = users.Logout()
    fmt.Println("Logout: ", response)
    
    validateUser := users.Me()
    fmt.Println("Me: ", validateUser)
    
    // get user
    user := users.Get(objectId)
    fmt.Println("Get: ", user)
    
    // update user
    user["email"] = "user@localhost.dev"
    updatedUser := users.Update(user)
    fmt.Println("Update: ", updatedUser)

    // get all users
    for _, user := range users.GetAll() {
        fmt.Println("user", user)
    }

    response = users.ResetPassword(user["email"])
    fmt.Println("ResetPassword:", response)
    
    // delete a user
    response = users.Delete(objectId)
    fmt.Println("Delete:", response)
}
