package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	fmt.Println(presentTime.Format("15-12-1992"))

	createDate := time.Date(2020, time.August, 12, 23, 4, 5, 3, time.UTC)
	fmt.Println(createDate)
}
