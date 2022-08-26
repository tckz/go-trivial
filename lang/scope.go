package main

import (
	"errors"
	"fmt"
	"os"
)

func some2() (string, error) {
	return "oooo", errors.New("from some2")
}

func some() (err error) {
	v, err := some2()
	_ = v
	// to be overridden
	// 返り値の宣言のところにあるerrと同じerrを指すか？
	// -> 同じ。
	err = errors.New("return's err is overridden?")
	return
}

func main() {
	ret := some()
	fmt.Fprintf(os.Stderr, "%v\n", ret)
}
