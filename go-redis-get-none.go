package main

import (
	"context"
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

// goredisのGETでキーが存在しないときの挙動

func main() {

	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()
	addr := mr.Addr()

	cl := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{addr},
	})
	defer cl.Close()

	// キー自体がない -> "", redis.Nil
	ret, err := cl.Get(context.Background(), "not_exist").Result()
	// ret=, err=redis: nil, redis.Nil?=true
	log.Printf("ret=%v, err=%v, redis.Nil?=%v", ret, err, err == redis.Nil)
}
