package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

/*
 * AWS_DEFAULT_REGION=ap-northeast-1 module --bucket ....
 */

func main() {
	bucket := flag.String("bucket", "", "Bucket to check")
	key := flag.String("key", "", "path/to/key")
	flag.Parse()

	if *bucket == "" {
		log.Fatalf("*** --bucket must be specified")
	}

	if *key == "" {
		log.Fatalf("*** --key must be specified")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := s3.New(sess)

	ret, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		if ae, ok := err.(awserr.Error); ok {
			// documentではno such keyが返りそうだが実際は"NotFound"でしかもSDK内に定義がない
			// https://github.com/aws/aws-sdk-go/issues/1208
			// PRも取り込まれず
			// https://github.com/aws/aws-sdk-go/pull/2948
			code := ae.Code()
			if code == "NotFound" {
				fmt.Fprintf(os.Stderr, "Code=%s, key=%s not found, err=%s\n", code, *key, err)
				return
			}
		}
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "%s\n", ret.GoString())
}
