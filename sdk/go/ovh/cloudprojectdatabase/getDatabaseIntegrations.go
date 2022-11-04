// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of integrations of a database cluster associated with a public cloud project.
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
//			integrations, err := CloudProjectDatabase.GetDatabaseIntegrations(ctx, &cloudprojectdatabase.GetDatabaseIntegrationsArgs{
//				ServiceName: "XXX",
//				Engine:      "YYY",
//				ClusterId:   "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("integrationIds", integrations.IntegrationIds)
//			return nil
//		})
//	}
//
// ```
func GetDatabaseIntegrations(ctx *pulumi.Context, args *GetDatabaseIntegrationsArgs, opts ...pulumi.InvokeOption) (*GetDatabaseIntegrationsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatabaseIntegrationsResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabaseIntegrations:getDatabaseIntegrations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseIntegrations.
type GetDatabaseIntegrationsArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want to list integrations. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`
	Engine string `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabaseIntegrations.
type GetDatabaseIntegrationsResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of integrations ids of the database cluster associated with the project.
	IntegrationIds []string `pulumi:"integrationIds"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetDatabaseIntegrationsOutput(ctx *pulumi.Context, args GetDatabaseIntegrationsOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseIntegrationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseIntegrationsResult, error) {
			args := v.(GetDatabaseIntegrationsArgs)
			r, err := GetDatabaseIntegrations(ctx, &args, opts...)
			var s GetDatabaseIntegrationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabaseIntegrationsResultOutput)
}

// A collection of arguments for invoking getDatabaseIntegrations.
type GetDatabaseIntegrationsOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The engine of the database cluster you want to list integrations. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`
	Engine pulumi.StringInput `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDatabaseIntegrationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseIntegrationsArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseIntegrations.
type GetDatabaseIntegrationsResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseIntegrationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseIntegrationsResult)(nil)).Elem()
}

func (o GetDatabaseIntegrationsResultOutput) ToGetDatabaseIntegrationsResultOutput() GetDatabaseIntegrationsResultOutput {
	return o
}

func (o GetDatabaseIntegrationsResultOutput) ToGetDatabaseIntegrationsResultOutputWithContext(ctx context.Context) GetDatabaseIntegrationsResultOutput {
	return o
}

// See Argument Reference above.
func (o GetDatabaseIntegrationsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseIntegrationsResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationsResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabaseIntegrationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of integrations ids of the database cluster associated with the project.
func (o GetDatabaseIntegrationsResultOutput) IntegrationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationsResult) []string { return v.IntegrationIds }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o GetDatabaseIntegrationsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseIntegrationsResultOutput{})
}