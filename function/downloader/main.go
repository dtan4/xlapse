package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	s3api "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/getsentry/sentry-go"

	"github.com/dtan4/xlapse/service/s3"
	v1 "github.com/dtan4/xlapse/types/v1"
	"github.com/dtan4/xlapse/version"
)

const (
	defaultTimezone = "UTC"
)

var (
	sentryEnabled = false
)

func init() {
	if os.Getenv("SENTRY_DSN") != "" {
		sentryEnabled = true
	}
}

func HandleRequest(ctx context.Context, entry *v1.Entry) error {
	log.Printf("entry: %#v", entry)

	timezone := entry.Timezone
	if timezone == "" {
		timezone = defaultTimezone
	}

	log.Printf("function version: %q", version.Version)
	log.Printf("function built commit: %q", version.Commit)
	log.Printf("function built date: %q", version.Date)

	log.Printf("url: %q", entry.Url)
	log.Printf("bucket: %q", entry.Bucket)
	log.Printf("key prefix: %q", entry.KeyPrefix)
	log.Printf("timezone: %q", timezone)

	if sentryEnabled {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: os.Getenv("SENTRY_DSN"),
			Transport: &sentry.HTTPSyncTransport{
				Timeout: 5 * time.Second,
			},

			Release: version.Version,
			// https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime
			ServerName: os.Getenv("AWS_LAMBDA_FUNCTION_NAME"),
		}); err != nil {
			return fmt.Errorf("cannot initialize Sentry client: %w", err)
		}

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("function", "downloader")
			// We can distinguish target images by bucket and key_prefix
			scope.SetTag("bucket", entry.Bucket)
			scope.SetTag("key_prefix", entry.KeyPrefix)

			scope.SetExtra("entry", entry)
		})
	}

	if err := do(ctx, entry.Url, entry.Bucket, entry.KeyPrefix, timezone); err != nil {
		if sentryEnabled {
			sentry.CaptureException(err)
		}

		return err
	}

	return nil
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
