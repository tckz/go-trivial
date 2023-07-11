package main

import (
	"go.uber.org/zap"
)

func main() {
	zl, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// 同じフィールドをWithすると複数になる（上書きしない）
	logger := zl.Sugar().With(zap.Any("f1", "val1")).With(zap.Any("f1", "val1"))
	logger.Infof("ya")
}
