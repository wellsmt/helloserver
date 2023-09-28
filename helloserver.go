package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a request handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Set the response header
		w.Header().Set("Content-Type", "text/plain")

		// Write a response to the client
		fmt.Fprintln(w, "Hello, World!")
	}

	// Register the handler function to handle all requests to "/"
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 80
	fmt.Println("Server is listening on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
