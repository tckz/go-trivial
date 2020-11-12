package main

import "fmt"

func r(f func() interface{}) (ret interface{}) {
	defer func() {
		if r := recover(); r != nil {
			ret = r
		}
	}()

	return f()
}

func main() {
	a := []int{11, 12, 13, 14, 15}

	// [:1]=[11]
	fmt.Printf("[:1]=%v\n", r(func() interface{} { return a[:1] }))
	// [:]=[11 12 13 14 15]
	fmt.Printf("[:5]=%v\n", r(func() interface{} { return a[:5] }))
	// [:6]=runtime error: slice bounds out of range [:6] with capacity 5
	fmt.Printf("[:6]=%v\n", r(func() interface{} { return a[:6] }))
	// [6:]=runtime error: slice bounds out of range [6:5]
	fmt.Printf("[6:]=%v\n", r(func() interface{} { return a[6:] }))
	// [:]=[11 12 13 14 15]
	fmt.Printf("[:]=%v\n", r(func() interface{} { return a[:] }))
	// [0:2]=[11 12]
	fmt.Printf("[0:2]=%v\n", r(func() interface{} { return a[0:2] }))
	// [3:3]=[]
	fmt.Printf("[3:3]=%v\n", r(func() interface{} { return a[3:3] }))
	// [3:2]=runtime error: slice bounds out of range [3:2]
	// To avoid compile error
	i := 2
	fmt.Printf("[3:2]=%v\n", r(func() interface{} { return a[3:i] }))
}
