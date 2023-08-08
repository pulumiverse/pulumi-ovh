// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a database of a database cluster associated with a public cloud project.
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
//			database, err := CloudProjectDatabase.GetDatabaseInstance(ctx, &cloudprojectdatabase.GetDatabaseInstanceArgs{
//				ServiceName: "XXX",
//				Engine:      "YYY",
//				ClusterId:   "ZZZ",
//				Name:        "UUU",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("databaseName", database.Name)
//			return nil
//		})
//	}
//
// ```
func LookupDatabaseInstance(ctx *pulumi.Context, args *LookupDatabaseInstanceArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseInstanceResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabaseInstance:getDatabaseInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseInstance.
type LookupDatabaseInstanceArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want database information. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine string `pulumi:"engine"`
	// Name of the database.
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabaseInstance.
type LookupDatabaseInstanceResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// Defines if the database has been created by default.
	Default bool `pulumi:"default"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the database.
	Name string `pulumi:"name"`
	// Current status of the database.
	ServiceName string `pulumi:"serviceName"`
}

func LookupDatabaseInstanceOutput(ctx *pulumi.Context, args LookupDatabaseInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseInstanceResult, error) {
			args := v.(LookupDatabaseInstanceArgs)
			r, err := LookupDatabaseInstance(ctx, &args, opts...)
			var s LookupDatabaseInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseInstanceResultOutput)
}

// A collection of arguments for invoking getDatabaseInstance.
type LookupDatabaseInstanceOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The engine of the database cluster you want database information. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine pulumi.StringInput `pulumi:"engine"`
	// Name of the database.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupDatabaseInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseInstance.
type LookupDatabaseInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseInstanceResult)(nil)).Elem()
}

func (o LookupDatabaseInstanceResultOutput) ToLookupDatabaseInstanceResultOutput() LookupDatabaseInstanceResultOutput {
	return o
}

func (o LookupDatabaseInstanceResultOutput) ToLookupDatabaseInstanceResultOutputWithContext(ctx context.Context) LookupDatabaseInstanceResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupDatabaseInstanceResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Defines if the database has been created by default.
func (o LookupDatabaseInstanceResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o LookupDatabaseInstanceResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatabaseInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the database.
func (o LookupDatabaseInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Current status of the database.
func (o LookupDatabaseInstanceResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseInstanceResultOutput{})
}
