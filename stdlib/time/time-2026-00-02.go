package main

import (
	"fmt"
	"time"
)

func main() {
	// こういう実装で先月を求めるコードがあり、2026/00/02がどうなるか確認するもの
	// -> 2025/12/01になる

	now := time.Date(2026, 1, 2, 20, 0, 0, 0, time.Local)
	prevMonth := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location())

	// TZ=JST
	// 2025-12-01 00:00:00 +0900 JST
	fmt.Printf("%s\n", prevMonth)
}
