package lambda

import (
	"context"
	"encoding/json"
	"fmt"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	lambdav2 "github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdav2types "github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"github.com/dtan4/xlapse/types"
	v1 "github.com/dtan4/xlapse/types/v1"
)

type APIV2 interface {
	Invoke(context.Context, *lambdav2.InvokeInput, ...func(*lambdav2.Options)) (*lambdav2.InvokeOutput, error)
}

type ClientV2 struct {
	api APIV2
}

func NewV2(api APIV2) *ClientV2 {
	return &ClientV2{
		api: api,
	}
}

func (c *ClientV2) InvokeDownloaderFuncs(ctx context.Context, es types.Entries, arn string) error {
	for _, e := range es {
		payload, err := json.Marshal(e)
		if err != nil {
			return fmt.Errorf("cannot decode entry %#v to JSON: %w", *e, err)
		}

		_, err = c.api.Invoke(ctx, &lambdav2.InvokeInput{
			FunctionName:   awsv2.String(arn),
			InvocationType: lambdav2types.InvocationTypeEvent,
			Payload:        payload,
		})
		if err != nil {
			return fmt.Errorf("cannot invoke lambda function %q with entry %#v: %w", arn, *e, err)
		}
	}

	return nil
}

func (c *ClientV2) InvokeGifMakerFuncs(ctx context.Context, req *v1.GifRequest, arn string) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("cannot decode entry %#v to JSON: %w", req, err)
	}

	_, err = c.api.Invoke(ctx, &lambdav2.InvokeInput{
		FunctionName:   awsv2.String(arn),
		InvocationType: lambdav2types.InvocationTypeEvent,
		Payload:        payload,
	})
	if err != nil {
		return fmt.Errorf("cannot invoke lambda function %q with request %#v: %w", arn, req, err)
	}

	return nil
}
