// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a budget action.
func (c *Client) CreateBudgetAction(ctx context.Context, params *CreateBudgetActionInput, optFns ...func(*Options)) (*CreateBudgetActionOutput, error) {
	if params == nil {
		params = &CreateBudgetActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBudgetAction", params, optFns, addOperationCreateBudgetActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBudgetActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBudgetActionInput struct {

	// The account ID of the user. It should be a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// The trigger threshold of the action.
	//
	// This member is required.
	ActionThreshold *types.ActionThreshold

	// The type of action. This defines the type of tasks that can be carried out by
	// this action. This field also determines the format for definition.
	//
	// This member is required.
	ActionType types.ActionType

	// This specifies if the action needs manual or automatic approval.
	//
	// This member is required.
	ApprovalModel types.ApprovalModel

	// A string that represents the budget name. The ":" and "\" characters aren't
	// allowed.
	//
	// This member is required.
	BudgetName *string

	// Specifies all of the type-specific parameters.
	//
	// This member is required.
	Definition *types.Definition

	// The role passed for action execution and reversion. Roles and actions must be in
	// the same account.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The type of a notification. It must be ACTUAL or FORECASTED.
	//
	// This member is required.
	NotificationType types.NotificationType

	// A list of subscribers.
	//
	// This member is required.
	Subscribers []*types.Subscriber
}

type CreateBudgetActionOutput struct {

	// The account ID of the user. It should be a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A system-generated universally unique identifier (UUID) for the action.
	//
	// This member is required.
	ActionId *string

	// A string that represents the budget name. The ":" and "\" characters aren't
	// allowed.
	//
	// This member is required.
	BudgetName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBudgetActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBudgetAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBudgetAction{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateBudgetActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBudgetAction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateBudgetAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "CreateBudgetAction",
	}
}