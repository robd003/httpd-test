package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        clientIP := r.RemoteAddr
	hostname := r.Host
	path := r.URL.Path
        fmt.Printf("Received request from %s to Hostname: %s Path: %s\n", clientIP, hostname, path)

	response := fmt.Sprintf("Test worked!\nHostname: %s\nPath: %s", hostname, path)
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
