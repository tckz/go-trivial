package main

import "fmt"

func main() {
	m := map[string]int{}

	// 存在しないキーを参照すると、ゼロ値が返る
	// 0
	v := m["not_exist"]
	fmt.Println(v)
}
