package main

import (
	"fmt"
	"os"
)

func main() {
	type SomeType struct {
		Key   string
		Value string
		Ptr   *string
	}

	ss := "sss"
	v := SomeType{Key: "k", Value: "v", Ptr: &ss}

	// %sはStringer非互換部分が書式間違い表現になる
	// {k v %!s(*string=0xc000014070)}
	fmt.Fprintf(os.Stderr, "%s\n", v)

	// %vは値だけ
	// {k v 0xc000014070}
	fmt.Fprintf(os.Stderr, "%v\n", v)

	// %+vはフィールド名と値
	// {Key:k Value:v Ptr:0xc000014070}
	fmt.Fprintf(os.Stderr, "%+v\n", v)

	// %#vはgoのコードに即した書式
	// main.SomeType{Key:"k", Value:"v", Ptr:(*string)(0xc000014070)}
	fmt.Fprintf(os.Stderr, "%#v\n", v)

	// いずれもポインターはポインター値のまま
}
