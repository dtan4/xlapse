package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/dtan4/remote-file-to-s3-function/types"
)

const (
	invocationType = "Event" // Event - Invoke the function asynchronously.
)

type LambdaClient struct {
	api lambdaiface.LambdaAPI
}

func newLambdaClient(api lambdaiface.LambdaAPI) *LambdaClient {
	return &LambdaClient{
		api: api,
	}
}

func (c *LambdaClient) InvokeGifMakerFuncs(ctx context.Context, req types.GifRequest, arn string) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("cannot decode entry %#v to JSON: %w", req, err)
	}

	_, err = c.api.InvokeWithContext(ctx, &lambda.InvokeInput{
		FunctionName:   aws.String(arn),
		InvocationType: aws.String("Event"),
		Payload:        payload,
	})
	if err != nil {
		return fmt.Errorf("cannot invoke lambda function %q with request %#v: %w", arn, req, err)
	}

	return nil
}
