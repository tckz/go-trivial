package context

import (
	"context"
	"fmt"
	"time"
)

type noCancelContext struct {
	context.Context
}

func (*noCancelContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*noCancelContext) Done() <-chan struct{} {
	return nil
}

func (*noCancelContext) Err() error {
	return nil
}

func (c *noCancelContext) String() string {
	return fmt.Sprintf("%s.WithoutCancel", c.Context)
}

// WithoutCancel ctxのValueを引継ぎつつキャンセルチェーンを切り離したContextを返す
func WithoutCancel(ctx context.Context) context.Context {
	return &noCancelContext{
		Context: ctx,
	}
}
