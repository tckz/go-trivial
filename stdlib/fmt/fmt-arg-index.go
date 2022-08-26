package main

import (
	"fmt"
	"os"
)

func main() {
	// 引数のインデックスを指定する書き方
	// s=hello
	fmt.Fprintf(os.Stderr, "s=%[2]s\n", 1, "hello")
}
