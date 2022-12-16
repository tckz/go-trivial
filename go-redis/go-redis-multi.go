package main

import (
	"context"
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

// go-redisのMULTIはTxPipeline/TxPipelinedを使う

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

	ctx := context.Background()
	rets, err := cl.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, "key1", "val1", 0)
		// typeが違うのでエラー
		pipe.ZCard(ctx, "key1")
		// MULTIの途中でエラーがあっても継続
		pipe.Set(ctx, "key2", "val2", 0)
		pipe.Get(ctx, "key2")
		return nil
	})

	// TxPipelined全体ではエラー
	// err(proto.RedisError)=WRONGTYPE Operation against a key holding the wrong kind of value
	log.Printf("err(%T)=%v", err, err)
	/*
		rets[0]: fullname=set, err=<nil>, String()=set key1 val1: OK
		rets[1]: fullname=zcard, err=WRONGTYPE Operation against a key holding the wrong kind of value, String()=zcard key1: WRONGTYPE Operation against a key holding the wrong kind of value
		rets[2]: fullname=set, err=<nil>, String()=set key2 val2: OK
		rets[3]: fullname=get, err=<nil>, String()=get key2: val2
	*/
	for i, e := range rets {
		log.Printf("rets[%d]: fullname=%s, err=%v, String()=%s", i, e.FullName(), e.Err(), e.String())
	}

	// エラーが起きる前のstepでsetした内容をgetできる
	// s=val1, err=<nil>
	s, err := cl.Get(ctx, "key1").Result()
	log.Printf("s=%s, err=%v", s, err)
}
