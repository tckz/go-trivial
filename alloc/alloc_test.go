package alloc

import (
	"testing"
)

// スライスをポインターで返して、デリファレンスするとallocateがある、という確認
// go test -bench . -benchmem

func createSlice1() *[]string {
	var s []string
	for i := 0; i < 10000; i++ {
		s = append(s, "hell")
	}

	return &s
}

func createSlice2() []string {
	var s []string
	for i := 0; i < 10000; i++ {
		s = append(s, "hell")
	}

	return s
}

func Benchmark_Pointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := createSlice1()
		l := 0
		for _, e := range *s {
			l += len(e)
		}
	}
}

func Benchmark_Slice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := createSlice2()
		l := 0
		for _, e := range s {
			l += len(e)
		}
	}
}
