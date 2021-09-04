package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	baselambda "github.com/aws/aws-lambda-go/lambda"
	configv2 "github.com/aws/aws-sdk-go-v2/config"
	lambdav2 "github.com/aws/aws-sdk-go-v2/service/lambda"
	s3v2 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/getsentry/sentry-go"

	"github.com/dtan4/xlapse/service/lambda"
	"github.com/dtan4/xlapse/service/s3"
	"github.com/dtan4/xlapse/types"
	v1 "github.com/dtan4/xlapse/types/v1"
	"github.com/dtan4/xlapse/version"
)

var (
	sentryEnabled = false
)

func init() {
	if os.Getenv("SENTRY_DSN") != "" {
		sentryEnabled = true
	}
}

func HandleRequest(ctx context.Context) error {
	bucket := os.Getenv("BUCKET")
	key := os.Getenv("KEY")
	farn := os.Getenv("GIF_MAKER_FUNCTION_ARN")

	log.Printf("function version: %q", version.Version)
	log.Printf("function built commit: %q", version.Commit)
	log.Printf("function built date: %q", version.Date)

	log.Printf("bucket: %q", bucket)
	log.Printf("key: %q", key)
	log.Printf("farn: %q", farn)

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
			scope.SetTag("function", "gif-distributor")
		})
	}

	if err := do(ctx, bucket, key, farn); err != nil {
		if sentryEnabled {
			sentry.CaptureException(err)
		}

		return err
	}

	return nil
}

func main() {
	baselambda.Start(HandleRequest)
}

func do(ctx context.Context, bucket, key, farn string) error {
	cfg, err := configv2.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("cannot load default AWS SDK config: %w", err)
	}

	awsv2.AWSV2Instrumentor(&cfg.APIOptions)

	s3Client := s3.NewV2(s3v2.NewFromConfig(cfg))

	body, err := s3Client.GetObject(ctx, bucket, key)
	if err != nil {
		return fmt.Errorf("cannot download file from S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	es, err := types.DecodeEntriesYAML(body)
	if err != nil {
		return fmt.Errorf("cannot decode YAML: %w", err)
	}

	api := lambdav2.NewFromConfig(cfg)
	lambdaClient := lambda.NewV2(api)

	now := time.Now()

	for _, e := range es {
		log.Printf("URL: %q, Bucket: %q, KeyPrefix: %q, Timezone: %q\n", e.GetUrl(), e.GetBucket(), e.GetKeyPrefix(), e.GetTimezone())

		loc, err := time.LoadLocation(e.Timezone)
		if err != nil {
			return fmt.Errorf("cannot load timezone %q: %w", e.Timezone, err)
		}

		yday := now.In(loc).Add(-24 * time.Hour)

		log.Printf("yesterday: %q", yday.String())

		req := &v1.GifRequest{
			Bucket:    e.GetBucket(),
			KeyPrefix: e.GetKeyPrefix(),
			Year:      int32(yday.Year()),
			Month:     int32(yday.Month()),
			Day:       int32(yday.Day()),
		}

		log.Printf("invoking gif-maker function for bucket %q key %q", e.Bucket, e.KeyPrefix)

		if err := lambdaClient.InvokeGifMakerFuncs(ctx, req, farn); err != nil {
			return fmt.Errorf("cannot invoke gif-maker function: %w", err)
		}
	}

	return nil
}
