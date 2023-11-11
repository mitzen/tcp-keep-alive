package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving a connection")
	fmt.Fprintf(w, "Hello, World!")
	fmt.Println("setting on the wait time: to 10 min.")
	time.Sleep(1 * time.Minute)
	fmt.Println("Finish response.")
}

func main() {

	http.HandleFunc("/", handler)
	//idleTimeout := 5 * time.Minute
	idleTimeout := 1 * time.Second
	readTimeout := 1 * time.Second

	// Create an HTTP server with keep-alives enabled
	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: readTimeout,
	}

	fmt.Println("Starting server on port 8080 - version 1.0b")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
