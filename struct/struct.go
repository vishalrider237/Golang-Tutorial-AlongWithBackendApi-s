package _struct

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}
type Contact struct {
	email string
	phone string
}
type Address struct {
	house_no int
	state    string
}
type Employee struct {
	employee_detail  Person
	employee_contact Contact
	employee_address Address
}

func Struct() {
	var vishal Person
	vishal.firstName = "Vishal"
	vishal.lastName = "Pandey"
	vishal.age = 24
	//fmt.Println("Vishal details is :", vishal)

	//person1 := Person{
	//	firstName: "Vickey",
	//	lastName:  "Pandey",
	//	age:       30,
	//}
	//fmt.Println("Vickey details is :", person1)

	var person2 = new(Person)
	person2.firstName = "Neha"
	person2.lastName = "Pandey"
	person2.age = 28
	//fmt.Println("Neha details is :", person2)
	var employee Employee
	employee.employee_detail = Person{
		firstName: "Vishal",
		lastName:  "Pandey",
		age:       24,
	}
	employee.employee_contact.phone = "123566"
	employee.employee_contact.email = "a@gmail.com"

	employee.employee_address = Address{
		state:    "Bihar",
		house_no: 23,
	}
	fmt.Println("Employee details is :", employee)
}
