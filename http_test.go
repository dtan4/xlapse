package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownload(t *testing.T) {
	testcases := map[string]struct {
		url        string
		body       string
		statusCode int
		want       string
		wantErr    error
	}{
		"success": {
			url:        "https://example.com/foo.jpg",
			body:       "foo",
			statusCode: http.StatusOK,
			want:       "foo\n",
			wantErr:    nil,
		},
		"404": {
			url:        "https://example.com/foo.jpg",
			body:       "foo",
			statusCode: http.StatusNotFound,
			want:       "",
			wantErr:    fmt.Errorf("invalid response (status code: 404)"),
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				fmt.Fprintln(w, "foo")
			}))
			defer ts.Close()

			got, err := download(ctx, *ts.Client(), ts.URL)
			if tc.wantErr == nil {
				if err != nil {
					t.Errorf("want no error, got %s", err)
				}

				if string(got) != tc.want {
					t.Errorf("want %q, got %q", tc.want, string(got))
				}
			} else {
				if err == nil {
					t.Errorf("want error %s, got no error", tc.wantErr)
				}

				if err.Error() != tc.wantErr.Error() {
					t.Errorf("want error %s, got %s", tc.wantErr, err)
				}
			}
		})
	}
}
