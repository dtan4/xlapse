package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dtan4/remote-file-to-s3-function/types"
)

func main() {
	if err := realMain(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func realMain(args []string) error {
	if len(args) < 6 {
		return fmt.Errorf("./gif-maker <bucket> <keyPrefix> <year> <month> <day>")
	}

	bucket, keyPrefix := args[1], args[2]

	year, err := strconv.Atoi(args[3])
	if err != nil {
		return fmt.Errorf("year is not integer")
	}

	month, err := strconv.Atoi(args[4])
	if err != nil {
		return fmt.Errorf("month is not integer")
	}

	day, err := strconv.Atoi(args[5])
	if err != nil {
		return fmt.Errorf("day is not integer")
	}

	ctx := context.Background()

	req := types.GifRequest{
		Bucket:    bucket,
		KeyPrefix: keyPrefix,
		Year:      year,
		Month:     month,
		Day:       day,
	}

	return HandleRequest(ctx, req)
}

func HandleRequest(ctx context.Context, req types.GifRequest) error {
	log.Printf("bucket: %q", req.Bucket)
	log.Printf("key prefix: %q", req.KeyPrefix)
	log.Printf("year: %d", req.Year)
	log.Printf("month: %d", req.Month)
	log.Printf("day: %d", req.Day)

	return do(ctx, req.Bucket, req.KeyPrefix, req.Year, req.Month, req.Day)
}

func do(ctx context.Context, bucket, keyPrefix string, year, month, day int) error {
	sess := session.New()
	api := s3.New(sess)
	s3Client := newS3Client(api)

	folder := composeFolder(keyPrefix, year, month, day)

	log.Printf("retrieving object list in bucket: %q folder: %q", bucket, folder)

	keys, err := s3Client.ListObjectKeys(ctx, bucket, folder)
	if err != nil {
		return fmt.Errorf("cannot retrieve object list from S3: %w", err)
	}

	for _, k := range keys {
		log.Println(k)
	}

	return nil
}
