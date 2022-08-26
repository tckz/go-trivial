package main

import (
	"fmt"
	"os"
)

type ReceiverTest struct {
	Some  int
	Some2 int
}

func (r *ReceiverTest) Pointer(x int) *ReceiverTest {
	r.Some = x
	return r
}

func (r ReceiverTest) Value(x int) *ReceiverTest {
	r.Some = x
	return &r
}

func main() {

	r := ReceiverTest{Some: 1, Some2: 10}
	// Before:   0xc00001a0d0, &{Some:1 Some2:10}
	fmt.Fprintf(os.Stderr, "Before:   %p, %+v\n", &r, &r)

	// retはrと同じもの
	ret := r.Pointer(2)
	// Pointer:  0xc00001a0d0, &{Some:2 Some2:10}
	fmt.Fprintf(os.Stderr, "Pointer:  %p, %+v\n", ret, ret)

	// ret2はコピー
	ret2 := r.Value(3)
	// Value:    0xc00001a120, &{Some:3 Some2:10}
	fmt.Fprintf(os.Stderr, "Value:    %p, %+v\n", ret2, ret2)

	// Pointer()により書き換わっている
	// Original: 0xc00001a0d0, &{Some:2 Some2:10}
	fmt.Fprintf(os.Stderr, "After:    %p, %+v\n", &r, &r)
}
