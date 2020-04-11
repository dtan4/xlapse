package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	s3api "github.com/aws/aws-sdk-go/service/s3"

	"github.com/dtan4/remote-file-to-s3-function/types"
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
	if err := HandleRequest(context.Background()); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func do(ctx context.Context, bucket, key, farn string) error {
	sess := session.New()
	s3API := s3api.New(sess)
	s3Client := newS3Client(s3API)

	body, err := s3Client.GetObject(ctx, bucket, key)
	if err != nil {
		return fmt.Errorf("cannot download file from S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	es, err := types.DecodeEntriesYAML(body)
	if err != nil {
		return fmt.Errorf("cannot decode YAML: %w", err)
	}

	now := time.Now()

	for _, e := range es {
		log.Printf("URL: %q, Bucket: %q, KeyPrefix: %q, Timezone: %q\n", e.URL, e.Bucket, e.KeyPrefix, e.Timezone)

		loc, err := time.LoadLocation(e.Timezone)
		if err != nil {
			return fmt.Errorf("cannot load timezone %q: %w", e.Timezone, err)
		}

		yesterday := now.In(loc).Add(-24 * time.Hour)

		log.Printf("yesterday: %s", yesterday.Format("2006-01-02"))
	}

	return nil
}
