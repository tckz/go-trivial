package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	firstErr := errors.New("1st error")
	wrapped1st := errors.Wrapf(firstErr, "*** Failed to 1st")

	secondErr := errors.New("2nd error")
	wrapped2nd := errors.Wrapf(secondErr, "*** Failed to 2nd")

	err1 := errors.Wrapf(wrapped1st, "%+v", wrapped2nd)
	err2 := errors.Wrapf(wrapped1st, "%v", wrapped2nd)

	fmt.Fprintf(os.Stderr, "====1.wrapped1st %%s\n%s\n", wrapped1st)
	fmt.Fprintf(os.Stderr, "====1.wrapped1st %%v\n%v\n", wrapped1st)
	fmt.Fprintf(os.Stderr, "====1.wrapped1st %%+v\n%+v\n", wrapped1st)

	fmt.Fprintf(os.Stderr, "====3.err1 %%+v\n%+v\n", err1)
	fmt.Fprintf(os.Stderr, "====4.err2 %%+v\n%+v\n", err2)
}
