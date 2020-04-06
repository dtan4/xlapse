package main

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type mockS3API struct {
	s3iface.S3API
	err error
}

func (m *mockS3API) PutObjectWithContext(ctx context.Context, input *s3.PutObjectInput, opts ...request.Option) (*s3.PutObjectOutput, error) {
	if m.err != nil {
		return nil, m.err
	}

	return nil, nil
}

func TestUploadToS3(t *testing.T) {
	testcases := map[string]struct {
		bucket    string
		key       string
		body      string
		uploadErr error
		expectErr error
	}{
		"success": {
			bucket:    "test",
			key:       "test.jpg",
			body:      "foo",
			uploadErr: nil,
			expectErr: nil,
		},
		"error": {
			bucket:    "test",
			key:       "test.jpg",
			body:      "foo",
			uploadErr: fmt.Errorf("cannot upload"),
			expectErr: fmt.Errorf("cannot upload file to S3: cannot upload"),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			reader := bytes.NewReader([]byte(tc.body))

			s3Client := &Client{api: &mockS3API{
				err: tc.uploadErr,
			}}

			err := s3Client.UploadToS3(ctx, tc.bucket, tc.key, reader)
			if tc.expectErr == nil {
				if err != nil {
					t.Errorf("want no error, got: %q", err.Error())
				}
			} else {
				if err == nil {
					t.Errorf("want error: %q, got nil", tc.expectErr.Error())
				}

				if err.Error() != tc.expectErr.Error() {
					t.Errorf("want error: %q, got: %q", tc.expectErr.Error(), err.Error())
				}
			}
		})
	}
}
