package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func main() {
	if err := realMain(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func realMain(args []string) error {
	if len(args) < 4 {
		return fmt.Errorf("bucket, key and/or function ARN is missing")
	}
	bucket, key, farn := args[1], args[2], args[3]

	log.Printf("bucket: %q", bucket)
	log.Printf("key: %q", key)

	ctx := context.Background()

	sess := session.New()
	s3API := s3.New(sess)
	s3Client := newS3Client(s3API)

	body, err := s3Client.GetObject(ctx, bucket, key)
	if err != nil {
		return fmt.Errorf("cannot download file from S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	es, err := decodeYAML(body)
	if err != nil {
		return fmt.Errorf("cannot decode YAML: %w", err)
	}

	for _, e := range es {
		fmt.Printf("URL: %q, Bucket: %q, KeyPrefix: %q, Timezone: %q\n", e.URL, e.Bucket, e.KeyPrefix, e.Timezone)
	}

	lambdaAPI := lambda.New(sess)
	lambdaClient := newLambdaClient(lambdaAPI)

	if err := lambdaClient.InvokeDownloaderFuncs(ctx, es, farn); err != nil {
		return fmt.Errorf("cannot invoke download functions: %w", err)
	}

	return nil
}
