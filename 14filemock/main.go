package mypackage

import (
	"fmt"
	"os"
)

// FileReader defines the interface for reading and closing files.
type FileReader interface {
	ReadFile(filename string) ([]byte, error)
	Close() error
}

// RealFileReader implements the FileReader interface using the actual file operations.
type RealFileReader struct{}

func (r *RealFileReader) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (r *RealFileReader) Close() error {
	// Close any resources associated with the file reader
	return nil
}

func ProcessFile(reader FileReader) error {
	data, err := reader.ReadFile("example.txt")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	err = reader.Close()
	if err != nil {
		return err
	}

	return nil
}
