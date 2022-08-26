package main

import (
	"fmt"
	"os"
)

func main() {
	// []byteを%dでログ出力しているコードがあり、どういうことなのか確認したもの
	// [0, 1, 2]となる。
	// けど、httpレスポンスbodyをこの調子で出力したらログ出すぎじゃないかな

	b := []byte{0, 1, 2}
	fmt.Fprintf(os.Stderr, "%d\n", b)

}
