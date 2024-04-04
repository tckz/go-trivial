package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func GojaErrorString(err error) string {
	if err == nil {
		return ""
	}

	var ge *goja.Exception
	if errors.As(err, &ge) {
		// こっちだとスタックトレースも含まれる
		return ge.String()
	}

	return err.Error()
}

func main() {
	vm := goja.New()

	// https://github.com/dop251/goja/issues/116
	require.NewRegistry().Enable(vm)
	vm.SetParserOptions(parser.WithDisableSourceMaps)
	console.Enable(vm)

	// jsが投げた例外は goja.Exception としてgo側のerrに返される

	v, err := vm.RunString(`
function jsFunc(v) {
	throw new Error("jsFunc error");
}
jsFunc();
`)

	/*
		return=<nil>, err(*goja.Exception)=Error: jsFunc error
			at jsFunc (<eval>:3:8(3))
			at <eval>:5:7(3)
	*/
	fmt.Fprintf(os.Stderr, "return=%+v, err(%T)=%+v\n", v, err, GojaErrorString(err))
}
