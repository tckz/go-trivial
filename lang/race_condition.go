package main

import (
	"fmt"
)

// -race付きで実行するとrace detectorが検知する
// GORACE="halt_on_error=1"

func main() {
	num := 0
	ch := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {
			num += 1
		}
		fmt.Printf("other: %d\n", num)
		ch <- 1
	}()
	for i := 0; i < 100000; i++ {
		num += 1
	}
	fmt.Printf("main: %d\n", num)
	<-ch
}
