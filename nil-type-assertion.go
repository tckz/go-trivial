package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

type someType struct{}
type someType2 struct{}

func main() {
	v1 := func() interface{} {
		var p *someType
		return p
	}

	// nilであってもpanicしない
	v := v1().(*someType)
	fmt.Fprintf(os.Stderr, "v=%v\n", v)

	// 型が合わないとpanic
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "panic=%v\n%s", r, string(debug.Stack()))
			}
		}()
		v := v1().(*someType2)
		fmt.Fprintf(os.Stderr, "v=%v\n", v)
	}()
}
