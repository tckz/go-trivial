package main

import (
	"fmt"
	"net/http"
	"os"
)

func do(status http.ConnState) {
	fmt.Fprintf(os.Stderr, "s=%d\n", int(status))
}

func main() {
	do(http.StatusBadRequest)
}
