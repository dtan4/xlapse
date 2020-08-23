package types

import (
	"testing"

	v1 "github.com/dtan4/xlapse/types/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestDecodeEntriesYAML(t *testing.T) {
	testcases := map[string]struct {
		body []byte
		want Entries
	}{
		"success": {
			body: []byte(`- url: https://example.co.jp/foo.jpg
  bucket: bucket
  key_prefix: prefix
  timezone: Asia/Tokyo
- url: https://example.com.sg/bar.png
  bucket: bucket-sg
  key_prefix: prefix-sg
  timezone: Asia/Singapore
`),
			want: Entries{
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
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got, err := DecodeEntriesYAML(tc.body)
			if err != nil {
				t.Errorf("want no error, got %s", err)
			}

			if len(got) != len(tc.want) {
				t.Errorf("want %d entries, got %d", len(tc.want), len(got))
			}

			for i := range got {
				opt := cmpopts.IgnoreUnexported(*tc.want[i])

				if diff := cmp.Diff(*tc.want[i], *got[i], opt); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			}
		})
	}
}
