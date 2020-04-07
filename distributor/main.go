package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
	if len(args) < 3 {
		return fmt.Errorf("bucket and/or key is missing")
	}
	bucket, key := args[1], args[2]

	log.Printf("bucket: %q", bucket)
	log.Printf("key: %q", key)

	ctx := context.Background()

	sess := session.New()
	api := s3.New(sess)
	s3Client := newS3Client(api)

	body, err := s3Client.GetObject(ctx, bucket, key)
	if err != nil {
		return fmt.Errorf("cannot download file from S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	fmt.Println(string(body))

	return nil
}
