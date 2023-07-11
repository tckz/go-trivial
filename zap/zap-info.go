package main

import (
	"errors"

	"go.uber.org/zap"
)

type SomeID2 string

func (i *SomeID2) String() string {
	if i == nil {
		return ""
	}
	return string(*i)
}

func main() {
	zl, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	type SomeID string

	id1 := SomeID("id1")
	id2 := SomeID2("id2")

	logger := zl.Sugar()
	// Infoだから連結される: [%s][%s]0xc0000a0e30 id2
	logger.Info("Infoだから連結される: [%s][%s]", &id1, &id2)
	// ポインタと%s、ポインタ側にString()実装があれば意図通りに: [%!s(*main.SomeID=0xc000112e30)][id2]
	logger.Infof("ポインタと%%s、ポインタ側にString()実装があれば意図通りに: [%s][%s]", &id1, &id2)
	// %wは使えない
	logger.Infof("err=%w", errors.New("err da"))
}
