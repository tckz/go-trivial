package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// ioutil.TempFileのコメントに書いてあるとおりだけど、
	// プロセスが終わったら一時ファイルが削除されるということはない、
	// という確認

	f, err := ioutil.TempFile("/tmp", "myprefix")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "tempfile: %s\n", f.Name())

}
