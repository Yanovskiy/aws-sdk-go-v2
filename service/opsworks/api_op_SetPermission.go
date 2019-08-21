// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/SetPermissionRequest
type SetPermissionInput struct {
	_ struct{} `type:"structure"`

	// The user is allowed to use SSH to communicate with the instance.
	AllowSsh *bool `type:"boolean"`

	// The user is allowed to use sudo to elevate privileges.
	AllowSudo *bool `type:"boolean"`

	// The user's IAM ARN. This can also be a federated user's ARN.
	//
	// IamUserArn is a required field
	IamUserArn *string `type:"string" required:"true"`

	// The user's permission level, which must be set to one of the following strings.
	// You cannot set your own permissions level.
	//
	//    * deny
	//
	//    * show
	//
	//    * deploy
	//
	//    * manage
	//
	//    * iam_only
	//
	// For more information about the permissions associated with these levels,
	// see Managing User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
	Level *string `type:"string"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetPermissionInput"}

	if s.IamUserArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("IamUserArn"))
	}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/SetPermissionOutput
type SetPermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetPermissionOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetPermission = "SetPermission"

// SetPermissionRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Specifies a user's permissions. For more information, see Security and Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingsecurity.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using SetPermissionRequest.
//    req := client.SetPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/SetPermission
func (c *Client) SetPermissionRequest(input *SetPermissionInput) SetPermissionRequest {
	op := &aws.Operation{
		Name:       opSetPermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetPermissionInput{}
	}

	req := c.newRequest(op, input, &SetPermissionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetPermissionRequest{Request: req, Input: input, Copy: c.SetPermissionRequest}
}

// SetPermissionRequest is the request type for the
// SetPermission API operation.
type SetPermissionRequest struct {
	*aws.Request
	Input *SetPermissionInput
	Copy  func(*SetPermissionInput) SetPermissionRequest
}

// Send marshals and sends the SetPermission API request.
func (r SetPermissionRequest) Send(ctx context.Context) (*SetPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetPermissionResponse{
		SetPermissionOutput: r.Request.Data.(*SetPermissionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetPermissionResponse is the response type for the
// SetPermission API operation.
type SetPermissionResponse struct {
	*SetPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetPermission request.
func (r *SetPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}