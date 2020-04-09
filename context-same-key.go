package main

import (
	"context"
	"fmt"
	"os"
)

const ctxKey = iota

func main() {
	ctx := context.Background()

	// 同じキーで複数回、値を保持した場合にどの値が取得できるか->最後の値（リスト構造のように最後のcontextからたどる）
	ctx = context.WithValue(ctx, ctxKey, "1")
	ctx = context.WithValue(ctx, ctxKey, "2")
	ctx = context.WithValue(ctx, ctxKey, "3")

	fmt.Fprintf(os.Stderr, "%s\n", ctx.Value(ctxKey))
}
