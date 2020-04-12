package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	s3api "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-xray-sdk-go/xray"

	"github.com/dtan4/remote-file-to-s3-function/service/s3"
	"github.com/dtan4/remote-file-to-s3-function/types"
)

const (
	defaultTimezone = "UTC"
)

func HandleRequest(ctx context.Context, entry types.Entry) error {
	log.Printf("entry: %#v", entry)

	timezone := entry.Timezone
	if timezone == "" {
		timezone = defaultTimezone
	}

	log.Printf("url: %q", entry.URL)
	log.Printf("bucket: %q", entry.Bucket)
	log.Printf("key prefix: %q", entry.KeyPrefix)
	log.Printf("timezone: %q", timezone)

	return do(ctx, entry.URL, entry.Bucket, entry.KeyPrefix, timezone)
}

func main() {
	lambda.Start(HandleRequest)
}

func do(ctx context.Context, url, bucket, keyPrefix, timezone string) error {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Errorf("cannot retrieve timezone %q; %w", timezone, err)
	}

	httpClient := xray.Client(&http.Client{
		Timeout: 5 * time.Second,
	})

	log.Printf("downloading %s", url)

	body, ext, err := download(ctx, httpClient, url)
	if err != nil {
		return fmt.Errorf("cannot download file from %q: %w", url, err)
	}

	sess := session.New()
	api := s3api.New(sess)
	xray.AWS(api.Client)
	s3Client := s3.New(api)

	now := time.Now().In(loc)
	key := s3.ComposeKey(keyPrefix, now, ext)

	log.Printf("uploading to bucket: %s key: %s", bucket, key)

	if err := s3Client.Upload(ctx, bucket, key, bytes.NewReader(body)); err != nil {
		return fmt.Errorf("cannot upload downloaded file to S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	return nil
}
