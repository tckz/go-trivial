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

func main() {
	vm := otto.New()
	MustSet(vm, "param", map[string]interface{}{
		"key1": 1,
		"key2": "val2",
	})
	// structはObjectになるようだ
	MustSet(vm, "now", time.Date(2021, 10, 17, 16, 42, 0, 0, time.UTC))

	v, err := vm.Run(`
		console.log("Unix()="+ now.Unix()); // Unix()=1634488920 goのメソッドが呼ばれる
	    abc = 2 + 2;
	    console.log("The value of abc is " + abc); // 4
		console.log(param.key1);	// 1
		console.log(param.key2);	// val2
		[abc, "jj"];
	`)

	// value=4,jj, err=
	fmt.Fprintf(os.Stderr, "value=%+v, err=%s\n", v, OttoErrorString(err))

	// vmが同じなので状態を保持している
	v, err = vm.Run(`abc;`)
	// value=4, err=
	fmt.Fprintf(os.Stderr, "value=%+v, err=%+v\n", v, OttoErrorString(err))

	v, err = vm.Run(`invalid_var.prop;`)
	/*
		value=undefined, err=ReferenceError: 'invalid_var' is not defined
		    at <anonymous>:1:1
	*/
	fmt.Fprintf(os.Stderr, "value=%+v, err=%+v\n", v, OttoErrorString(err))
}
