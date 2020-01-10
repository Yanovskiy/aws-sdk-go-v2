// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteImageRecipeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image recipe to delete.
	//
	// ImageRecipeArn is a required field
	ImageRecipeArn *string `location:"querystring" locationName:"imageRecipeArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteImageRecipeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteImageRecipeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteImageRecipeInput"}

	if s.ImageRecipeArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageRecipeArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteImageRecipeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ImageRecipeArn != nil {
		v := *s.ImageRecipeArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "imageRecipeArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteImageRecipeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image recipe that was deleted.
	ImageRecipeArn *string `locationName:"imageRecipeArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteImageRecipeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteImageRecipeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ImageRecipeArn != nil {
		v := *s.ImageRecipeArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageRecipeArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteImageRecipe = "DeleteImageRecipe"

// DeleteImageRecipeRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Deletes an image recipe.
//
//    // Example sending a request using DeleteImageRecipeRequest.
//    req := client.DeleteImageRecipeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/DeleteImageRecipe
func (c *Client) DeleteImageRecipeRequest(input *DeleteImageRecipeInput) DeleteImageRecipeRequest {
	op := &aws.Operation{
		Name:       opDeleteImageRecipe,
		HTTPMethod: "DELETE",
		HTTPPath:   "/DeleteImageRecipe",
	}

	if input == nil {
		input = &DeleteImageRecipeInput{}
	}

	req := c.newRequest(op, input, &DeleteImageRecipeOutput{})
	return DeleteImageRecipeRequest{Request: req, Input: input, Copy: c.DeleteImageRecipeRequest}
}

// DeleteImageRecipeRequest is the request type for the
// DeleteImageRecipe API operation.
type DeleteImageRecipeRequest struct {
	*aws.Request
	Input *DeleteImageRecipeInput
	Copy  func(*DeleteImageRecipeInput) DeleteImageRecipeRequest
}

// Send marshals and sends the DeleteImageRecipe API request.
func (r DeleteImageRecipeRequest) Send(ctx context.Context) (*DeleteImageRecipeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteImageRecipeResponse{
		DeleteImageRecipeOutput: r.Request.Data.(*DeleteImageRecipeOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteImageRecipeResponse is the response type for the
// DeleteImageRecipe API operation.
type DeleteImageRecipeResponse struct {
	*DeleteImageRecipeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteImageRecipe request.
func (r *DeleteImageRecipeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}