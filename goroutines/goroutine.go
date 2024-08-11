package goroutines

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello,World")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Program Ended")
}
func SayHii() {
	fmt.Println("Hii all!!")
}
