package main

import (
	"reflect"
	"testing"
)

func TestDecodeYAML(t *testing.T) {
	testcases := map[string]struct {
		body []byte
		want Entries
	}{
		"success": {
			body: []byte(`- url: https://example.co.jp/foo.jpg
  bucket: bucket
  key_prefix: prefix
  timezone: Asia/Tokyo
- url: https://example.com.so/bar.png
  bucket: bucket-sg
  key_prefix: prefix-sg
  timezone: Asia/Singapore
`),
			want: Entries{
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
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got, err := decodeYAML(tc.body)
			if err != nil {
				t.Errorf("want no error, got %s", err)
			}

			if len(got) != len(tc.want) {
				t.Errorf("want %d entries, got %d", len(tc.want), len(got))
			}

			if reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %#v, got %#v", tc.want, got)
			}
		})
	}
}
