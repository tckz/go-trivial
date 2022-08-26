package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stderr, "len(args)=%d\n", len(os.Args))
	for i, e := range os.Args {
		fmt.Fprintf(os.Stderr, "args[%d]=%s\n", i, e)
	}

}
