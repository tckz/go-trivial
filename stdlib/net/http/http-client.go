package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"os"
)

var (
	optURL    = flag.String("url", "https://www.yahoo.co.jp", "")
	optBearer = flag.String("bearer", "", "")
)

// GODEBUG=http2debug=2
//  -> HTTP/2でやり取りすればstderrに出てくる
// GODEBUG=http2debug=2,http2client=0
//  -> http2クライアント無効

func main() {
	flag.Parse()

	req, err := http.NewRequest("GET", *optURL, nil)
	if err != nil {
		panic(err)
	}

	if *optBearer != "" {
		req.Header.Set("Authorization", "Bearer "+*optBearer)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(b)
	os.Stdout.WriteString("\n")

	b, err = httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(b)
	os.Stdout.WriteString("\n")
}
