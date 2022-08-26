package main

import (
	"fmt"
	"os"
)

func main() {
	// nilなchanからのrange読み取り
	ch := make(chan struct{})
	ch = nil
	fmt.Fprintf(os.Stderr, "ch=%v\n", ch)

	for e := range ch {
		fmt.Fprintf(os.Stderr, "here: %v\n", e)
	}
	//  all goroutines are asleepでここには到達しない
	fmt.Fprintf(os.Stderr, "here\n")
}
