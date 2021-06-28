package main

import "testing"

// go test -bench . -benchmem

/*
goos: linux
goarch: amd64
pkg: github.com/tckz/go-trivial/slice-copy
cpu: AMD Ryzen 9 3950X 16-Core Processor
Benchmark_Copy-8                        15184999                78.88 ns/op            0 B/op          0 allocs/op
Benchmark_AppendOneTime-8                1300498               921.9 ns/op          4096 B/op          1 allocs/op
Benchmark_SetEach-8                      9600442               124.1 ns/op             0 B/op          0 allocs/op
Benchmark_AppendEach-8                   5433493               222.5 ns/op             0 B/op          0 allocs/op
Benchmark_AppendEachFromNil-8             512550              2258 ns/op            8184 B/op         10 allocs/op
*/
const size = 500

// 同じ要素数のsliceをmakeしてcopy
func Benchmark_Copy(b *testing.B) {
	src := make([]int, size)
	for i, _ := range src {
		src[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]int, size)
		copy(dst, src)
	}
}

// nilスライスにsrcを可変引数で展開してappend
func Benchmark_AppendOneTime(b *testing.B) {
	src := make([]int, size)
	for i, _ := range src {
		src[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := append(([]int)(nil), src...)
		_ = dst
	}
}

// 同じ要素数のsliceをmakeして各要素を自前で設定
func Benchmark_SetEach(b *testing.B) {
	src := make([]int, size)
	for i, _ := range src {
		src[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]int, size)
		for j, e := range src {
			dst[j] = e
		}
	}
}

// Benchmark_AppendEach 事前に同キャパシティのスライスを用意して1要素ずつappend
func Benchmark_AppendEach(b *testing.B) {
	src := make([]int, size)
	for i, _ := range src {
		src[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]int, 0, size)
		for _, e := range src {
			dst = append(dst, e)
		}
	}
}

// Benchmark_AppendEachFromNil おまけ。nilスライスにappendしつつ拡張
func Benchmark_AppendEachFromNil(b *testing.B) {
	src := make([]int, size)
	for i, _ := range src {
		src[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst []int
		for _, e := range src {
			dst = append(dst, e)
		}
	}
}
