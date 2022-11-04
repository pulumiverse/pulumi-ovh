// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a vrack network for your IP Loadbalancing service.
type VrackNetwork struct {
	pulumi.CustomResourceState

	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVHcloud API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds pulumi.IntArrayOutput `pulumi:"farmIds"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp pulumi.StringOutput `pulumi:"natIp"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// IP block of the private network in the vRack
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan pulumi.IntOutput `pulumi:"vlan"`
	// (Required) Internal Load Balancer identifier of the vRack private network
	VrackNetworkId pulumi.IntOutput `pulumi:"vrackNetworkId"`
}

// NewVrackNetwork registers a new resource with the given unique name, arguments, and options.
func NewVrackNetwork(ctx *pulumi.Context,
	name string, args *VrackNetworkArgs, opts ...pulumi.ResourceOption) (*VrackNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NatIp == nil {
		return nil, errors.New("invalid value for required argument 'NatIp'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VrackNetwork
	err := ctx.RegisterResource("ovh:IpLoadBalancing/vrackNetwork:VrackNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVrackNetwork gets an existing VrackNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVrackNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VrackNetworkState, opts ...pulumi.ResourceOption) (*VrackNetwork, error) {
	var resource VrackNetwork
	err := ctx.ReadResource("ovh:IpLoadBalancing/vrackNetwork:VrackNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VrackNetwork resources.
type vrackNetworkState struct {
	// Human readable name for your vrack network
	DisplayName *string `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVHcloud API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds []int `pulumi:"farmIds"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp *string `pulumi:"natIp"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// IP block of the private network in the vRack
	Subnet *string `pulumi:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan *int `pulumi:"vlan"`
	// (Required) Internal Load Balancer identifier of the vRack private network
	VrackNetworkId *int `pulumi:"vrackNetworkId"`
}

type VrackNetworkState struct {
	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrInput
	// This attribute is there for documentation purpose only and isnt passed to the OVHcloud API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds pulumi.IntArrayInput
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// IP block of the private network in the vRack
	Subnet pulumi.StringPtrInput
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan pulumi.IntPtrInput
	// (Required) Internal Load Balancer identifier of the vRack private network
	VrackNetworkId pulumi.IntPtrInput
}

func (VrackNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackNetworkState)(nil)).Elem()
}

type vrackNetworkArgs struct {
	// Human readable name for your vrack network
	DisplayName *string `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVHcloud API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds []int `pulumi:"farmIds"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp string `pulumi:"natIp"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// IP block of the private network in the vRack
	Subnet string `pulumi:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan *int `pulumi:"vlan"`
}

// The set of arguments for constructing a VrackNetwork resource.
type VrackNetworkArgs struct {
	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrInput
	// This attribute is there for documentation purpose only and isnt passed to the OVHcloud API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds pulumi.IntArrayInput
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp pulumi.StringInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// IP block of the private network in the vRack
	Subnet pulumi.StringInput
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan pulumi.IntPtrInput
}

func (VrackNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackNetworkArgs)(nil)).Elem()
}

type VrackNetworkInput interface {
	pulumi.Input

	ToVrackNetworkOutput() VrackNetworkOutput
	ToVrackNetworkOutputWithContext(ctx context.Context) VrackNetworkOutput
}

func (*VrackNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackNetwork)(nil)).Elem()
}

func (i *VrackNetwork) ToVrackNetworkOutput() VrackNetworkOutput {
	return i.ToVrackNetworkOutputWithContext(context.Background())
}

func (i *VrackNetwork) ToVrackNetworkOutputWithContext(ctx context.Context) VrackNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackNetworkOutput)
}

// VrackNetworkArrayInput is an input type that accepts VrackNetworkArray and VrackNetworkArrayOutput values.
// You can construct a concrete instance of `VrackNetworkArrayInput` via:
//
//	VrackNetworkArray{ VrackNetworkArgs{...} }
type VrackNetworkArrayInput interface {
	pulumi.Input

	ToVrackNetworkArrayOutput() VrackNetworkArrayOutput
	ToVrackNetworkArrayOutputWithContext(context.Context) VrackNetworkArrayOutput
}

type VrackNetworkArray []VrackNetworkInput

func (VrackNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackNetwork)(nil)).Elem()
}

func (i VrackNetworkArray) ToVrackNetworkArrayOutput() VrackNetworkArrayOutput {
	return i.ToVrackNetworkArrayOutputWithContext(context.Background())
}

func (i VrackNetworkArray) ToVrackNetworkArrayOutputWithContext(ctx context.Context) VrackNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackNetworkArrayOutput)
}

// VrackNetworkMapInput is an input type that accepts VrackNetworkMap and VrackNetworkMapOutput values.
// You can construct a concrete instance of `VrackNetworkMapInput` via:
//
//	VrackNetworkMap{ "key": VrackNetworkArgs{...} }
type VrackNetworkMapInput interface {
	pulumi.Input

	ToVrackNetworkMapOutput() VrackNetworkMapOutput
	ToVrackNetworkMapOutputWithContext(context.Context) VrackNetworkMapOutput
}

type VrackNetworkMap map[string]VrackNetworkInput

func (VrackNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackNetwork)(nil)).Elem()
}

func (i VrackNetworkMap) ToVrackNetworkMapOutput() VrackNetworkMapOutput {
	return i.ToVrackNetworkMapOutputWithContext(context.Background())
}

func (i VrackNetworkMap) ToVrackNetworkMapOutputWithContext(ctx context.Context) VrackNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackNetworkMapOutput)
}

type VrackNetworkOutput struct{ *pulumi.OutputState }

func (VrackNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackNetwork)(nil)).Elem()
}

func (o VrackNetworkOutput) ToVrackNetworkOutput() VrackNetworkOutput {
	return o
}

func (o VrackNetworkOutput) ToVrackNetworkOutputWithContext(ctx context.Context) VrackNetworkOutput {
	return o
}

// Human readable name for your vrack network
func (o VrackNetworkOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VrackNetwork) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// This attribute is there for documentation purpose only and isnt passed to the OVHcloud API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
func (o VrackNetworkOutput) FarmIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *VrackNetwork) pulumi.IntArrayOutput { return v.FarmIds }).(pulumi.IntArrayOutput)
}

// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
func (o VrackNetworkOutput) NatIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackNetwork) pulumi.StringOutput { return v.NatIp }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o VrackNetworkOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackNetwork) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// IP block of the private network in the vRack
func (o VrackNetworkOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackNetwork) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
func (o VrackNetworkOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v *VrackNetwork) pulumi.IntOutput { return v.Vlan }).(pulumi.IntOutput)
}

// (Required) Internal Load Balancer identifier of the vRack private network
func (o VrackNetworkOutput) VrackNetworkId() pulumi.IntOutput {
	return o.ApplyT(func(v *VrackNetwork) pulumi.IntOutput { return v.VrackNetworkId }).(pulumi.IntOutput)
}

type VrackNetworkArrayOutput struct{ *pulumi.OutputState }

func (VrackNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackNetwork)(nil)).Elem()
}

func (o VrackNetworkArrayOutput) ToVrackNetworkArrayOutput() VrackNetworkArrayOutput {
	return o
}

func (o VrackNetworkArrayOutput) ToVrackNetworkArrayOutputWithContext(ctx context.Context) VrackNetworkArrayOutput {
	return o
}

func (o VrackNetworkArrayOutput) Index(i pulumi.IntInput) VrackNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VrackNetwork {
		return vs[0].([]*VrackNetwork)[vs[1].(int)]
	}).(VrackNetworkOutput)
}

type VrackNetworkMapOutput struct{ *pulumi.OutputState }

func (VrackNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackNetwork)(nil)).Elem()
}

func (o VrackNetworkMapOutput) ToVrackNetworkMapOutput() VrackNetworkMapOutput {
	return o
}

func (o VrackNetworkMapOutput) ToVrackNetworkMapOutputWithContext(ctx context.Context) VrackNetworkMapOutput {
	return o
}

func (o VrackNetworkMapOutput) MapIndex(k pulumi.StringInput) VrackNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VrackNetwork {
		return vs[0].(map[string]*VrackNetwork)[vs[1].(string)]
	}).(VrackNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VrackNetworkInput)(nil)).Elem(), &VrackNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrackNetworkArrayInput)(nil)).Elem(), VrackNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrackNetworkMapInput)(nil)).Elem(), VrackNetworkMap{})
	pulumi.RegisterOutputType(VrackNetworkOutput{})
	pulumi.RegisterOutputType(VrackNetworkArrayOutput{})
	pulumi.RegisterOutputType(VrackNetworkMapOutput{})
}
