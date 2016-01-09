package parse

import (
    "fmt"
)


func Example() {
    client := Client{
        AppId: "",
        RestKey: "",
        MasterKey: "",
    }
    
    // Objects
    usersObject := client.Object("Users")

    // create
    newUser := make(map[string]interface{})
    newUser["name"] = "Gopher"
    newUser["age"] = 10
    response := usersObject.Create(newUser)
    objectId := response["objectId"].(string)
    newUser["objectId"] = objectId
    fmt.Println(newUser["objectId"])

    // get
    user1 := usersObject.Get(objectId)
    fmt.Println(user1)

    // update
    user2 := usersObject.Get(objectId)
    user2["name"] = "Gopher - Updated"
    user2["age"] = 20
    user3 := usersObject.Update(user2)
    fmt.Println(user3["name"], user3["age"])

    // fetch
    usersObjects := usersObject.GetAll()
    for _, user := range usersObjects {
        fmt.Println("user", user, "\n\r")
    }

    // delete
    usersObject.Delete(objectId)

}