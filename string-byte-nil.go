package main

import (
	"fmt"
	"os"
)

func main() {
	var b []byte
	// []byteのnilをstringにtype conversionすると空文字になる
	// s=
	fmt.Fprintf(os.Stderr, "s=%s\n", string(b))
}
