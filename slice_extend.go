package main

import "fmt"

func main() {
	var a []int

	// sliceのキャパシティの増え方
	// len=1, cap=1
	// len=2, cap=2
	// len=3, cap=4
	// len=5, cap=8
	// len=9, cap=16
	// len=17, cap=32
	// len=33, cap=64
	// len=65, cap=128
	// len=129, cap=256
	// len=257, cap=512
	// len=513, cap=1024
	// len=1025, cap=1280
	// len=1281, cap=1696
	// len=1697, cap=2304
	// len=2305, cap=3072
	// len=3073, cap=4096
	// len=4097, cap=5120
	prevCap := cap(a)
	for i := 0; i < 5000; i++ {
		a = append(a, i)
		if c := cap(a); c != prevCap {
			fmt.Printf("len=%d, cap=%d\n", len(a), c)
			prevCap = c
		}
	}
}
