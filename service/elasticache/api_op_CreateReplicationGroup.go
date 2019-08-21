// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateReplicationGroup operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateReplicationGroupMessage
type CreateReplicationGroupInput struct {
	_ struct{} `type:"structure"`

	// A flag that enables encryption at rest when set to true.
	//
	// You cannot modify the value of AtRestEncryptionEnabled after the replication
	// group is created. To enable encryption at rest on a replication group you
	// must set AtRestEncryptionEnabled to true when you create the replication
	// group.
	//
	// Required: Only available when creating a replication group in an Amazon VPC
	// using redis version 3.2.6, 4.x or later.
	//
	// Default: false
	AtRestEncryptionEnabled *bool `type:"boolean"`

	// Reserved parameter. The password used to access a password protected server.
	//
	// AuthToken can be specified only on replication groups where TransitEncryptionEnabled
	// is true.
	//
	// For HIPAA compliance, you must specify TransitEncryptionEnabled as true,
	// an AuthToken, and a CacheSubnetGroup.
	//
	// Password constraints:
	//
	//    * Must be only printable ASCII characters.
	//
	//    * Must be at least 16 characters and no more than 128 characters in length.
	//
	//    * Cannot contain any of the following characters: '/', '"', or '@'.
	//
	// For more information, see AUTH password (http://redis.io/commands/AUTH) at
	// http://redis.io/commands/AUTH.
	AuthToken *string `type:"string"`

	// This parameter is currently disabled.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// Specifies whether a read-only replica is automatically promoted to read/write
	// primary if the existing primary fails.
	//
	// If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ
	// is disabled for this replication group.
	//
	// AutomaticFailoverEnabled must be enabled for Redis (cluster mode enabled)
	// replication groups.
	//
	// Default: false
	//
	// Amazon ElastiCache for Redis does not support Multi-AZ with automatic failover
	// on:
	//
	//    * Redis versions earlier than 2.8.6.
	//
	//    * Redis (cluster mode disabled): T1 node types.
	//
	//    * Redis (cluster mode enabled): T1 node types.
	AutomaticFailoverEnabled *bool `type:"boolean"`

	// The compute and memory capacity of the nodes in the node group (shard).
	//
	// The following node types are supported by ElastiCache. Generally speaking,
	// the current generation types provide more memory and computational power
	// at lower cost when compared to their equivalent previous generation counterparts.
	//
	//    * General purpose: Current generation: M5 node types: cache.m5.large,
	//    cache.m5.xlarge, cache.m5.2xlarge, cache.m5.4xlarge, cache.m5.12xlarge,
	//    cache.m5.24xlarge M4 node types: cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge,
	//    cache.m4.4xlarge, cache.m4.10xlarge T2 node types: cache.t2.micro, cache.t2.small,
	//    cache.t2.medium Previous generation: (not recommended) T1 node types:
	//    cache.t1.micro M1 node types: cache.m1.small, cache.m1.medium, cache.m1.large,
	//    cache.m1.xlarge M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge,
	//    cache.m3.2xlarge
	//
	//    * Compute optimized: Previous generation: (not recommended) C1 node types:
	//    cache.c1.xlarge
	//
	//    * Memory optimized: Current generation: R5 node types: cache.r5.large,
	//    cache.r5.xlarge, cache.r5.2xlarge, cache.r5.4xlarge, cache.r5.12xlarge,
	//    cache.r5.24xlarge R4 node types: cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge,
	//    cache.r4.4xlarge, cache.r4.8xlarge, cache.r4.16xlarge Previous generation:
	//    (not recommended) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge
	//    R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge, cache.r3.4xlarge,
	//    cache.r3.8xlarge
	//
	// Additional node type info
	//
	//    * All current generation instance types are created in Amazon VPC by default.
	//
	//    * Redis append-only files (AOF) are not supported for T1 or T2 instances.
	//
	//    * Redis Multi-AZ with automatic failover is not supported on T1 instances.
	//
	//    * Redis configuration variables appendonly and appendfsync are not supported
	//    on Redis version 2.8.22 and later.
	CacheNodeType *string `type:"string"`

	// The name of the parameter group to associate with this replication group.
	// If this argument is omitted, the default cache parameter group for the specified
	// engine is used.
	//
	// If you are restoring to an engine version that is different than the original,
	// you must specify the default version of that version. For example, CacheParameterGroupName=default.redis4.0.
	//
	// If you are running Redis version 3.2.4 or later, only one node group (shard),
	// and want to use a default parameter group, we recommend that you specify
	// the parameter group by name.
	//
	//    * To create a Redis (cluster mode disabled) replication group, use CacheParameterGroupName=default.redis3.2.
	//
	//    * To create a Redis (cluster mode enabled) replication group, use CacheParameterGroupName=default.redis3.2.cluster.on.
	CacheParameterGroupName *string `type:"string"`

	// A list of cache security group names to associate with this replication group.
	CacheSecurityGroupNames []string `locationNameList:"CacheSecurityGroupName" type:"list"`

	// The name of the cache subnet group to be used for the replication group.
	//
	// If you're going to launch your cluster in an Amazon VPC, you need to create
	// a subnet group before you start creating a cluster. For more information,
	// see Subnets and Subnet Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html).
	CacheSubnetGroupName *string `type:"string"`

	// The name of the cache engine to be used for the clusters in this replication
	// group.
	Engine *string `type:"string"`

	// The version number of the cache engine to be used for the clusters in this
	// replication group. To view the supported cache engine versions, use the DescribeCacheEngineVersions
	// operation.
	//
	// Important: You can upgrade to a newer engine version (see Selecting a Cache
	// Engine and Version (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement))
	// in the ElastiCache User Guide, but you cannot downgrade to an earlier engine
	// version. If you want to use an earlier engine version, you must delete the
	// existing cluster or replication group and create it anew with the earlier
	// engine version.
	EngineVersion *string `type:"string"`

	// A list of node group (shard) configuration options. Each node group (shard)
	// configuration has the following members: PrimaryAvailabilityZone, ReplicaAvailabilityZones,
	// ReplicaCount, and Slots.
	//
	// If you're creating a Redis (cluster mode disabled) or a Redis (cluster mode
	// enabled) replication group, you can use this parameter to individually configure
	// each node group (shard), or you can omit this parameter. However, when seeding
	// a Redis (cluster mode enabled) cluster from a S3 rdb file, you must configure
	// each node group (shard) using this parameter because you must specify the
	// slots for each node group.
	NodeGroupConfiguration []NodeGroupConfiguration `locationNameList:"NodeGroupConfiguration" type:"list"`

	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service
	// (SNS) topic to which notifications are sent.
	//
	// The Amazon SNS topic owner must be the same as the cluster owner.
	NotificationTopicArn *string `type:"string"`

	// The number of clusters this replication group initially has.
	//
	// This parameter is not used if there is more than one node group (shard).
	// You should use ReplicasPerNodeGroup instead.
	//
	// If AutomaticFailoverEnabled is true, the value of this parameter must be
	// at least 2. If AutomaticFailoverEnabled is false you can omit this parameter
	// (it will default to 1), or you can explicitly set it to a value between 2
	// and 6.
	//
	// The maximum permitted value for NumCacheClusters is 6 (1 primary plus 5 replicas).
	NumCacheClusters *int64 `type:"integer"`

	// An optional parameter that specifies the number of node groups (shards) for
	// this Redis (cluster mode enabled) replication group. For Redis (cluster mode
	// disabled) either omit this parameter or set it to 1.
	//
	// Default: 1
	NumNodeGroups *int64 `type:"integer"`

	// The port number on which each member of the replication group accepts connections.
	Port *int64 `type:"integer"`

	// A list of EC2 Availability Zones in which the replication group's clusters
	// are created. The order of the Availability Zones in the list is the order
	// in which clusters are allocated. The primary cluster is created in the first
	// AZ in the list.
	//
	// This parameter is not used if there is more than one node group (shard).
	// You should use NodeGroupConfiguration instead.
	//
	// If you are creating your replication group in an Amazon VPC (recommended),
	// you can only locate clusters in Availability Zones associated with the subnets
	// in the selected subnet group.
	//
	// The number of Availability Zones listed must equal the value of NumCacheClusters.
	//
	// Default: system chosen Availability Zones.
	PreferredCacheClusterAZs []string `locationNameList:"AvailabilityZone" type:"list"`

	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Valid
	// values for ddd are:
	//
	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	//
	// Valid values for ddd are:
	//
	//    * sun
	//
	//    * mon
	//
	//    * tue
	//
	//    * wed
	//
	//    * thu
	//
	//    * fri
	//
	//    * sat
	//
	// Example: sun:23:00-mon:01:30
	PreferredMaintenanceWindow *string `type:"string"`

	// The identifier of the cluster that serves as the primary for this replication
	// group. This cluster must already exist and have a status of available.
	//
	// This parameter is not required if NumCacheClusters, NumNodeGroups, or ReplicasPerNodeGroup
	// is specified.
	PrimaryClusterId *string `type:"string"`

	// An optional parameter that specifies the number of replica nodes in each
	// node group (shard). Valid values are 0 to 5.
	ReplicasPerNodeGroup *int64 `type:"integer"`

	// A user-created description for the replication group.
	//
	// ReplicationGroupDescription is a required field
	ReplicationGroupDescription *string `type:"string" required:"true"`

	// The replication group identifier. This parameter is stored as a lowercase
	// string.
	//
	// Constraints:
	//
	//    * A name must contain from 1 to 20 alphanumeric characters or hyphens.
	//
	//    * The first character must be a letter.
	//
	//    * A name cannot end with a hyphen or contain two consecutive hyphens.
	//
	// ReplicationGroupId is a required field
	ReplicationGroupId *string `type:"string" required:"true"`

	// One or more Amazon VPC security groups associated with this replication group.
	//
	// Use this parameter only when you are creating a replication group in an Amazon
	// Virtual Private Cloud (Amazon VPC).
	SecurityGroupIds []string `locationNameList:"SecurityGroupId" type:"list"`

	// A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB
	// snapshot files stored in Amazon S3. The snapshot files are used to populate
	// the new replication group. The Amazon S3 object name in the ARN cannot contain
	// any commas. The new replication group will have the number of node groups
	// (console: shards) specified by the parameter NumNodeGroups or the number
	// of node groups configured by NodeGroupConfiguration regardless of the number
	// of ARNs specified here.
	//
	// Example of an Amazon S3 ARN: arn:aws:s3:::my_bucket/snapshot1.rdb
	SnapshotArns []string `locationNameList:"SnapshotArn" type:"list"`

	// The name of a snapshot from which to restore data into the new replication
	// group. The snapshot status changes to restoring while the new replication
	// group is being created.
	SnapshotName *string `type:"string"`

	// The number of days for which ElastiCache retains automatic snapshots before
	// deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot
	// that was taken today is retained for 5 days before being deleted.
	//
	// Default: 0 (i.e., automatic backups are disabled for this cluster).
	SnapshotRetentionLimit *int64 `type:"integer"`

	// The daily time range (in UTC) during which ElastiCache begins taking a daily
	// snapshot of your node group (shard).
	//
	// Example: 05:00-09:00
	//
	// If you do not specify this parameter, ElastiCache automatically chooses an
	// appropriate time range.
	SnapshotWindow *string `type:"string"`

	// A list of cost allocation tags to be added to this resource. Tags are comma-separated
	// key,value pairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple
	// tags as shown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster
	// is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled
	// to true when you create a cluster.
	//
	// This parameter is valid only if the Engine parameter is redis, the EngineVersion
	// parameter is 3.2.6 or 4.x, and the cluster is being created in an Amazon
	// VPC.
	//
	// If you enable in-transit encryption, you must also specify a value for CacheSubnetGroup.
	//
	// Required: Only available when creating a replication group in an Amazon VPC
	// using redis version 3.2.6, 4.x or later.
	//
	// Default: false
	//
	// For HIPAA compliance, you must specify TransitEncryptionEnabled as true,
	// an AuthToken, and a CacheSubnetGroup.
	TransitEncryptionEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s CreateReplicationGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateReplicationGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateReplicationGroupInput"}

	if s.ReplicationGroupDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupDescription"))
	}

	if s.ReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupId"))
	}
	if s.NodeGroupConfiguration != nil {
		for i, v := range s.NodeGroupConfiguration {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "NodeGroupConfiguration", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateReplicationGroupResult
type CreateReplicationGroupOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *ReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s CreateReplicationGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateReplicationGroup = "CreateReplicationGroup"

// CreateReplicationGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Creates a Redis (cluster mode disabled) or a Redis (cluster mode enabled)
// replication group.
//
// A Redis (cluster mode disabled) replication group is a collection of clusters,
// where one of the clusters is a read/write primary and the others are read-only
// replicas. Writes to the primary are asynchronously propagated to the replicas.
//
// A Redis (cluster mode enabled) replication group is a collection of 1 to
// 90 node groups (shards). Each node group (shard) has one read/write primary
// node and up to 5 read-only replica nodes. Writes to the primary are asynchronously
// propagated to the replicas. Redis (cluster mode enabled) replication groups
// partition the data across node groups (shards).
//
// When a Redis (cluster mode disabled) replication group has been successfully
// created, you can add one or more read replicas to it, up to a total of 5
// read replicas. You cannot alter a Redis (cluster mode enabled) replication
// group after it has been created. However, if you need to increase or decrease
// the number of node groups (console: shards), you can avail yourself of ElastiCache
// for Redis' enhanced backup and restore. For more information, see Restoring
// From a Backup with Cluster Resizing (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-restoring.html)
// in the ElastiCache User Guide.
//
// This operation is valid for Redis only.
//
//    // Example sending a request using CreateReplicationGroupRequest.
//    req := client.CreateReplicationGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateReplicationGroup
func (c *Client) CreateReplicationGroupRequest(input *CreateReplicationGroupInput) CreateReplicationGroupRequest {
	op := &aws.Operation{
		Name:       opCreateReplicationGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateReplicationGroupInput{}
	}

	req := c.newRequest(op, input, &CreateReplicationGroupOutput{})
	return CreateReplicationGroupRequest{Request: req, Input: input, Copy: c.CreateReplicationGroupRequest}
}

// CreateReplicationGroupRequest is the request type for the
// CreateReplicationGroup API operation.
type CreateReplicationGroupRequest struct {
	*aws.Request
	Input *CreateReplicationGroupInput
	Copy  func(*CreateReplicationGroupInput) CreateReplicationGroupRequest
}

// Send marshals and sends the CreateReplicationGroup API request.
func (r CreateReplicationGroupRequest) Send(ctx context.Context) (*CreateReplicationGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReplicationGroupResponse{
		CreateReplicationGroupOutput: r.Request.Data.(*CreateReplicationGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReplicationGroupResponse is the response type for the
// CreateReplicationGroup API operation.
type CreateReplicationGroupResponse struct {
	*CreateReplicationGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReplicationGroup request.
func (r *CreateReplicationGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}