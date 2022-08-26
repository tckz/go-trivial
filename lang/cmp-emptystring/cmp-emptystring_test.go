package cmp_emptystring

import (
	"testing"
)

// 空文字比較でcmp2の比較を書いてるのがいて優位性があるのかを調べた
// いずれも突出してるわけではないが、わざわざcmp2, cmp3にする理由はない。
/*
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/tckz/trivial/cmp-emptystring
Benchmark_EmptyString-8         1000000000               0.244 ns/op
Benchmark_LenLess1-8            1000000000               0.475 ns/op
Benchmark_LenEq0-8              1000000000               0.251 ns/op
*/

/* go1.17.1
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/tckz/go-trivial/cmp-emptystring
cpu: AMD Ryzen 9 3950X 16-Core Processor
Benchmark_EmptyString-8         1000000000               0.2350 ns/op
Benchmark_EmptyStringFalse-8    1000000000               0.2348 ns/op
Benchmark_LenLess1-8            1000000000               0.2393 ns/op
Benchmark_LenLess1False-8       1000000000               0.2393 ns/op
Benchmark_LenEq0-8              1000000000               0.2388 ns/op
Benchmark_LenEq0False-8         1000000000               0.2362 ns/op
*/

func cmp1(s string) bool {
	return s == ""
}

func cmp2(s string) bool {
	return len(s) < 1
}

func cmp3(s string) bool {
	return len(s) == 0
}

func Benchmark_EmptyString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := cmp1("")
		_ = b
	}
}

func Benchmark_EmptyStringFalse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := cmp1("not empty")
		_ = b
	}
}

func Benchmark_LenLess1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := cmp2("")
		_ = b
	}
}

func Benchmark_LenLess1False(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := cmp2("not empty")
		_ = b
	}
}

func Benchmark_LenEq0(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := cmp3("")
		_ = b
	}
}

func Benchmark_LenEq0False(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := cmp3("not empty")
		_ = b
	}
}
