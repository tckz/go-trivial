package main

import (
	"fmt"
	"time"
)

func main() {
	// こういう実装で翌日を求めるコードがあり、32日がどうなるか確認するもの
	// -> 2020/01/01になる

	now := time.Date(2019, 12, 31, 20, 0, 0, 0, time.Local)
	expireAt := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local).Unix()

	fmt.Printf("%d\n", expireAt)
}
