package main

import "fmt"

func main() {
	// [a b c]
	fmt.Printf("%s\n", []string{"a", "b", "c"})

	// [a b c]
	fmt.Printf("%v\n", []string{"a", "b", "c"})

	// []string{"a", "b", "c"}
	fmt.Printf("%#v\n", []string{"a", "b", "c"})
}
