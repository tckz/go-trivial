package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	outFile := flag.String("out-file", "", "path/to/out/file")
	flag.Parse()

	// 0400で作成しつつ書き込めるか：書ける
	// ただし、すでに存在するとROなので書けない（openできない）
	f, err := os.OpenFile(*outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0400)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "hello!\n")

}
