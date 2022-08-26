package main

import (
	"fmt"
)

// 範囲外のインデックスを指定した部分文字列（s2のケース）

func main() {
	s1 := "123"
	// s1[1:]='23'
	fmt.Printf("s1[1:]='%s'\n", s1[1:])

	s2 := "1"
	// こっちは空文字
	// s2[1:]=''
	fmt.Printf("s2[1:]='%s'\n", s2[1:])
	// こっちはpanic
	// panic: runtime error: slice bounds out of range [2:1]
	fmt.Printf("s2[2:]='%s'\n", s2[2:])
}
