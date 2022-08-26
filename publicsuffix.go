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

	// domain=www.google.com, publicSuffix=com, icann=true
	// etld=google.com, err=<nil>
	fmt.Fprintf(os.Stderr, "domain=%s, publicSuffix=%s, icann=%t\n", domain, s, icann)
	etld, err := publicsuffix.EffectiveTLDPlusOne(domain)
	fmt.Fprintf(os.Stderr, "etld=%s, err=%v\n", etld, err)
}
