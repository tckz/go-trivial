package main

import (
	"fmt"
	"os"
)

func main() {
	// nilなchanからのselect
	ch := make(chan struct{})
	ch = nil
	fmt.Fprintf(os.Stderr, "ch=%v\n", ch)

	select {
	case e := <-ch:
		fmt.Fprintf(os.Stderr, "case: %v\n", e)
	default:
		// 読み取り可能なデータがなくこっちを通る
		fmt.Fprintf(os.Stderr, "default\n")
	}
}
