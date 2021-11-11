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

	// 123
	fmt.Println(bps.NewFromAmount(123).Amounts())
	// 123.0
	fmt.Println(bps.NewFromAmount(123).FloatString(1))
	// 1233
	fmt.Println(bps.MustFromString("1233.778").Amounts())

	// cmp=0 (means equal)
	fmt.Printf("cmp=%d\n", bps.MustFromString("1233").Cmp(bps.NewFromAmount(1233)))

	// cmp=1
	fmt.Printf("cmp=%d\n", bps.MustFromString("1233").Cmp(bps.NewFromBaseUnit(1233)))

	// panic: can't convert  to BPS
	fmt.Println(bps.MustFromString("").Amounts())
}
