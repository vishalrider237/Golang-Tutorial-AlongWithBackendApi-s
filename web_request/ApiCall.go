package web_request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ApiCall() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Error in reading the data", err)
		return
	}
	defer res.Body.Close()
	data, error := ioutil.ReadAll(res.Body)
	if error != nil {
		fmt.Println("Error in reading the data", error)
		return
	}
	fmt.Println("Json data:", string(data))
}
