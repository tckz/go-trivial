package main

import (
	"fmt"
	"os"

	"golang.org/x/net/publicsuffix"
)

func main() {

	domain := "www.google.com"
	if len(os.Args) >= 2 {
		domain = os.Args[1]
	}
	s, icann := publicsuffix.PublicSuffix(domain)

	fmt.Fprintf(os.Stderr, "domain=%s, publicSuffix=%s, icann=%t\n", domain, s, icann)
}
