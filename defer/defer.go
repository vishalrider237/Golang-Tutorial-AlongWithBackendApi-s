package _defer

import "fmt"

func Defer() {
	fmt.Println("Starting of the program")
	defer fmt.Println("Middle of the program")
	defer fmt.Println("Middle of the program again")
	fmt.Println("End of the program")
}
