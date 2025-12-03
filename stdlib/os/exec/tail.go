package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/samber/lo"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("file must be specified")
	}

	ctx := context.Background()
	{
		c, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
		defer cancel()
		ctx = c
	}
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	f := os.Args[1]
	cmd := exec.CommandContext(ctx, "tail", "-c", "0", "-F", f)

	i := lo.Must(cmd.StdinPipe())
	i.Close()
	o := lo.Must(cmd.StdoutPipe())
	e := lo.Must(cmd.StderrPipe())

	lo.Must0(cmd.Start())

	w := func(name string, r io.ReadCloser) func() error {
		return func() (retErr error) {
			defer func() { cancel(retErr) }()

			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				log.Printf("%s: %s", name, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				if err != io.EOF {
					return fmt.Errorf("scanner.%s: %w", name, err)
				}
			}
			return
		}
	}

	go func() { _ = w("stdout", o)() }()
	go func() { _ = w("stderr", e)() }()

	lo.Must0(cmd.Wait())
}
