package file

import (
	"fmt"
	"os"
)

func File() {
	// creating file
	//file, err := os.Create("example.txt")
	//if err != nil {
	//	fmt.Println("Error in creating file!!", err)
	//	return
	//}
	//defer file.Close()
	//content := "Hello World by vishal!!"
	//_, error := io.WriteString(file, content+"\n")
	//if error != nil {
	//	fmt.Println("Error in writing in file!!", error)
	//	return
	//}
	//fmt.Println("Successfully file created!!")

	// Reading file first method
	//file, err := os.Open("example.txt")
	//if err != nil {
	//	fmt.Println("Error in openning file!!", err)
	//	return
	//}
	//defer file.Close()
	//buffer := make([]byte, 1024)
	//for {
	//	n, error := file.Read(buffer)
	//	if error == io.EOF {
	//		break
	//	}
	//	if error != nil {
	//		fmt.Println("Error in reading file!!", error)
	//		return
	//	}
	//	fmt.Println(string(buffer[:n]))
	//}

	// 2nd method
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error in reading file!!", err)
		return
	}
	fmt.Println(string(content))
}
