package main

import "fmt"

func main() {

	var one int
	var ptr *int
	ptr = &one

	fmt.Println(ptr, one, *ptr)
}
