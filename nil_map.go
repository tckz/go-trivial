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
	// nil?=true, v=false
	fmt.Fprintf(os.Stderr, "nil?=%v, v=%v\n", s.MapBool == nil, s.MapBool["who?"])
	// nil?=true, v=
	fmt.Fprintf(os.Stderr, "nil?=%v, v=%v\n", s.MapString == nil, s.MapString["wao!"])

	// mapのmap
	m := map[string]map[string]string{
		"key": map[string]string{},
	}
	// empty?=true, v=
	fmt.Fprintf(os.Stderr, "empty?=%v, v=%v\n", m["key"]["not_exist"] == "", m["key"]["not_exist"])
	// 最初の階層で存在しないキーを参照してさらにその配下を参照->普通に参照できる。指定キーが存在しない扱いのようだ
	// empty?=true, v=
	fmt.Fprintf(os.Stderr, "empty?=%v, v=%v\n", m["not_exist"]["not_exist2"] == "", m["not_exist"]["not_exist2"])
}
