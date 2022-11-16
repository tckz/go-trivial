package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

/*
 * AWS_DEFAULT_REGION=ap-northeast-1 module --bucket ....
 */

func main() {
	bucket := flag.String("bucket", "", "Bucket to list")
	maxKeys := flag.Int("max-keys", 3, "Number of keys within one page")
	prefix := flag.String("prefix", "", "Key prefix to list like 'path/to/key'")
	delimiter := flag.String("delimiter", "", "")
	flag.Parse()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := s3.New(sess)
	marker := ""
	isTruncated := true
	for isTruncated {
		fmt.Fprintf(os.Stderr, "List with Marker: %s\n", marker)
		resp, err := svc.ListObjects(&s3.ListObjectsInput{
			Bucket:    bucket,
			Prefix:    prefix,
			Delimiter: delimiter,
			MaxKeys:   aws.Int64(int64(*maxKeys)),
			Marker:    aws.String(marker),
		})
		if err != nil {
			panic(err)
		}

		if len(resp.Contents) == 0 {
			break
		}

		// delimiterが指定されたときだけresp.NextMarkerはセットされる
		// nextMarkerの値は最後のkeyと同じらしく。
		// delimiter指定なしの場合も考慮すると自前で最後のキーを指定することに。
		// といよりは、ListObjectsPagesを使うのが正しそう
		nextMarker := *(resp.Contents[len(resp.Contents)-1].Key)

		fmt.Fprintf(os.Stderr, "IsTruncated: %v, NextMarker: %s\n", *resp.IsTruncated, nextMarker)

		for _, e := range resp.Contents {
			fmt.Fprintf(os.Stdout, "%s\n", *e.Key)
		}

		isTruncated = *resp.IsTruncated
		marker = nextMarker
	}

}
