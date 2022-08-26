package main

import (
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

// goredisのDELでキーが存在しないときの挙動

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

	// キー自体がない -> 0, nil
	ret, err := cl.Del("not_exist").Result()
	// ret=0, err=<nil>, redis.Nil?=false
	log.Printf("ret=%v, err=%v, redis.Nil?=%v", ret, err, err == redis.Nil)
}
