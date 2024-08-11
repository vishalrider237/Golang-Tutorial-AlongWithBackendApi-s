package json

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func JSON() {
	person := Person{Name: "John", Age: 24, IsAdult: true}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error in Marshalling!!", err)
		return
	}
	fmt.Println("Json data:", string(jsonData))
	var personData Person
	err1 := json.Unmarshal(jsonData, &personData)
	if err1 != nil {
		fmt.Println("Error in Marshalling!!", err)
		return
	}
	fmt.Println("Object data:", personData)
}
