package main

import (
	"log"
)

// deferはfuncスコープ、登録した逆順で実行

func main() {
	log.Print("begin func")
	for _, e := range []string{"a", "b", "c"} {
		ee := e
		defer log.Printf("ee=%s\n", ee)
	}
	log.Print("end func")
}
