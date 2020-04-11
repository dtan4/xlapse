package s3

import (
	"context"
	"fmt"
	"io/ioutil"

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
