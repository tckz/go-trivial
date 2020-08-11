package main

import (
	"fmt"
	"os"
)

// closeしたchanから読み込むと特に何も起こらず

func main() {

	ch := make(chan int)

	close(ch)

	v, ok := <-ch
	// v=0, ok=false
	fmt.Fprintf(os.Stderr, "v=%v, ok=%t\n", v, ok)
}
