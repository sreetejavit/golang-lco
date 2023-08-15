package main

import "fmt"

type Val interface {
	int | float64 | string | rune
}

func AddUniversal[T Val](a T, b T) T {
	return a + b
}

func main() {

	//Add anything universal

	fmt.Println(AddUniversal(1, 2))
	fmt.Println(AddUniversal(1.0, 2.9))
	fmt.Println(AddUniversal('a', 'b'))
	fmt.Println(AddUniversal("a", "b"))

}
