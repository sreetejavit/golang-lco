package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type fileOperations interface {
	WriteValue([]byte, string)
	Close(file *os.File) error
	Create(name string) (*os.File, error)
}

type fileo struct {
}

// func main() {
// 	fileop := &fileo{
// 		filename: "Teju",
// 	}

// 	Actual(fileop)

// }

func Actual(fileop fileOperations) {

	file, err := fileop.Create("fileop.filename")

	if err != nil {
		log.Fatal("Fatal error in creating file.")
	}

	content := []byte("Hello there..!")
	fmt.Println("Inside function Actaul")

	fileop.WriteValue(content, "fileop.filename")

	defer fileop.Close(file)
}
func (f *fileo) Create(name string) (*os.File, error) {
	file, err := os.Create(name)
	fmt.Println("Actial")
	return file, err
}

func (f *fileo) WriteValue(b []byte, filename string) {
	err := ioutil.WriteFile(filename, b, 0644)

	if err != nil {
		log.Fatal("Fatal error in creating file.")
		return
	}
	fmt.Println("Content written to file")

}

func (f *fileo) Close(file *os.File) error {
	err := file.Close()
	if err != nil {
		fmt.Println("error")
	}
	return nil
}
