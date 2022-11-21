package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

// pathパラメーターをMWの中でも取得できるのことの確認

func main() {
	optBind := flag.String("bind", ":8080", "addr:port")
	flag.Parse()

	r := mux.NewRouter()

	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := mux.Vars(r)
			log.Printf("mux.Vars(r)=%v", v)
			h.ServeHTTP(w, r)
		})
	})

	v1 := r.PathPrefix("/v1/something/").Subrouter()
	v1.HandleFunc("/{pathparam}", func(w http.ResponseWriter, r *http.Request) {
		b, _ := httputil.DumpRequest(r, true)
		log.Printf("request: %s", string(b))
	}).Methods(http.MethodGet)

	if err := http.ListenAndServe(*optBind, r); err != nil {
		log.Fatalf("*** http.ListenAndServe: %s", err)
	}
}
