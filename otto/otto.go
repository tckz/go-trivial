package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/parser"
)

func OttoErrorString(err error) string {
	if err == nil {
		return ""
	}

	// ErrorListの場合複数のエラーが(and 1 more errors)のようになってしまうので展開して連結
	/*
		(anonymous): Line 2:14 Unexpected token { (and 1 more errors)
	*/
	var el parser.ErrorList
	if errors.As(err, &el) {
		var m []string
		for _, e := range el {
			m = append(m, e.Error())
		}
		return strings.Join(m, "\n")
	}

	var oe *otto.Error
	if errors.As(err, &oe) {
		// こっちだとスタックトレースも含まれる
		return oe.String()
	}

	return err.Error()
}

func MustSet(vm *otto.Otto, name string, v interface{}) {
	err := vm.Set(name, v)
	if err != nil {
		panic(err)
	}
}

func Run(vm *otto.Otto, s string) otto.Value {
	v, err := vm.Run(s)
	fmt.Fprintf(os.Stderr, "value(%T)=%s, err(%T)=%+v\n", v, v, err, OttoErrorString(err))
	return v
}

func main() {
	vm := otto.New()
	MustSet(vm, "param", map[string]interface{}{
		"key1": 1,
		"key2": "val2",
	})
	// structはObjectになるようだ
	MustSet(vm, "now", time.Date(2021, 10, 17, 16, 42, 0, 0, time.UTC))

	// value(otto.Value)=4,jj, err(<nil>)=
	Run(vm, `
		console.log("Unix()="+ now.Unix()); // Unix()=1634488920 goのメソッドが呼ばれる
	    abc = 2 + 2;
	    console.log("The value of abc is " + abc); // 4
		console.log(param.key1);	// 1
		console.log(param.key2);	// val2
		[abc, "jj"];
	`)

	// vmが同じなので状態を保持している
	// value(otto.Value)=4, err(<nil>)=
	Run(vm, `abc;`)

	/*
		value(otto.Value)=undefined, err(*otto.Error)=ReferenceError: 'invalid_var' is not defined
		    at <anonymous>:1:1
	*/
	Run(vm, `invalid_var.prop;`)

	/*
		value(otto.Value)=undefined, err(parser.ErrorList)=(anonymous): Line 2:14 Unexpected token {
		(anonymous): Line 4:1 Unexpected token }
	*/
	Run(vm, `
!vvfunction(){
some.method();
}()
`)
}
