package lambda

import (
	"context"
	"fmt"
	"strings"
	"testing"

	lambdav2 "github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/google/go-cmp/cmp"

	"github.com/dtan4/xlapse/types"
	v1 "github.com/dtan4/xlapse/types/v1"
)

var (
	gotPayloads = [][]byte{}
)

type mockAPIV2 struct {
	err error
}

func (m *mockAPIV2) Invoke(ctx context.Context, params *lambdav2.InvokeInput, optFns ...func(*lambdav2.Options)) (*lambdav2.InvokeOutput, error) {
	if m.err != nil {
		return nil, m.err
	}

	gotPayloads = append(gotPayloads, params.Payload)

	return &lambdav2.InvokeOutput{}, nil
}

func TestInvokeDownloaderFuncsV2(t *testing.T) {
	testcases := map[string]struct {
		es        types.Entries
		arn       string
		want      [][]byte
		invokeErr error
		wantErr   error
	}{
		"success": {
			es: types.Entries{
				&v1.Entry{
					Url:       "https://example.co.jp/foo.jpg",
					Bucket:    "bucket",
					KeyPrefix: "prefix",
					Timezone:  "Asia/Tokyo",
				},
				&v1.Entry{
					Url:       "https://example.com.sg/bar.png",
					Bucket:    "bucket-sg",
					KeyPrefix: "prefix-sg",
					Timezone:  "Asia/Singapore",
				},
			},
			arn: "foo",
			want: [][]byte{
				[]byte(`{"url":"https://example.co.jp/foo.jpg","bucket":"bucket","key_prefix":"prefix","timezone":"Asia/Tokyo"}`),
				[]byte(`{"url":"https://example.com.sg/bar.png","bucket":"bucket-sg","key_prefix":"prefix-sg","timezone":"Asia/Singapore"}`),
			},
			invokeErr: nil,
			wantErr:   nil,
		},
		"error": {
			es: types.Entries{
				&v1.Entry{
					Url:       "https://example.co.jp/foo.jpg",
					Bucket:    "bucket",
					KeyPrefix: "prefix",
					Timezone:  "Asia/Tokyo",
				},
				&v1.Entry{
					Url:       "https://example.com.sg/bar.png",
					Bucket:    "bucket-sg",
					KeyPrefix: "prefix-sg",
					Timezone:  "Asia/Singapore",
				},
			},
			arn:       "foo",
			want:      [][]byte{},
			invokeErr: fmt.Errorf("cannot invoke function"),
			wantErr:   fmt.Errorf(`cannot invoke lambda function "foo" with entry v1.Entry`),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			client := &ClientV2{api: &mockAPIV2{
				err: tc.invokeErr,
			}}

			err := client.InvokeDownloaderFuncs(ctx, tc.es, tc.arn)
			if tc.wantErr == nil {
				if err != nil {
					t.Errorf("want no error, got %q", err)
				}

				if diff := cmp.Diff(tc.want, gotPayloads); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			} else {
				if err == nil {
					t.Errorf("want error %q, got nil", tc.wantErr)
				}

				if !strings.HasPrefix(err.Error(), tc.wantErr.Error()) {
					t.Errorf("want error %q, got %q", tc.wantErr, err)
				}
			}

			t.Cleanup(func() {
				gotPayloads = [][]byte{}
			})
		})
	}
}

func TestInvokeGifMakerFuncsV2(t *testing.T) {
	testcases := map[string]struct {
		req       *v1.GifRequest
		arn       string
		want      [][]byte
		invokeErr error
		wantErr   error
	}{
		"success": {
			req: &v1.GifRequest{
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
			req: &v1.GifRequest{
				Bucket:    "bucket",
				KeyPrefix: "prefix",
				Year:      2020,
				Month:     4,
				Day:       11,
			},
			arn:       "foo",
			want:      [][]byte{},
			invokeErr: fmt.Errorf("cannot invoke function"),
			wantErr:   fmt.Errorf(`cannot invoke lambda function "foo" with request &v1.GifRequest{`),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			client := &ClientV2{api: &mockAPIV2{
				err: tc.invokeErr,
			}}

			err := client.InvokeGifMakerFuncs(ctx, tc.req, tc.arn)
			if tc.wantErr == nil {
				if err != nil {
					t.Errorf("want no error, got %q", err)
				}

				if diff := cmp.Diff(tc.want, gotPayloads); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			} else {
				if err == nil {
					t.Errorf("want error %q, got nil", tc.wantErr)
				}

				if !strings.HasPrefix(err.Error(), tc.wantErr.Error()) {
					t.Errorf("want error %q, got %q", tc.wantErr, err)
				}
			}

			t.Cleanup(func() {
				gotPayloads = [][]byte{}
			})
		})
	}
}

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
		es        types.Entries
		arn       string
		want      [][]byte
		invokeErr error
		wantErr   error
	}{
		"success": {
			es: types.Entries{
				&v1.Entry{
					Url:       "https://example.co.jp/foo.jpg",
					Bucket:    "bucket",
					KeyPrefix: "prefix",
					Timezone:  "Asia/Tokyo",
				},
				&v1.Entry{
					Url:       "https://example.com.sg/bar.png",
					Bucket:    "bucket-sg",
					KeyPrefix: "prefix-sg",
					Timezone:  "Asia/Singapore",
				},
			},
			arn: "foo",
			want: [][]byte{
				[]byte(`{"url":"https://example.co.jp/foo.jpg","bucket":"bucket","key_prefix":"prefix","timezone":"Asia/Tokyo"}`),
				[]byte(`{"url":"https://example.com.sg/bar.png","bucket":"bucket-sg","key_prefix":"prefix-sg","timezone":"Asia/Singapore"}`),
			},
			invokeErr: nil,
			wantErr:   nil,
		},
		"error": {
			es: types.Entries{
				&v1.Entry{
					Url:       "https://example.co.jp/foo.jpg",
					Bucket:    "bucket",
					KeyPrefix: "prefix",
					Timezone:  "Asia/Tokyo",
				},
				&v1.Entry{
					Url:       "https://example.com.sg/bar.png",
					Bucket:    "bucket-sg",
					KeyPrefix: "prefix-sg",
					Timezone:  "Asia/Singapore",
				},
			},
			arn:       "foo",
			want:      [][]byte{},
			invokeErr: fmt.Errorf("cannot invoke function"),
			wantErr:   fmt.Errorf(`cannot invoke lambda function "foo" with entry v1.Entry`),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			client := &Client{api: &mockLambdaAPI{
				err: tc.invokeErr,
			}}

			err := client.InvokeDownloaderFuncs(ctx, tc.es, tc.arn)
			if tc.wantErr == nil {
				if err != nil {
					t.Errorf("want no error, got %q", err)
				}

				if diff := cmp.Diff(tc.want, gotPayloads); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			} else {
				if err == nil {
					t.Errorf("want error %q, got nil", tc.wantErr)
				}

				if !strings.HasPrefix(err.Error(), tc.wantErr.Error()) {
					t.Errorf("want error %q, got %q", tc.wantErr, err)
				}
			}

			t.Cleanup(func() {
				gotPayloads = [][]byte{}
			})
		})
	}
}

func TestInvokeGifMakerFuncs(t *testing.T) {
	testcases := map[string]struct {
		req       *v1.GifRequest
		arn       string
		want      [][]byte
		invokeErr error
		wantErr   error
	}{
		"success": {
			req: &v1.GifRequest{
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
			req: &v1.GifRequest{
				Bucket:    "bucket",
				KeyPrefix: "prefix",
				Year:      2020,
				Month:     4,
				Day:       11,
			},
			arn:       "foo",
			want:      [][]byte{},
			invokeErr: fmt.Errorf("cannot invoke function"),
			wantErr:   fmt.Errorf(`cannot invoke lambda function "foo" with request &v1.GifRequest{`),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			client := &Client{api: &mockLambdaAPI{
				err: tc.invokeErr,
			}}

			err := client.InvokeGifMakerFuncs(ctx, tc.req, tc.arn)
			if tc.wantErr == nil {
				if err != nil {
					t.Errorf("want no error, got %q", err)
				}

				if diff := cmp.Diff(tc.want, gotPayloads); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			} else {
				if err == nil {
					t.Errorf("want error %q, got nil", tc.wantErr)
				}

				if !strings.HasPrefix(err.Error(), tc.wantErr.Error()) {
					t.Errorf("want error %q, got %q", tc.wantErr, err)
				}
			}

			t.Cleanup(func() {
				gotPayloads = [][]byte{}
			})
		})
	}
}
