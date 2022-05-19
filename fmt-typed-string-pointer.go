package main

import "fmt"

type SomeID2 string

func (i *SomeID2) String() string {
	if i == nil {
		return ""
	}
	return string(*i)
}

func main() {
	type SomeID string

	id1 := SomeID("id1")
	id2 := SomeID2("id2")

	// id1=id1
	fmt.Printf("id1=%s\n", id1)
	// &id1=%!s(*main.SomeID=0xc000010250)
	fmt.Printf("&id1=%s\n", &id1)
	// &id2=id2
	// ポインタ側にString()実装があればポインタの値ではなく指し示す文字列出力に
	fmt.Printf("&id2=%s\n", &id2)
}
