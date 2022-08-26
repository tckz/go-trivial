package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// 開いたままファイルを削除してもfdは開いたまま、ということを確認するためのもの
// 実行してSleepしてる間に、/proc/PID/fd/配下をls -lする

func main() {
	file := flag.String("file", "", "/file/path/to/delete")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	err = os.Remove(f.Name())
	if err != nil {
		panic(err)
	}

	pid := os.Getpid()

	fmt.Fprintf(os.Stderr, "PID=%d, Sleeping....", pid)
	time.Sleep(100 * time.Second)
}
