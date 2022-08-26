package main

import (
	"fmt"
	"os"
)

type Some struct {
}

func (s *Some) Get(v string) string {
	return "ya"
}

func main() {
	some := Some{}
	// format=0x499020
	fmt.Fprintf(os.Stderr, "format=%+v\n", some.Get)
}
