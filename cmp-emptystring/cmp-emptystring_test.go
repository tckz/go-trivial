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

func Benchmark_LenLess1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := cmp2("")
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
