package main

import (
	"fmt"
)

// 範囲外のインデックスを指定した部分文字列（s2のケース）
// ->空文字になる

func main() {
	s1 := "123"
	s2 := "1"
	fmt.Printf("s1[1:]='%s'\n", s1[1:])
	fmt.Printf("s2[1:]='%s'\n", s2[1:])
}
