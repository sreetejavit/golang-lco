package main

import "fmt"

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

func doSomething() {
	defer recoverFunc()

	fmt.Println("Doing something...")
	panic("Oops! Something went wrong.")
	fmt.Println("This line will not be executed.")
}

func main() {
	doSomething()
	fmt.Println("Program continues executing after panic.")
}
