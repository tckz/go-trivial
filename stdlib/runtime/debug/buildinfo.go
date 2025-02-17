package main

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/tckz/go-trivial"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Println("No build info available")
		return
	}

	trivial.OutYaml(os.Stdout, bi)
}
