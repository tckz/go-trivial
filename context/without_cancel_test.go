package context

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type dummyKey string

func TestNoCancel(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	ctx = context.WithValue(ctx, dummyKey("key"), "val")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctxWithoutCancel := WithoutCancel(ctx)
	ctxWithoutCancel = context.WithValue(ctxWithoutCancel, dummyKey("key2"), "val2")

	cancel()
	canceled := false
	select {
	case <-ctx.Done():
		canceled = true
	default:
	}
	assert.True(canceled)

	assert.Equal("val", ctxWithoutCancel.Value(dummyKey("key")))
	assert.Equal("val2", ctxWithoutCancel.Value(dummyKey("key2")))
	canceled2 := false
	select {
	case <-ctxWithoutCancel.Done():
		canceled2 = true
	default:
	}
	assert.False(canceled2)

	// go1.13からContext.String()の出力が変わった
	assert.Equal(`context.Background.WithValue(type context.dummyKey, val val).WithCancel.WithoutCancel.WithValue(type context.dummyKey, val val2)`, fmt.Sprintf("%s", ctxWithoutCancel))

	dead, ok := ctxWithoutCancel.Deadline()
	assert.False(ok, "no deadline is set")
	assert.Equal(time.Time{}, dead)

	assert.Nil(ctxWithoutCancel.Err(), "no error is set")
}
