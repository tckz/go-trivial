package main

import (
	"fmt"
	"os"
)

func main() {
	// %dにスライスを渡すと%vと同じようにスライスっぽい表現になる

	// []
	var s []int
	fmt.Fprintf(os.Stderr, "%d\n", s)

	// [1 2 3]
	s2 := []int{1, 2, 3}
	fmt.Fprintf(os.Stderr, "%d\n", s2)
}
