// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListTagsRequest
type ListTagsInput struct {
	_ struct{} `type:"structure"`

	// The function's Amazon Resource Name (ARN).
	//
	// Resource is a required field
	Resource *string `location:"uri" locationName:"ARN" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsInput"}

	if s.Resource == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resource"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Resource != nil {
		v := *s.Resource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListTagsResponse
type ListTagsOutput struct {
	_ struct{} `type:"structure"`

	// The function's tags.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s ListTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opListTags = "ListTags"

// ListTagsRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns a function's tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html).
// You can also view tags with GetFunction.
//
//    // Example sending a request using ListTagsRequest.
//    req := client.ListTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListTags
func (c *Client) ListTagsRequest(input *ListTagsInput) ListTagsRequest {
	op := &aws.Operation{
		Name:       opListTags,
		HTTPMethod: "GET",
		HTTPPath:   "/2017-03-31/tags/{ARN}",
	}

	if input == nil {
		input = &ListTagsInput{}
	}

	req := c.newRequest(op, input, &ListTagsOutput{})
	return ListTagsRequest{Request: req, Input: input, Copy: c.ListTagsRequest}
}

// ListTagsRequest is the request type for the
// ListTags API operation.
type ListTagsRequest struct {
	*aws.Request
	Input *ListTagsInput
	Copy  func(*ListTagsInput) ListTagsRequest
}

// Send marshals and sends the ListTags API request.
func (r ListTagsRequest) Send(ctx context.Context) (*ListTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsResponse{
		ListTagsOutput: r.Request.Data.(*ListTagsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsResponse is the response type for the
// ListTags API operation.
type ListTagsResponse struct {
	*ListTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTags request.
func (r *ListTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}