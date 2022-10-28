package main

import (
	"fmt"
	"io"
)

type someValue struct{}

type myCloser struct {
	name string
}

func (c *myCloser) Close() error {
	fmt.Printf("It's %s\n", c.name)
	return nil
}

func newSomeValue1() (*someValue, io.Closer, error) {
	return &someValue{}, &myCloser{name: "newSomeValue1"}, nil
}

func newSomeValue2() (*someValue, io.Closer, error) {
	return &someValue{}, &myCloser{name: "newSomeValue2"}, nil
}

func newSomeValue3() (*someValue, io.Closer, error) {
	return &someValue{}, &myCloser{name: "newSomeValue3"}, nil
}

/*
It's newSomeValue3
It's newSomeValue2
It's newSomeValue1
*/
func main() {
	v1, closer, err := newSomeValue1()
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	_ = v1

	v2, closer, err := newSomeValue2()
	if err != nil {
		panic(err)
	}
	// このdeferのところのcloserは変数としてはv1のところのcloserと同じものでnewSomeValue2の戻りで上書きされているが
	//  （さらにこの後newSomeValue3の返り値で上書きされる）
	// deferの仕様によりclose.Closer()の呼び出し直前の状態で値がcaptureされるため
	// ここのdeferはnewSomeValue2が返したcloserの値でClose()が呼び出される。
	// v1のところのdeferはnewSomeValue1が返したcloserの値でClose()が呼び出される。
	defer closer.Close()
	_ = v2

	v3, closer, err := newSomeValue3()
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	_ = v3
}
