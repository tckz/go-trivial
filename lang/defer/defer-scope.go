package main

import (
	"log"
)

// deferはfuncスコープ、登録した逆順で実行
/*
2022/10/28 14:55:58 begin func
2022/10/28 14:55:58 e=a
2022/10/28 14:55:58 e=b
2022/10/28 14:55:58 e=c
2022/10/28 14:55:58 end func
2022/10/28 14:55:58 defer ee=c
2022/10/28 14:55:58 defer ee=b
2022/10/28 14:55:58 defer ee=a
*/
func main() {
	log.Print("begin func")
	for _, e := range []string{"a", "b", "c"} {
		ee := e
		log.Printf("e=%s", e)
		defer log.Printf("defer ee=%s\n", ee)
	}
	log.Print("end func")
}
