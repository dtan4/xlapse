package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"path/filepath"
	"sort"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/dtan4/remote-file-to-s3-function/types"
)

const (
	defaultDelay   = 10 // 100ms per frame == 10fps
	defaultGifName = "movie.gif"
)

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

	return do(ctx, req.Bucket, req.KeyPrefix, req.Year, req.Month, req.Day, delay)
}

func do(ctx context.Context, bucket, keyPrefix string, year, month, day, delay int) error {
	sess := session.New()
	api := s3.New(sess)
	xray.AWS(api.Client)
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
