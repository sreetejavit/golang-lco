package main

import (
	"log"
	"os"
)

func main() {
	// Create a new logger
	logger := log.New(os.Stdout, "[MyApp] ", log.Ldate|log.Ltime)

	// Log messages with different severity levels
	logger.Println("This is a regular log message")
	logger.Printf("This is a formatted log message: %s", "Hello, World!")
	logger.Fatal("This is a fatal log message")

	// The program will exit after the fatal log message
}

