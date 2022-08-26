package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// signal.Notifyに使っているchanをcloseしてしまうとsignalを受けたときにchanに書き込んでpanicする

func main() {
	sig := make(chan os.Signal, 10)
	signal.Notify(sig, syscall.SIGINT)

	close(sig)
	// sleepの間にCTRL+Cを押す
	log.Printf("Waiting SIGINT for 10 seconds ...")
	// ここで待ちを入れてsignalハンドラーが処理をする余地を持たせる。いきなりmainを抜けるとハンドラーがchanに書き込む前にプロセス終了してしまう
	time.Sleep(10 * time.Second)
	// chanがcloseしているので panic: send on closed channel になる
}
