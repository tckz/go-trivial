package main

import (
	"errors"
	"fmt"
	"log"
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

func Run(vm *otto.Otto, s string) {
	v, err := vm.Run(s)
	fmt.Fprintf(os.Stderr, "value(%T)=%s, err(%T)=%+v\n", v, v, err, OttoErrorString(err))
}

func logDur(s string) func() {
	from := time.Now()
	return func() {
		log.Printf("%s: dur=%s", s, time.Since(from))
	}
}

func main() {
	defer logDur("main()")()

	vm := func() *otto.Otto {
		defer logDur("otto.New()")()
		return otto.New()
	}()
	for _, e := range os.Args[1:] {
		b, err := os.ReadFile(e)
		if err != nil {
			log.Fatalf("ReadFile: %v", err)
		}

		func() {
			s := string(b)
			defer logDur(fmt.Sprintf("Run(%s)", e))()
			Run(vm, s)
		}()
	}
}
