package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/samber/lo"
	"github.com/tckz/go-trivial"
)

func main() {
	ctx := context.Background()

	// アクセスキーが間違っている場合でもLoadDefaultAWSConfig時点ではエラーにならない
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	stsClient := sts.NewFromConfig(cfg)

	// 間違ったアクセスキーが設定されているとここでInvalidClientTokenIdになる
	res, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "*** GetCallerIdentity: err=%v\n", err)
		return
	}

	lo.Must0(trivial.OutYaml(os.Stdout, res))
}
