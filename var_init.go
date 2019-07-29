package main

import (
	"fmt"
	"os"

	"github.com/tckz/trivial/subpack"
)

// importしたpackageにあるvarの初期化でpanicが出るときに
// mainのコードが走る前にpanicすることを確認

func main() {
	fmt.Fprintf(os.Stderr, "before call\n")
	s := subpack.Hello()
	fmt.Fprintf(os.Stderr, "ret=%s\n", s)
}
