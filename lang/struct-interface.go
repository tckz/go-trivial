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

func newClient() MyInterface {
	return client{MyValue: 4}
}

func (c client) do() error {
	fmt.Fprintf(os.Stderr, "myVal: %d\n", c.MyValue)
	return nil
}

func main() {
	c := newClient()
	c.do()
}
