package main

import (
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

// goredisのZSCOREで要素がないときの挙動

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

	// キー自体がない -> 0.0f, redis.Nil
	// ZSETに要素が1つもないときはキー自体がなくなるので存在しないキーに対して確認すればok
	ret, err := cl.ZScore("not_exist", "no_member").Result()
	log.Printf("ret=%f, err=%v, redis.Nil?=%v", ret, err, err == redis.Nil)

	// 1つ以上の要素が存在し、かつ指定したmemberが存在しない -> 0.0f, redis.Nil
	cl.ZAdd("exist", redis.Z{
		Score:  100,
		Member: "m1",
	})
	ret, err = cl.ZScore("exist", "no_member").Result()
	log.Printf("ret=%f, err=%v, redis.Nil?=%v", ret, err, err == redis.Nil)
}
