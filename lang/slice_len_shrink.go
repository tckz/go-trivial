package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	fmt.Printf("len=%d, cap=%d: %v\n", len(a), cap(a), a)

	// 確保した領域はそのままでサイズ（len）だけ縮める
	a = a[:0]
	fmt.Printf("len=%d, cap=%d: %v\n", len(a), cap(a), a)

	a = append(a, "d")
	fmt.Printf("len=%d, cap=%d: %v\n", len(a), cap(a), a)

	// 領域の内容は元のままなのでlenだけ延ばすと以前入れた値が見える
	a = a[:3]
	fmt.Printf("len=%d, cap=%d: %v\n", len(a), cap(a), a)
}
