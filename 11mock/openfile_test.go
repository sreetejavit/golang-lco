package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

type mockObject struct {
	file    *os.File
	err     error
	content []byte
}

func (m *mockObject) Create(name string) (*os.File, error) {

	fmt.Println("mock")
	return m.file, nil
}

func (m *mockObject) WriteValue(b []byte, filename string) {
	fmt.Println("New mockery writevalue")
}

func (m *mockObject) Close(file *os.File) error {
	return errors.New("Ting Ting mock error")
}

func TestActual(t *testing.T) {
	mockery := &mockObject{
		file: &os.File{},
	}

	Actual(mockery)

}
