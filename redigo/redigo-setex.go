package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 古いredigoとredis（2.8.6）でSETにEXオプション使えることの確認用。。。

func main() {
	addr := flag.String("addr", "", "redis addr host:port")
	flag.Parse()

	if *addr == "" {
		log.Fatal("--addr must be specified")
	}

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", *addr,
				redis.DialReadTimeout(3*time.Second),
				redis.DialWriteTimeout(3*time.Second),
				redis.DialConnectTimeout(3*time.Second),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			log.Printf("TestOnBorrow: t=%s", t)
			_, err := c.Do("PING")
			return err
		},
		IdleTimeout: 10 * time.Second,
		Wait:        true,
		MaxIdle:     100,
	}
	defer pool.Close()

	func() {
		conn := pool.Get()
		defer conn.Close()

		err := conn.Send("SET", "key1", "val1")
		if err != nil {
			panic(err)
		}
		i, err := redis.Int(conn.Do("EXPIRE", "key1", 30))
		fmt.Fprintf(os.Stderr, "i=%d, err=%v\n", i, err)
	}()

	func() {
		conn := pool.Get()
		defer conn.Close()

		s, err := redis.String(conn.Do("SET", "key2", "val2", "EX", 30))
		fmt.Fprintf(os.Stderr, "s=%s, err=%v\n", s, err)
	}()
}
