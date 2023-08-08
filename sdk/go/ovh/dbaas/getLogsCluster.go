// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve informations about a DBaas logs cluster tenant.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dbaas.GetLogsCluster(ctx, &dbaas.GetLogsClusterArgs{
//				ServiceName: "ldp-xx-xxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLogsCluster(ctx *pulumi.Context, args *LookupLogsClusterArgs, opts ...pulumi.InvokeOption) (*LookupLogsClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLogsClusterResult
	err := ctx.Invoke("ovh:Dbaas/getLogsCluster:getLogsCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogsCluster.
type LookupLogsClusterArgs struct {
	// The service name. It's the ID of your Logs Data Platform instance.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getLogsCluster.
type LookupLogsClusterResult struct {
	// is allowed networks for ARCHIVE flow type
	ArchiveAllowedNetworks []string `pulumi:"archiveAllowedNetworks"`
	// is type of cluster (DEDICATED, PRO or TRIAL)
	ClusterType string `pulumi:"clusterType"`
	// is PEM for dedicated inputs
	DedicatedInputPem string `pulumi:"dedicatedInputPem"`
	// is allowed networks for DIRECT_INPUT flow type
	DirectInputAllowedNetworks []string `pulumi:"directInputAllowedNetworks"`
	// is PEM for direct inputs
	DirectInputPem string `pulumi:"directInputPem"`
	// is cluster hostname hosting the tenant
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// is true if all content generated by given service will be placed on this cluster
	IsDefault bool `pulumi:"isDefault"`
	// is true if given service can perform advanced operations on cluster
	IsUnlocked bool `pulumi:"isUnlocked"`
	// is allowed networks for QUERY flow type
	QueryAllowedNetworks []string `pulumi:"queryAllowedNetworks"`
	// is datacenter localization
	Region      string `pulumi:"region"`
	ServiceName string `pulumi:"serviceName"`
	// is the URN of the DBaas logs instance
	Urn string `pulumi:"urn"`
}

func LookupLogsClusterOutput(ctx *pulumi.Context, args LookupLogsClusterOutputArgs, opts ...pulumi.InvokeOption) LookupLogsClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLogsClusterResult, error) {
			args := v.(LookupLogsClusterArgs)
			r, err := LookupLogsCluster(ctx, &args, opts...)
			var s LookupLogsClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLogsClusterResultOutput)
}

// A collection of arguments for invoking getLogsCluster.
type LookupLogsClusterOutputArgs struct {
	// The service name. It's the ID of your Logs Data Platform instance.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupLogsClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogsClusterArgs)(nil)).Elem()
}

// A collection of values returned by getLogsCluster.
type LookupLogsClusterResultOutput struct{ *pulumi.OutputState }

func (LookupLogsClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogsClusterResult)(nil)).Elem()
}

func (o LookupLogsClusterResultOutput) ToLookupLogsClusterResultOutput() LookupLogsClusterResultOutput {
	return o
}

func (o LookupLogsClusterResultOutput) ToLookupLogsClusterResultOutputWithContext(ctx context.Context) LookupLogsClusterResultOutput {
	return o
}

// is allowed networks for ARCHIVE flow type
func (o LookupLogsClusterResultOutput) ArchiveAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) []string { return v.ArchiveAllowedNetworks }).(pulumi.StringArrayOutput)
}

// is type of cluster (DEDICATED, PRO or TRIAL)
func (o LookupLogsClusterResultOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) string { return v.ClusterType }).(pulumi.StringOutput)
}

// is PEM for dedicated inputs
func (o LookupLogsClusterResultOutput) DedicatedInputPem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) string { return v.DedicatedInputPem }).(pulumi.StringOutput)
}

// is allowed networks for DIRECT_INPUT flow type
func (o LookupLogsClusterResultOutput) DirectInputAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) []string { return v.DirectInputAllowedNetworks }).(pulumi.StringArrayOutput)
}

// is PEM for direct inputs
func (o LookupLogsClusterResultOutput) DirectInputPem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) string { return v.DirectInputPem }).(pulumi.StringOutput)
}

// is cluster hostname hosting the tenant
func (o LookupLogsClusterResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLogsClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// is true if all content generated by given service will be placed on this cluster
func (o LookupLogsClusterResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// is true if given service can perform advanced operations on cluster
func (o LookupLogsClusterResultOutput) IsUnlocked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) bool { return v.IsUnlocked }).(pulumi.BoolOutput)
}

// is allowed networks for QUERY flow type
func (o LookupLogsClusterResultOutput) QueryAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) []string { return v.QueryAllowedNetworks }).(pulumi.StringArrayOutput)
}

// is datacenter localization
func (o LookupLogsClusterResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupLogsClusterResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// is the URN of the DBaas logs instance
func (o LookupLogsClusterResultOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsClusterResult) string { return v.Urn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLogsClusterResultOutput{})
}
