package types

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	v1 "github.com/dtan4/xlapse/types/v1"
)

func TestDecodeEntriesYAML(t *testing.T) {
	testcases := map[string]struct {
		body []byte
		want *v1.Entries
	}{
		"success": {
			body: []byte(`entries:
- url: https://example.co.jp/foo.jpg
  bucket: bucket
  key_prefix: prefix
  timezone: Asia/Tokyo
- url: https://example.com.sg/bar.png
  bucket: bucket-sg
  key_prefix: prefix-sg
  timezone: Asia/Singapore
`),
			want: &v1.Entries{
				Entries: []*v1.Entry{
					{
						Url:       "https://example.co.jp/foo.jpg",
						Bucket:    "bucket",
						KeyPrefix: "prefix",
						Timezone:  "Asia/Tokyo",
					},
					{
						Url:       "https://example.com.sg/bar.png",
						Bucket:    "bucket-sg",
						KeyPrefix: "prefix-sg",
						Timezone:  "Asia/Singapore",
					},
				},
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got, err := DecodeEntriesYAML(tc.body)
			if err != nil {
				t.Errorf("want no error, got %s", err)
			}

			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
