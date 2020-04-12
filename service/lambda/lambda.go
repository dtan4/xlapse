package lambda

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"

	"github.com/dtan4/xlapse/types"
)

const (
	invocationType = "Event" // Event - Invoke the function asynchronously.
)

type Client struct {
	api lambdaiface.LambdaAPI
}

func New(api lambdaiface.LambdaAPI) *Client {
	return &Client{
		api: api,
	}
}

func (c *Client) InvokeDownloaderFuncs(ctx context.Context, es types.Entries, arn string) error {
	for _, e := range es {
		payload, err := json.Marshal(e)
		if err != nil {
			return fmt.Errorf("cannot decode entry %#v to JSON: %w", *e, err)
		}

		_, err = c.api.InvokeWithContext(ctx, &lambda.InvokeInput{
			FunctionName:   aws.String(arn),
			InvocationType: aws.String(invocationType),
			Payload:        payload,
		})
		if err != nil {
			return fmt.Errorf("cannot invoke lambda function %q with entry %#v: %w", arn, *e, err)
		}
	}

	return nil
}

func (c *Client) InvokeGifMakerFuncs(ctx context.Context, req types.GifRequest, arn string) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("cannot decode entry %#v to JSON: %w", req, err)
	}

	_, err = c.api.InvokeWithContext(ctx, &lambda.InvokeInput{
		FunctionName:   aws.String(arn),
		InvocationType: aws.String(invocationType),
		Payload:        payload,
	})
	if err != nil {
		return fmt.Errorf("cannot invoke lambda function %q with request %#v: %w", arn, req, err)
	}

	return nil
}
