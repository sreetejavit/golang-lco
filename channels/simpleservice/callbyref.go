// Program to illustrate call by reference

package main
import "fmt"

// call by value
func callByValue(num int) {

	num = 30
	fmt.Printf( "call by val: %d \n",num) // 30

}

// call by reference
func callByReference(num *int) {

	*num = 10
	fmt.Printf( "call by ref: %d \n",*num) // 10

}

func main() {

	number := 90

	// passing value
	callByValue(number)

	fmt.Printf( "call by main: %d \n",number)

	// passing a reference (address)
	callByReference(&number)
	fmt.Printf( "call by main: %d \n",number)

}

