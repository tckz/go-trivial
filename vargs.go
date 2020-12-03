package main

import (
	"fmt"
	"os"
)

func sub(name string, args ...string) {
	fmt.Fprintf(os.Stderr, "name=%s, args.nil?=%t, len=%d, args=%v\n", name, args == nil, len(args), args)
}

func main() {
	// name=case1, args.nil?=false, len=1, args=[a]
	sub("case1", "a")
	// name=case2, args.nil?=false, len=2, args=[a b]
	sub("case2", "a", "b")
	// name=case3, args.nil?=true, len=0, args=[]
	// argsはnilになる。空のsliceではない
	sub("case3")
}
