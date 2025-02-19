package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	serverURL := "http://localhost:8080"

	fmt.Println("Mesuring download speed...")
	start := time.Now()
	resp, err := http.Get(serverURL + "/download")
	if err != nil {
		fmt.Println(("Error"), err)
		return
	}

	defer resp.Body.Close()
	bytesDownloaded, _ := io.Copy(io.Discard, resp.Body)
	duration := time.Since(start).Seconds()
	downloadSpeed := (float64(bytesDownloaded) * 8) / (1024 * 1024) / duration
	fmt.Printf("Download speed: %.2f Mbps\n", downloadSpeed)

	fmt.Println("Measuring upload speed...")
	start = time.Now()
	data := bytes.Repeat([]byte("A"), 10*1024*1024)
	resp, err = http.Post(serverURL+"/upload", "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		fmt.Println(("Error"), err)
		return
	}

	defer resp.Body.Close()
	duration = time.Since(start).Seconds()
	uploadSpeed := (float64(len(data)) * 8) / (1024 * 1024) / duration
	fmt.Printf("Upload speed: %.2f Mbps\n", uploadSpeed)
}
