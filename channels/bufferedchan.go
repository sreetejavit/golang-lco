package main

import "fmt"

func main(){
	Bfchannel := make(chan int, 2)

	go func() {
		for i := 1; i <= 5; i++ {
			Bfchannel <- i
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println(<- Bfchannel)
	}
}
