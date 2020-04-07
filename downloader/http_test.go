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
		url         string
		body        string
		contentType string
		statusCode  int
		want        string
		wantExt     string
		wantErr     error
	}{
		"success with png": {
			url:         "https://example.com/foo.png",
			body:        "foo",
			contentType: "image/png",
			statusCode:  http.StatusOK,
			want:        "foo\n",
			wantExt:     "png",
			wantErr:     nil,
		},
		"success with jpg": {
			url:         "https://example.com/foo.jpg",
			body:        "foo",
			contentType: "image/jpeg",
			statusCode:  http.StatusOK,
			want:        "foo\n",
			wantExt:     "jpg",
			wantErr:     nil,
		},
		"success with gif": {
			url:         "https://example.com/foo.gif",
			body:        "foo",
			contentType: "image/gif",
			statusCode:  http.StatusOK,
			want:        "foo\n",
			wantExt:     "gif",
			wantErr:     nil,
		},
		"success with text": {
			url:         "https://example.com/foo.gif",
			body:        "foo",
			contentType: "text/html",
			statusCode:  http.StatusOK,
			want:        "foo\n",
			wantExt:     "",
			wantErr:     nil,
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
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(tc.statusCode)
				fmt.Fprintln(w, "foo")
			}))
			defer ts.Close()

			body, ext, err := download(ctx, ts.Client(), ts.URL)
			if tc.wantErr == nil {
				if err != nil {
					t.Errorf("want no error, got %s", err)
				}

				if string(body) != tc.want {
					t.Errorf("want body %q, got %q", tc.want, string(body))
				}

				if ext != tc.wantExt {
					t.Errorf("want ext %q, got %q", tc.wantExt, ext)
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
