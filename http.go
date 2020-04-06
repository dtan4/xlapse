package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

var extensions = map[string]string{
	"image/png":  "png",
	"image/jpeg": "jpg",
	"image/gif":  "gif",
}

func download(ctx context.Context, client http.Client, url string) ([]byte, string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return []byte{}, "", fmt.Errorf("cannot make HTTP request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, "", fmt.Errorf("cannot get response from server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, "", fmt.Errorf("invalid response (status code: %d)", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, "", fmt.Errorf("cannot read response: %w", err)
	}

	contentType := resp.Header.Get("Content-Type")

	return body, extensions[contentType], nil
}
