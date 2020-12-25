package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type SomeType struct {
		Name string
		Age  int
	}

	// 両方: HTML絡みがUNICODEエスケープされる
	// Encoder: 最後に改行がつく
	/*
		'{"Name":"\u003cmyname\u003e","Age":25}'
		'{"Name":"\u003cmyname\u003e","Age":25}
		'<-改行
	*/

	v := SomeType{
		Name: "<myname>",
		Age:  25,
	}

	b, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "'%s'\n", string(b))

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err = enc.Encode(&v)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "'%s'\n", string(buf.Bytes()))
}
