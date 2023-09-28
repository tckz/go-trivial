package main

import (
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
)

func main() {
	type Error struct {
		Code int
	}

	type SomeType struct {
		Name string

		// 埋め込みされている場合、json構造としては {"Code":111}となる。
		// Marshalも同様で、Error階層は含まれず{"Name":"xxx","Code":1}となる。
		*Error
	}

	{
		var v SomeType
		lo.Must0(json.Unmarshal([]byte(`{"Name":"xxx","Code":111}`), &v))

		b := lo.Must(json.Marshal(v))
		// {"Name":"xxx","Code":111}
		fmt.Println(string(b))
	}

	type Error2 struct {
		Code int
	}
	type SomeType2 struct {
		Name string

		// 同じプロパティがmapされる型が複数埋め込まれている場合、無視されるっぽい
		*Error
		*Error2
	}

	{
		var v SomeType2
		lo.Must0(json.Unmarshal([]byte(`{"Name":"xxx","Code":111}`), &v))
		// Error=<nil>
		fmt.Printf("Error=%v\n", v.Error)
		// Error2=<nil>
		fmt.Printf("Error2=%v\n", v.Error2)

		b := lo.Must(json.Marshal(v))
		// {"Name":"xxx"}
		fmt.Println(string(b))
	}
}
