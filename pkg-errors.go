package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("1st error")
	werr := errors.Wrapf(err, "*** Failed to 1st")

	// *** Failed to 1st: 1st error
	fmt.Fprintf(os.Stderr, "====%%s\n%s\n", werr)
	// *** Failed to 1st: 1st error
	fmt.Fprintf(os.Stderr, "====%%v\n%v\n", werr)
	// *** Failed to 1st: 1st error
	fmt.Fprintf(os.Stderr, "====%%#v\n%v\n", werr)
	/*
		1st error
		main.main
		        /home/aki/go/src/github.com/tckz/go-trivial/errors.go:11
		runtime.main
		        /home/aki/go1.13/src/runtime/proc.go:203
		runtime.goexit
		        /home/aki/go1.13/src/runtime/asm_amd64.s:1357
		*** Failed to 1st
		main.main
		        /home/aki/go/src/github.com/tckz/go-trivial/errors.go:12
		runtime.main
		        /home/aki/go1.13/src/runtime/proc.go:203
		runtime.goexit
		        /home/aki/go1.13/src/runtime/asm_amd64.s:1357
	*/
	fmt.Fprintf(os.Stderr, "====%%+v\n%+v\n", werr)
}
