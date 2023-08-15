package main

import "fmt"

func main(){

	a := 3

	channel := make(chan int)

	//go func() {
	//	a += 2
	//	channel <- 1
	//	close(channel)
	//	//fmt.Println(a)
	//}()

	go func() {
		a *=2
		channel <- 2
		close(channel)
		//fmt.Println(a)
	}()

	<- channel

	fmt.Println(a)
}
