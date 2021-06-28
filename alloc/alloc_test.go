package alloc

import (
	"testing"
)

// スライスをポインターで返して、デリファレンスするとallocateがある、という確認
// go test -bench . -benchmem
/*
go1.15.12 test -bench . -benchmem
Benchmark_Pointer-8         5859            213174 ns/op          826007 B/op         21 allocs/op
Benchmark_Slice-8           7021            177595 ns/op          825973 B/op         20 allocs/op
*/
/*
go1.16だとallocateされない
Benchmark_Pointer-8         6999            174979 ns/op          825975 B/op         20 allocs/op
Benchmark_Slice-8           5743            181383 ns/op          825976 B/op         20 allocs/op
*/

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
