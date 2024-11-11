package main

import (
	"fmt"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/samber/lo"
	"github.com/tckz/go-trivial/goja/common"
)

func main() {
	vm := goja.New()

	// https://github.com/dop251/goja/issues/116
	require.NewRegistry().Enable(vm)
	vm.SetParserOptions(parser.WithDisableSourceMaps)
	console.Enable(vm)

	pg := lo.Must(common.Compile(`
function jsFunc(v) {
	// 2024/04/02 12:00:22 {"Code":1,"Msg":"hello","Map":{}}
	console.log(JSON.stringify(v));
	// vはポインターで渡しているので値の変更を反映できる
	v.Code = 999;
	// go側で変更後の値を取得できる
	v.Map["key3"] = "val3";
	const a = {ResultCode: 123, Status: "ok", unexportedValue: 77, AnotherCode: 66, notExistValue: 55};
	return a;
}
`))

	lo.Must(common.RunProgram(vm, pg))

	type Param struct {
		// デフォルトではjsonタグは関係ないがTagFieldNameMapperを明示することでjsonタグが反映される
		Code int `json:"code"`
		Msg  string
		Map  map[string]any

		// js側で見えない
		unexportedValue int
	}
	type Result struct {
		ResultCode int
		Status     string
		// デフォルトではjsonタグは関係ないがTagFieldNameMapperを明示することでjsonタグが反映される
		AnotherCode int `json:"-"`

		// jsから返してもgo側に反映されない
		unexportedValue int
	}

	// パラメーターをポインターで渡せばJS側の値の変更が反映される
	var jsFunc func(*Param) (Result, error)

	if f := vm.Get("jsFunc"); f == nil {
		panic("jsFunc not found")
	} else if err := vm.ExportTo(f, &jsFunc); err != nil {
		panic(fmt.Errorf("ExportTo: %w", err))
	}

	param := Param{Code: 1, Msg: "hello", Map: map[string]any{}, unexportedValue: 88}
	v, err := jsFunc(&param)

	// return={ResultCode:123 Status:ok unexportedValue:0}, err(<nil>)=
	fmt.Fprintf(os.Stderr, "return=%+v, err(%T)=%+v\n", v, err, common.GojaErrorString(err))

	// js側の設定値が反映されている
	// param={Code:999 Msg:hello Map:map[key3:val3] unexportedValue:88}
	fmt.Fprintf(os.Stderr, "param=%+v\n", param)
}
