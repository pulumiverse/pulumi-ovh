// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an ip service.
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
//			_, err := ovh.LookupIpService(ctx, &GetIpServiceArgs{
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
func LookupIpService(ctx *pulumi.Context, args *LookupIpServiceArgs, opts ...pulumi.InvokeOption) (*LookupIpServiceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupIpServiceResult
	err := ctx.Invoke("ovh:index/getIpService:getIpService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpService.
type LookupIpServiceArgs struct {
	// The service name
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getIpService.
type LookupIpServiceResult struct {
	// can be terminated
	CanBeTerminated bool `pulumi:"canBeTerminated"`
	// country
	Country string `pulumi:"country"`
	// Custom description on your ip
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ip block
	Ip string `pulumi:"ip"`
	// IP block organisation Id
	OrganisationId string `pulumi:"organisationId"`
	// Routage information
	RoutedTos []GetIpServiceRoutedTo `pulumi:"routedTos"`
	// Service where ip is routed to
	ServiceName string `pulumi:"serviceName"`
	// Possible values for ip type (    "cdn", "cloud", "dedicated", "failover", "hostedSsl", "housing", "loadBalancing", "mail", "overthebox", "pcc", "pci", "private", "vpn", "vps", "vrack", "xdsl")
	Type string `pulumi:"type"`
}

func LookupIpServiceOutput(ctx *pulumi.Context, args LookupIpServiceOutputArgs, opts ...pulumi.InvokeOption) LookupIpServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpServiceResult, error) {
			args := v.(LookupIpServiceArgs)
			r, err := LookupIpService(ctx, &args, opts...)
			var s LookupIpServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpServiceResultOutput)
}

// A collection of arguments for invoking getIpService.
type LookupIpServiceOutputArgs struct {
	// The service name
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupIpServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpServiceArgs)(nil)).Elem()
}

// A collection of values returned by getIpService.
type LookupIpServiceResultOutput struct{ *pulumi.OutputState }

func (LookupIpServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpServiceResult)(nil)).Elem()
}

func (o LookupIpServiceResultOutput) ToLookupIpServiceResultOutput() LookupIpServiceResultOutput {
	return o
}

func (o LookupIpServiceResultOutput) ToLookupIpServiceResultOutputWithContext(ctx context.Context) LookupIpServiceResultOutput {
	return o
}

// can be terminated
func (o LookupIpServiceResultOutput) CanBeTerminated() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpServiceResult) bool { return v.CanBeTerminated }).(pulumi.BoolOutput)
}

// country
func (o LookupIpServiceResultOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpServiceResult) string { return v.Country }).(pulumi.StringOutput)
}

// Custom description on your ip
func (o LookupIpServiceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpServiceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIpServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// ip block
func (o LookupIpServiceResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpServiceResult) string { return v.Ip }).(pulumi.StringOutput)
}

// IP block organisation Id
func (o LookupIpServiceResultOutput) OrganisationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpServiceResult) string { return v.OrganisationId }).(pulumi.StringOutput)
}

// Routage information
func (o LookupIpServiceResultOutput) RoutedTos() GetIpServiceRoutedToArrayOutput {
	return o.ApplyT(func(v LookupIpServiceResult) []GetIpServiceRoutedTo { return v.RoutedTos }).(GetIpServiceRoutedToArrayOutput)
}

// Service where ip is routed to
func (o LookupIpServiceResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpServiceResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Possible values for ip type (    "cdn", "cloud", "dedicated", "failover", "hostedSsl", "housing", "loadBalancing", "mail", "overthebox", "pcc", "pci", "private", "vpn", "vps", "vrack", "xdsl")
func (o LookupIpServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpServiceResultOutput{})
}
