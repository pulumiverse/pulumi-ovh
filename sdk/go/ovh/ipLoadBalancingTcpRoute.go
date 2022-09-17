// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage tcp route for a loadbalancer service
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ovh.NewIpLoadBalancingTcpRoute(ctx, "tcpreject", &ovh.IpLoadBalancingTcpRouteArgs{
//				Action: &IpLoadBalancingTcpRouteActionArgs{
//					Type: pulumi.String("reject"),
//				},
//				ServiceName: pulumi.String("loadbalancer-xxxxxxxxxxxxxxxxxx"),
//				Weight:      pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpLoadBalancingTcpRoute struct {
	pulumi.CustomResourceState

	// Action triggered when all rules match
	Action IpLoadBalancingTcpRouteActionOutput `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId pulumi.IntOutput `pulumi:"frontendId"`
	// List of rules to match to trigger action
	Rules IpLoadBalancingTcpRouteRuleTypeArrayOutput `pulumi:"rules"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Route status. Routes in "ok" state are ready to operate
	Status pulumi.StringOutput `pulumi:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewIpLoadBalancingTcpRoute registers a new resource with the given unique name, arguments, and options.
func NewIpLoadBalancingTcpRoute(ctx *pulumi.Context,
	name string, args *IpLoadBalancingTcpRouteArgs, opts ...pulumi.ResourceOption) (*IpLoadBalancingTcpRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpLoadBalancingTcpRoute
	err := ctx.RegisterResource("ovh:index/ipLoadBalancingTcpRoute:IpLoadBalancingTcpRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpLoadBalancingTcpRoute gets an existing IpLoadBalancingTcpRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpLoadBalancingTcpRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpLoadBalancingTcpRouteState, opts ...pulumi.ResourceOption) (*IpLoadBalancingTcpRoute, error) {
	var resource IpLoadBalancingTcpRoute
	err := ctx.ReadResource("ovh:index/ipLoadBalancingTcpRoute:IpLoadBalancingTcpRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpLoadBalancingTcpRoute resources.
type ipLoadBalancingTcpRouteState struct {
	// Action triggered when all rules match
	Action *IpLoadBalancingTcpRouteAction `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId *int `pulumi:"frontendId"`
	// List of rules to match to trigger action
	Rules []IpLoadBalancingTcpRouteRuleType `pulumi:"rules"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Route status. Routes in "ok" state are ready to operate
	Status *string `pulumi:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight *int `pulumi:"weight"`
}

type IpLoadBalancingTcpRouteState struct {
	// Action triggered when all rules match
	Action IpLoadBalancingTcpRouteActionPtrInput
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrInput
	// Route traffic for this frontend
	FrontendId pulumi.IntPtrInput
	// List of rules to match to trigger action
	Rules IpLoadBalancingTcpRouteRuleTypeArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Route status. Routes in "ok" state are ready to operate
	Status pulumi.StringPtrInput
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight pulumi.IntPtrInput
}

func (IpLoadBalancingTcpRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingTcpRouteState)(nil)).Elem()
}

type ipLoadBalancingTcpRouteArgs struct {
	// Action triggered when all rules match
	Action IpLoadBalancingTcpRouteAction `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId *int `pulumi:"frontendId"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a IpLoadBalancingTcpRoute resource.
type IpLoadBalancingTcpRouteArgs struct {
	// Action triggered when all rules match
	Action IpLoadBalancingTcpRouteActionInput
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrInput
	// Route traffic for this frontend
	FrontendId pulumi.IntPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight pulumi.IntPtrInput
}

func (IpLoadBalancingTcpRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingTcpRouteArgs)(nil)).Elem()
}

type IpLoadBalancingTcpRouteInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpRouteOutput() IpLoadBalancingTcpRouteOutput
	ToIpLoadBalancingTcpRouteOutputWithContext(ctx context.Context) IpLoadBalancingTcpRouteOutput
}

func (*IpLoadBalancingTcpRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingTcpRoute)(nil)).Elem()
}

func (i *IpLoadBalancingTcpRoute) ToIpLoadBalancingTcpRouteOutput() IpLoadBalancingTcpRouteOutput {
	return i.ToIpLoadBalancingTcpRouteOutputWithContext(context.Background())
}

func (i *IpLoadBalancingTcpRoute) ToIpLoadBalancingTcpRouteOutputWithContext(ctx context.Context) IpLoadBalancingTcpRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpRouteOutput)
}

// IpLoadBalancingTcpRouteArrayInput is an input type that accepts IpLoadBalancingTcpRouteArray and IpLoadBalancingTcpRouteArrayOutput values.
// You can construct a concrete instance of `IpLoadBalancingTcpRouteArrayInput` via:
//
//	IpLoadBalancingTcpRouteArray{ IpLoadBalancingTcpRouteArgs{...} }
type IpLoadBalancingTcpRouteArrayInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpRouteArrayOutput() IpLoadBalancingTcpRouteArrayOutput
	ToIpLoadBalancingTcpRouteArrayOutputWithContext(context.Context) IpLoadBalancingTcpRouteArrayOutput
}

type IpLoadBalancingTcpRouteArray []IpLoadBalancingTcpRouteInput

func (IpLoadBalancingTcpRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingTcpRoute)(nil)).Elem()
}

func (i IpLoadBalancingTcpRouteArray) ToIpLoadBalancingTcpRouteArrayOutput() IpLoadBalancingTcpRouteArrayOutput {
	return i.ToIpLoadBalancingTcpRouteArrayOutputWithContext(context.Background())
}

func (i IpLoadBalancingTcpRouteArray) ToIpLoadBalancingTcpRouteArrayOutputWithContext(ctx context.Context) IpLoadBalancingTcpRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpRouteArrayOutput)
}

// IpLoadBalancingTcpRouteMapInput is an input type that accepts IpLoadBalancingTcpRouteMap and IpLoadBalancingTcpRouteMapOutput values.
// You can construct a concrete instance of `IpLoadBalancingTcpRouteMapInput` via:
//
//	IpLoadBalancingTcpRouteMap{ "key": IpLoadBalancingTcpRouteArgs{...} }
type IpLoadBalancingTcpRouteMapInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpRouteMapOutput() IpLoadBalancingTcpRouteMapOutput
	ToIpLoadBalancingTcpRouteMapOutputWithContext(context.Context) IpLoadBalancingTcpRouteMapOutput
}

type IpLoadBalancingTcpRouteMap map[string]IpLoadBalancingTcpRouteInput

func (IpLoadBalancingTcpRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingTcpRoute)(nil)).Elem()
}

func (i IpLoadBalancingTcpRouteMap) ToIpLoadBalancingTcpRouteMapOutput() IpLoadBalancingTcpRouteMapOutput {
	return i.ToIpLoadBalancingTcpRouteMapOutputWithContext(context.Background())
}

func (i IpLoadBalancingTcpRouteMap) ToIpLoadBalancingTcpRouteMapOutputWithContext(ctx context.Context) IpLoadBalancingTcpRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpRouteMapOutput)
}

type IpLoadBalancingTcpRouteOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingTcpRoute)(nil)).Elem()
}

func (o IpLoadBalancingTcpRouteOutput) ToIpLoadBalancingTcpRouteOutput() IpLoadBalancingTcpRouteOutput {
	return o
}

func (o IpLoadBalancingTcpRouteOutput) ToIpLoadBalancingTcpRouteOutputWithContext(ctx context.Context) IpLoadBalancingTcpRouteOutput {
	return o
}

// Action triggered when all rules match
func (o IpLoadBalancingTcpRouteOutput) Action() IpLoadBalancingTcpRouteActionOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpRoute) IpLoadBalancingTcpRouteActionOutput { return v.Action }).(IpLoadBalancingTcpRouteActionOutput)
}

// Human readable name for your route, this field is for you
func (o IpLoadBalancingTcpRouteOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpRoute) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Route traffic for this frontend
func (o IpLoadBalancingTcpRouteOutput) FrontendId() pulumi.IntOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpRoute) pulumi.IntOutput { return v.FrontendId }).(pulumi.IntOutput)
}

// List of rules to match to trigger action
func (o IpLoadBalancingTcpRouteOutput) Rules() IpLoadBalancingTcpRouteRuleTypeArrayOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpRoute) IpLoadBalancingTcpRouteRuleTypeArrayOutput { return v.Rules }).(IpLoadBalancingTcpRouteRuleTypeArrayOutput)
}

// The internal name of your IP load balancing
func (o IpLoadBalancingTcpRouteOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpRoute) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Route status. Routes in "ok" state are ready to operate
func (o IpLoadBalancingTcpRouteOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpRoute) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
func (o IpLoadBalancingTcpRouteOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpRoute) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type IpLoadBalancingTcpRouteArrayOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingTcpRoute)(nil)).Elem()
}

func (o IpLoadBalancingTcpRouteArrayOutput) ToIpLoadBalancingTcpRouteArrayOutput() IpLoadBalancingTcpRouteArrayOutput {
	return o
}

func (o IpLoadBalancingTcpRouteArrayOutput) ToIpLoadBalancingTcpRouteArrayOutputWithContext(ctx context.Context) IpLoadBalancingTcpRouteArrayOutput {
	return o
}

func (o IpLoadBalancingTcpRouteArrayOutput) Index(i pulumi.IntInput) IpLoadBalancingTcpRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpLoadBalancingTcpRoute {
		return vs[0].([]*IpLoadBalancingTcpRoute)[vs[1].(int)]
	}).(IpLoadBalancingTcpRouteOutput)
}

type IpLoadBalancingTcpRouteMapOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingTcpRoute)(nil)).Elem()
}

func (o IpLoadBalancingTcpRouteMapOutput) ToIpLoadBalancingTcpRouteMapOutput() IpLoadBalancingTcpRouteMapOutput {
	return o
}

func (o IpLoadBalancingTcpRouteMapOutput) ToIpLoadBalancingTcpRouteMapOutputWithContext(ctx context.Context) IpLoadBalancingTcpRouteMapOutput {
	return o
}

func (o IpLoadBalancingTcpRouteMapOutput) MapIndex(k pulumi.StringInput) IpLoadBalancingTcpRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpLoadBalancingTcpRoute {
		return vs[0].(map[string]*IpLoadBalancingTcpRoute)[vs[1].(string)]
	}).(IpLoadBalancingTcpRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpRouteInput)(nil)).Elem(), &IpLoadBalancingTcpRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpRouteArrayInput)(nil)).Elem(), IpLoadBalancingTcpRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpRouteMapInput)(nil)).Elem(), IpLoadBalancingTcpRouteMap{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpRouteOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpRouteArrayOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpRouteMapOutput{})
}
