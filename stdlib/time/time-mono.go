package main

import (
	"fmt"
	"time"
)

func main() {
	// 2019-06-03 10:28:37.992787873 +0900 JST m=+0.000211999
	// Go1.9で導入されたmonotonic clockがString()の結果に含まれる
	fmt.Println(time.Now().String())
}
