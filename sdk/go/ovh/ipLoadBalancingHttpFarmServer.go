// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backend server entry linked to http loadbalancing group (farm)
type IpLoadBalancingHttpFarmServer struct {
	pulumi.CustomResourceState

	// Address of the backend server (IP from either internal or OVH network)
	Address pulumi.StringOutput `pulumi:"address"`
	// is it a backup server used in case of failure of all the non-backup backends
	Backup pulumi.BoolPtrOutput   `pulumi:"backup"`
	Chain  pulumi.StringPtrOutput `pulumi:"chain"`
	// Value of the stickiness cookie used for this backend.
	Cookie pulumi.StringOutput `pulumi:"cookie"`
	// Label for the server
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// ID of the farm this server is attached to
	FarmId pulumi.IntOutput `pulumi:"farmId"`
	// Port that backend will respond on
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe pulumi.BoolPtrOutput `pulumi:"probe"`
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to recieving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion pulumi.StringPtrOutput `pulumi:"proxyProtocolVersion"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// is the connection ciphered with SSL (TLS)
	Ssl pulumi.BoolPtrOutput `pulumi:"ssl"`
	// backend status - `active` or `inactive`
	Status pulumi.StringOutput `pulumi:"status"`
	// used in loadbalancing algorithm
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewIpLoadBalancingHttpFarmServer registers a new resource with the given unique name, arguments, and options.
func NewIpLoadBalancingHttpFarmServer(ctx *pulumi.Context,
	name string, args *IpLoadBalancingHttpFarmServerArgs, opts ...pulumi.ResourceOption) (*IpLoadBalancingHttpFarmServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.FarmId == nil {
		return nil, errors.New("invalid value for required argument 'FarmId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpLoadBalancingHttpFarmServer
	err := ctx.RegisterResource("ovh:index/ipLoadBalancingHttpFarmServer:IpLoadBalancingHttpFarmServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpLoadBalancingHttpFarmServer gets an existing IpLoadBalancingHttpFarmServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpLoadBalancingHttpFarmServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpLoadBalancingHttpFarmServerState, opts ...pulumi.ResourceOption) (*IpLoadBalancingHttpFarmServer, error) {
	var resource IpLoadBalancingHttpFarmServer
	err := ctx.ReadResource("ovh:index/ipLoadBalancingHttpFarmServer:IpLoadBalancingHttpFarmServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpLoadBalancingHttpFarmServer resources.
type ipLoadBalancingHttpFarmServerState struct {
	// Address of the backend server (IP from either internal or OVH network)
	Address *string `pulumi:"address"`
	// is it a backup server used in case of failure of all the non-backup backends
	Backup *bool   `pulumi:"backup"`
	Chain  *string `pulumi:"chain"`
	// Value of the stickiness cookie used for this backend.
	Cookie *string `pulumi:"cookie"`
	// Label for the server
	DisplayName *string `pulumi:"displayName"`
	// ID of the farm this server is attached to
	FarmId *int `pulumi:"farmId"`
	// Port that backend will respond on
	Port *int `pulumi:"port"`
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe *bool `pulumi:"probe"`
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to recieving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion *string `pulumi:"proxyProtocolVersion"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// is the connection ciphered with SSL (TLS)
	Ssl *bool `pulumi:"ssl"`
	// backend status - `active` or `inactive`
	Status *string `pulumi:"status"`
	// used in loadbalancing algorithm
	Weight *int `pulumi:"weight"`
}

type IpLoadBalancingHttpFarmServerState struct {
	// Address of the backend server (IP from either internal or OVH network)
	Address pulumi.StringPtrInput
	// is it a backup server used in case of failure of all the non-backup backends
	Backup pulumi.BoolPtrInput
	Chain  pulumi.StringPtrInput
	// Value of the stickiness cookie used for this backend.
	Cookie pulumi.StringPtrInput
	// Label for the server
	DisplayName pulumi.StringPtrInput
	// ID of the farm this server is attached to
	FarmId pulumi.IntPtrInput
	// Port that backend will respond on
	Port pulumi.IntPtrInput
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe pulumi.BoolPtrInput
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to recieving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// is the connection ciphered with SSL (TLS)
	Ssl pulumi.BoolPtrInput
	// backend status - `active` or `inactive`
	Status pulumi.StringPtrInput
	// used in loadbalancing algorithm
	Weight pulumi.IntPtrInput
}

func (IpLoadBalancingHttpFarmServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingHttpFarmServerState)(nil)).Elem()
}

type ipLoadBalancingHttpFarmServerArgs struct {
	// Address of the backend server (IP from either internal or OVH network)
	Address string `pulumi:"address"`
	// is it a backup server used in case of failure of all the non-backup backends
	Backup *bool   `pulumi:"backup"`
	Chain  *string `pulumi:"chain"`
	// Label for the server
	DisplayName *string `pulumi:"displayName"`
	// ID of the farm this server is attached to
	FarmId int `pulumi:"farmId"`
	// Port that backend will respond on
	Port *int `pulumi:"port"`
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe *bool `pulumi:"probe"`
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to recieving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion *string `pulumi:"proxyProtocolVersion"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// is the connection ciphered with SSL (TLS)
	Ssl *bool `pulumi:"ssl"`
	// backend status - `active` or `inactive`
	Status string `pulumi:"status"`
	// used in loadbalancing algorithm
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a IpLoadBalancingHttpFarmServer resource.
type IpLoadBalancingHttpFarmServerArgs struct {
	// Address of the backend server (IP from either internal or OVH network)
	Address pulumi.StringInput
	// is it a backup server used in case of failure of all the non-backup backends
	Backup pulumi.BoolPtrInput
	Chain  pulumi.StringPtrInput
	// Label for the server
	DisplayName pulumi.StringPtrInput
	// ID of the farm this server is attached to
	FarmId pulumi.IntInput
	// Port that backend will respond on
	Port pulumi.IntPtrInput
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe pulumi.BoolPtrInput
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to recieving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// is the connection ciphered with SSL (TLS)
	Ssl pulumi.BoolPtrInput
	// backend status - `active` or `inactive`
	Status pulumi.StringInput
	// used in loadbalancing algorithm
	Weight pulumi.IntPtrInput
}

func (IpLoadBalancingHttpFarmServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingHttpFarmServerArgs)(nil)).Elem()
}

type IpLoadBalancingHttpFarmServerInput interface {
	pulumi.Input

	ToIpLoadBalancingHttpFarmServerOutput() IpLoadBalancingHttpFarmServerOutput
	ToIpLoadBalancingHttpFarmServerOutputWithContext(ctx context.Context) IpLoadBalancingHttpFarmServerOutput
}

func (*IpLoadBalancingHttpFarmServer) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingHttpFarmServer)(nil)).Elem()
}

func (i *IpLoadBalancingHttpFarmServer) ToIpLoadBalancingHttpFarmServerOutput() IpLoadBalancingHttpFarmServerOutput {
	return i.ToIpLoadBalancingHttpFarmServerOutputWithContext(context.Background())
}

func (i *IpLoadBalancingHttpFarmServer) ToIpLoadBalancingHttpFarmServerOutputWithContext(ctx context.Context) IpLoadBalancingHttpFarmServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingHttpFarmServerOutput)
}

// IpLoadBalancingHttpFarmServerArrayInput is an input type that accepts IpLoadBalancingHttpFarmServerArray and IpLoadBalancingHttpFarmServerArrayOutput values.
// You can construct a concrete instance of `IpLoadBalancingHttpFarmServerArrayInput` via:
//
//	IpLoadBalancingHttpFarmServerArray{ IpLoadBalancingHttpFarmServerArgs{...} }
type IpLoadBalancingHttpFarmServerArrayInput interface {
	pulumi.Input

	ToIpLoadBalancingHttpFarmServerArrayOutput() IpLoadBalancingHttpFarmServerArrayOutput
	ToIpLoadBalancingHttpFarmServerArrayOutputWithContext(context.Context) IpLoadBalancingHttpFarmServerArrayOutput
}

type IpLoadBalancingHttpFarmServerArray []IpLoadBalancingHttpFarmServerInput

func (IpLoadBalancingHttpFarmServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingHttpFarmServer)(nil)).Elem()
}

func (i IpLoadBalancingHttpFarmServerArray) ToIpLoadBalancingHttpFarmServerArrayOutput() IpLoadBalancingHttpFarmServerArrayOutput {
	return i.ToIpLoadBalancingHttpFarmServerArrayOutputWithContext(context.Background())
}

func (i IpLoadBalancingHttpFarmServerArray) ToIpLoadBalancingHttpFarmServerArrayOutputWithContext(ctx context.Context) IpLoadBalancingHttpFarmServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingHttpFarmServerArrayOutput)
}

// IpLoadBalancingHttpFarmServerMapInput is an input type that accepts IpLoadBalancingHttpFarmServerMap and IpLoadBalancingHttpFarmServerMapOutput values.
// You can construct a concrete instance of `IpLoadBalancingHttpFarmServerMapInput` via:
//
//	IpLoadBalancingHttpFarmServerMap{ "key": IpLoadBalancingHttpFarmServerArgs{...} }
type IpLoadBalancingHttpFarmServerMapInput interface {
	pulumi.Input

	ToIpLoadBalancingHttpFarmServerMapOutput() IpLoadBalancingHttpFarmServerMapOutput
	ToIpLoadBalancingHttpFarmServerMapOutputWithContext(context.Context) IpLoadBalancingHttpFarmServerMapOutput
}

type IpLoadBalancingHttpFarmServerMap map[string]IpLoadBalancingHttpFarmServerInput

func (IpLoadBalancingHttpFarmServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingHttpFarmServer)(nil)).Elem()
}

func (i IpLoadBalancingHttpFarmServerMap) ToIpLoadBalancingHttpFarmServerMapOutput() IpLoadBalancingHttpFarmServerMapOutput {
	return i.ToIpLoadBalancingHttpFarmServerMapOutputWithContext(context.Background())
}

func (i IpLoadBalancingHttpFarmServerMap) ToIpLoadBalancingHttpFarmServerMapOutputWithContext(ctx context.Context) IpLoadBalancingHttpFarmServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingHttpFarmServerMapOutput)
}

type IpLoadBalancingHttpFarmServerOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingHttpFarmServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingHttpFarmServer)(nil)).Elem()
}

func (o IpLoadBalancingHttpFarmServerOutput) ToIpLoadBalancingHttpFarmServerOutput() IpLoadBalancingHttpFarmServerOutput {
	return o
}

func (o IpLoadBalancingHttpFarmServerOutput) ToIpLoadBalancingHttpFarmServerOutputWithContext(ctx context.Context) IpLoadBalancingHttpFarmServerOutput {
	return o
}

// Address of the backend server (IP from either internal or OVH network)
func (o IpLoadBalancingHttpFarmServerOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// is it a backup server used in case of failure of all the non-backup backends
func (o IpLoadBalancingHttpFarmServerOutput) Backup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.BoolPtrOutput { return v.Backup }).(pulumi.BoolPtrOutput)
}

func (o IpLoadBalancingHttpFarmServerOutput) Chain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.StringPtrOutput { return v.Chain }).(pulumi.StringPtrOutput)
}

// Value of the stickiness cookie used for this backend.
func (o IpLoadBalancingHttpFarmServerOutput) Cookie() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.StringOutput { return v.Cookie }).(pulumi.StringOutput)
}

// Label for the server
func (o IpLoadBalancingHttpFarmServerOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// ID of the farm this server is attached to
func (o IpLoadBalancingHttpFarmServerOutput) FarmId() pulumi.IntOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.IntOutput { return v.FarmId }).(pulumi.IntOutput)
}

// Port that backend will respond on
func (o IpLoadBalancingHttpFarmServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// defines if backend will be probed to determine health and keep as active in farm if healthy
func (o IpLoadBalancingHttpFarmServerOutput) Probe() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.BoolPtrOutput { return v.Probe }).(pulumi.BoolPtrOutput)
}

// version of the PROXY protocol used to pass origin connection information from loadbalancer to recieving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
func (o IpLoadBalancingHttpFarmServerOutput) ProxyProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.StringPtrOutput { return v.ProxyProtocolVersion }).(pulumi.StringPtrOutput)
}

// The internal name of your IP load balancing
func (o IpLoadBalancingHttpFarmServerOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// is the connection ciphered with SSL (TLS)
func (o IpLoadBalancingHttpFarmServerOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

// backend status - `active` or `inactive`
func (o IpLoadBalancingHttpFarmServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// used in loadbalancing algorithm
func (o IpLoadBalancingHttpFarmServerOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingHttpFarmServer) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

type IpLoadBalancingHttpFarmServerArrayOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingHttpFarmServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingHttpFarmServer)(nil)).Elem()
}

func (o IpLoadBalancingHttpFarmServerArrayOutput) ToIpLoadBalancingHttpFarmServerArrayOutput() IpLoadBalancingHttpFarmServerArrayOutput {
	return o
}

func (o IpLoadBalancingHttpFarmServerArrayOutput) ToIpLoadBalancingHttpFarmServerArrayOutputWithContext(ctx context.Context) IpLoadBalancingHttpFarmServerArrayOutput {
	return o
}

func (o IpLoadBalancingHttpFarmServerArrayOutput) Index(i pulumi.IntInput) IpLoadBalancingHttpFarmServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpLoadBalancingHttpFarmServer {
		return vs[0].([]*IpLoadBalancingHttpFarmServer)[vs[1].(int)]
	}).(IpLoadBalancingHttpFarmServerOutput)
}

type IpLoadBalancingHttpFarmServerMapOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingHttpFarmServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingHttpFarmServer)(nil)).Elem()
}

func (o IpLoadBalancingHttpFarmServerMapOutput) ToIpLoadBalancingHttpFarmServerMapOutput() IpLoadBalancingHttpFarmServerMapOutput {
	return o
}

func (o IpLoadBalancingHttpFarmServerMapOutput) ToIpLoadBalancingHttpFarmServerMapOutputWithContext(ctx context.Context) IpLoadBalancingHttpFarmServerMapOutput {
	return o
}

func (o IpLoadBalancingHttpFarmServerMapOutput) MapIndex(k pulumi.StringInput) IpLoadBalancingHttpFarmServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpLoadBalancingHttpFarmServer {
		return vs[0].(map[string]*IpLoadBalancingHttpFarmServer)[vs[1].(string)]
	}).(IpLoadBalancingHttpFarmServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingHttpFarmServerInput)(nil)).Elem(), &IpLoadBalancingHttpFarmServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingHttpFarmServerArrayInput)(nil)).Elem(), IpLoadBalancingHttpFarmServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingHttpFarmServerMapInput)(nil)).Elem(), IpLoadBalancingHttpFarmServerMap{})
	pulumi.RegisterOutputType(IpLoadBalancingHttpFarmServerOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingHttpFarmServerArrayOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingHttpFarmServerMapOutput{})
}
