package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dtan4/remote-file-to-s3-function/types"
)

const (
	defaultDelay   = 10 // 100ms per frame == 10fps
	defaultGifName = "movie.gif"
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

	delay := defaultDelay

	return do(ctx, req.Bucket, req.KeyPrefix, req.Year, req.Month, req.Day, delay)
}

func do(ctx context.Context, bucket, keyPrefix string, year, month, day, delay int) error {
	sess := session.New()
	api := s3.New(sess)
	s3Client := newS3Client(api)

	folder := composeFolder(keyPrefix, year, month, day)

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

	if g.SaveToFile(defaultGifName); err != nil {
		return fmt.Errorf("cannot save GIF image to %q: %w", defaultGifName, err)
	}

	f, err := os.Open(defaultGifName)
	if err != nil {
		return fmt.Errorf("cannot open %q", defaultGifName)
	}
	defer f.Close()

	outKey := filepath.Join(folder, defaultGifName)
	if err := s3Client.Upload(ctx, bucket, outKey, f); err != nil {
		return fmt.Errorf("cannot upload animated GIF to S3 bucket: %q key: %q, %w", bucket, outKey, err)
	}

	return nil
}
