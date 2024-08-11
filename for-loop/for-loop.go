package for_loop

import "fmt"

func For_Loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("value of i:", i)
	}
	counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 3 {
			fmt.Println("Counter is :", counter)
			break
		}
	}
	number := []int{1, 2, 3, 4, 5, 6}
	for index, value := range number {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
