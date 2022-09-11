// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of Vrack network ids available for your IPLoadbalancer associated with your OVH account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ovh.GetIpLoadbalancingVrackNetworks(ctx, &GetIpLoadbalancingVrackNetworksArgs{
//				ServiceName: "XXXXXX",
//				Subnet:      pulumi.StringRef("10.0.0.0/24"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIpLoadbalancingVrackNetworks(ctx *pulumi.Context, args *GetIpLoadbalancingVrackNetworksArgs, opts ...pulumi.InvokeOption) (*GetIpLoadbalancingVrackNetworksResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIpLoadbalancingVrackNetworksResult
	err := ctx.Invoke("ovh:index/getIpLoadbalancingVrackNetworks:getIpLoadbalancingVrackNetworks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpLoadbalancingVrackNetworks.
type GetIpLoadbalancingVrackNetworksArgs struct {
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Filters networks on the subnet.
	Subnet *string `pulumi:"subnet"`
	// Filters networks on the vlan id.
	VlanId *int `pulumi:"vlanId"`
}

// A collection of values returned by getIpLoadbalancingVrackNetworks.
type GetIpLoadbalancingVrackNetworksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of vrack network ids.
	Results     []int   `pulumi:"results"`
	ServiceName string  `pulumi:"serviceName"`
	Subnet      *string `pulumi:"subnet"`
	VlanId      *int    `pulumi:"vlanId"`
}

func GetIpLoadbalancingVrackNetworksOutput(ctx *pulumi.Context, args GetIpLoadbalancingVrackNetworksOutputArgs, opts ...pulumi.InvokeOption) GetIpLoadbalancingVrackNetworksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpLoadbalancingVrackNetworksResult, error) {
			args := v.(GetIpLoadbalancingVrackNetworksArgs)
			r, err := GetIpLoadbalancingVrackNetworks(ctx, &args, opts...)
			var s GetIpLoadbalancingVrackNetworksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpLoadbalancingVrackNetworksResultOutput)
}

// A collection of arguments for invoking getIpLoadbalancingVrackNetworks.
type GetIpLoadbalancingVrackNetworksOutputArgs struct {
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Filters networks on the subnet.
	Subnet pulumi.StringPtrInput `pulumi:"subnet"`
	// Filters networks on the vlan id.
	VlanId pulumi.IntPtrInput `pulumi:"vlanId"`
}

func (GetIpLoadbalancingVrackNetworksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpLoadbalancingVrackNetworksArgs)(nil)).Elem()
}

// A collection of values returned by getIpLoadbalancingVrackNetworks.
type GetIpLoadbalancingVrackNetworksResultOutput struct{ *pulumi.OutputState }

func (GetIpLoadbalancingVrackNetworksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpLoadbalancingVrackNetworksResult)(nil)).Elem()
}

func (o GetIpLoadbalancingVrackNetworksResultOutput) ToGetIpLoadbalancingVrackNetworksResultOutput() GetIpLoadbalancingVrackNetworksResultOutput {
	return o
}

func (o GetIpLoadbalancingVrackNetworksResultOutput) ToGetIpLoadbalancingVrackNetworksResultOutputWithContext(ctx context.Context) GetIpLoadbalancingVrackNetworksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpLoadbalancingVrackNetworksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworksResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of vrack network ids.
func (o GetIpLoadbalancingVrackNetworksResultOutput) Results() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworksResult) []int { return v.Results }).(pulumi.IntArrayOutput)
}

func (o GetIpLoadbalancingVrackNetworksResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworksResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetIpLoadbalancingVrackNetworksResultOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworksResult) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o GetIpLoadbalancingVrackNetworksResultOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworksResult) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpLoadbalancingVrackNetworksResultOutput{})
}
