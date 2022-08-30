// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a vps associated with your OVH Account.
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
//			_, err := ovh.GetVps(ctx, &GetVpsArgs{
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
func GetVps(ctx *pulumi.Context, args *GetVpsArgs, opts ...pulumi.InvokeOption) (*GetVpsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetVpsResult
	err := ctx.Invoke("ovh:index/getVps:getVps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVps.
type GetVpsArgs struct {
	// The serviceName of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getVps.
type GetVpsResult struct {
	// The ovh cluster the vps is in
	Cluster string `pulumi:"cluster"`
	// The datacenter in which the vps is located
	// * `datacenter.longname` - The fullname of the datacenter (ex: "Strasbourg SBG1")
	// * `datacenter.name` - The short name of the datacenter (ex: "sbg1)
	Datacenter map[string]string `pulumi:"datacenter"`
	// The displayed name in the ovh web admin
	Displayname string `pulumi:"displayname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of IPs addresses attached to the vps
	Ips []string `pulumi:"ips"`
	// The keymap for the ip kvm, valid values "", "fr", "us"
	Keymap string `pulumi:"keymap"`
	// The amount of memory in MB of the vps.
	Memory int `pulumi:"memory"`
	// A dict describing the type of vps.
	// * `model.name` - The model name (ex: model1)
	// * `model.offer` - The model human description (ex: "VPS 2016 SSD 1")
	// * `model.version` - The model version (ex: "2017v2")
	Model map[string]string `pulumi:"model"`
	Name  string            `pulumi:"name"`
	// The source of the boot kernel
	Netbootmode string `pulumi:"netbootmode"`
	// The type of offer (ssd, cloud, classic)
	Offertype   string `pulumi:"offertype"`
	ServiceName string `pulumi:"serviceName"`
	// A boolean to indicate if OVH sla monitoring is active.
	Slamonitoring bool `pulumi:"slamonitoring"`
	// The state of the vps
	State string `pulumi:"state"`
	// The type of server
	Type string `pulumi:"type"`
	// The number of vcore of the vps
	Vcore int `pulumi:"vcore"`
	// The OVH zone where the vps is
	Zone string `pulumi:"zone"`
}

func GetVpsOutput(ctx *pulumi.Context, args GetVpsOutputArgs, opts ...pulumi.InvokeOption) GetVpsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpsResult, error) {
			args := v.(GetVpsArgs)
			r, err := GetVps(ctx, &args, opts...)
			var s GetVpsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpsResultOutput)
}

// A collection of arguments for invoking getVps.
type GetVpsOutputArgs struct {
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetVpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpsArgs)(nil)).Elem()
}

// A collection of values returned by getVps.
type GetVpsResultOutput struct{ *pulumi.OutputState }

func (GetVpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpsResult)(nil)).Elem()
}

func (o GetVpsResultOutput) ToGetVpsResultOutput() GetVpsResultOutput {
	return o
}

func (o GetVpsResultOutput) ToGetVpsResultOutputWithContext(ctx context.Context) GetVpsResultOutput {
	return o
}

// The ovh cluster the vps is in
func (o GetVpsResultOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Cluster }).(pulumi.StringOutput)
}

// The datacenter in which the vps is located
// * `datacenter.longname` - The fullname of the datacenter (ex: "Strasbourg SBG1")
// * `datacenter.name` - The short name of the datacenter (ex: "sbg1)
func (o GetVpsResultOutput) Datacenter() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetVpsResult) map[string]string { return v.Datacenter }).(pulumi.StringMapOutput)
}

// The displayed name in the ovh web admin
func (o GetVpsResultOutput) Displayname() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Displayname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of IPs addresses attached to the vps
func (o GetVpsResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpsResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

// The keymap for the ip kvm, valid values "", "fr", "us"
func (o GetVpsResultOutput) Keymap() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Keymap }).(pulumi.StringOutput)
}

// The amount of memory in MB of the vps.
func (o GetVpsResultOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v GetVpsResult) int { return v.Memory }).(pulumi.IntOutput)
}

// A dict describing the type of vps.
// * `model.name` - The model name (ex: model1)
// * `model.offer` - The model human description (ex: "VPS 2016 SSD 1")
// * `model.version` - The model version (ex: "2017v2")
func (o GetVpsResultOutput) Model() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetVpsResult) map[string]string { return v.Model }).(pulumi.StringMapOutput)
}

func (o GetVpsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Name }).(pulumi.StringOutput)
}

// The source of the boot kernel
func (o GetVpsResultOutput) Netbootmode() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Netbootmode }).(pulumi.StringOutput)
}

// The type of offer (ssd, cloud, classic)
func (o GetVpsResultOutput) Offertype() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Offertype }).(pulumi.StringOutput)
}

func (o GetVpsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// A boolean to indicate if OVH sla monitoring is active.
func (o GetVpsResultOutput) Slamonitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v GetVpsResult) bool { return v.Slamonitoring }).(pulumi.BoolOutput)
}

// The state of the vps
func (o GetVpsResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.State }).(pulumi.StringOutput)
}

// The type of server
func (o GetVpsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Type }).(pulumi.StringOutput)
}

// The number of vcore of the vps
func (o GetVpsResultOutput) Vcore() pulumi.IntOutput {
	return o.ApplyT(func(v GetVpsResult) int { return v.Vcore }).(pulumi.IntOutput)
}

// The OVH zone where the vps is
func (o GetVpsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpsResultOutput{})
}
