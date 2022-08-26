package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	optName = flag.String("name", "", "some name")
)

func main() {
	flag.Parse()

	fmt.Printf("name=%s\n", *optName)

	fmt.Fprintf(os.Stderr, "len(flag.Args)=%d\n", len(flag.Args()))
	for i, e := range flag.Args() {
		fmt.Fprintf(os.Stderr, "flag.Args[%d]=%s\n", i, e)
	}

	fmt.Fprintf(os.Stderr, "len(os.Args)=%d\n", len(os.Args))
	for i, e := range os.Args {
		fmt.Fprintf(os.Stderr, "os.Args[%d]=%s\n", i, e)
	}

}
