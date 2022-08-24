package main

import (
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

// goredisのSPOPで要素がないときの挙動

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
	// SETに要素が1つもないときはキー自体がなくなるので存在しないキーに対して確認すればok
	ret, err := cl.SPop("not_exist").Result()
	log.Printf("ret=%s, err=%v, redis.Nil?=%v", ret, err, err == redis.Nil)
}
