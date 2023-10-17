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
	b := []int{}

	// a=[11 12 13 14 15]
	fmt.Printf("a=%v\n", r(func() interface{} { return a }))
	// a[:1]=[11]
	fmt.Printf("a[:1]=%v\n", r(func() interface{} { return a[:1] }))
	// a[:]=[11 12 13 14 15]
	fmt.Printf("a[:5]=%v\n", r(func() interface{} { return a[:5] }))
	// a[:6]=runtime error: slice bounds out of range [:6] with capacity 5
	fmt.Printf("a[:6]=%v\n", r(func() interface{} { return a[:6] }))
	// a[6:]=runtime error: slice bounds out of range [6:5]
	fmt.Printf("a[6:]=%v\n", r(func() interface{} { return a[6:] }))
	// a[:]=[11 12 13 14 15]
	fmt.Printf("a[:]=%v\n", r(func() interface{} { return a[:] }))
	// a[:]とした場合に別のsliceになるかどうか -> 同じ
	// a=0xc000122060, [:]=0xc000122060
	fmt.Printf("a=%p, a[:]=%p\n", a, a[:])
	// nilにappendして無理やり別のsliceに
	// a=0xc000122060, append(nil, a...)=0xc0001220c0
	fmt.Printf("a=%p, append(nil, a...)=%p\n", a, append([]int(nil), a...))
	// a[0:2]=[11 12]
	fmt.Printf("a[0:2]=%v\n", r(func() interface{} { return a[0:2] }))
	// a[1:2]=[12]
	fmt.Printf("a[1:2]=%v\n", r(func() interface{} { return a[1:2] }))
	// a[3:3]=[]
	fmt.Printf("a[3:3]=%v\n", r(func() interface{} { return a[3:3] }))
	// a[3:2]=runtime error: slice bounds out of range [3:2]
	// To avoid compile error
	i := 2
	fmt.Printf("a[3:2]=%v\n", r(func() interface{} { return a[3:i] }))
	// 長さ0のスライス化、capは維持される
	// a[:0]=[], cap(a)=5
	fmt.Printf("a[:0]=%v, cap(a)=%d\n", r(func() interface{} { return a[:0] }), cap(a))

	// b[1:]=runtime error: slice bounds out of range [1:0]
	fmt.Printf("b[1:]=%v\n", r(func() interface{} { return b[1:] }))
	// b[0:0]=[]
	fmt.Printf("b[0:0]=%v\n", r(func() interface{} { return b[0:0] }))

	// nilを長さ0のスライスに
	// c[:0]=[]
	var c []int
	fmt.Printf("c[:0]=%v\n", r(func() interface{} { return c[:0] }))
}
