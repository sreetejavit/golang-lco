package main

import "fmt"

func main() {
	fmt.Println("Hello to multi argumnet")
	do(4, 5, 6)
}

func do(c, a, b int) {
	fmt.Println(c, b, a)
}
