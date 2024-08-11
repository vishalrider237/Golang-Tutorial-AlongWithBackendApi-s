package string12

import (
	"fmt"
	"strings"
)

func String() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println("New Data is:", parts)

	str := "one two three one three two two"
	count := strings.Count(str, "two")
	fmt.Println("Count Data is:", count)
	str = "  Hello,Go!!  "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed Data is:", trimmed)
	str1 := "Vishal"
	str2 := "Pandey"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println("New String is:", result)

}
