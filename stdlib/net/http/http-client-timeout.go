package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	var testURL string

	ctx, cancelExternalClient := context.WithCancel(context.Background())
	defer cancelExternalClient()

	m := &http.ServeMux{}
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("root handler begin")
		defer log.Printf("root handler done")

		time.Sleep(time.Second * 2)
	})

	testServer := httptest.NewServer(m)
	defer testServer.Close()

	testURL = testServer.URL

	req, err := http.NewRequest("GET", testURL, nil)
	if err != nil {
		panic(err)
	}

	req = req.WithContext(ctx)

	cl := &http.Client{Timeout: 1 * time.Second}
	log.Printf("external-client.Do with timeout:%s", cl.Timeout)
	res, err := cl.Do(req)
	if err != nil {
		err := err
		for err != nil {
			// external-client: type=*url.Error, Get "http://127.0.0.1:45203": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
			// external-client: type=*http.httpError, context deadline exceeded (Client.Timeout exceeded while awaiting headers)
			log.Printf("external-client: type=%T, %v", err, err)

			if ue, ok := err.(interface{ Unwrap() error }); ok {
				err = ue.Unwrap()
			} else {
				break
			}
		}
		return
	}
	defer res.Body.Close()

	log.Printf("external-client done")
}
