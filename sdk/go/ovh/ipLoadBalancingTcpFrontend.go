// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backend server group (frontend) to be used by loadbalancing frontend(s)
type IpLoadBalancingTcpFrontend struct {
	pulumi.CustomResourceState

	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources pulumi.StringArrayOutput `pulumi:"allowedSources"`
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayOutput `pulumi:"dedicatedIpfos"`
	// Default TCP Farm of your frontend
	DefaultFarmId pulumi.IntOutput `pulumi:"defaultFarmId"`
	// Default ssl served to your customer
	DefaultSslId pulumi.IntOutput `pulumi:"defaultSslId"`
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Human readable name for your frontend, this field is for you
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringOutput `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrOutput `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewIpLoadBalancingTcpFrontend registers a new resource with the given unique name, arguments, and options.
func NewIpLoadBalancingTcpFrontend(ctx *pulumi.Context,
	name string, args *IpLoadBalancingTcpFrontendArgs, opts ...pulumi.ResourceOption) (*IpLoadBalancingTcpFrontend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpLoadBalancingTcpFrontend
	err := ctx.RegisterResource("ovh:index/ipLoadBalancingTcpFrontend:IpLoadBalancingTcpFrontend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpLoadBalancingTcpFrontend gets an existing IpLoadBalancingTcpFrontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpLoadBalancingTcpFrontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpLoadBalancingTcpFrontendState, opts ...pulumi.ResourceOption) (*IpLoadBalancingTcpFrontend, error) {
	var resource IpLoadBalancingTcpFrontend
	err := ctx.ReadResource("ovh:index/ipLoadBalancingTcpFrontend:IpLoadBalancingTcpFrontend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpLoadBalancingTcpFrontend resources.
type ipLoadBalancingTcpFrontendState struct {
	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources []string `pulumi:"allowedSources"`
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	// Default TCP Farm of your frontend
	DefaultFarmId *int `pulumi:"defaultFarmId"`
	// Default ssl served to your customer
	DefaultSslId *int `pulumi:"defaultSslId"`
	// Disable your frontend. Default: 'false'
	Disabled *bool `pulumi:"disabled"`
	// Human readable name for your frontend, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port *string `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl *bool `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone *string `pulumi:"zone"`
}

type IpLoadBalancingTcpFrontendState struct {
	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources pulumi.StringArrayInput
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayInput
	// Default TCP Farm of your frontend
	DefaultFarmId pulumi.IntPtrInput
	// Default ssl served to your customer
	DefaultSslId pulumi.IntPtrInput
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolPtrInput
	// Human readable name for your frontend, this field is for you
	DisplayName pulumi.StringPtrInput
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrInput
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringPtrInput
}

func (IpLoadBalancingTcpFrontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingTcpFrontendState)(nil)).Elem()
}

type ipLoadBalancingTcpFrontendArgs struct {
	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources []string `pulumi:"allowedSources"`
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	// Default TCP Farm of your frontend
	DefaultFarmId *int `pulumi:"defaultFarmId"`
	// Default ssl served to your customer
	DefaultSslId *int `pulumi:"defaultSslId"`
	// Disable your frontend. Default: 'false'
	Disabled *bool `pulumi:"disabled"`
	// Human readable name for your frontend, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port string `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl *bool `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a IpLoadBalancingTcpFrontend resource.
type IpLoadBalancingTcpFrontendArgs struct {
	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources pulumi.StringArrayInput
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayInput
	// Default TCP Farm of your frontend
	DefaultFarmId pulumi.IntPtrInput
	// Default ssl served to your customer
	DefaultSslId pulumi.IntPtrInput
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolPtrInput
	// Human readable name for your frontend, this field is for you
	DisplayName pulumi.StringPtrInput
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrInput
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringInput
}

func (IpLoadBalancingTcpFrontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingTcpFrontendArgs)(nil)).Elem()
}

type IpLoadBalancingTcpFrontendInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpFrontendOutput() IpLoadBalancingTcpFrontendOutput
	ToIpLoadBalancingTcpFrontendOutputWithContext(ctx context.Context) IpLoadBalancingTcpFrontendOutput
}

func (*IpLoadBalancingTcpFrontend) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingTcpFrontend)(nil)).Elem()
}

func (i *IpLoadBalancingTcpFrontend) ToIpLoadBalancingTcpFrontendOutput() IpLoadBalancingTcpFrontendOutput {
	return i.ToIpLoadBalancingTcpFrontendOutputWithContext(context.Background())
}

func (i *IpLoadBalancingTcpFrontend) ToIpLoadBalancingTcpFrontendOutputWithContext(ctx context.Context) IpLoadBalancingTcpFrontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpFrontendOutput)
}

// IpLoadBalancingTcpFrontendArrayInput is an input type that accepts IpLoadBalancingTcpFrontendArray and IpLoadBalancingTcpFrontendArrayOutput values.
// You can construct a concrete instance of `IpLoadBalancingTcpFrontendArrayInput` via:
//
//	IpLoadBalancingTcpFrontendArray{ IpLoadBalancingTcpFrontendArgs{...} }
type IpLoadBalancingTcpFrontendArrayInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpFrontendArrayOutput() IpLoadBalancingTcpFrontendArrayOutput
	ToIpLoadBalancingTcpFrontendArrayOutputWithContext(context.Context) IpLoadBalancingTcpFrontendArrayOutput
}

type IpLoadBalancingTcpFrontendArray []IpLoadBalancingTcpFrontendInput

func (IpLoadBalancingTcpFrontendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingTcpFrontend)(nil)).Elem()
}

func (i IpLoadBalancingTcpFrontendArray) ToIpLoadBalancingTcpFrontendArrayOutput() IpLoadBalancingTcpFrontendArrayOutput {
	return i.ToIpLoadBalancingTcpFrontendArrayOutputWithContext(context.Background())
}

func (i IpLoadBalancingTcpFrontendArray) ToIpLoadBalancingTcpFrontendArrayOutputWithContext(ctx context.Context) IpLoadBalancingTcpFrontendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpFrontendArrayOutput)
}

// IpLoadBalancingTcpFrontendMapInput is an input type that accepts IpLoadBalancingTcpFrontendMap and IpLoadBalancingTcpFrontendMapOutput values.
// You can construct a concrete instance of `IpLoadBalancingTcpFrontendMapInput` via:
//
//	IpLoadBalancingTcpFrontendMap{ "key": IpLoadBalancingTcpFrontendArgs{...} }
type IpLoadBalancingTcpFrontendMapInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpFrontendMapOutput() IpLoadBalancingTcpFrontendMapOutput
	ToIpLoadBalancingTcpFrontendMapOutputWithContext(context.Context) IpLoadBalancingTcpFrontendMapOutput
}

type IpLoadBalancingTcpFrontendMap map[string]IpLoadBalancingTcpFrontendInput

func (IpLoadBalancingTcpFrontendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingTcpFrontend)(nil)).Elem()
}

func (i IpLoadBalancingTcpFrontendMap) ToIpLoadBalancingTcpFrontendMapOutput() IpLoadBalancingTcpFrontendMapOutput {
	return i.ToIpLoadBalancingTcpFrontendMapOutputWithContext(context.Background())
}

func (i IpLoadBalancingTcpFrontendMap) ToIpLoadBalancingTcpFrontendMapOutputWithContext(ctx context.Context) IpLoadBalancingTcpFrontendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpFrontendMapOutput)
}

type IpLoadBalancingTcpFrontendOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpFrontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingTcpFrontend)(nil)).Elem()
}

func (o IpLoadBalancingTcpFrontendOutput) ToIpLoadBalancingTcpFrontendOutput() IpLoadBalancingTcpFrontendOutput {
	return o
}

func (o IpLoadBalancingTcpFrontendOutput) ToIpLoadBalancingTcpFrontendOutputWithContext(ctx context.Context) IpLoadBalancingTcpFrontendOutput {
	return o
}

// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
func (o IpLoadBalancingTcpFrontendOutput) AllowedSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.StringArrayOutput { return v.AllowedSources }).(pulumi.StringArrayOutput)
}

// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
func (o IpLoadBalancingTcpFrontendOutput) DedicatedIpfos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.StringArrayOutput { return v.DedicatedIpfos }).(pulumi.StringArrayOutput)
}

// Default TCP Farm of your frontend
func (o IpLoadBalancingTcpFrontendOutput) DefaultFarmId() pulumi.IntOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.IntOutput { return v.DefaultFarmId }).(pulumi.IntOutput)
}

// Default ssl served to your customer
func (o IpLoadBalancingTcpFrontendOutput) DefaultSslId() pulumi.IntOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.IntOutput { return v.DefaultSslId }).(pulumi.IntOutput)
}

// Disable your frontend. Default: 'false'
func (o IpLoadBalancingTcpFrontendOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Human readable name for your frontend, this field is for you
func (o IpLoadBalancingTcpFrontendOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Port(s) attached to your frontend. Supports single port (numerical value),
// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
// and/or 'range'. Each port must be in the [1;49151] range
func (o IpLoadBalancingTcpFrontendOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o IpLoadBalancingTcpFrontendOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// SSL deciphering. Default: 'false'
func (o IpLoadBalancingTcpFrontendOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
func (o IpLoadBalancingTcpFrontendOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFrontend) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type IpLoadBalancingTcpFrontendArrayOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpFrontendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingTcpFrontend)(nil)).Elem()
}

func (o IpLoadBalancingTcpFrontendArrayOutput) ToIpLoadBalancingTcpFrontendArrayOutput() IpLoadBalancingTcpFrontendArrayOutput {
	return o
}

func (o IpLoadBalancingTcpFrontendArrayOutput) ToIpLoadBalancingTcpFrontendArrayOutputWithContext(ctx context.Context) IpLoadBalancingTcpFrontendArrayOutput {
	return o
}

func (o IpLoadBalancingTcpFrontendArrayOutput) Index(i pulumi.IntInput) IpLoadBalancingTcpFrontendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpLoadBalancingTcpFrontend {
		return vs[0].([]*IpLoadBalancingTcpFrontend)[vs[1].(int)]
	}).(IpLoadBalancingTcpFrontendOutput)
}

type IpLoadBalancingTcpFrontendMapOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpFrontendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingTcpFrontend)(nil)).Elem()
}

func (o IpLoadBalancingTcpFrontendMapOutput) ToIpLoadBalancingTcpFrontendMapOutput() IpLoadBalancingTcpFrontendMapOutput {
	return o
}

func (o IpLoadBalancingTcpFrontendMapOutput) ToIpLoadBalancingTcpFrontendMapOutputWithContext(ctx context.Context) IpLoadBalancingTcpFrontendMapOutput {
	return o
}

func (o IpLoadBalancingTcpFrontendMapOutput) MapIndex(k pulumi.StringInput) IpLoadBalancingTcpFrontendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpLoadBalancingTcpFrontend {
		return vs[0].(map[string]*IpLoadBalancingTcpFrontend)[vs[1].(string)]
	}).(IpLoadBalancingTcpFrontendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpFrontendInput)(nil)).Elem(), &IpLoadBalancingTcpFrontend{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpFrontendArrayInput)(nil)).Elem(), IpLoadBalancingTcpFrontendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpFrontendMapInput)(nil)).Elem(), IpLoadBalancingTcpFrontendMap{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpFrontendOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpFrontendArrayOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpFrontendMapOutput{})
}
