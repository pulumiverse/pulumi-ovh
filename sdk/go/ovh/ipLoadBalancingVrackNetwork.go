// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a vrack network for your IP Loadbalancing service.
type IpLoadBalancingVrackNetwork struct {
	pulumi.CustomResourceState

	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
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

// NewIpLoadBalancingVrackNetwork registers a new resource with the given unique name, arguments, and options.
func NewIpLoadBalancingVrackNetwork(ctx *pulumi.Context,
	name string, args *IpLoadBalancingVrackNetworkArgs, opts ...pulumi.ResourceOption) (*IpLoadBalancingVrackNetwork, error) {
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
	var resource IpLoadBalancingVrackNetwork
	err := ctx.RegisterResource("ovh:index/ipLoadBalancingVrackNetwork:IpLoadBalancingVrackNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpLoadBalancingVrackNetwork gets an existing IpLoadBalancingVrackNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpLoadBalancingVrackNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpLoadBalancingVrackNetworkState, opts ...pulumi.ResourceOption) (*IpLoadBalancingVrackNetwork, error) {
	var resource IpLoadBalancingVrackNetwork
	err := ctx.ReadResource("ovh:index/ipLoadBalancingVrackNetwork:IpLoadBalancingVrackNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpLoadBalancingVrackNetwork resources.
type ipLoadBalancingVrackNetworkState struct {
	// Human readable name for your vrack network
	DisplayName *string `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
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

type IpLoadBalancingVrackNetworkState struct {
	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrInput
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
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

func (IpLoadBalancingVrackNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingVrackNetworkState)(nil)).Elem()
}

type ipLoadBalancingVrackNetworkArgs struct {
	// Human readable name for your vrack network
	DisplayName *string `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
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

// The set of arguments for constructing a IpLoadBalancingVrackNetwork resource.
type IpLoadBalancingVrackNetworkArgs struct {
	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrInput
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
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

func (IpLoadBalancingVrackNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingVrackNetworkArgs)(nil)).Elem()
}

type IpLoadBalancingVrackNetworkInput interface {
	pulumi.Input

	ToIpLoadBalancingVrackNetworkOutput() IpLoadBalancingVrackNetworkOutput
	ToIpLoadBalancingVrackNetworkOutputWithContext(ctx context.Context) IpLoadBalancingVrackNetworkOutput
}

func (*IpLoadBalancingVrackNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingVrackNetwork)(nil)).Elem()
}

func (i *IpLoadBalancingVrackNetwork) ToIpLoadBalancingVrackNetworkOutput() IpLoadBalancingVrackNetworkOutput {
	return i.ToIpLoadBalancingVrackNetworkOutputWithContext(context.Background())
}

func (i *IpLoadBalancingVrackNetwork) ToIpLoadBalancingVrackNetworkOutputWithContext(ctx context.Context) IpLoadBalancingVrackNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingVrackNetworkOutput)
}

// IpLoadBalancingVrackNetworkArrayInput is an input type that accepts IpLoadBalancingVrackNetworkArray and IpLoadBalancingVrackNetworkArrayOutput values.
// You can construct a concrete instance of `IpLoadBalancingVrackNetworkArrayInput` via:
//
//	IpLoadBalancingVrackNetworkArray{ IpLoadBalancingVrackNetworkArgs{...} }
type IpLoadBalancingVrackNetworkArrayInput interface {
	pulumi.Input

	ToIpLoadBalancingVrackNetworkArrayOutput() IpLoadBalancingVrackNetworkArrayOutput
	ToIpLoadBalancingVrackNetworkArrayOutputWithContext(context.Context) IpLoadBalancingVrackNetworkArrayOutput
}

type IpLoadBalancingVrackNetworkArray []IpLoadBalancingVrackNetworkInput

func (IpLoadBalancingVrackNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingVrackNetwork)(nil)).Elem()
}

func (i IpLoadBalancingVrackNetworkArray) ToIpLoadBalancingVrackNetworkArrayOutput() IpLoadBalancingVrackNetworkArrayOutput {
	return i.ToIpLoadBalancingVrackNetworkArrayOutputWithContext(context.Background())
}

func (i IpLoadBalancingVrackNetworkArray) ToIpLoadBalancingVrackNetworkArrayOutputWithContext(ctx context.Context) IpLoadBalancingVrackNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingVrackNetworkArrayOutput)
}

// IpLoadBalancingVrackNetworkMapInput is an input type that accepts IpLoadBalancingVrackNetworkMap and IpLoadBalancingVrackNetworkMapOutput values.
// You can construct a concrete instance of `IpLoadBalancingVrackNetworkMapInput` via:
//
//	IpLoadBalancingVrackNetworkMap{ "key": IpLoadBalancingVrackNetworkArgs{...} }
type IpLoadBalancingVrackNetworkMapInput interface {
	pulumi.Input

	ToIpLoadBalancingVrackNetworkMapOutput() IpLoadBalancingVrackNetworkMapOutput
	ToIpLoadBalancingVrackNetworkMapOutputWithContext(context.Context) IpLoadBalancingVrackNetworkMapOutput
}

type IpLoadBalancingVrackNetworkMap map[string]IpLoadBalancingVrackNetworkInput

func (IpLoadBalancingVrackNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingVrackNetwork)(nil)).Elem()
}

func (i IpLoadBalancingVrackNetworkMap) ToIpLoadBalancingVrackNetworkMapOutput() IpLoadBalancingVrackNetworkMapOutput {
	return i.ToIpLoadBalancingVrackNetworkMapOutputWithContext(context.Background())
}

func (i IpLoadBalancingVrackNetworkMap) ToIpLoadBalancingVrackNetworkMapOutputWithContext(ctx context.Context) IpLoadBalancingVrackNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingVrackNetworkMapOutput)
}

type IpLoadBalancingVrackNetworkOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingVrackNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingVrackNetwork)(nil)).Elem()
}

func (o IpLoadBalancingVrackNetworkOutput) ToIpLoadBalancingVrackNetworkOutput() IpLoadBalancingVrackNetworkOutput {
	return o
}

func (o IpLoadBalancingVrackNetworkOutput) ToIpLoadBalancingVrackNetworkOutputWithContext(ctx context.Context) IpLoadBalancingVrackNetworkOutput {
	return o
}

// Human readable name for your vrack network
func (o IpLoadBalancingVrackNetworkOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingVrackNetwork) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
func (o IpLoadBalancingVrackNetworkOutput) FarmIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *IpLoadBalancingVrackNetwork) pulumi.IntArrayOutput { return v.FarmIds }).(pulumi.IntArrayOutput)
}

// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
func (o IpLoadBalancingVrackNetworkOutput) NatIp() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingVrackNetwork) pulumi.StringOutput { return v.NatIp }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o IpLoadBalancingVrackNetworkOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingVrackNetwork) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// IP block of the private network in the vRack
func (o IpLoadBalancingVrackNetworkOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingVrackNetwork) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
func (o IpLoadBalancingVrackNetworkOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v *IpLoadBalancingVrackNetwork) pulumi.IntOutput { return v.Vlan }).(pulumi.IntOutput)
}

// (Required) Internal Load Balancer identifier of the vRack private network
func (o IpLoadBalancingVrackNetworkOutput) VrackNetworkId() pulumi.IntOutput {
	return o.ApplyT(func(v *IpLoadBalancingVrackNetwork) pulumi.IntOutput { return v.VrackNetworkId }).(pulumi.IntOutput)
}

type IpLoadBalancingVrackNetworkArrayOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingVrackNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingVrackNetwork)(nil)).Elem()
}

func (o IpLoadBalancingVrackNetworkArrayOutput) ToIpLoadBalancingVrackNetworkArrayOutput() IpLoadBalancingVrackNetworkArrayOutput {
	return o
}

func (o IpLoadBalancingVrackNetworkArrayOutput) ToIpLoadBalancingVrackNetworkArrayOutputWithContext(ctx context.Context) IpLoadBalancingVrackNetworkArrayOutput {
	return o
}

func (o IpLoadBalancingVrackNetworkArrayOutput) Index(i pulumi.IntInput) IpLoadBalancingVrackNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpLoadBalancingVrackNetwork {
		return vs[0].([]*IpLoadBalancingVrackNetwork)[vs[1].(int)]
	}).(IpLoadBalancingVrackNetworkOutput)
}

type IpLoadBalancingVrackNetworkMapOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingVrackNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingVrackNetwork)(nil)).Elem()
}

func (o IpLoadBalancingVrackNetworkMapOutput) ToIpLoadBalancingVrackNetworkMapOutput() IpLoadBalancingVrackNetworkMapOutput {
	return o
}

func (o IpLoadBalancingVrackNetworkMapOutput) ToIpLoadBalancingVrackNetworkMapOutputWithContext(ctx context.Context) IpLoadBalancingVrackNetworkMapOutput {
	return o
}

func (o IpLoadBalancingVrackNetworkMapOutput) MapIndex(k pulumi.StringInput) IpLoadBalancingVrackNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpLoadBalancingVrackNetwork {
		return vs[0].(map[string]*IpLoadBalancingVrackNetwork)[vs[1].(string)]
	}).(IpLoadBalancingVrackNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingVrackNetworkInput)(nil)).Elem(), &IpLoadBalancingVrackNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingVrackNetworkArrayInput)(nil)).Elem(), IpLoadBalancingVrackNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingVrackNetworkMapInput)(nil)).Elem(), IpLoadBalancingVrackNetworkMap{})
	pulumi.RegisterOutputType(IpLoadBalancingVrackNetworkOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingVrackNetworkArrayOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingVrackNetworkMapOutput{})
}
