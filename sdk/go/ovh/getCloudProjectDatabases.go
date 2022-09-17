// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of managed databases of a public cloud project.
//
// ## Example Usage
//
// To get the list of database clusters service for a given engine:
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			dbs, err := ovh.GetCloudProjectDatabases(ctx, &GetCloudProjectDatabasesArgs{
//				ServiceName: "XXXXXX",
//				Engine:      "YYYY",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("clusterIds", dbs.ClusterIds)
//			return nil
//		})
//	}
//
// ```
func GetCloudProjectDatabases(ctx *pulumi.Context, args *GetCloudProjectDatabasesArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectDatabasesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCloudProjectDatabasesResult
	err := ctx.Invoke("ovh:index/getCloudProjectDatabases:getCloudProjectDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectDatabases.
type GetCloudProjectDatabasesArgs struct {
	// The database engine you want to list. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectDatabases.
type GetCloudProjectDatabasesResult struct {
	// The list of managed databases ids of the project.
	ClusterIds []string `pulumi:"clusterIds"`
	Engine     string   `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	ServiceName string `pulumi:"serviceName"`
}

func GetCloudProjectDatabasesOutput(ctx *pulumi.Context, args GetCloudProjectDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectDatabasesResult, error) {
			args := v.(GetCloudProjectDatabasesArgs)
			r, err := GetCloudProjectDatabases(ctx, &args, opts...)
			var s GetCloudProjectDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectDatabasesResultOutput)
}

// A collection of arguments for invoking getCloudProjectDatabases.
type GetCloudProjectDatabasesOutputArgs struct {
	// The database engine you want to list. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetCloudProjectDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectDatabases.
type GetCloudProjectDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectDatabasesResult)(nil)).Elem()
}

func (o GetCloudProjectDatabasesResultOutput) ToGetCloudProjectDatabasesResultOutput() GetCloudProjectDatabasesResultOutput {
	return o
}

func (o GetCloudProjectDatabasesResultOutput) ToGetCloudProjectDatabasesResultOutputWithContext(ctx context.Context) GetCloudProjectDatabasesResultOutput {
	return o
}

// The list of managed databases ids of the project.
func (o GetCloudProjectDatabasesResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCloudProjectDatabasesResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

func (o GetCloudProjectDatabasesResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabasesResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloudProjectDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabasesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabasesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectDatabasesResultOutput{})
}
