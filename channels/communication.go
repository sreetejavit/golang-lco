package main

import (
	"fmt"
	"sync"
)

func main(){

	thread := new(sync.WaitGroup)

	thread.Add(1)

	channel := make(chan int)

	defer close(channel)

	go communicate(thread, channel)

	for i := 1; i <= 5; i++ {
		fmt.Println(<- channel)
	}


	thread.Wait()
}


func communicate(wg *sync.WaitGroup, channel chan int){
	for i := 1; i <= 5; i++ {
		channel <- i
	}
	defer wg.Done()
}