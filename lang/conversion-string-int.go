package main

import (
	"fmt"
	"os"
)

func main() {
	s := string(40 + 1)

	// ")" runeで41
	// "41"にはならないよ、という話。
	fmt.Fprintf(os.Stderr, "%s\n", s)
}
