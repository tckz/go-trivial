package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

func main() {

	m := &http.ServeMux{}
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "hello\n")
	})

	testServer := httptest.NewServer(m)
	defer testServer.Close()

	vals := url.Values{}
	vals.Add("key", "param")

	// 500応答でもerrではない、という確認
	res, err := http.PostForm(testServer.URL, vals)
	if err != nil {
		log.Printf("err=%v", err)
		return
	}

	defer res.Body.Close()

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatalf("*** httputil.DumpResponse: %v", err)
	}

	fmt.Println(string(dump))
}
