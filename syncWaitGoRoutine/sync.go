package syncWaitGoRoutine

import (
	"fmt"
	"sync"
)

func Woker(i int, group sync.WaitGroup) {
	group.Done()
	fmt.Printf("Worker %d started \n", i)
	fmt.Printf("Worker %d ended \n", i)
}
