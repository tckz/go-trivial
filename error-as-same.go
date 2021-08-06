package main

import (
	"errors"
	"log"
	"net/url"
	"sync"
)

// GORACE="halt_on_error=1" go run -race error-as-same.go

// errors.Asで同じ場所を指しているコードがあったのでrace検知されるという確認

func main() {
	wg := &sync.WaitGroup{}
	var sameVar *url.Error
	for i := 0; i <= 30; i++ {
		n := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				var err error
				if n%2 == 0 {
					err = &url.Error{
						Op:  "some",
						URL: "http://localhost",
						Err: errors.New("wao"),
					}
				} else {
					err = errors.New("something")
				}

				if errors.As(err, &sameVar) {
					log.Printf("[%d]true", n)
				} else {
					log.Printf("[%d]false", n)
				}
			}
		}()
	}

	wg.Wait()
}
