package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int

	// 64bit向けビルドだと8, 32bit向けだと4
	// GOARCH=386 go run sizeof.go
	fmt.Printf("sizeof(int)=%d\n", unsafe.Sizeof(i))
}
