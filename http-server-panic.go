package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

// httpサーバーのハンドラからpanicが出たときの挙動
// サーバーごと落ちない、という確認

func main() {

	m := &http.ServeMux{}
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
	})
	m.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Wao!")
	})

	testServer := httptest.NewServer(m)
	defer testServer.Close()

	// panic
	func() {
		res, err := http.Get(testServer.URL + "/panic")
		if err != nil {
			log.Printf("*** http.Get: %v", err)
			return
		}
		defer res.Body.Close()

		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Printf("*** httputil.DumpResponse: %v", err)
			return
		}

		fmt.Println(string(dump))
	}()

	// panic出た後も普通にコードは進む
	func() {
		res, err := http.Get(testServer.URL)
		if err != nil {
			log.Printf("*** http.Get: %v", err)
			return
		}
		defer res.Body.Close()

		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Printf("*** httputil.DumpResponse: %v", err)
			return
		}

		fmt.Println(string(dump))
	}()
}
