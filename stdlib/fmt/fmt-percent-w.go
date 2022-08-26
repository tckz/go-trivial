package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// %wをfmt.Errorf以外で使ってるコードがあったので確認
	err := errors.New("error desu")

	// 1:err=%!w(*errors.errorString=&{error desu}) dayo
	fmt.Fprintf(os.Stderr, "1:err=%w dayo\n", err)
	// 2:err=error desu dayo
	fmt.Fprintf(os.Stderr, "2:err=%v dayo\n", err)
	// 3:err=&errors.errorString{s:"error desu"} dayo
	fmt.Fprintf(os.Stderr, "3:err=%#v dayo\n", err)
	//4:err=error desu dayo
	fmt.Fprintf(os.Stderr, "4:err=%+v dayo\n", err)
	// 5:err=error desu dayo
	fmt.Fprintf(os.Stderr, "5:err=%s dayo\n", err)
}
