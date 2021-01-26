package main

import (
	"fmt"
	"os"
)

func main() {
	var s []string
	fmt.Fprintf(os.Stderr, "s is nil? = %v\n", s == nil)

	// nilをrangeしても何も問題ない。0要素と同じ
	for i, e := range s {
		fmt.Fprintf(os.Stderr, "i, v = (%d, %s)\n", i, e)
	}
}
