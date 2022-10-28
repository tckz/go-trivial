package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	optWait := flag.Duration("wait", time.Millisecond*100, "minimum optWait")
	optJitter := flag.Duration("jitter", time.Millisecond*50, "random optWait")
	optBind := flag.String("bind", ":8082", "Listen addr:port")
	flag.Parse()

	log.Printf("Bind: %s", *optBind)

	lis, err := net.Listen("tcp", *optBind)
	if err != nil {
		panic(err)
	}

	genJitter := func() time.Duration {
		return 0
	}
	if *optJitter > 0 {
		j := int64(*optJitter)
		genJitter = func() time.Duration {
			return time.Duration(rand.Int63n(j))
		}
	}
	m := &http.ServeMux{}
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("DumpRequest: %s", err)
			return
		}

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("ParseForm: %s", err)
			return
		}

		time.Sleep(*optWait + genJitter())

		fmt.Fprintf(os.Stderr, "%s\n", dump)
		fmt.Fprintf(os.Stderr, "PostForm=%v\n", r.PostForm)
		fmt.Fprintf(os.Stderr, "Form=%v\n", r.Form)
	})

	srv := &http.Server{
		Handler: m,
	}
	go func() {
		if err := srv.Serve(lis); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	s := <-sig
	log.Printf("Received signal: %s", s)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	srv.Shutdown(ctx)
	cancel()
	log.Println("done")
}
