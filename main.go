package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	timeFormat = "2006-01-02-15-04-00"
)

func HandleRequest(ctx context.Context) error {
	url := os.Getenv("URL")
	bucket := os.Getenv("BUCKET")
	keyPrefix := os.Getenv("KEY_PREFIX")

	log.Printf("url: %q", url)
	log.Printf("bucket: %q", bucket)
	log.Printf("key prefix: %q", keyPrefix)

	return do(ctx, url, bucket, keyPrefix)
}

func main() {
	lambda.Start(HandleRequest)
}

func do(ctx context.Context, url, bucket, keyPrefix string) error {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	log.Printf("downloading %s", url)

	body, ext, err := download(ctx, httpClient, url)
	if err != nil {
		return fmt.Errorf("cannot download file from %q: %w", url, err)
	}

	sess := session.New()
	api := s3.New(sess)
	s3Client := newS3Client(api)

	key := filepath.Join(keyPrefix, time.Now().Format(timeFormat))
	if ext != "" {
		key += "." + ext
	}

	log.Printf("uploading to bucket: %s key: %s", bucket, key)

	if err := s3Client.UploadToS3(ctx, bucket, key, bytes.NewReader(body)); err != nil {
		return fmt.Errorf("cannot upload downloaded file to S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	return nil
}
