package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type some0 struct {
	error
}

type some1 struct {
	error
}

func (s some1) Error() string {
	return "some1's:" + s.error.Error()
}

func main() {
	s0 := some0{error: errors.New("myerror0")}
	s1 := some1{error: errors.New("myerror1")}

	// myerror0
	fmt.Printf("some0.Error()=%s\n", s0.Error())
	// some1's:myerror1
	fmt.Printf("some1.Error()=%s\n", s1.Error())
}
