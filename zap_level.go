package main

import "go.uber.org/zap"

// zapのlevelの値を確認
// -> info, warn, error, panic, fatal

func main() {
	zl, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	zl.Info("I am info")
	zl.Warn("I am warn")
	zl.Error("I am warn")
	func() {
		defer func() {
			_ = recover()
		}()
		zl.Panic("I am panic")
	}()
	zl.Fatal("I am fatal")
}
