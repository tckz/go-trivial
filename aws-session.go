package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err=%v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "sess=%#v\n", sess)
	}
}
