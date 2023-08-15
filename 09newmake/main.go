package main

import "fmt"

func main() {

	l := new(int)
	fmt.Print(*l)

	p := make(map[any]any)
	// p[1] = "jskasa"
	// p["ssa"] = 9
	fmt.Println(p)
}
 