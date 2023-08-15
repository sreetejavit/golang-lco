package main

import "fmt"




type animal struct{
	name string
}

func (a *animal) display() {
	fmt.Println("Animal name is: " , a.name)
}


type Dog struct{
	animal
}
func main(){



	pup := animal{name: "Shitzu"}

	pup.display()

	Dog := Dog{animal{
		name: "Dolmashiom",
	}}
	Dog.display()




}

