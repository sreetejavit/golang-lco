package main

import "fmt"

type student struct {
	name string
	id   int
}

func (s student) changeName() {

	s.name = "bharath"
}

func main() {
	new := student{
		name: "teju",
		id:   1221,
	}

	mapi := map[string]int{"one": 1, "two": 2}
	fmt.Printf("The value of struct %+v\n", mapi)
	changemapi(mapi)
	fmt.Printf("The value of struct %+v\n", mapi)
	fmt.Printf("The value of struct %+v\n", new)
	new.changeName()
	fmt.Printf("The value of struct %+v\n", new)

}

func changemapi(mapi map[string]int) {
	mapi["two"] = 3
}
