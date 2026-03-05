package main

import (
	"fmt"
	"os"
)

func main() {
	// "a"
	fmt.Fprintf(os.Stderr, "%q\n", "a")
	// "a\"b\\c"
	fmt.Fprintf(os.Stderr, "%q\n", `a"b\c`)
	// '{' ASCII 123
	fmt.Fprintf(os.Stderr, "%q\n", 123)
}
