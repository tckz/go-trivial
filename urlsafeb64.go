package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "*** arg(plain-text) must be specified\n")
		return
	}
	s := base64.RawURLEncoding.EncodeToString([]byte(os.Args[1]))
	fmt.Println(s)

	d, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))
}
