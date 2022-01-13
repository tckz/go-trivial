package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	/*
		デフォルトでは
		- 最初の列数と後続の列数が同じでなければならない
		  -> FieldsPerRecordで変わる
		- フィールドの二重引用符はあってもなくてもいいし、混在できる
		  -> 値に二重引用符が含まれる場合は、フィールドの二重引用符必要
			 -> LazyQuotesで変えられる
		- 値に二重引用符を含める場合は""2つ重ねる
		- 最終行の行末改行はなくてもいい
		- 完全に空の行は無視される
		- クォートしていれば値に改行を含められる
	*/
	r := strings.NewReader(`col1,col2,col3

"quoted,value",c22,c23

"c31","c""quote""32","c3

3"`)

	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, rec := range records {
		fmt.Printf("[%d]len=%d, %#v\n", i, len(rec), rec)
	}
}
