package main

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

// WithoutCancelInheritDeadline キャンセルチェーンを切り離しつつ、deadlineがあれば引き継ぐ
// 例えば、クライアントの接続断からcancelされる状況で接続断されても処理継続したいが、
// Lambdaのように処理時間が決められていて上流からdeadlineを指定されている場合は従いたい、という状況。
func WithoutCancelInheritDeadline(ctx context.Context) (ret context.Context, cancel context.CancelFunc) {
	ret = WithoutCancel(ctx)
	if d, ok := ctx.Deadline(); ok {
		return context.WithDeadline(ret, d)
	} else {
		cancel = func() {}
	}
	return
}
