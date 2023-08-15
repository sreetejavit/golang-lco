package main

import "fmt"

func main() {
	newmap := make(map[string]string)

	newmap["1"] = "sw"
	newmap["2"] = "dwdw"

	fmt.Println(newmap)

	for i, k := range newmap {
		fmt.Println(i, k)
	}

}
