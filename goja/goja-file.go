package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func logDur(s string) func() {
	from := time.Now()
	return func() {
		log.Printf("%s: dur=%s", s, time.Since(from))
	}
}

func main() {
	defer logDur("main()")()

	vm := func() *goja.Runtime {
		defer logDur("goja.New()")()
		vm := goja.New()
		vm.SetParserOptions(parser.WithDisableSourceMaps)
		// https://github.com/dop251/goja/issues/116
		require.NewRegistry().Enable(vm)
		console.Enable(vm)

		return vm
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
