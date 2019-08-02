package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// goroutine内でpanicが出たときのdeferの呼び出され方
// panicを出してないgoroutineのdeferは呼び出されない

func main() {
	defer fmt.Fprintf(os.Stderr, "main defer\n")

	fmt.Fprintf(os.Stderr, "start\n")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Fprintf(os.Stderr, "go-1 defer\n")
		time.Sleep(time.Second * 10)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Fprintf(os.Stderr, "go-2 defer\n")
		time.Sleep(time.Second * 10)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Fprintf(os.Stderr, "go-3 defer\n")
		time.Sleep(time.Second * 2)
		panic("Waoooo!")
	}()
	wg.Wait()
}
