package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Code that may potentially panic
	fmt.Errorf("New error")

	// ...
	panic("Something went wrong!")
}

