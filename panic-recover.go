package main

import "fmt"

func deferFunc(index int) {
	if r := recover(); r != nil {
		fmt.Printf("[%d]recover: %v\n", index, r)
	} else {
		fmt.Printf("[%d]no recover\n", index)
	}
}

func main() {
	defer deferFunc(0)
	defer deferFunc(1)

	panic("Nooooo")
}
