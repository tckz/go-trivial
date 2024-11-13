package gojahelper

import (
	"errors"
	"fmt"
	"os"

	"github.com/dop251/goja"
)

func RunProgram(vm *goja.Runtime, pg *goja.Program) (goja.Value, error) {
	v, err := vm.RunProgram(pg)
	fmt.Fprintf(os.Stderr, "RunProgram: return(%T)=%s, err(%T)=%+v\n", v, v, err, GojaErrorString(err))
	return v, err
}

func Compile(s string) (*goja.Program, error) {
	pg, err := goja.Compile("", s, true)
	fmt.Fprintf(os.Stderr, "Compile: err(%T)=%+v\n", err, GojaErrorString(err))
	return pg, err
}

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

func Run(vm *goja.Runtime, s string) goja.Value {
	v, err := vm.RunString(s)
	fmt.Fprintf(os.Stderr, "value(%T)=%s, err(%T)=%+v\n", v, v, err, GojaErrorString(err))
	return v
}

func MustSet(vm *goja.Runtime, name string, v any) {
	err := vm.Set(name, v)
	if err != nil {
		panic(err)
	}
}
