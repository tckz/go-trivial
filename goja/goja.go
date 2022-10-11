package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func MustSet(vm *goja.Runtime, name string, v interface{}) {
	err := vm.Set(name, v)
	if err != nil {
		panic(err)
	}
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

func main() {
	vm := goja.New()

	// https://github.com/dop251/goja/issues/116
	require.NewRegistry().Enable(vm)
	console.Enable(vm)

	MustSet(vm, "param", map[string]interface{}{
		"key1": 1,
		"key2": "val2",
	})
	// structはObjectになるようだ
	MustSet(vm, "now", time.Date(2021, 10, 17, 16, 42, 0, 0, time.UTC))

	// value(*goja.Object)=4,jj, err(<nil>)=
	Run(vm, `
		console.log("Unix()="+ now.Unix()); // Unix()=1634488920 goのメソッドが呼ばれる
	    abc = 2 + 2;
	    console.log("The value of abc is " + abc); // 4
		console.log(param.key1);	// 1
		console.log(param.key2);	// val2
		[abc, "jj"];
	`)

	// vmが同じなので状態を保持している
	// value(goja.valueInt)=4, err(<nil>)=
	Run(vm, `abc;`)

	/*
		value(<nil>)=%!s(<nil>), err(*goja.Exception)=ReferenceError: invalid_var is not defined
		        at <eval>:1:1(0)
	*/
	Run(vm, `invalid_var.prop;`)

	// ottoと違って(and 1 more errors)の中身を知る方法がなさそう
	// parserのところではotto同様にErrorListで保持しているが途中でError()が呼び出されstringだけ返ってくる
	/*
		value(<nil>)=%!s(<nil>), err(*goja.Exception)=SyntaxError: SyntaxError: (anonymous): Line 2:14 Unexpected token { (and 1 more errors)
	*/
	Run(vm, `
!vvfunction(){
some.method();
}()
`)
}
