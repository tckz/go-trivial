package main

import (
	"errors"
	"fmt"
	"os"
)

func some() (err error) {
	// returnで指定した値は、deferの段階で返り値のところで宣言した変数に代入されているか->される
	defer func() {
		fmt.Fprintf(os.Stderr, "err=%v\n", err)
	}()
	return errors.New("error@here")
}

func main() {
	ret := some()
	fmt.Fprintf(os.Stderr, "ret=%v\n", ret)
}
