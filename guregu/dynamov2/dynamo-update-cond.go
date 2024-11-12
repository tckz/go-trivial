package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/guregu/dynamo/v2"
	"github.com/samber/lo"
	"github.com/tckz/go-trivial"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("ap-northeast-1"),
		config.WithRetryer(func() aws.Retryer {
			return retry.NewStandard(dynamo.RetryTxConflicts)
		}))
	if err != nil {
		log.Fatalf("config.LoadDefaultConfig: %v", err)
	}

	n := lo.Must(strconv.Atoi(os.Args[1]))

	db := dynamo.New(cfg)

	type Cache struct {
		Key string `dynamo:"key,hash"`

		Count     int       `dynamo:"count"`
		ExpiresAt time.Time `dynamo:"expires_at,unixtime"`
	}

	/*
		- 属性countに対して
		  - 存在しなければ初期値設定
		  - 存在すれば加算
		  - ただし結果が0以上になること
		- 更新後の項目を取得
	*/
	var ret Cache
	tbl := db.Table("cache")
	var cc dynamo.ConsumedCapacity
	err = tbl.Update("key", "key1").
		SetExpr("$ = if_not_exists($, ?) + ?", "count", "count", n, n).
		If("attribute_not_exists($) and ? >= ? or attribute_exists($) and $ >= ?",
			"count", n, 0,
			"count", "count", -n, 0).
		ConsumedCapacity(&cc).
		Value(ctx, &ret)

	if err != nil {
		lo.Must0(trivial.OutYaml(os.Stdout, cc))
		log.Fatalf("Update: err=%T: %v", err, err)
	}
	lo.Must0(trivial.OutYaml(os.Stdout, ret, cc))
}
