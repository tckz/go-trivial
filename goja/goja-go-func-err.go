package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/tckz/go-trivial/goja/gojahelper"
)

type ExampleError struct {
	s string
}

func (e *ExampleError) Error() string {
	return e.s
}

func main() {
	vm := goja.New()

	// https://github.com/dop251/goja/issues/116
	require.NewRegistry().Enable(vm)
	vm.SetParserOptions(parser.WithDisableSourceMaps)
	console.Enable(vm)

	// goの関数を渡してjs側で呼び出しそれがエラーになったときどんな例外になるか
	// goja.GoErrorでwrapされ、元のgoのerrorとして取り出せる

	gojahelper.MustSet(vm, "goFunc", func(v int) error {
		return &ExampleError{s: fmt.Sprintf("goFunc error: %v", v)}
	})
	v, err := vm.RunString(`
goFunc(123);
`)
	fmt.Fprintf(os.Stderr, "value(%T)=%s, err(%T)=%+v\n", v, v, err, gojahelper.GojaErrorString(err))

	var ee *ExampleError
	ok := errors.As(err, &ee)
	// ok=true, ee=goFunc error: 123
	fmt.Fprintf(os.Stderr, "ok=%t, ee=%+v\n", ok, ee)
}
