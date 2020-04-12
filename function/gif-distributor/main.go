package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	baselambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaapi "github.com/aws/aws-sdk-go/service/lambda"
	s3api "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-xray-sdk-go/xray"

	"github.com/dtan4/xlapse/service/lambda"
	"github.com/dtan4/xlapse/service/s3"
	"github.com/dtan4/xlapse/types"
)

func HandleRequest(ctx context.Context) error {
	bucket := os.Getenv("BUCKET")
	key := os.Getenv("KEY")
	farn := os.Getenv("GIF_MAKER_FUNCTION_ARN")

	log.Printf("bucket: %q", bucket)
	log.Printf("key: %q", key)
	log.Printf("farn: %q", farn)

	return do(ctx, bucket, key, farn)
}

func main() {
	baselambda.Start(HandleRequest)
}

func do(ctx context.Context, bucket, key, farn string) error {
	sess := session.New()
	s3API := s3api.New(sess)
	xray.AWS(s3API.Client)
	s3Client := s3.New(s3API)

	body, err := s3Client.GetObject(ctx, bucket, key)
	if err != nil {
		return fmt.Errorf("cannot download file from S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	es, err := types.DecodeEntriesYAML(body)
	if err != nil {
		return fmt.Errorf("cannot decode YAML: %w", err)
	}

	lambdaAPI := lambdaapi.New(sess)
	xray.AWS(lambdaAPI.Client)
	lambdaClient := lambda.New(lambdaAPI)

	now := time.Now()

	for _, e := range es {
		log.Printf("URL: %q, Bucket: %q, KeyPrefix: %q, Timezone: %q\n", e.URL, e.Bucket, e.KeyPrefix, e.Timezone)

		loc, err := time.LoadLocation(e.Timezone)
		if err != nil {
			return fmt.Errorf("cannot load timezone %q: %w", e.Timezone, err)
		}

		yday := now.In(loc).Add(-24 * time.Hour)

		log.Printf("yesterday: %q", yday.String())

		req := types.GifRequest{
			Bucket:    e.Bucket,
			KeyPrefix: e.KeyPrefix,
			Year:      yday.Year(),
			Month:     int(yday.Month()),
			Day:       yday.Day(),
		}

		log.Printf("invoking gif-maker function for bucket %q key %q", e.Bucket, e.KeyPrefix)

		if err := lambdaClient.InvokeGifMakerFuncs(ctx, req, farn); err != nil {
			return fmt.Errorf("cannot invoke gif-maker function: %w", err)
		}
	}

	return nil
}
