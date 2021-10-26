package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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
	from := time.Now()
	v, err := vm.Run(s)
	fmt.Fprintf(os.Stderr, "value(%T)=%s, err(%T)=%+v, dur=%s\n", v, v, err, OttoErrorString(err), time.Since(from))
}

func logDur(from time.Time, s string) func() {
	return func() {
		log.Printf("%s: dur=%s", s, time.Since(from))
	}
}

func main() {
	defer logDur(time.Now(), "main()")()

	vm := func() *otto.Otto {
		defer logDur(time.Now(), "otto.New()")()
		return otto.New()
	}()
	for _, e := range os.Args[1:] {
		b, err := ioutil.ReadFile(e)
		if err != nil {
			log.Fatalf("ReadFile: %v", err)
		}

		Run(vm, string(b))
	}
}
