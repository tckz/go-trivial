package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
)

// 1つのredigo.Connを複数のgRPCハンドラから共有しているコードがあり
// まずいんじゃないかということを確認したもの。
// -> 同じtcpセッション上で複数のcommandやり取りを書き込むのでエラーになる
// こういうのはPoolを使うべき、という確認

func main() {

	parallel := flag.Int("parallel", 8, "Number of parralelism")
	optAddr := flag.String("addr", "", "Addr of redis")
	count := flag.Int("count", 100000, "Number of trials")
	flag.Parse()

	addr := ""
	if *optAddr != "" {
		addr = *optAddr
	} else {
		mr, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		defer mr.Close()
		addr = mr.Addr()
	}

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr,
				redis.DialReadTimeout(3*time.Second),
				redis.DialWriteTimeout(3*time.Second),
				redis.DialConnectTimeout(3*time.Second),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		IdleTimeout: 10 * time.Second,
		Wait:        true,
		MaxIdle:     100,
	}
	defer pool.Close()

	wg := &sync.WaitGroup{}
	for i := 0; i < *parallel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < *count; j++ {
				conn := pool.Get()
				ret, err := redis.Int(conn.Do("INCR", "same-conn-check"))
				conn.Close()
				if err != nil {
					panic(err)
				}
				if ret%1000 == 0 {
					fmt.Fprintf(os.Stderr, "incr: %d\n", ret)
				}
			}
		}()
	}

	wg.Wait()
}
