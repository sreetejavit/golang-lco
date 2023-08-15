package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Give ratimg for online practise..")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println(input)
	fmt.Printf("%T", input)

	convInput, _ := strconv.ParseInt(input, 10, 1)
	fmt.Printf("%T", convInput)

}
