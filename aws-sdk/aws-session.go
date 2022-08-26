package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {

	// アクセスキーが間違っている場合でもNewSession時点ではエラーにならない

	sess, err := session.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err=%v\n", err)
		return
	}

	stsClient := sts.New(sess)

	// 間違ったアクセスキーが設定されているとここでInvalidClientTokenIdになる

	out, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "*** GetCallerIdentity: err=%v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "%#v\n", out)
}
