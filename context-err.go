package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

// Contextがtimeoutしたときのctx.Err()はcontext.DeadlineExceededと同じか->同じ

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	fmt.Fprintf(os.Stderr, "1st: %v, %v\n", ctx.Err(), ctx.Err() == context.DeadlineExceeded)

	_ = <-ctx.Done()

	fmt.Fprintf(os.Stderr, "2nd: %v, %v\n", ctx.Err(), ctx.Err() == context.DeadlineExceeded)
}
