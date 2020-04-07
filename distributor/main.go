package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	s3api "github.com/aws/aws-sdk-go/service/s3"
	lambdaapi "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
)

func HandleRequest(ctx context.Context) error {
	bucket := os.Getenv("BUCKET")
	key := os.Getenv("KEY")
	farn := os.Getenv("DOWNLOADER_FUNCTION_ARN")

	log.Printf("bucket: %q", bucket)
	log.Printf("key: %q", key)
	log.Printf("farn: %q", farn)

	return do(ctx, bucket, key, farn)
}

func main() {
	lambda.Start(HandleRequest)
}

func do(ctx context.Context, bucket, key, farn string) error {
	sess := session.New()
	s3API := s3api.New(sess)
	xray.AWS(s3API.Client)
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

	lambdaAPI := lambdaapi.New(sess)
	xray.AWS(lambdaAPI.Client)
	lambdaClient := newLambdaClient(lambdaAPI)

	if err := lambdaClient.InvokeDownloaderFuncs(ctx, es, farn); err != nil {
		return fmt.Errorf("cannot invoke download functions: %w", err)
	}

	return nil
}
