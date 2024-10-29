package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/goccy/go-yaml"
	"github.com/samber/lo"
)

func main() {
	ctx := context.Background()

	// アクセスキーが間違っている場合でもLoadDefaultAWSConfig時点ではエラーにならない
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	stsClient := sts.NewFromConfig(cfg)

	// 間違ったアクセスキーが設定されているとここでInvalidClientTokenIdになる
	res, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "*** GetCallerIdentity: err=%v\n", err)
		return
	}

	lo.Must0(outYaml(res, os.Stdout))
}

func outYaml(src any, w io.Writer) error {
	// yaml.Marshal which compliant with encoding/yaml with types without yaml tag such as GetScheduleOutput outputs keys as lowercase.
	// To avoid it, we marshal it to JSON and decode it again.
	js, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	dec := json.NewDecoder(bytes.NewReader(js))
	dec.UseNumber()
	var v any
	err = dec.Decode(&v)
	if err != nil {
		return fmt.Errorf("json.Decode: %w", err)
	}

	enc := yaml.NewEncoder(w, yaml.UseLiteralStyleIfMultiline(true))
	if err = enc.Encode(v); err != nil {
		return fmt.Errorf("yaml.Encode: %w", err)
	}
	return nil
}
