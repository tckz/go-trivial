package main

import (
	"flag"
	"fmt"
)

var (
	optBool        = flag.Bool("bool", false, "boolean option")
	optDefaultTrue = flag.Bool("default-true", true, "true option")
)

func main() {
	flag.Parse()

	// Bool型のflagでdefault=falseの場合、
	// 「-bool」「-bool=true」でtrueになる
	// 「当該オプション未指定」「-bool=false」でfalseになる
	// 「-bool true」は、「true」はargsに含まれフラグ値にはならない（非フラグの最初の引数以降はargs）

	// Bool型のflagでdefault=trueの場合、
	// 「当該オプション未指定」「-default-true」「-default-true=true」でtrueになる
	// 「-default-true=false」でfalseになる

	fmt.Printf("--bool=%t\n", *optBool)
	fmt.Printf("--default-true=%t\n", *optDefaultTrue)
	fmt.Printf("flag.Args=%#v\n", flag.Args())
}
