package main

import (
	"context"
	"errors"
	"log"
	"os"
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

	db := dynamo.New(cfg)

	type Cache struct {
		Key string `dynamo:"key,hash"`

		ExpiresAt time.Time `dynamo:"expires_at,unixtime"`
	}

	var ret Cache
	var cc dynamo.ConsumedCapacity
	err = db.Table("cache").Get("key", "not-exist").
		ConsumedCapacity(&cc).
		One(ctx, &ret)
	if err != nil {
		lo.Must0(trivial.OutYaml(os.Stdout, cc))
		// Get: err=*errors.errorString: dynamo.ErrNotFound?=true, dynamo: no item found
		log.Fatalf("Get: err=%T: dynamo.ErrNotFound?=%t, %v", err, errors.Is(err, dynamo.ErrNotFound), err)
	}
	lo.Must0(trivial.OutYaml(os.Stdout, ret, cc))
}
