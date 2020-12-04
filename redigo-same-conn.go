package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/alicebob/miniredis"
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

	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < *parallel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < *count; j++ {
				ret, err := redis.Int(conn.Do("INCR", "same-conn-check"))
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
