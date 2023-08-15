package main

import "fmt"

func main() {
	var r rune = 'A'      // Assigning a rune literal 'A' to the variable r
	fmt.Printf("%c\n", r) // Printing the rune value as a character

	// Iterating over a string using range, which returns the rune value of each character
	str := "Hello, 世界"
	for _, char := range str {
		fmt.Printf("%c ", char) // Printing each rune value as a character
	}
}
