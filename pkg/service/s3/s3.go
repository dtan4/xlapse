package s3

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"time"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	s3v2 "github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	timeFormat = "2006-01-02-15-04-05"
)

type APIV2 interface {
	GetObject(context.Context, *s3v2.GetObjectInput, ...func(*s3v2.Options)) (*s3v2.GetObjectOutput, error)
	ListObjectsV2(context.Context, *s3v2.ListObjectsV2Input, ...func(*s3v2.Options)) (*s3v2.ListObjectsV2Output, error)
	PutObject(context.Context, *s3v2.PutObjectInput, ...func(*s3v2.Options)) (*s3v2.PutObjectOutput, error)
}

type ListObjectV2Pager interface {
	HasMorePages() bool
	NextPage(context.Context, ...func(*s3v2.Options)) (*s3v2.ListObjectsV2Output, error)
}

type ListObjectsV2PagerFactory func(client s3v2.ListObjectsV2APIClient, params *s3v2.ListObjectsV2Input, optFns ...func(*s3v2.ListObjectsV2PaginatorOptions)) ListObjectV2Pager

// ClientV2 represents the wrapper of S3 API Client using AWS SDK V2
type ClientV2 struct {
	api                       APIV2
	listObjectsV2PagerFactory ListObjectsV2PagerFactory
}

// NewV2 creates new ClientV2
func NewV2(api APIV2) *ClientV2 {
	return &ClientV2{
		api: api,
		listObjectsV2PagerFactory: func(client s3v2.ListObjectsV2APIClient, params *s3v2.ListObjectsV2Input, optFns ...func(*s3v2.ListObjectsV2PaginatorOptions)) ListObjectV2Pager {
			return s3v2.NewListObjectsV2Paginator(client, params)
		},
	}
}

// ListObjectKeys retrieves the list of keys in the given S3 bucket and folder
func (c *ClientV2) ListObjectKeys(ctx context.Context, bucket, folder string) ([]string, error) {
	keys := []string{}

	paginator := c.listObjectsV2PagerFactory(c.api, &s3v2.ListObjectsV2Input{
		Bucket: awsv2.String(bucket),
		Prefix: awsv2.String(folder),
	})

	for paginator.HasMorePages() {
		out, err := paginator.NextPage(ctx)
		if err != nil {
			return []string{}, fmt.Errorf("cannot retrieve object list from S3 (bucket: %q, folder: %q): %w", bucket, folder, err)
		}

		for _, c := range out.Contents {
			keys = append(keys, awsv2.ToString(c.Key))
		}
	}

	return keys, nil
}

// GetObject downloads an object from the specified S3 location
func (c *ClientV2) GetObject(ctx context.Context, bucket, key string) ([]byte, error) {
	out, err := c.api.GetObject(ctx, &s3v2.GetObjectInput{
		Bucket: awsv2.String(bucket),
		Key:    awsv2.String(key),
	})
	if err != nil {
		return []byte{}, fmt.Errorf("cannot download S3 object from bucket: %q, key: %q: %w", bucket, key, err)
	}
	defer out.Body.Close()

	body, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read S3 object from bucket: %q, key: %q: %w", bucket, key, err)
	}

	return body, nil
}

// Upload uploads the given stream to the given S3 location
func (c *ClientV2) Upload(ctx context.Context, bucket, key string, reader io.ReadSeeker) error {
	_, err := c.api.PutObject(ctx, &s3v2.PutObjectInput{
		Bucket: awsv2.String(bucket),
		Key:    awsv2.String(key),
		Body:   reader,
	})
	if err != nil {
		return fmt.Errorf("cannot upload file to S3: %w", err)
	}

	return nil
}

// {prefix}/2006/01/02/
func ComposeFolder(prefix string, year, month, day int) string {
	return composeFolder(prefix, year, month, day)
}

func composeFolder(prefix string, year, month, day int) string {
	return filepath.Join(prefix, fmt.Sprintf("%04d/%02d/%02d", year, month, day)) + "/"
}

// {prefix}/2006/01/02/2006-01-02-15-04-00.png
func ComposeKey(prefix string, now time.Time, ext string) string {
	key := filepath.Join(composeFolder(prefix, now.Year(), int(now.Month()), now.Day()), now.Format(timeFormat))
	if ext != "" {
		key += "." + ext
	}

	return key
}
