package array

import "fmt"

func Array() {
	var name [5]string
	name[0] = "Vishal"
	name[1] = "Vickey"
	fmt.Println("Name is:", name)
	number := []int{1, 2, 3, 4, 5}
	// Slice array
	fmt.Println("Length is:", len(number))
	number1 := make([]int, 3, 5)
	fmt.Println("Capacity is:", cap(number1))
}
