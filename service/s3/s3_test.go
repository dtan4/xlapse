package s3

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/google/go-cmp/cmp"
)

type mockS3API struct {
	s3iface.S3API
	body []byte
	keys []string
	err  error
}

func (m *mockS3API) ListObjectsV2PagesWithContext(ctx aws.Context, input *s3.ListObjectsV2Input, fn func(*s3.ListObjectsV2Output, bool) bool, opts ...request.Option) error {
	if m.err != nil {
		return m.err
	}

	objects := []*s3.Object{}

	for _, k := range m.keys {
		objects = append(objects, &s3.Object{
			Key: aws.String(k),
		})
	}

	_ = fn(&s3.ListObjectsV2Output{
		Contents: objects,
	}, true)

	return nil
}

func (m *mockS3API) GetObjectWithContext(ctx context.Context, input *s3.GetObjectInput, opts ...request.Option) (*s3.GetObjectOutput, error) {
	if m.err != nil {
		return nil, m.err
	}

	return &s3.GetObjectOutput{
		Body: ioutil.NopCloser(bytes.NewReader(m.body)),
	}, nil
}

func (m *mockS3API) PutObjectWithContext(ctx context.Context, input *s3.PutObjectInput, opts ...request.Option) (*s3.PutObjectOutput, error) {
	if m.err != nil {
		return nil, m.err
	}

	return nil, nil
}

func TestListObjectKeys(t *testing.T) {
	testcases := map[string]struct {
		bucket    string
		folder    string
		keys      []string
		want      []string
		listErr   error
		expectErr error
	}{
		"success": {
			bucket: "test",
			folder: "2020/04/01/",
			keys: []string{
				"2020/04/01/foo",
				"2020/04/01/bar",
				"2020/04/01/baz",
			},
			want: []string{
				"2020/04/01/foo",
				"2020/04/01/bar",
				"2020/04/01/baz",
			},
			listErr:   nil,
			expectErr: nil,
		},
		"error": {
			bucket: "test",
			folder: "2020/04/01/",
			keys: []string{
				"2020/04/01/foo",
				"2020/04/01/bar",
				"2020/04/01/baz",
			},
			want: []string{
				"2020/04/01/foo",
				"2020/04/01/bar",
				"2020/04/01/baz",
			},
			listErr:   fmt.Errorf("failed"),
			expectErr: fmt.Errorf(`cannot retrieve object list from S3 (bucket: "test", folder: "2020/04/01/"): failed`),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			client := &Client{api: &mockS3API{
				keys: tc.keys,
				err:  tc.listErr,
			}}

			got, err := client.ListObjectKeys(ctx, tc.bucket, tc.folder)
			if tc.expectErr == nil {
				if err != nil {
					t.Errorf("want no error, got %q", err.Error())
				}

				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			} else {
				if err == nil {
					t.Errorf("want error %q, got nil", tc.expectErr.Error())
				}

				if err.Error() != tc.expectErr.Error() {
					t.Errorf("want error %q, got %q", tc.expectErr.Error(), err.Error())
				}
			}
		})
	}
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

			client := &Client{api: &mockS3API{
				body: []byte(tc.body),
				err:  tc.getErr,
			}}

			got, err := client.GetObject(ctx, tc.bucket, tc.key)
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

func TestUpload(t *testing.T) {
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

			client := &Client{api: &mockS3API{
				err: tc.uploadErr,
			}}

			err := client.Upload(ctx, tc.bucket, tc.key, reader)
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

func TestComposeFolder(t *testing.T) {
	testcases := map[string]struct {
		prefix string
		year   int
		month  int
		day    int
		want   string
	}{
		"2020/04/01/": {
			prefix: "",
			year:   2020,
			month:  4,
			day:    1,
			want:   "2020/04/01/",
		},
		"2020/04/11/": {
			prefix: "",
			year:   2020,
			month:  4,
			day:    11,
			want:   "2020/04/11/",
		},
		"2020/10/11/": {
			prefix: "",
			year:   2020,
			month:  10,
			day:    11,
			want:   "2020/10/11/",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got := ComposeFolder(tc.prefix, tc.year, tc.month, tc.day)
			if got != tc.want {
				t.Errorf("want: %q, got: %q", tc.want, got)
			}
		})
	}
}

func TestComposeKey(t *testing.T) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("cannot retrieve Asia/Tokyo timezone: %s", err)
	}

	testcases := map[string]struct {
		prefix string
		now    time.Time
		ext    string
		want   string
	}{
		"no prefix but ext": {
			prefix: "",
			now:    time.Date(2020, 4, 6, 11, 22, 33, 0, time.UTC),
			ext:    "png",
			want:   "2020/04/06/2020-04-06-11-22-33.png",
		},
		"no prefix and no ext": {
			prefix: "",
			now:    time.Date(2020, 4, 6, 11, 22, 33, 0, time.UTC),
			ext:    "",
			want:   "2020/04/06/2020-04-06-11-22-33",
		},
		"prefix and ext": {
			prefix: "awesome",
			now:    time.Date(2020, 4, 6, 11, 22, 33, 0, time.UTC),
			ext:    "png",
			want:   "awesome/2020/04/06/2020-04-06-11-22-33.png",
		},
		"prefix but no ext": {
			prefix: "awesome",
			now:    time.Date(2020, 4, 6, 11, 22, 33, 0, time.UTC),
			ext:    "",
			want:   "awesome/2020/04/06/2020-04-06-11-22-33",
		},
		"JST": {
			prefix: "awesome",
			now:    time.Date(2020, 4, 6, 11, 22, 33, 0, jst),
			ext:    "png",
			want:   "awesome/2020/04/06/2020-04-06-11-22-33.png",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got := ComposeKey(tc.prefix, tc.now, tc.ext)
			if got != tc.want {
				t.Errorf("want: %q, got: %q", tc.want, got)
			}
		})
	}
}
