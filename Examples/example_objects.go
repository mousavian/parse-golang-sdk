package parse

import (
    "fmt"
)


func ExampleObjects() {
    client := Client{
        AppId: "",
        RestKey: "",
        MasterKey: "",
    }

    var response map[string]interface{}
    
    // Objects
    objects := client.Object("Collection")

    // create
    newObject := make(map[string]interface{})
    newObject["name"] = "Gopher"
    newObject["age"] = 10
    response = objects.Create(newObject)
    objectId := response["objectId"].(string)
    newObject["objectId"] = objectId
    fmt.Println("Create: ", newObject["objectId"])

    // get
    obj1 := objects.Get(objectId)
    fmt.Println("Get: ", obj1)

    // update
    obj2 := objects.Get(objectId)
    obj2["name"] = "Gopher - Updated"
    obj2["age"] = 20
    obj3 := objects.Update(obj2)
    fmt.Println("Update: ", obj3["name"], obj3["age"])

    // fetch
    allObjects := objects.GetAll()
    for _, object := range allObjects {
        fmt.Println("object", object, "\n\r")
    }

    // delete
    response = objects.Delete(objectId)
    fmt.Println("Delete: ", response)
}
