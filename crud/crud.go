package crud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Get() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error in reading the data", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in reading the data", res.Status)
	}
	//data, error := ioutil.ReadAll(res.Body)
	//if error != nil {
	//	fmt.Println("Error in reading the data", error)
	//	return
	//}
	//fmt.Println("Json data:", string(data))
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error in decoding the data", err)
		return
	}
	fmt.Println("Decoding data", todo)
}
func Post() {
	todo := Todo{
		UserId:    23,
		Title:     "Vishal Pandey",
		Completed: false,
	}
	// convert to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in Marshalling the data", err)
		return
	}
	// convert json data to string
	jsonString := string(jsonData)
	// convert string into reader
	jsonReader := strings.NewReader(jsonString)
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error in Sending the response", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Json response", string(data))
}
func Update() {
	todo := Todo{
		UserId:    23,
		Title:     "Vishal Pandey",
		Completed: false,
	}
	// convert to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in Marshalling the data", err)
		return
	}
	// convert json data to string
	jsonString := string(jsonData)
	// convert string into reader
	jsonReader := strings.NewReader(jsonString)
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", jsonReader)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, _ := client.Do(req)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Json response", string(data))
}
func Delete() {
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error in Deleting the data", err)
		return
	}
	client := http.Client{}
	res, _ := client.Do(req)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Json response", string(data))
	fmt.Println("Json response status", res.Status)
}
