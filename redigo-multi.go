package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/gomodule/redigo/redis"
)

// redigoのMULTI/EXECはどう値を返すのか確認

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

	func() {
		conn := pool.Get()
		defer conn.Close()

		conn.Do("SET", "key1", "val1")
		conn.Do("RPUSH", "key2", "val21", "val22")
		conn.Do("RPUSH", "key3", "val31", "val32", "val33")
	}()

	func() {
		conn := pool.Get()
		defer conn.Close()

		err := conn.Send("MULTI")
		if err != nil {
			panic(err)
		}

		err = conn.Send("LRANGE", "key2", "0", "-1")
		if err != nil {
			panic(err)
		}

		err = conn.Send("LRANGE", "key3", "0", "-1")
		if err != nil {
			panic(err)
		}

		vals, err := redis.Values(conn.Do("EXEC"))
		// vals=[[[118 97 108 50 49] [118 97 108 50 50]] [[118 97 108 51 49] [118 97 108 51 50] [118 97 108 51 51]]], len=2, err=<nil>
		fmt.Fprintf(os.Stderr, "vals=%+v, len=%d, err=%v\n", vals, len(vals), err)

		// vals[0]=[val21 val22], len=2
		// vals[1]=[val31 val32 val33], len=3
		for i, v := range vals {
			s, _ := redis.Strings(v, nil)
			fmt.Fprintf(os.Stderr, "vals[%d]=%+v, len=%d\n", i, s, len(s))
		}
	}()

	func() {
		conn := pool.Get()
		defer conn.Close()

		err := conn.Send("LRANGE", "key2", "0", "-1")
		if err != nil {
			panic(err)
		}

		// redigoの話ではなくredisの挙動だが、パイプラインのどこかでエラーがあった場合も途中の操作は取り消されない
		err = conn.Send("INCR", "keyinc")
		if err != nil {
			panic(err)
		}

		// typeが違うのでエラーになる
		err = conn.Send("ZCARD", "key3")
		if err != nil {
			panic(err)
		}

		vals, err := redis.Values(conn.Do("EXEC"))

		// コマンドのどれかがエラーになるとEXECの結果がそのコマンドのエラーになる。
		// どれがエラーかはわからない。

		// vals=[], len=0, err=WRONGTYPE Operation against a key holding the wrong kind of value
		fmt.Fprintf(os.Stderr, "vals=%+v, len=%d, err=%v\n", vals, len(vals), err)
	}()

	// redigoの話ではないが、MULTIの途中の操作は取り消されない
	func() {
		conn := pool.Get()
		defer conn.Close()

		// INCRした結果、1が得られる
		v, err := redis.Int(conn.Do("GET", "keyinc"))
		// v=1, err=<nil>
		fmt.Fprintf(os.Stderr, "v=%+v, err=%v\n", v, err)
	}()
}
