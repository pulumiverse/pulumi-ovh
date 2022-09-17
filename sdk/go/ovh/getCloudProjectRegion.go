// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a region associated with a public cloud project. The region must be associated with the project.
//
// ## Example Usage
//
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
//			_, err := ovh.GetCloudProjectRegion(ctx, &GetCloudProjectRegionArgs{
//				Name:        "GRA1",
//				ServiceName: "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCloudProjectRegion(ctx *pulumi.Context, args *GetCloudProjectRegionArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectRegionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCloudProjectRegionResult
	err := ctx.Invoke("ovh:index/getCloudProjectRegion:getCloudProjectRegion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectRegion.
type GetCloudProjectRegionArgs struct {
	// The name of the region associated with the public cloud
	// project.
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectRegion.
type GetCloudProjectRegionResult struct {
	// the code of the geographic continent the region is running.
	// E.g.: EU for Europe, US for America...
	ContinentCode string `pulumi:"continentCode"`
	// The location code of the datacenter.
	// E.g.: "GRA", meaning Gravelines, for region "GRA1"
	DatacenterLocation string `pulumi:"datacenterLocation"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// the name of the public cloud service
	Name        string `pulumi:"name"`
	ServiceName string `pulumi:"serviceName"`
	// The list of public cloud services running within the region
	Services []GetCloudProjectRegionService `pulumi:"services"`
}

func GetCloudProjectRegionOutput(ctx *pulumi.Context, args GetCloudProjectRegionOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectRegionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectRegionResult, error) {
			args := v.(GetCloudProjectRegionArgs)
			r, err := GetCloudProjectRegion(ctx, &args, opts...)
			var s GetCloudProjectRegionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectRegionResultOutput)
}

// A collection of arguments for invoking getCloudProjectRegion.
type GetCloudProjectRegionOutputArgs struct {
	// The name of the region associated with the public cloud
	// project.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetCloudProjectRegionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectRegionArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectRegion.
type GetCloudProjectRegionResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectRegionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectRegionResult)(nil)).Elem()
}

func (o GetCloudProjectRegionResultOutput) ToGetCloudProjectRegionResultOutput() GetCloudProjectRegionResultOutput {
	return o
}

func (o GetCloudProjectRegionResultOutput) ToGetCloudProjectRegionResultOutputWithContext(ctx context.Context) GetCloudProjectRegionResultOutput {
	return o
}

// the code of the geographic continent the region is running.
// E.g.: EU for Europe, US for America...
func (o GetCloudProjectRegionResultOutput) ContinentCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.ContinentCode }).(pulumi.StringOutput)
}

// The location code of the datacenter.
// E.g.: "GRA", meaning Gravelines, for region "GRA1"
func (o GetCloudProjectRegionResultOutput) DatacenterLocation() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.DatacenterLocation }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloudProjectRegionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.Id }).(pulumi.StringOutput)
}

// the name of the public cloud service
func (o GetCloudProjectRegionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCloudProjectRegionResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// The list of public cloud services running within the region
func (o GetCloudProjectRegionResultOutput) Services() GetCloudProjectRegionServiceArrayOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) []GetCloudProjectRegionService { return v.Services }).(GetCloudProjectRegionServiceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectRegionResultOutput{})
}
