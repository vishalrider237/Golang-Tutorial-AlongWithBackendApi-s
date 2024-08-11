package userInput

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	fmt.Println("What's Your Name?")
	//var name string
	//fmt.Scan(&name)
	//fmt.Println("Hello Mr,", name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello Mr,", name)
}
