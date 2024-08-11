package time

import (
	"fmt"
	"time"
)

func Time() {
	currentTime := time.Now()
	fmt.Println("Current Time is:", currentTime)
	formatted_date := currentTime.Format("2006-02-01 3:04 PM")
	fmt.Println("formatted time is:", formatted_date)
}
