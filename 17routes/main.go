package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "http://127.0.0.1:8081"

var count int32

func main() {

	server := &http.Server{
		Addr:    ":8081",
		Handler: nil, // Use the default router
	}
	// Define a handler function for processing incoming HTTP requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	handlerpost := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Post Request")
		count += 1
		fmt.Fprintf(w, "Hello"+string(count)+"for post")
	}

	// Register the handler function with the default HTTP server
	http.HandleFunc("/", handler)

	http.HandleFunc("/post", handlerpost)

	// Start the HTTP server on a specific port
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
	resp, err := http.Get(url)

	content, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(content))

	PerformPostReq()

	err = server.Shutdown(nil)
	if err != nil {
		fmt.Println("Server shutdown error:", err)
	} else {
		fmt.Println("Server stopped.")
	}

}

func PerformPostReq() {

	requestBody := strings.NewReader(`{
		"course": "wdwdww",
		"wswsw": "ssas"
	}`)

	resp, err := http.Post(url+"/post", "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	cont, _ := io.ReadAll(resp.Body)

	fmt.Println(string(cont))

}
