package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/go-redis/redis"
)

func main() {
	optAddr := flag.String("addr", "localhost:6379", "addr:port")
	flag.Parse()

	cl := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{*optAddr},
	})
	defer cl.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	// redisへの接続がどのタイミングであるか確認するもの
	// この段階では接続establishedがない
	log.Printf("wait SIGINT 1")
	<-sigCh

	// 存在しない想定のキーをGET -> "", redis.Nil
	ret, err := cl.Get("52e10b46-267f-4608-a1cd-7d818939d988").Result()

	// poolの分が1つ残る
	log.Printf("wait SIGINT 2")
	<-sigCh

	// ret=, err=redis: nil redis.Nil?=true
	log.Printf("ret=%v, err=%v, redis.Nil?=%v", ret, err, err == redis.Nil)
}
