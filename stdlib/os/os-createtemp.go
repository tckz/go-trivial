package main

import (
	"fmt"
	"os"
)

func main() {

	// CreateTempのコメントに書いてあるとおりだけど、
	// プロセスが終わったら一時ファイルが削除されるということはない、
	// という確認

	f, err := os.CreateTemp("/tmp", "myprefix")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "tempfile: %s\n", f.Name())

}
