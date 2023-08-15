package main

import "fmt"


type teacher struct{
	name string
	age	int
	salary int
	ID int
}

type student struct{
	name string
	rank int
	Age int
}

type profiles interface {
	display()
}

func (a student)display(){
	fmt.Printf("Student name is: %s and rank is:  %d \n" ,a.name,a.rank)
}

func (b teacher)display(){
	fmt.Printf("Teacher name is: %s and salary is: %d \n", b.name, b.salary)
}

func common(p profiles){
	p.display()
}

func main(){

	Topper := student{name: "Teju", rank: 1, Age: 30}

	Mentor := teacher{name: "Ramu Sir", salary: 1000, age: 50, ID: 90}

	common(Topper)
	common(Mentor)



}

