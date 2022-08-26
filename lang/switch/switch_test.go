package sw

import (
	"testing"
)

// go test -bench .
// これぐらいの件数で値がわかっているならswitchの方が効率がいい
/*
Benchmark_Switch-8      1000000000               1.07 ns/op
Benchmark_Map-8         187241085                5.99 ns/op
*/

type level int

const (
	DebugLevel level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
	Last
)

var tbl = map[level]string{
	DebugLevel:  "DEBUG",
	InfoLevel:   "INFO",
	WarnLevel:   "WARNING",
	ErrorLevel:  "ERROR",
	DPanicLevel: "CRITICAL",
	PanicLevel:  "ALERT",
	FatalLevel:  "EMERGENCY",
}

func lookupByMap(n level) string {
	v, _ := tbl[n]
	return v
}

func lookupBySwitch(n level) string {
	switch n % Last {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARNING"
	case ErrorLevel:
		return "ERROR"
	case DPanicLevel:
		return "CRITICAL"
	case PanicLevel:
		return "ALERT"
	case FatalLevel:
		return "EMERGENCY"
	}

	return ""
}

func Benchmark_Switch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lookupBySwitch(level(i))
	}
}

func Benchmark_Map(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lookupByMap(level(i))
	}
}
