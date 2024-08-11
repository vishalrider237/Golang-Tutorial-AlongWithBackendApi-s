package data_conv

import (
	"fmt"
	"strconv"
)

func Data() {
	num := 2
	decimal := float64(num)
	fmt.Println("Converted number is:", decimal)
	str := strconv.Itoa(num)
	str += "abc"
	fmt.Println("String data is:", str)

	num_str := "1234"
	num_int, _ := strconv.Atoi(num_str)
	fmt.Println("Integer data is:", num_int)
}
