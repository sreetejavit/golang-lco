package main

import "fmt"

func main() {
	i := 0

	for {
		if i < 10 {

			goto lco
			i++

		}

	}

lco:
	fmt.Println("Label executed")
}
