package main

import (
    "fmt"
    "encoding/json"
)

type product struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Price float64 `json:"price"`
    Tags []string `json:"tags"`
}

func main() {
    // struct to JSON
    p1 := product{1, "Notebook", 1899.22, []string{"Mac", "i5"}}
    p1Json, _ := json.Marshal(p1)
    fmt.Println(string(p1Json))

    // JSON to struct
    var p2 product
    jsonString := `{"id":1,"name":"Pen","price":1.22,"tags":["x"]}`
    json.Unmarshal([]byte(jsonString), &p2)
    fmt.Println(p2)
}