package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	type SomeType struct {
		Name string
		Age  int
	}

	// 1つのDecodeが終わったところで一回戻ってくるのでio.EOFまで繰り返し入力できる
	dec := json.NewDecoder(os.Stdin)
	for {
		var v SomeType
		err := dec.Decode(&v)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		fmt.Fprintf(os.Stderr, "%+v\n", v)
	}
}
