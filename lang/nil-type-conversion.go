package main

import (
	"fmt"
	"os"
)

func main() {
	// []byte(nil)をstringに変換すると空文字列になる
	var b []byte
	fmt.Fprintf(os.Stderr, "string(nil)=%v\n", string(b))
}
