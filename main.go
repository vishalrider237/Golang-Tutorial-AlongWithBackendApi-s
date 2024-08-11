package main

import (
	"fmt"
	"mylearning/CrudWithMySql"
	"mylearning/array"
	"mylearning/crud"
	data_conv "mylearning/data-conv"
	_defer "mylearning/defer"
	"mylearning/error_handling"
	"mylearning/file"
	for_loop "mylearning/for-loop"
	"mylearning/function"
	"mylearning/goroutines"
	if_else "mylearning/if-else"
	"mylearning/json"
	_map "mylearning/map"
	"mylearning/myutil"
	"mylearning/pointer"
	string12 "mylearning/string"
	_struct "mylearning/struct"
	switch_case "mylearning/switch-case"
	"mylearning/syncWaitGoRoutine"
	"mylearning/time"
	"mylearning/url"
	"mylearning/web_request"
	"sync"
)

func main() {
	fmt.Println("Learning Go programming language!!")
	myutil.PrintMessage("Hello World")
	var myName string = "Vishal"
	fmt.Println(myName)
	var number = 1
	fmt.Println(number)
	persion := 12
	fmt.Println(persion)
	var public = "data is important"
	fmt.Println(public)
	var private = "data is private"
	fmt.Println(private)
	height := 5.765567
	fmt.Printf("height is %.3f\n", height)
	fmt.Printf("Age is %d\n", persion)
	fmt.Printf("Type of height is %T\n", height)
	fmt.Println("-----------------------------------------------")
	//userInput.UserInput()
	fmt.Println("-----------------------------------------------")
	ans := function.Add(3, 4)
	fmt.Println("Addition is:", ans)
	mul := function.Multiply(2, 5, 8)
	fmt.Println("Multiplication is:", mul)
	div, _ := error_handling.Divide(10, 0)
	fmt.Println("Division is:", div)
	array.Array()
	fmt.Println("-----------------------------------------------")
	var x = if_else.If_Else(4)
	fmt.Println("Condition  is:", x)

	var days = switch_case.Swtch_case(4)
	fmt.Println("Day name  is:", days)
	for_loop.For_Loop()
	fmt.Println("-----------------------------------------------")
	_map.Map()
	fmt.Println("-----------------------------------------------")
	_struct.Struct()
	fmt.Println("-----------------------------------------------")
	pointer.Pointer()
	fmt.Println("-----------------------------------------------")
	data_conv.Data()
	fmt.Println("-----------------------------------------------")
	string12.String()
	fmt.Println("-----------------------------------------------")
	time.Time()
	fmt.Println("-----------------------------------------------")
	_defer.Defer()
	fmt.Println("-----------------------------------------------")
	file.File()
	fmt.Println("-----------------------------------------------")
	web_request.ApiCall()
	fmt.Println("-----------------------------------------------")
	url.URL()
	fmt.Println("-----------------------------------------------")
	json.JSON()
	fmt.Println("-----------------------------------------------")
	crud.Get()
	crud.Post()
	crud.Update()
	crud.Delete()
	fmt.Println("-----------------------------------------------")
	go goroutines.SayHello()
	go goroutines.SayHii()
	fmt.Println("-----------------------------------------------")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go syncWaitGoRoutine.Woker(i, wg)
	}
	wg.Wait()
	fmt.Println("Worker task ended")
	fmt.Println("-----------------------------------------------")
	CrudWithMySql.Main()
}
