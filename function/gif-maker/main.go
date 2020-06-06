package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	s3api "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/getsentry/sentry-go"

	"github.com/dtan4/xlapse/service/s3"
	"github.com/dtan4/xlapse/types"
)

const (
	defaultDelay   = 10 // 100ms per frame == 10fps
	defaultGifName = "movie.gif"
)

var (
	sentryEnabled = false
)

func init() {
	if os.Getenv("SENTRY_DSN") != "" {
		sentryEnabled = true
	}
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, req types.GifRequest) error {
	log.Printf("bucket: %q", req.Bucket)
	log.Printf("key prefix: %q", req.KeyPrefix)
	log.Printf("year: %d", req.Year)
	log.Printf("month: %d", req.Month)
	log.Printf("day: %d", req.Day)

	delay := defaultDelay

	if sentryEnabled {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: os.Getenv("SENTRY_DSN"),
			Transport: &sentry.HTTPSyncTransport{
				Timeout: 5 * time.Second,
			},

			// https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime
			Release:    os.Getenv("AWS_LAMBDA_FUNCTION_VERSION"),
			ServerName: os.Getenv("AWS_LAMBDA_FUNCTION_NAME"),
		}); err != nil {
			return fmt.Errorf("cannot initialize Sentry client: %w", err)
		}

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("function", "gif-maker")
			// We can distinguish target images by bucket and key_prefix
			scope.SetTag("bucket", req.Bucket)
			scope.SetTag("key_prefix", req.KeyPrefix)

			scope.SetExtra("gif_request", req)
		})
	}

	if err := do(ctx, req.Bucket, req.KeyPrefix, req.Year, req.Month, req.Day, delay); err != nil {
		if sentryEnabled {
			sentry.CaptureException(err)
		}

		return err
	}

	return nil
}

func do(ctx context.Context, bucket, keyPrefix string, year, month, day, delay int) error {
	sess := session.New()
	api := s3api.New(sess)
	xray.AWS(api.Client)
	s3Client := s3.New(api)

	folder := s3.ComposeFolder(keyPrefix, year, month, day)

	log.Printf("retrieving object list in bucket: %q folder: %q", bucket, folder)

	keys, err := s3Client.ListObjectKeys(ctx, bucket, folder)
	if err != nil {
		return fmt.Errorf("cannot retrieve object list from S3: %w", err)
	}

	if len(keys) == 0 {
		return fmt.Errorf("no files found in bucket: %q folder: %q", bucket, folder)
	}

	sort.Strings(keys)

	g := NewGif()

	for _, k := range keys {
		if filepath.Base(k) == defaultGifName {
			log.Printf("skip %q", k)
			continue
		}

		log.Printf("appending image %q to animated GIF", k)

		body, err := s3Client.GetObject(ctx, bucket, k)
		if err != nil {
			return fmt.Errorf("cannot download object from S3: %q", err)
		}

		if err := g.Append(body, delay); err != nil {
			return fmt.Errorf("cannot append image %q to animated GIF: %w", k, err)
		}
	}

	log.Printf("saving animated GIF to %q", defaultGifName)

	var b bytes.Buffer

	if g.Save(&b); err != nil {
		return fmt.Errorf("cannot save GIF image to %q: %w", defaultGifName, err)
	}

	outKey := filepath.Join(folder, defaultGifName)
	if err := s3Client.Upload(ctx, bucket, outKey, bytes.NewReader(b.Bytes())); err != nil {
		return fmt.Errorf("cannot upload animated GIF to S3 bucket: %q key: %q, %w", bucket, outKey, err)
	}

	return nil
}
