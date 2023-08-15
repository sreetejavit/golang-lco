package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string `json: "course"`
	ID    int    `json: "uniqueID"`
	price int    `json: "price, omitempty"`
}

func main() {

	lister := course{
		Name:  "dd",
		ID:    1,
		price: 100,
	}
	data, _ := json.MarshalIndent(lister, "", "\t")

	fmt.Printf("%s\n", data)

	DecodeJSon(data)

}

func DecodeJSon(content []byte) {
	mar := new(course)
	err := json.Unmarshal(content, &mar)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", content)
	fmt.Printf("%+v", mar)

}
