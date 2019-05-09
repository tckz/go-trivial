package main

import (
	"fmt"
	"os"
)

type MyInterface interface {
	do() error
}

type client struct {
	MyValue int
}

func newClient() (MyInterface, error) {
	return client{MyValue: 3}, fmt.Errorf("this is error")
}

func (c client) do() error {
	fmt.Fprintf(os.Stderr, "myVal: %d\n", c.MyValue)
	return nil
}

func main() {
	// structを返してるので当然nilにならないけど、
	// こういうエラー判定をしているコードがあったので一応確認
	c, err := newClient()
	if c != nil {
		fmt.Fprintf(os.Stderr, "not nil desu\n")
	} else {
		fmt.Fprintf(os.Stderr, "nil desu\n")
	}

	_ = err
}
