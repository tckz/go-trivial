package main

import (
	"fmt"
	"os"
)

type some struct {
	field int
}

func (s *some) someFunc() {
	// レシーバーがnilでも普通に呼び出せる
	fmt.Fprintf(os.Stderr, "someFunc: %v\n", s)
}

func main() {

	s1 := &some{}
	s1.someFunc()
	fmt.Fprintf(os.Stderr, "s1: %v\n", s1.field)

	var s2 *some
	s2.someFunc()
	// nilに対してフィールドを参照するのはpanic
	fmt.Fprintf(os.Stderr, "s2: %v\n", s2.field)
}
