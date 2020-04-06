package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if err := realMain(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func realMain(args []string) error {
	if len(args) < 4 {
		return fmt.Errorf("insufficient arguments")
	}
	url, bucket, key := args[1], args[2], args[3]

	ctx := context.Background()

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("cannot make HTTP request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("cannot get response from server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response (status code: %d)", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read response: %w", err)
	}

	sess := session.New()
	api := s3.New(sess)
	s3Client := newS3Client(api)

	if err := s3Client.UploadToS3(ctx, bucket, key, bytes.NewReader(body)); err != nil {
		return fmt.Errorf("cannot upload downloaded file to S3 (bucket: %q, key: %q): %w", bucket, key, err)
	}

	return nil
}
