package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func measureDownloadSpeed(serverURL string) float64 {
	start := time.Now()
	resp, err := http.Get(serverURL + "/download")
	if err != nil {
		fmt.Println("Download error:", err)
		return 0
	}
	defer resp.Body.Close()

	bytesDownloaded, _ := io.Copy(io.Discard, resp.Body)
	duration := time.Since(start).Seconds()

	downloadSpeed := (float64(bytesDownloaded) * 8) / (1024 * 1024) / duration
	return downloadSpeed
}

func measureUploadSpeed(serverURL string) float64 {
	data := bytes.Repeat([]byte("A"), 100*1024*1024) // 100MB に変更
	start := time.Now()

	resp, err := http.Post(serverURL+"/upload", "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		fmt.Println("Upload error:", err)
		return 0
	}
	defer resp.Body.Close()

	duration := time.Since(start).Seconds()
	uploadSpeed := (float64(len(data)) * 8) / (1024 * 1024) / duration
	return uploadSpeed
}

func main() {
	serverURL := "http://localhost:8080"

	// 3回測定して平均を取る
	downloadSpeeds := []float64{}
	uploadSpeeds := []float64{}

	fmt.Println("Measuring download speed...")
	for i := 0; i < 3; i++ {
		speed := measureDownloadSpeed(serverURL)
		downloadSpeeds = append(downloadSpeeds, speed)
		fmt.Printf("Download speed %d: %.2f Mbps\n", i+1, speed)
	}
	averageDownload := (downloadSpeeds[0] + downloadSpeeds[1] + downloadSpeeds[2]) / 3
	fmt.Printf("Average Download Speed: %.2f Mbps\n", averageDownload)

	fmt.Println("Measuring upload speed...")
	for i := 0; i < 3; i++ {
		speed := measureUploadSpeed(serverURL)
		uploadSpeeds = append(uploadSpeeds, speed)
		fmt.Printf("Upload speed %d: %.2f Mbps\n", i+1, speed)
	}
	averageUpload := (uploadSpeeds[0] + uploadSpeeds[1] + uploadSpeeds[2]) / 3
	fmt.Printf("Average Upload Speed: %.2f Mbps\n", averageUpload)
}
