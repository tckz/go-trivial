package main

import (
	"fmt"
	"os"
)

func main() {
	// nilなchanからの読み取り
	ch := make(chan struct{})
	ch = nil
	fmt.Fprintf(os.Stderr, "ch=%v\n", ch)
	<-ch
	//  all goroutines are asleepでここには到達しない
	fmt.Fprintf(os.Stderr, "here\n")
}
