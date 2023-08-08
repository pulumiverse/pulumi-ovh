// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a topic of a kafka cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			topic, err := CloudProjectDatabase.GetKafkaTopic(ctx, &cloudprojectdatabase.GetKafkaTopicArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//				Id:          "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("topicName", topic.Name)
//			return nil
//		})
//	}
//
// ```
func LookupKafkaTopic(ctx *pulumi.Context, args *LookupKafkaTopicArgs, opts ...pulumi.InvokeOption) (*LookupKafkaTopicResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKafkaTopicResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getKafkaTopic:getKafkaTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKafkaTopic.
type LookupKafkaTopicArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// Topic ID
	Id string `pulumi:"id"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKafkaTopic.
type LookupKafkaTopicResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// See Argument Reference above.
	Id string `pulumi:"id"`
	// Minimum insync replica accepted for this topic.
	MinInsyncReplicas int `pulumi:"minInsyncReplicas"`
	// Name of the topic.
	Name string `pulumi:"name"`
	// Number of partitions for this topic.
	Partitions int `pulumi:"partitions"`
	// Number of replication for this topic.
	Replication int `pulumi:"replication"`
	// Number of bytes for the retention of the data for this topic. Inferior to 0 mean Unlimited
	RetentionBytes int `pulumi:"retentionBytes"`
	// Number of hours for the retention of the data for this topic. Inferior to 0 mean Unlimited
	RetentionHours int `pulumi:"retentionHours"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func LookupKafkaTopicOutput(ctx *pulumi.Context, args LookupKafkaTopicOutputArgs, opts ...pulumi.InvokeOption) LookupKafkaTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKafkaTopicResult, error) {
			args := v.(LookupKafkaTopicArgs)
			r, err := LookupKafkaTopic(ctx, &args, opts...)
			var s LookupKafkaTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKafkaTopicResultOutput)
}

// A collection of arguments for invoking getKafkaTopic.
type LookupKafkaTopicOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Topic ID
	Id pulumi.StringInput `pulumi:"id"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupKafkaTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaTopicArgs)(nil)).Elem()
}

// A collection of values returned by getKafkaTopic.
type LookupKafkaTopicResultOutput struct{ *pulumi.OutputState }

func (LookupKafkaTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaTopicResult)(nil)).Elem()
}

func (o LookupKafkaTopicResultOutput) ToLookupKafkaTopicResultOutput() LookupKafkaTopicResultOutput {
	return o
}

func (o LookupKafkaTopicResultOutput) ToLookupKafkaTopicResultOutputWithContext(ctx context.Context) LookupKafkaTopicResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupKafkaTopicResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupKafkaTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// Minimum insync replica accepted for this topic.
func (o LookupKafkaTopicResultOutput) MinInsyncReplicas() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) int { return v.MinInsyncReplicas }).(pulumi.IntOutput)
}

// Name of the topic.
func (o LookupKafkaTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of partitions for this topic.
func (o LookupKafkaTopicResultOutput) Partitions() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) int { return v.Partitions }).(pulumi.IntOutput)
}

// Number of replication for this topic.
func (o LookupKafkaTopicResultOutput) Replication() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) int { return v.Replication }).(pulumi.IntOutput)
}

// Number of bytes for the retention of the data for this topic. Inferior to 0 mean Unlimited
func (o LookupKafkaTopicResultOutput) RetentionBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) int { return v.RetentionBytes }).(pulumi.IntOutput)
}

// Number of hours for the retention of the data for this topic. Inferior to 0 mean Unlimited
func (o LookupKafkaTopicResultOutput) RetentionHours() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) int { return v.RetentionHours }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupKafkaTopicResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaTopicResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKafkaTopicResultOutput{})
}
