package s3

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

const (
	timeFormat = "2006-01-02-15-04-05"
)

// Client represents the wrapper of S3 API Client
type Client struct {
	api s3iface.S3API
}

// New creates new Client
func New(api s3iface.S3API) *Client {
	return &Client{
		api: api,
	}
}

// ListObjectKeys retrieves the list of keys in the given S3 bucket and folder
func (c *Client) ListObjectKeys(ctx context.Context, bucket, folder string) ([]string, error) {
	keys := []string{}

	err := c.api.ListObjectsV2PagesWithContext(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(folder),
	}, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, c := range page.Contents {
			keys = append(keys, aws.StringValue(c.Key))
		}

		return true
	})
	if err != nil {
		return []string{}, fmt.Errorf("cannot retrieve object list from S3 (bucket: %q, folder: %q): %w", bucket, folder, err)
	}

	return keys, nil
}

// UploadToS3 uploads local file to the specified S3 location
func (c *Client) GetObject(ctx context.Context, bucket, key string) ([]byte, error) {
	out, err := c.api.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
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
func (c *Client) Upload(ctx context.Context, bucket, key string, reader io.ReadSeeker) error {
	_, err := c.api.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   reader,
	})
	if err != nil {
		return fmt.Errorf("cannot upload file to S3: %w", err)
	}

	return nil
}

// {prefix}/2006/01/02/
func ComposeFolder(prefix string, year, month, day int) string {
	return filepath.Join(prefix, fmt.Sprintf("%04d/%02d/%02d", year, month, day)) + "/"
}

// {prefix}/2006/01/02/2006-01-02-15-04-00.png
func ComposeKey(prefix string, now time.Time, ext string) string {
	key := filepath.Join(prefix, fmt.Sprintf("%04d/%02d/%02d", now.Year(), now.Month(), now.Day()), now.Format(timeFormat))
	if ext != "" {
		key += "." + ext
	}

	return key
}
