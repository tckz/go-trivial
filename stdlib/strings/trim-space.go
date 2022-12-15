package main

import (
	"fmt"
	"strings"
)

func main() {
	// #0:'全角スペース'
	fmt.Printf("#0:'%s'\n", strings.TrimSpace("　全角スペース"))
	// #1:'改行'
	fmt.Printf("#1:'%s'\n", strings.TrimSpace("\n改行\n\n\n"))
}
