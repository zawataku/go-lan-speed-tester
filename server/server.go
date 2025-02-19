package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		data := make([]byte, 10*1024*1024)
		w.Write(data)
	})
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received %d bytes\n", r.ContentLength)
		w.WriteHeader(http.StatusOK)
	})
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
