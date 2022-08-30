// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a OVH Managed Kubernetes Service cluster IP restrictions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-ovh/sdk/go/ovh"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			iprestrictions, err := ovh.LookupCloudProjectKubeIpRestrictions(ctx, &GetCloudProjectKubeIpRestrictionsArgs{
//				ServiceName: "XXXXXX",
//				KubeId:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ips", iprestrictions.Ips)
//			return nil
//		})
//	}
//
// ```
func LookupCloudProjectKubeIpRestrictions(ctx *pulumi.Context, args *LookupCloudProjectKubeIpRestrictionsArgs, opts ...pulumi.InvokeOption) (*LookupCloudProjectKubeIpRestrictionsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCloudProjectKubeIpRestrictionsResult
	err := ctx.Invoke("ovh:index/getCloudProjectKubeIpRestrictions:getCloudProjectKubeIpRestrictions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectKubeIpRestrictions.
type LookupCloudProjectKubeIpRestrictionsArgs struct {
	// The id of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectKubeIpRestrictions.
type LookupCloudProjectKubeIpRestrictionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of CIDRs that restricts the access to the API server.
	Ips []string `pulumi:"ips"`
	// See Argument Reference above.
	KubeId string `pulumi:"kubeId"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func LookupCloudProjectKubeIpRestrictionsOutput(ctx *pulumi.Context, args LookupCloudProjectKubeIpRestrictionsOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProjectKubeIpRestrictionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProjectKubeIpRestrictionsResult, error) {
			args := v.(LookupCloudProjectKubeIpRestrictionsArgs)
			r, err := LookupCloudProjectKubeIpRestrictions(ctx, &args, opts...)
			var s LookupCloudProjectKubeIpRestrictionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudProjectKubeIpRestrictionsResultOutput)
}

// A collection of arguments for invoking getCloudProjectKubeIpRestrictions.
type LookupCloudProjectKubeIpRestrictionsOutputArgs struct {
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringInput `pulumi:"kubeId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCloudProjectKubeIpRestrictionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectKubeIpRestrictionsArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectKubeIpRestrictions.
type LookupCloudProjectKubeIpRestrictionsResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProjectKubeIpRestrictionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectKubeIpRestrictionsResult)(nil)).Elem()
}

func (o LookupCloudProjectKubeIpRestrictionsResultOutput) ToLookupCloudProjectKubeIpRestrictionsResultOutput() LookupCloudProjectKubeIpRestrictionsResultOutput {
	return o
}

func (o LookupCloudProjectKubeIpRestrictionsResultOutput) ToLookupCloudProjectKubeIpRestrictionsResultOutputWithContext(ctx context.Context) LookupCloudProjectKubeIpRestrictionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudProjectKubeIpRestrictionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectKubeIpRestrictionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of CIDRs that restricts the access to the API server.
func (o LookupCloudProjectKubeIpRestrictionsResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudProjectKubeIpRestrictionsResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o LookupCloudProjectKubeIpRestrictionsResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectKubeIpRestrictionsResult) string { return v.KubeId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupCloudProjectKubeIpRestrictionsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectKubeIpRestrictionsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProjectKubeIpRestrictionsResultOutput{})
}
