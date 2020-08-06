package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {

	// アクセスキーが間違っている場合でもLoadDefaultAWSConfig時点ではエラーにならない
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}

	stsClient := sts.New(cfg)

	// 間違ったアクセスキーが設定されているとここでInvalidClientTokenIdになる
	req := stsClient.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
	res, err := req.Send(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "*** GetCallerIdentity: err=%v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "%s\n", res)
}
