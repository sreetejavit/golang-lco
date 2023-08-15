package main

import (
	"fmt"
	"sync"
)
var global = 0
func main(){

	thread := new(sync.WaitGroup)
	thread.Add(1)
	go routine(thread)
	thread.Wait()
	thread.Add(1)
	go func(wg *sync.WaitGroup){
		fmt.Println("jdjsds")
		wg.Done()

	}(wg *sy)
	thread.Wait()
	thread.Add(1)
	go routine(thread)
	thread.Wait()
	thread.Add(1)
	go routine(thread)
	thread.Wait()








}



func routine(wg *sync.WaitGroup){
	global += 1
	routine := global
	fmt.Println("This is go routine", routine)
	defer wg.Done()
	fmt.Println("Ending this routine", routine)
}
