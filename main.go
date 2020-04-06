package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if err := realMain(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func realMain(args []string) error {
	if len(args) < 4 {
		return fmt.Errorf("insufficient arguments")
	}
	url, bucket, key := args[1], args[2], args[3]

	ctx := context.Background()

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	body, _, err := download(ctx, httpClient, url)
	if err != nil {
		return fmt.Errorf("cannot download file from %q: %w", url, err)
	}

	sess := session.New()
	api := s3.New(sess)
	s3Client := newS3Client(api)

	if err := s3Client.UploadToS3(ctx, bucket, key, bytes.NewReader(body)); err != nil {
		return fmt.Errorf("cannot upload downloaded file to S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	return nil
}
