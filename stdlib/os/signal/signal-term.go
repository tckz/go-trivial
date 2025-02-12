package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM)
	defer cancel()

	log.Printf("waiting for signal")

	// never done
	ctx = context.WithoutCancel(ctx)
	<-ctx.Done()

	log.Printf("signal received: %v", ctx.Err())
}
