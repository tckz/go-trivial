package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

func main() {
	sess := session.Must(session.NewSession())

	cli := dynamodb.New(sess, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	db := dynamo.NewFromIface(cli)

	type Cache struct {
		Key       string    `dynamo:"key"`
		ExpiredAt time.Time `dynamo:"expired_at"`
	}

	// guregu/dynamoにおいて
	// 0件でbatch writeするとどうなるか。
	// -> dynamo: no input items
	// でエラーが返る

	ctx := context.Background()
	recs := make([]interface{}, 0, 25)
	count, err := db.Table("cache").Batch().Write().Put(recs...).RunWithContext(ctx)
	if err != nil {
		panic(err)
	}

	log.Printf("count=%d", count)
}
