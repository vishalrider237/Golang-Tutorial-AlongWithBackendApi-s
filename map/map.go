package _map

import (
	"fmt"
)

func Map() {
	studentGrade := make(map[string]int)
	studentGrade["Vishal"] = 89
	studentGrade["Vickey"] = 65
	studentGrade["Neha"] = 87
	fmt.Println("Marks of Vishal is:", studentGrade["Vishal"])
	studentGrade["Vishal"] = 76
	fmt.Println("New Marks of Vishal is:", studentGrade["Vishal"])
	delete(studentGrade, "Vishal")
	fmt.Println("Marks Check of Vishal is:", studentGrade["Vishal"])
	grade, exist := studentGrade["Bob"]
	fmt.Println("Marks  of David is:", grade)
	fmt.Println("Marks Check of David is:", exist)
	for index, value := range studentGrade {
		fmt.Printf("Key is %s and Value is %d\n", index, value)
	}
}
