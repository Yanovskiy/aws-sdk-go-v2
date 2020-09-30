// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthorizationType string

// Enum values for AuthorizationType
const (
	AuthorizationTypeNone    AuthorizationType = "NONE"
	AuthorizationTypeAws_iam AuthorizationType = "AWS_IAM"
	AuthorizationTypeCustom  AuthorizationType = "CUSTOM"
	AuthorizationTypeJwt     AuthorizationType = "JWT"
)

type AuthorizerType string

// Enum values for AuthorizerType
const (
	AuthorizerTypeRequest AuthorizerType = "REQUEST"
	AuthorizerTypeJwt     AuthorizerType = "JWT"
)

type ConnectionType string

// Enum values for ConnectionType
const (
	ConnectionTypeInternet ConnectionType = "INTERNET"
	ConnectionTypeVpc_link ConnectionType = "VPC_LINK"
)

type ContentHandlingStrategy string

// Enum values for ContentHandlingStrategy
const (
	ContentHandlingStrategyConvert_to_binary ContentHandlingStrategy = "CONVERT_TO_BINARY"
	ContentHandlingStrategyConvert_to_text   ContentHandlingStrategy = "CONVERT_TO_TEXT"
)

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusPending  DeploymentStatus = "PENDING"
	DeploymentStatusFailed   DeploymentStatus = "FAILED"
	DeploymentStatusDeployed DeploymentStatus = "DEPLOYED"
)

type DomainNameStatus string

// Enum values for DomainNameStatus
const (
	DomainNameStatusAvailable DomainNameStatus = "AVAILABLE"
	DomainNameStatusUpdating  DomainNameStatus = "UPDATING"
)

type EndpointType string

// Enum values for EndpointType
const (
	EndpointTypeRegional EndpointType = "REGIONAL"
	EndpointTypeEdge     EndpointType = "EDGE"
)

type IntegrationType string

// Enum values for IntegrationType
const (
	IntegrationTypeAws        IntegrationType = "AWS"
	IntegrationTypeHttp       IntegrationType = "HTTP"
	IntegrationTypeMock       IntegrationType = "MOCK"
	IntegrationTypeHttp_proxy IntegrationType = "HTTP_PROXY"
	IntegrationTypeAws_proxy  IntegrationType = "AWS_PROXY"
)

type LoggingLevel string

// Enum values for LoggingLevel
const (
	LoggingLevelError LoggingLevel = "ERROR"
	LoggingLevelInfo  LoggingLevel = "INFO"
	LoggingLevelOff   LoggingLevel = "OFF"
)

type PassthroughBehavior string

// Enum values for PassthroughBehavior
const (
	PassthroughBehaviorWhen_no_match     PassthroughBehavior = "WHEN_NO_MATCH"
	PassthroughBehaviorNever             PassthroughBehavior = "NEVER"
	PassthroughBehaviorWhen_no_templates PassthroughBehavior = "WHEN_NO_TEMPLATES"
)

type ProtocolType string

// Enum values for ProtocolType
const (
	ProtocolTypeWebsocket ProtocolType = "WEBSOCKET"
	ProtocolTypeHttp      ProtocolType = "HTTP"
)

type SecurityPolicy string

// Enum values for SecurityPolicy
const (
	SecurityPolicyTls_1_0 SecurityPolicy = "TLS_1_0"
	SecurityPolicyTls_1_2 SecurityPolicy = "TLS_1_2"
)

type VpcLinkStatus string

// Enum values for VpcLinkStatus
const (
	VpcLinkStatusPending   VpcLinkStatus = "PENDING"
	VpcLinkStatusAvailable VpcLinkStatus = "AVAILABLE"
	VpcLinkStatusDeleting  VpcLinkStatus = "DELETING"
	VpcLinkStatusFailed    VpcLinkStatus = "FAILED"
	VpcLinkStatusInactive  VpcLinkStatus = "INACTIVE"
)

type VpcLinkVersion string

// Enum values for VpcLinkVersion
const (
	VpcLinkVersionV2 VpcLinkVersion = "V2"
)