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
	fmt.Fprintf(os.Stderr, "[0]v=%v\n", v)

	// nilであっても型があえばok=true
	v, ok := v1().(*someType)
	fmt.Fprintf(os.Stderr, "[1]v=%v, ok=%t\n", v, ok)

	// 型が合わなくてもokを受ければpanicせず判定できる
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "panic=%v\n%s", r, string(debug.Stack()))
			}
		}()
		v, ok := v1().(*someType2)
		fmt.Fprintf(os.Stderr, "[3]v=%v, ok=%t\n", v, ok)
	}()

	// 型が合わないとpanic
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "panic=%v\n%s", r, string(debug.Stack()))
			}
		}()
		v := v1().(*someType2)
		fmt.Fprintf(os.Stderr, "[4]v=%v\n", v)
	}()
}
