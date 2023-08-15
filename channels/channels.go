package main

import (
	"fmt"
	"time"
)

var hello string

func main(){


	func(){

		fmt.Println("Hello go programmer")
	}()

	fmt.Println("Hello Teju")
	Govinda := "Tirumal Vasa Govinda"


	Hari := fmt.Sprintf("%s", Govinda)
	read := make(chan string)

	go func(){
		channel := Hari
		read <- channel

		fmt.Println("read last")
	}()

	go func(){
		channel := "new value"
		fmt.Println("read 1")
		read <- channel

		fmt.Println(<- read)
	}()

	fmt.Println("time sleep done")
	go func(){
		fmt.Println("read 2")

		hello:=  <- read
		fmt.Println(hello)
	}()

	time.Sleep(120 * time.Second)







}