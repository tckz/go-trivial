package main

import (
	"fmt"
	"time"
)

// -race付きで実行するとrace detectorが検知する
// GORACE="halt_on_error=1"

func main() {
	num := 10
	go func() {
		num += 1
		fmt.Println(num)
	}()
	num += 1
	fmt.Println(num)
	time.Sleep(time.Second)
}
