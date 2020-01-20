package main

import (
	"fmt"
	"os"
)

func main() {
	var m map[string]interface{}
	fmt.Fprintf(os.Stderr, "m is nil? = %v\n", m == nil)

	// nilをrangeしても何も問題ない。0要素と同じ
	for k, v := range m {
		fmt.Fprintf(os.Stderr, "k,v = (%s, %s)\n", k, v)
	}
}
