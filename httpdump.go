package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {

	bind := flag.String("bind", ":8082", "Listen addr:port")
	flag.Parse()

	fmt.Fprintf(os.Stderr, "Bind: %s\n", *bind)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(os.Stderr, "%q\n", dump)
	})
	http.ListenAndServe(*bind, nil)
}
