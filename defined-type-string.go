package main

import (
	"fmt"
	"os"
)

type MyString string

func main() {
	s := MyString("hello")
	// s=hello
	fmt.Fprintf(os.Stderr, "s=%s\n", s)
}
