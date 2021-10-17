package main

import (
	"fmt"

	"go.mercari.io/go-bps/bps"
)

// https://engineering.mercari.com/blog/entry/20201203-basis-point/

func main() {
	n := bps.MustFromString("12.456")
	n = n.Add(bps.MustFromString("1.111"))
	// 13.567
	fmt.Println(n.FloatString(3))

	// 内部的にはbig.Intなので内部で1,000,000,000倍しても保持できる
	n2 := bps.MustFromString("9223372036854775808.123")
	// 9223372036854775808.123
	fmt.Println(n2.FloatString(3))
}
