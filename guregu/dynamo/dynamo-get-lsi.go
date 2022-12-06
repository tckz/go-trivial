package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamo.NewFromIface(dynamodb.New(sess))

	type MyTable struct {
		Key     string `dynamo:"key,hash"`
		Content string `dynamo:"content,range"`

		Seq int64 `dynamo:"seq" localIndex:"seq-index,range"`
	}

	ctx := context.Background()
	var m MyTable
	err := db.Table("stock").
		Get("key", "k1").
		Range("seq", dynamo.Equal, 1).
		Index("seq-index").
		Limit(1).
		OneWithContext(ctx, &m)

	if err != nil {
		panic(err)
	}

	log.Printf("rec=%+v", m)
}
