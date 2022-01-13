package main

import (
	"fmt"
	"os"
)

// ループ変数のメンバのポインタが示すものが全部同じになる件

func main() {
	type LoopVar struct {
		Value string
	}

	vars := []LoopVar{
		{Value: "1"},
		{Value: "2"},
		{Value: "3"},
	}

	var p []*string
	for _, e := range vars {
		p = append(p, &e.Value)
	}

	// [0xc000010230 0xc000010230 0xc000010230]
	fmt.Fprintf(os.Stderr, "%+v\n", p)

	/*
		[0]:3
		[1]:3
		[2]:3
	*/
	for i, e := range p {
		fmt.Fprintf(os.Stderr, "[%d]:%s\n", i, *e)
	}
}
