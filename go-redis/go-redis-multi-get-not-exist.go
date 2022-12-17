package main

import (
	"context"
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

// go-redisのMULTIはTxPipeline/TxPipelinedを使う
// GETでnot exist（redis.Nil）もTxPipelinedをエラーにするか？→する

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
		// 存在しない。multiじゃないただのGetの場合redis.Nilでエラーになる。TxPipelinedの戻りもエラー
		pipe.Get(ctx, "key1")
		pipe.Set(ctx, "key2", "val2", 0)
		pipe.Get(ctx, "key2")
		return nil
	})

	// TxPipelined全体ではエラー
	// err(proto.RedisError)=redis: nil
	log.Printf("err(%T)=%v", err, err)
	/*
		2022/12/16 19:01:31 rets[0]: fullname=get, err=redis: nil, String()=get key1: redis: nil
		2022/12/16 19:01:31 rets[1]: fullname=set, err=<nil>, String()=set key2 val2: OK
		2022/12/16 19:01:31 rets[2]: fullname=get, err=<nil>, String()=get key2: val2
	*/
	for i, e := range rets {
		log.Printf("rets[%d]: fullname=%s, err=%v, String()=%s", i, e.FullName(), e.Err(), e.String())
	}

	// エラーの後のSet分をgetできる
	// s=val2, err=<nil>
	s, err := cl.Get(ctx, "key2").Result()
	log.Printf("s=%s, err=%v", s, err)
}
