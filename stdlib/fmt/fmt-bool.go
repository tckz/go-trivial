package main

import (
	"fmt"
	"os"
)

func main() {
	// boolは%tって話

	// true
	fmt.Fprintf(os.Stderr, "%t\n", true)
	// %!s(bool=true)
	fmt.Fprintf(os.Stderr, "%s\n", true)
}
