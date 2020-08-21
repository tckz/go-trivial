package main

import (
	"context"
	"fmt"
	"os"
)

const (
	keyIntZero int = iota
)

type keyType int

const (
	keyTypeZero keyType = iota
	keyTypeOne
)

func main() {
	ctx := context.WithValue(context.Background(), keyIntZero, "zero")
	fmt.Fprintf(os.Stderr, "1-1: key(%v)=%v\n", keyIntZero, ctx.Value(keyIntZero))
	fmt.Fprintf(os.Stderr, "1-2: key(%v)=%v\n", 0, ctx.Value(0))
	// int(0)でWithValueしたものはkeyType(0)ではマッチしない
	fmt.Fprintf(os.Stderr, "1-3: key(%v)=%v\n", keyTypeZero, ctx.Value(keyTypeZero))
	fmt.Fprintf(os.Stderr, "1-4: %s\n", ctx)

	// typeとして個別の型を定義すれば実体int(1)であっても別のkeyとみなされる
	ctx = context.WithValue(context.Background(), keyTypeOne, "one")
	fmt.Fprintf(os.Stderr, "2-1: key(%v)=%v\n", keyTypeOne, ctx.Value(keyTypeOne))
	// keyType(1)でWithValueするとint(1)ではマッチしない。
	fmt.Fprintf(os.Stderr, "2-2: key(%v)=%v\n", 1, ctx.Value(1))
	fmt.Fprintf(os.Stderr, "2-3: %s\n", ctx)
}
