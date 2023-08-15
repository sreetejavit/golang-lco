package main
import "fmt"

// interface
//type Shape interface {
//	area() float32
//	perimeter() float32
//}

// Rectangle struct implements the interface
type Rectangle struct {
	length, breadth float32
}

type Square struct {
	length, breadth float32
}

// Rectangle provides implementation for area()
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func (s Square) area() float32 {
	return s.length * s.breadth
}

//// access method of the interface
//func calculate(s Shape) float32 {
//	return s.area()
//}

// main function
func main() {

	// assigns value to struct members
	r := Rectangle{7, 4}
	s := Square{2, 2}

	// call calculate() with struct variable rect
	fmt.Println(r.area())
	fmt.Println(s.area())
}
