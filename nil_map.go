package main

import (
	"fmt"
	"os"
)

type Some struct {
	MapBool   map[string]bool
	MapString map[string]string
}

// mapがnilのときにキー指定で値を参照するとどうなるか。 -> 普通に参照できる。指定キーが存在しない扱いになるようだ

func main() {
	s := &Some{}
	fmt.Fprintf(os.Stderr, "nil?=%v, v=%v\n", s.MapBool == nil, s.MapBool["who?"])
	fmt.Fprintf(os.Stderr, "nil?=%v, v=%v\n", s.MapString == nil, s.MapString["wao!"])
}
