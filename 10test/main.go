package main

import (
	"flag"
	"fmt"

	S "test/sum"
)

var a, b int

func init() {

	flag.IntVar(&a, "a", 0, "a value")
	flag.IntVar(&b, "b", 0, "b value")
	flag.Parse()
	fmt.Println(a, b)
}

func main() {

	fmt.Println("This is the main function where code starts")

	fmt.Println("The sum of passed values is ", S.Sum(a, b))
}
