package main

import "fmt"

type addvalue struct {
	one int
	two int
}

func main() {

	fmt.Println("Write a smiple table test for add function")
	newstruct := addvalue{one: 2, two: 3}

	Add(newstruct.one, newstruct.two)
}

func Add(a int, b int) int {
	return a + b
}
