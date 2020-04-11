package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type mockS3API struct {
	s3iface.S3API
	body []byte
	err  error
}

func (m *mockS3API) GetObjectWithContext(ctx context.Context, input *s3.GetObjectInput, opts ...request.Option) (*s3.GetObjectOutput, error) {
	if m.err != nil {
		return nil, m.err
	}

	return &s3.GetObjectOutput{
		Body: ioutil.NopCloser(bytes.NewReader(m.body)),
	}, nil
}

func TestGetObject(t *testing.T) {
	testcases := map[string]struct {
		bucket    string
		key       string
		body      string
		want      []byte
		getErr    error
		expectErr error
	}{
		"success": {
			bucket:    "test",
			key:       "foo",
			body:      "bar",
			want:      []byte("bar"),
			getErr:    nil,
			expectErr: nil,
		},
		"error": {
			bucket:    "test",
			key:       "foo",
			body:      "",
			want:      []byte{},
			getErr:    fmt.Errorf("cannot upload"),
			expectErr: fmt.Errorf(`cannot download S3 object from bucket: "test", key: "foo": cannot upload`),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			s3Client := &S3Client{api: &mockS3API{
				body: []byte(tc.body),
				err:  tc.getErr,
			}}

			got, err := s3Client.GetObject(ctx, tc.bucket, tc.key)
			if tc.expectErr == nil {
				if err != nil {
					t.Errorf("want no error, got: %q", err.Error())
				}

				if bytes.Compare(got, tc.want) != 0 {
					t.Errorf("want %q, got %q", string(tc.want), string(got))
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
