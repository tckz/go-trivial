package main

import "fmt"

func main() {
	var s interface{} = "some text"
	// 多値で受け取ってればtype assertionできなくても判別できる
	err, ok := s.(error)
	fmt.Printf("ok=%t\n", ok)
	// panic: interface conversion: string is not error: missing method Error
	err = s.(error)
	fmt.Printf("err=%v\n", err)
}
