package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		data := make([]byte, 100*1024*1024) // 100MB に変更
		w.Write(data)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		bytesReceived, err := io.Copy(io.Discard, r.Body) // 実際に受信したバイト数を計測
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Received %d bytes\n", bytesReceived)
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
