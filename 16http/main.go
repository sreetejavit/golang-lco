package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("LCO web request")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	databytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))

	defer resp.Body.Close()

}
