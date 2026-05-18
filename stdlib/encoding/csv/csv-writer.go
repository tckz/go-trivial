package main

import (
	"encoding/csv"
	"os"
)

func main() {
	/*
		val1,val2,val3
		"val2,1","val2,2","val2,3"
		"val3""1""","val3""2""","val3""3"""
	*/
	w := csv.NewWriter(os.Stdout)
	w.WriteAll([][]string{
		{"val1", "val2", "val3"},
		// 値にクォート対象が含まれると列値が二重引用符で囲まれる
		{"val2,1", "val2,2", "val2,3"},
		// 列値に二重引用符が含まれると列値は二重引用符で囲まれ、値としての二重引用符部分に二重引用符が前置されて出力上は2つになる
		{`val3"1"`, `val3"2"`, `val3"3"`},
	})
	w.Flush()
}
