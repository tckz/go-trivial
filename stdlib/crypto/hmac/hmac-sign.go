package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
)

var (
	optData = flag.String("data", "", "data to sign")
	optKey  = flag.String("key", "", "key to sign")
)

func main() {
	hm := hmac.New(sha256.New, []byte(*optKey))
	hm.Write([]byte(*optData))
	sig := hm.Sum(nil)

	fmt.Printf("%s\n", hex.EncodeToString(sig))
}
