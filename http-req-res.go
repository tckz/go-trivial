package main

import (
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://example.jp", nil)
	if err != nil {
		log.Fatalf("*** http.NewRequest: %v", err)
	}

	// Newした段階でreq.Responseがnilかどうかを知りたい -> nilになってる
	log.Printf("req=%#v\b", req)
}
