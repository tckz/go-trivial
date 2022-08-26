package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type MyType struct {
	Name     string
	SomeTime time.Time
}

// nilのsliceへのポインタをMarshalすると普通に「null」

// なんかこういうコードがあったので挙動確認した

func main() {
	var m []MyType
	p := &m

	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	// Marshal=null
	fmt.Fprintf(os.Stderr, "Marshal=%s\n", string(b))
}
