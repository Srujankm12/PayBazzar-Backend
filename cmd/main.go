package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	e := InitializeApp()
	resp, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		log.Println("⚠️ Failed to get public IP:", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		log.Println("🌐 Render public IP (send this to RechargeKit to whitelist):", string(body))
	}

	e.Logger.Fatal(e.Start(":8080"))
	log.Println("🚀 Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
