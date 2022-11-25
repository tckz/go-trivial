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
		Code  string `dynamo:"code"`
		Count int64  `dynamo:"count"`
	}

	// 重複した属性の更新はエラーになる
	// ValidationException: Invalid UpdateExpression: Two document paths overlap with each other; must remove or rewrite one of these paths; path one: [count], path two: [count]
	ctx := context.Background()
	var m MyTable
	err := db.Table("mytable").
		Update("code", "xxxxx").
		Add("count", 1).
		SetExpr("'count' = 'count' + ?", 1).
		OnlyUpdatedValueWithContext(ctx, &m)
	if err != nil {
		panic(err)
	}

	log.Printf("rec=%+v", m)
}
