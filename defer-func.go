package main

import "log"

func gen() func() {
	log.Print("gen() called")
	return func() {
		log.Print("generated func called")
	}
}

// deferのところに実際deferで実行する関数を返す関数xを書いた場合にxがいつ呼び出されるか->defer登録時点
// こういう順番
/*
2020/12/25 18:20:35 begin
2020/12/25 18:20:35 gen() called
2020/12/25 18:20:35 end of main
2020/12/25 18:20:35 generated func called
*/
func main() {
	log.Print("begin")
	defer gen()()
	log.Print("end of main")
}
