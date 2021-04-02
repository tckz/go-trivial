package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// %wをfmt.Errorf以外で使ってるコードがあったので確認
	err := errors.New("error desu")

	// err=%!w(*errors.errorString=&{error desu}) dayo
	fmt.Fprintf(os.Stderr, "err=%w dayo\n", err)
}
