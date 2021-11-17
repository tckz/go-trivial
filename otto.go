package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/robertkrimen/otto"
)

func OttoErrorString(err error) string {
	if err == nil {
		return ""
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

	// value=4,jj, err=
	Run(vm, `
		console.log("Unix()="+ now.Unix()); // Unix()=1634488920 goのメソッドが呼ばれる
	    abc = 2 + 2;
	    console.log("The value of abc is " + abc); // 4
		console.log(param.key1);	// 1
		console.log(param.key2);	// val2
		[abc, "jj"];
	`)

	// vmが同じなので状態を保持している
	// value=4, err=
	Run(vm, `abc;`)

	/*
		value=undefined, err=ReferenceError: 'invalid_var' is not defined
		    at <anonymous>:1:1
	*/
	Run(vm, `invalid_var.prop;`)
}
