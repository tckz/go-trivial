package main

import (
	"context"
	"errors"
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
	defer lis.Close()

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
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("DumpRequest: %s", err)
			return
		}
		fmt.Fprintf(os.Stderr, "%s\n", dump)

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("ParseForm: %s", err)
			return
		}

		time.Sleep(*optWait + genJitter())

		fmt.Fprintf(os.Stderr, "PostForm=%v\n", r.PostForm)
		fmt.Fprintf(os.Stderr, "Form=%v\n", r.Form)
	})

	srv := &http.Server{
		Handler: m,
	}
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	go func() {
		defer cancel()
		if err := srv.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ctx.Done()
	log.Printf("signal: %v", context.Cause(ctx))

	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Shutdown: %s", err)
		}
	}
	log.Println("done")
}
