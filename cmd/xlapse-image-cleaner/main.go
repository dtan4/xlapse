package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.uber.org/zap"
)

const (
	usage              = "xlapse-image-cleaner S3_BUCKET"
	maxObjectsToDelete = 1000
)

var (
	filesToSkip = []string{
		"images.yaml",
		"movie.gif",
	}
)

type s3Client struct {
	client *s3.Client
	logger *zap.Logger
}

func (c *s3Client) deleteObjects(ctx context.Context, bucket string, objects []types.ObjectIdentifier, dryRun bool) ([]types.DeletedObject, error) {
	c.logger.Info(fmt.Sprintf("deleting %d objects", len(objects)))

	deleted := []types.DeletedObject{}

	if dryRun {
		for _, obj := range objects {
			c.logger.Info("(dry-run) deleted", zap.String("key", aws.ToString(obj.Key)))
			deleted = append(deleted, types.DeletedObject{
				Key: obj.Key,
			})
		}
	} else {
		resp, err := c.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String(bucket),
			Delete: &types.Delete{
				Objects: objects,
			},
		})
		if err != nil {
			return []types.DeletedObject{}, fmt.Errorf("cannot delete objects: %w", err)
		}

		deleted = resp.Deleted

		for _, obj := range resp.Deleted {
			c.logger.Info("deleted", zap.String("key", aws.ToString(obj.Key)))
		}

		if len(resp.Errors) > 0 {
			for _, failed := range resp.Errors {
				c.logger.Error(
					"cannot delete object",
					zap.String("key", aws.ToString(failed.Key)),
					zap.String("code", aws.ToString(failed.Code)),
					zap.String("message", aws.ToString(failed.Message)),
				)
			}

			return deleted, fmt.Errorf("cannot delete some objects")
		}
	}

	return deleted, nil
}

func (c *s3Client) AsyncDeleteObjects(ctx context.Context, bucket string, ch <-chan string, dryRun bool) {
	totalDeleted := 0
	objs := make([]types.ObjectIdentifier, 0, maxObjectsToDelete)

	for key := range ch {
		objs = append(objs, types.ObjectIdentifier{
			Key: aws.String(key),
		})

		if len(objs) == maxObjectsToDelete {
			deleted, err := c.deleteObjects(ctx, bucket, objs, dryRun)
			if err != nil {
				c.logger.Error("cannot delete objects", zap.Error(err))
			}

			totalDeleted += len(deleted)
			objs = make([]types.ObjectIdentifier, 0, maxObjectsToDelete)
		}
	}

	// delete remaining objects
	if len(objs) > 0 {
		deleted, err := c.deleteObjects(ctx, bucket, objs, dryRun)
		if err != nil {
			c.logger.Error("cannot delete objects", zap.Error(err))
		}

		totalDeleted += len(deleted)
		c.logger.Info(fmt.Sprintf("deleted %d objects", totalDeleted))
	}
}

func isSkippable(key string) bool {
	for _, skip := range filesToSkip {
		if path.Base(key) == skip {
			return true
		}
	}

	return false
}

func realMain(args []string, logger *zap.Logger, dryRun bool) error {
	if len(args) < 2 {
		return fmt.Errorf("S3_BUCKET is missing. Usage: %s", usage)
	}
	bucket := args[1]

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("cannot load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}

	p := s3.NewListObjectsV2Paginator(client, params)

	ch := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		c := &s3Client{
			client: client,
			logger: logger,
		}
		c.AsyncDeleteObjects(ctx, bucket, ch, dryRun)
	}()

	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			close(ch)
			wg.Wait()

			return fmt.Errorf("cannot get page: %w", err)
		}

		for _, obj := range page.Contents {
			key := aws.ToString(obj.Key)

			if isSkippable(key) {
				logger.Info("skipping", zap.String("key", key))
				continue
			}

			ch <- key
		}
	}

	close(ch)
	wg.Wait()

	return nil
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot initialize zap loger: %v", err)
	}

	defer logger.Sync()

	dryRun := false
	if os.Getenv("DRY_RUN") == "true" {
		dryRun = true
		logger.Info("dry-run mode enabled")
	}

	if err := realMain(os.Args, logger, dryRun); err != nil {
		logger.Error("error", zap.Error(err))
		os.Exit(1)
	}
}
