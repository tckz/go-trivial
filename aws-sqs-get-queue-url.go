package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	queueName := flag.String("queue", "", "Name of SQS queue")
	flag.Parse()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)

	// キューが存在しなかったときのエラー
	/*
		err=AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.
		        status code: 400, request id: e2d3a452-7a0a-5fcf-a52c-ff49f3cc565d
	*/
	out, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(*queueName)})
	fmt.Printf("err=%v, out=%s\n", err, out)
}
