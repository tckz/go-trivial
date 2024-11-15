package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
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

	tbl := db.Table("cache")
	var cc dynamo.ConsumedCapacity
	err = db.WriteTx().
		// クライアントトークンを乱数生成して設定してくれる
		Idempotent(true).
		ConsumedCapacity(&cc).
		Put(tbl.Put(&Cache{Key: "key3", ExpiresAt: time.Now().Add(time.Minute * 1)}).
			If("attribute_not_exists($)", "key")).
		Put(tbl.Put(&Cache{Key: "key4", ExpiresAt: time.Now().Add(time.Minute * 1)}).
			If("attribute_not_exists($)", "key")).
		Put(tbl.Put(&Cache{Key: "key5", ExpiresAt: time.Now().Add(time.Minute * 1)}).
			If("attribute_not_exists($)", "key")).
		Run(ctx)
	if err != nil {
		var oe *smithy.OperationError
		var re *http.ResponseError
		var te *types.TransactionCanceledException
		if errors.As(err, &oe) {
			// smithy.OperationError: Service=DynamoDB, Operation=TransactWriteItems
			log.Printf("smithy.OperationError: Service=%s, Operation=%s", oe.Service(), oe.Operation())
		}
		if errors.As(err, &re) {
			// ServiceRequestID=9J6QIC349A6432O9VSRSTS82MBVV4KQNSO5AEMVJF66Q9ASUAAJG
			log.Printf("http.ResponseError: ServiceRequestID=%s", re.ServiceRequestID())
		}
		if errors.As(err, &te) {
			var reasons []string
			for _, r := range te.CancellationReasons {
				reasons = append(reasons, fmt.Sprintf("Code=%s, Message=%s, Item=%#v", *r.Code, *r.Message, r.Item))
			}
			// types.TransactionCanceledException: ErrorCode=TransactionCanceledException, CancellationReasons=[
			//   Code=ConditionalCheckFailed, Message=The conditional request failed, Item=map[string]types.AttributeValue(nil)
			//   Code=ConditionalCheckFailed, Message=The conditional request failed, Item=map[string]types.AttributeValue(nil)
			// ]
			log.Printf("types.TransactionCanceledException: ErrorCode=%s, CancellationReasons=%v", te.ErrorCode(), reasons)
		}
		log.Fatalf("WriteTx: err=%T: %v", err, err)
	}

	lo.Must0(trivial.OutYaml(os.Stdout, cc))
}
