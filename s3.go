package main

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// Client represents the wrapper of S3 API Client
type Client struct {
	api s3iface.S3API
}

// New creates new Client
func newS3Client(api s3iface.S3API) *Client {
	return &Client{
		api: api,
	}
}

// UploadToS3 uploads local file to the specified S3 location
func (c *Client) UploadToS3(ctx context.Context, bucket, key string, reader io.ReadSeeker) error {
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
