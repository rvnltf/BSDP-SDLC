package main

import (
	"fmt"
	"net/http"
)

// handler untuk HTTP request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Mengirimkan response ke browser
	fmt.Fprintf(w, "Hello, World wkwk!")
}

func main() {
	// Menetapkan route '/hello' ke handler helloHandler
	http.HandleFunc("/", helloHandler)

	// Menjalankan HTTP server pada port 8080
	fmt.Println("Server is running on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
