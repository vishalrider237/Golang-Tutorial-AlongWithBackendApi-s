package pointer

import "fmt"

func Pointer() {
	var num int
	num = 2
	var ptr *int
	ptr = &num
	fmt.Println("Addres of variable ", num, " is ", ptr)
	val := 2
	ptr1 := &val
	fmt.Println("Data contains of this pointer is:", *ptr1)
	value := 20
	ModifyValue(&value)
	fmt.Println("Modified value is:", value)
}

func ModifyValue(num *int) {
	*num = *num * 10
}
