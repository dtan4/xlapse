package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

var (
	gotPayloads = [][]byte{}
)

type mockLambdaAPI struct {
	lambdaiface.LambdaAPI
	err error
}

func (m *mockLambdaAPI) InvokeWithContext(ctx context.Context, input *lambda.InvokeInput, opts ...request.Option) (*lambda.InvokeOutput, error) {
	if m.err != nil {
		return nil, m.err
	}

	gotPayloads = append(gotPayloads, input.Payload)

	return &lambda.InvokeOutput{}, nil
}

func TestInvokeDownloaderFuncs(t *testing.T) {
	testcases := map[string]struct {
		es        Entries
		arn       string
		want      [][]byte
		invokeErr error
		wantErr   error
	}{
		"success": {
			es: Entries{
				&Entry{
					URL:       "https://example.co.jp/foo.jpg",
					Bucket:    "bucket",
					KeyPrefix: "prefix",
					Timezone:  "Asia/Tokyo",
				},
				&Entry{
					URL:       "https://example.com.sg/bar.png",
					Bucket:    "bucket-sg",
					KeyPrefix: "prefix-sg",
					Timezone:  "Asia/Singapore",
				},
			},
			arn:       "foo",
			want:      [][]byte{
				[]byte(`{"url":"https://example.co.jp/foo.jpg","bucket":"bucket","key_prefix":"prefix","timezone":"Asia/Tokyo"}`),
				[]byte(`{"url":"https://example.com.sg/bar.png","bucket":"bucket-sg","key_prefix":"prefix-sg","timezone":"Asia/Singapore"}`),
			},
			invokeErr: nil,
			wantErr:   nil,
		},
		"error": {
			es: Entries{
				&Entry{
					URL:       "https://example.co.jp/foo.jpg",
					Bucket:    "bucket",
					KeyPrefix: "prefix",
					Timezone:  "Asia/Tokyo",
				},
				&Entry{
					URL:       "https://example.com.sg/bar.png",
					Bucket:    "bucket-sg",
					KeyPrefix: "prefix-sg",
					Timezone:  "Asia/Singapore",
				},
			},
			arn:       "foo",
			want:      [][]byte{},
			invokeErr: fmt.Errorf("cannot invoke function"),
			wantErr:   fmt.Errorf(`cannot invoke lambda function "foo" with entry main.Entry{URL:"https://example.co.jp/foo.jpg", Bucket:"bucket", KeyPrefix:"prefix", Timezone:"Asia/Tokyo"}: cannot invoke function`),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			lambdaClient := &LambdaClient{api: &mockLambdaAPI{
				err: tc.invokeErr,
			}}

			err := lambdaClient.InvokeDownloaderFuncs(ctx, tc.es, tc.arn)
			if tc.wantErr == nil {
				if err != nil {
					t.Errorf("want no error, got %q", err)
				}

				if !reflect.DeepEqual(gotPayloads, tc.want) {
					t.Errorf("want %q, got %q", tc.want, gotPayloads)
				}
			} else {
				if err == nil {
					t.Errorf("want error %q, got nil", tc.wantErr)
				}

				if err.Error() != tc.wantErr.Error() {
					t.Errorf("want error %q, got %q", tc.wantErr, err)
				}
			}

			t.Cleanup(func() {
				gotPayloads = [][]byte{}
			})
		})
	}
}
