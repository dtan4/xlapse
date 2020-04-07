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

func (c *LambdaClient) InvokeDownloaderFuncs(ctx context.Context, es types.Entries, arn string) error {
	for _, e := range es {
		payload, err := json.Marshal(e)
		if err != nil {
			return fmt.Errorf("cannot decode entry %#v to JSON: %w", *e, err)
		}

		_, err = c.api.InvokeWithContext(ctx, &lambda.InvokeInput{
			FunctionName:   aws.String(arn),
			InvocationType: aws.String("Event"),
			Payload:        payload,
		})
		if err != nil {
			return fmt.Errorf("cannot invoke lambda function %q with entry %#v: %w", arn, *e, err)
		}
	}

	return nil
}
