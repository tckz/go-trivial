package main

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
	// 親のctxがキャンセルされていることを確認
	assert.True(canceled)

	// 値は最初のctxも含めてつながる
	assert.Equal("val", ctxWithoutCancel.Value(dummyKey("key")))
	assert.Equal("val2", ctxWithoutCancel.Value(dummyKey("key2")))
	canceled2 := false
	select {
	case <-ctxWithoutCancel.Done():
		canceled2 = true
	default:
	}
	// 親がキャンセルされても自分はキャンセルされてない
	assert.False(canceled2)

	// go1.13からContext.String()の出力が変わった
	assert.Equal(`context.Background.WithValue(type context.dummyKey, val val).WithCancel.WithoutCancel.WithValue(type context.dummyKey, val val2)`, fmt.Sprintf("%s", ctxWithoutCancel))

	dead, ok := ctxWithoutCancel.Deadline()
	assert.False(ok, "no deadline is set")
	assert.Equal(time.Time{}, dead)

	assert.Nil(ctxWithoutCancel.Err(), "no error is set")
}

func TestNoCancelInheritDeadline(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	ctx = context.WithValue(ctx, dummyKey("key"), "val")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dead := time.Now().UTC().Add(time.Millisecond * 200)
	ctx, cancelDeadline := context.WithDeadline(ctx, dead)
	defer cancelDeadline()

	ctxWithoutCancel, cancelInheritDeadline := WithoutCancelInheritDeadline(ctx)
	defer cancelInheritDeadline()
	ctxWithoutCancel = context.WithValue(ctxWithoutCancel, dummyKey("key2"), "val2")

	cancel()
	canceled := false
	select {
	case <-ctx.Done():
		canceled = true
	default:
	}
	// 親のctxがキャンセルされていることを確認
	assert.True(canceled)

	// 値は最初のctxも含めてつながる
	assert.Equal("val", ctxWithoutCancel.Value(dummyKey("key")))
	assert.Equal("val2", ctxWithoutCancel.Value(dummyKey("key2")))
	canceled2 := false
	select {
	case <-ctxWithoutCancel.Done():
		canceled2 = true
	default:
	}
	// 親がキャンセルされても自分はキャンセルされてない
	assert.False(canceled2)

	d, ok := ctxWithoutCancel.Deadline()
	assert.True(ok, "deadline is set")
	assert.Equal(dead, d)
	assert.Nil(ctxWithoutCancel.Err(), "no error is set")

	canceled3 := false
	select {
	case <-time.After(time.Millisecond * 500):
	case <-ctxWithoutCancel.Done():
		canceled3 = true
	}
	// deadline経過でキャンセルされる
	assert.True(canceled3, "should deadline exceeded")
	assert.EqualError(ctxWithoutCancel.Err(), "context deadline exceeded")
}

func TestNoCancelNoInheritDeadline(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	ctx = context.WithValue(ctx, dummyKey("key"), "val")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 親のctxがdeadlineを持っていない場合

	ctxWithoutCancel, cancelInheritDeadline := WithoutCancelInheritDeadline(ctx)
	defer cancelInheritDeadline()
	ctxWithoutCancel = context.WithValue(ctxWithoutCancel, dummyKey("key2"), "val2")

	// 親のctxをキャンセルしキャンセルされていることを確認
	cancel()
	{
		canceled := false
		select {
		case <-ctx.Done():
			canceled = true
		default:
		}
		assert.True(canceled)
	}

	// 値は最初のctxも含めてつながる
	assert.Equal("val", ctxWithoutCancel.Value(dummyKey("key")))
	assert.Equal("val2", ctxWithoutCancel.Value(dummyKey("key2")))
	// 親がキャンセルされても自分はキャンセルされてない
	{
		canceled := false
		select {
		case <-ctxWithoutCancel.Done():
			canceled = true
		default:
		}
		assert.False(canceled)
	}

	d, ok := ctxWithoutCancel.Deadline()
	assert.False(ok, "no deadline is set")
	assert.Equal(time.Time{}, d)

	assert.Nil(ctxWithoutCancel.Err(), "no error is set")

	timeout := false
	select {
	case <-time.After(time.Millisecond * 200):
		timeout = true
	case <-ctxWithoutCancel.Done():
	}
	assert.True(timeout, "should be timed out")
	// timeout経過で抜けてきて、キャンセルはされていない
	{
		canceled := false
		select {
		case <-ctxWithoutCancel.Done():
			canceled = true
		default:
		}
		assert.False(canceled)
	}
}
