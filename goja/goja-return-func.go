package main

import (
	"fmt"
	"log"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/tckz/go-trivial/goja/gojahelper"
)

func main() {
	vm := goja.New()

	// https://github.com/dop251/goja/issues/116
	require.NewRegistry().Enable(vm)
	vm.SetParserOptions(parser.WithDisableSourceMaps)
	console.Enable(vm)

	// functionオブジェクトを返して、それをgo側で呼び出す

	v := gojahelper.Run(vm, `
(function(x) {
	return x;
});
`)

	var jsFunc func(int) (int, error)
	if err := vm.ExportTo(v, &jsFunc); err != nil {
		panic(fmt.Errorf("ExportTo: %w", err))
	}

	vv, err := jsFunc(456)
	if err != nil {
		panic(err)
	}
	log.Printf("call(456)=%+v", vv)
}
