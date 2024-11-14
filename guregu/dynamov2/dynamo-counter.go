package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"
	"github.com/samber/lo"
	"github.com/tckz/go-trivial"
)

var optWorkers = flag.Int("workers", 10, "Number of workers")
var optCount = flag.Int("count", 100, "Number of reqs")

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

	// 同じitemのcount列を同時に複数のworkerから競合排除しつつ加算

	type Cache struct {
		Key string `dynamo:"key,hash"`

		Count   int `dynamo:"count"`
		Version int `dynamo:"version"`
	}

	wg := &sync.WaitGroup{}
	tbl := db.Table("cache")
	key := "key2"
	ch := make(chan int, *optWorkers)
	capacity := float64(0)
	n := 1
	for range *optWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := range ch {
				for attempt := 0; ; attempt++ {
					var ret Cache
					{
						var cc dynamo.ConsumedCapacity
						err := tbl.Get("key", key).ConsumedCapacity(&cc).One(ctx, &ret)
						capacity += cc.Total
						if err != nil {
							if !errors.Is(err, dynamo.ErrNotFound) {
								log.Fatalf("Get: err=%T: %v", err, err)
							}
						}
					}

					q := tbl.Update("key", key)
					if ret.Key == "" {
						q = q.If("attribute_not_exists($)", "key").
							SetExpr("$ = ?", "count", n).
							Set("version", 1)
					} else {
						q = q.If("attribute_exists($)", "key").
							If("version = ?", ret.Version).
							SetExpr("$ = $ + ?", "count", "count", n).
							SetExpr("version = version + ?", 1)
					}

					{
						var cc dynamo.ConsumedCapacity
						err := q.ConsumedCapacity(&cc).Run(ctx)
						capacity += cc.Total
						if err != nil {
							var cce *types.ConditionalCheckFailedException
							if errors.As(err, &cce) {
								log.Printf("[%d-%d]%v", i, attempt, cce)
								continue
							}
							log.Fatalf("Update: err=%T: %v", err, err)
						}
					}
					break
				}
			}
		}()
	}

	for i := range *optCount {
		if i%10 == 0 {
			log.Printf("%d reqs", i)
		}
		ch <- i
	}

	close(ch)
	log.Printf("Waiting for workers to finish")
	wg.Wait()

	log.Printf("Total consumed capacity: %.1f", capacity)

	var ret Cache
	if err := tbl.Get("key", key).One(ctx, &ret); err != nil {
		log.Fatalf("Get: err=%T: %v", err, err)
	}
	lo.Must0(trivial.OutYaml(os.Stdout, ret))
}
