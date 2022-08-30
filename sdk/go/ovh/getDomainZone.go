// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a domain zone.
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
//			_, err := ovh.LookupDomainZone(ctx, &GetDomainZoneArgs{
//				Name: "mysite.ovh",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDomainZone(ctx *pulumi.Context, args *LookupDomainZoneArgs, opts ...pulumi.InvokeOption) (*LookupDomainZoneResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupDomainZoneResult
	err := ctx.Invoke("ovh:index/getDomainZone:getDomainZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainZone.
type LookupDomainZoneArgs struct {
	// The name of the domain zone.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDomainZone.
type LookupDomainZoneResult struct {
	// Is DNSSEC supported by this zone
	DnssecSupported bool `pulumi:"dnssecSupported"`
	// hasDnsAnycast flag of the DNS zone
	HasDnsAnycast bool `pulumi:"hasDnsAnycast"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Last update date of the DNS zone
	LastUpdate string `pulumi:"lastUpdate"`
	Name       string `pulumi:"name"`
	// Name servers that host the DNS zone
	NameServers []string `pulumi:"nameServers"`
}

func LookupDomainZoneOutput(ctx *pulumi.Context, args LookupDomainZoneOutputArgs, opts ...pulumi.InvokeOption) LookupDomainZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainZoneResult, error) {
			args := v.(LookupDomainZoneArgs)
			r, err := LookupDomainZone(ctx, &args, opts...)
			var s LookupDomainZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainZoneResultOutput)
}

// A collection of arguments for invoking getDomainZone.
type LookupDomainZoneOutputArgs struct {
	// The name of the domain zone.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDomainZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainZoneArgs)(nil)).Elem()
}

// A collection of values returned by getDomainZone.
type LookupDomainZoneResultOutput struct{ *pulumi.OutputState }

func (LookupDomainZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainZoneResult)(nil)).Elem()
}

func (o LookupDomainZoneResultOutput) ToLookupDomainZoneResultOutput() LookupDomainZoneResultOutput {
	return o
}

func (o LookupDomainZoneResultOutput) ToLookupDomainZoneResultOutputWithContext(ctx context.Context) LookupDomainZoneResultOutput {
	return o
}

// Is DNSSEC supported by this zone
func (o LookupDomainZoneResultOutput) DnssecSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainZoneResult) bool { return v.DnssecSupported }).(pulumi.BoolOutput)
}

// hasDnsAnycast flag of the DNS zone
func (o LookupDomainZoneResultOutput) HasDnsAnycast() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainZoneResult) bool { return v.HasDnsAnycast }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDomainZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// Last update date of the DNS zone
func (o LookupDomainZoneResultOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainZoneResult) string { return v.LastUpdate }).(pulumi.StringOutput)
}

func (o LookupDomainZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// Name servers that host the DNS zone
func (o LookupDomainZoneResultOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainZoneResult) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainZoneResultOutput{})
}
