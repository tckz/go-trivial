package main

import (
	"log"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/garyburd/redigo/redis"
)

// redigoのAPIはpoolからGetできなかったときもnilを返さない、常にClose出来る確認

func main() {

	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()
	addr := mr.Addr()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr,
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

	// おまけのTestOnBorrowがいつどういう値で呼び出されるか確認
	for i := 0; i < 3; i++ {
		func() {
			log.Printf("PING[%d]: before Get()", i)
			conn := pool.Get()
			defer func() {
				if err := conn.Close(); err != nil {
					log.Printf("PING[%d]: Close: err=%v", i, err)
				}
			}()
			ret, err := redis.String(conn.Do("PING"))
			if err != nil {
				log.Printf("PING[%d]: err=%v", i, err)
			} else {
				log.Printf("PING[%d]: ret=%s", i, ret)
			}
		}()
	}

	// redisを停止することでクライアントからのアクセスをエラーにする
	mr.Close()
	func() {
		log.Printf("PING[ex]: before Get()")
		conn := pool.Get()
		defer func() {
			if err := conn.Close(); err != nil {
				// Getの時点でエラーになってるがCloseを呼び出せて、ここに入ってくる
				log.Printf("PING[ex]: Close: err=%v", err)
			}
		}()
		ret, err := redis.String(conn.Do("PING"))
		if err != nil {
			log.Printf("PING[ex]: err=%v", err)
		} else {
			log.Printf("PING[ex]: ret=%s", ret)
		}
	}()
}
