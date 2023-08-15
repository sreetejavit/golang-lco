package main

import "fmt"

func main() {
	// array := [3]string{"banana", "apple", "jack"}

	// for _, i := range array {
	// 	fmt.Println(i)
	// }

	// slice := []string{}
	// slice = append(slice, "banny", "tommy", "djsndjs")
	// for _, i := range slice {
	// 	fmt.Println(i)
	// }
	// fmt.Println(slice[1:3])

	var ptr *[]int

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 235
	highscore[2] = 236
	highscore[3] = 237

	numbers := []int{1, 2, 3, 4, 5}
	indexToDelete := 1

	// Remove the element at indexToDelete
	numbers = append(numbers[:indexToDelete], numbers[indexToDelete+1:]...)

	fmt.Println(numbers)

	ptr = &highscore

	fmt.Printf("%p\n", ptr)

	//u cant give value directly like this instead u can use append as it is slcie
	// highscore[4] = 237
	highscore = append(highscore, 238)
	highscore = append(highscore, 238)
	ptr = &highscore

	fmt.Printf("%p\n", ptr)

	fmt.Println(highscore)

	var stringp *string
	new := "yeh"
	stringp = &new

	fmt.Printf("%p", stringp)

	new = "yedwjkdw" + "mksmds"
	stringp = &new

	fmt.Printf("%p", stringp)

}
