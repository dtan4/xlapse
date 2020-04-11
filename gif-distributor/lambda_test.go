package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"

	"github.com/dtan4/remote-file-to-s3-function/types"
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

func TestInvokeGifMakerFuncs(t *testing.T) {
	testcases := map[string]struct {
		req       types.GifRequest
		arn       string
		want      [][]byte
		invokeErr error
		wantErr   error
	}{
		"success": {
			req: types.GifRequest{
				Bucket:    "bucket",
				KeyPrefix: "prefix",
				Year:      2020,
				Month:     4,
				Day:       11,
			},
			arn: "foo",
			want: [][]byte{
				[]byte(`{"bucket":"bucket","key_prefix":"prefix","year":2020,"month":4,"day":11}`),
			},
			invokeErr: nil,
			wantErr:   nil,
		},
		"error": {
			req: types.GifRequest{
				Bucket:    "bucket",
				KeyPrefix: "prefix",
				Year:      2020,
				Month:     4,
				Day:       11,
			},
			arn:       "foo",
			want:      [][]byte{},
			invokeErr: fmt.Errorf("cannot invoke function"),
			wantErr:   fmt.Errorf(`cannot invoke lambda function "foo" with request types.GifRequest{Bucket:"bucket", KeyPrefix:"prefix", Year:2020, Month:4, Day:11}: cannot invoke function`),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			lambdaClient := &LambdaClient{api: &mockLambdaAPI{
				err: tc.invokeErr,
			}}

			err := lambdaClient.InvokeGifMakerFuncs(ctx, tc.req, tc.arn)
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
