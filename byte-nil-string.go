package main

import (
	"fmt"
	"os"
)

func main() {
	var b []byte

	// 空文字になる
	s := string(b)
	fmt.Fprintf(os.Stderr, "nil?=%t, %s, empty?=%t\n", b == nil, s, s == "")
}
