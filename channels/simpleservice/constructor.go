package main


import (
	"fmt"
	"time"
)


type WellsFargo struct{
	User string
	Amount int
}

type BankAccount interface {
	amount() int
	creditcard() string
}

//func  NewWellsFargo() *WellsFargo{
//	return &WellsFargo{User: "Teju", Amount: 100}
//}

func (w *WellsFargo)amount()int{
	w.Amount *= 100
	return w.Amount
}
func (w *WellsFargo)creditcard()string{
	w.User += "credit"
	return w.User
}

func main(){

	Fargo := new(WellsFargo)

	fmt.Printf("%+v\n", Fargo)




	BankAcc := WellsFargo{
		User:   "Teju",
		Amount: 100,
	}

	fmt.Println(BankAcc.amount())


	go fmt.Println("go 1:",BankAcc.amount())

	time.Sleep(60*890)

	go fmt.Println("go 2:", BankAcc.amount())

	fmt.Println(BankAcc.creditcard())
	fmt.Println(BankAcc.Amount, BankAcc.amount())



}