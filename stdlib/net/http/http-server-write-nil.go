package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

// http.ResponseWriterにnilをWriteしたらどうなるか
// 何も起きない。（長さ0のbody）

func main() {

	m := &http.ServeMux{}
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(nil)
	})

	testServer := httptest.NewServer(m)
	defer testServer.Close()

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
