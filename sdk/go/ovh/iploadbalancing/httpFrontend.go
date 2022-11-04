// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backend HTTP server group (frontend) to be used by loadbalancing frontend(s)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/IpLoadBalancing"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			lb, err := IpLoadBalancing.GetIpLoadBalancing(ctx, &iploadbalancing.GetIpLoadBalancingArgs{
//				ServiceName: pulumi.StringRef("ip-1.2.3.4"),
//				State:       pulumi.StringRef("ok"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			farm80, err := IpLoadBalancing.NewHttpFarm(ctx, "farm80", &IpLoadBalancing.HttpFarmArgs{
//				DisplayName: pulumi.String("ingress-8080-gra"),
//				Port:        pulumi.Int(80),
//				ServiceName: pulumi.String(lb.ServiceName),
//				Zone:        pulumi.String("all"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = IpLoadBalancing.NewHttpFrontend(ctx, "testfrontend", &IpLoadBalancing.HttpFrontendArgs{
//				DefaultFarmId: farm80.ID(),
//				DisplayName:   pulumi.String("ingress-8080-gra"),
//				Port:          pulumi.String("80,443"),
//				ServiceName:   pulumi.String(lb.ServiceName),
//				Zone:          pulumi.String("all"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type HttpFrontend struct {
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
	// Redirection HTTP'
	RedirectLocation pulumi.StringPtrOutput `pulumi:"redirectLocation"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrOutput `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewHttpFrontend registers a new resource with the given unique name, arguments, and options.
func NewHttpFrontend(ctx *pulumi.Context,
	name string, args *HttpFrontendArgs, opts ...pulumi.ResourceOption) (*HttpFrontend, error) {
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
	var resource HttpFrontend
	err := ctx.RegisterResource("ovh:IpLoadBalancing/httpFrontend:HttpFrontend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpFrontend gets an existing HttpFrontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpFrontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpFrontendState, opts ...pulumi.ResourceOption) (*HttpFrontend, error) {
	var resource HttpFrontend
	err := ctx.ReadResource("ovh:IpLoadBalancing/httpFrontend:HttpFrontend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpFrontend resources.
type httpFrontendState struct {
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
	// Redirection HTTP'
	RedirectLocation *string `pulumi:"redirectLocation"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl *bool `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone *string `pulumi:"zone"`
}

type HttpFrontendState struct {
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
	// Redirection HTTP'
	RedirectLocation pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrInput
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringPtrInput
}

func (HttpFrontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpFrontendState)(nil)).Elem()
}

type httpFrontendArgs struct {
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
	// Redirection HTTP'
	RedirectLocation *string `pulumi:"redirectLocation"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl *bool `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a HttpFrontend resource.
type HttpFrontendArgs struct {
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
	// Redirection HTTP'
	RedirectLocation pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrInput
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringInput
}

func (HttpFrontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpFrontendArgs)(nil)).Elem()
}

type HttpFrontendInput interface {
	pulumi.Input

	ToHttpFrontendOutput() HttpFrontendOutput
	ToHttpFrontendOutputWithContext(ctx context.Context) HttpFrontendOutput
}

func (*HttpFrontend) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpFrontend)(nil)).Elem()
}

func (i *HttpFrontend) ToHttpFrontendOutput() HttpFrontendOutput {
	return i.ToHttpFrontendOutputWithContext(context.Background())
}

func (i *HttpFrontend) ToHttpFrontendOutputWithContext(ctx context.Context) HttpFrontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFrontendOutput)
}

// HttpFrontendArrayInput is an input type that accepts HttpFrontendArray and HttpFrontendArrayOutput values.
// You can construct a concrete instance of `HttpFrontendArrayInput` via:
//
//	HttpFrontendArray{ HttpFrontendArgs{...} }
type HttpFrontendArrayInput interface {
	pulumi.Input

	ToHttpFrontendArrayOutput() HttpFrontendArrayOutput
	ToHttpFrontendArrayOutputWithContext(context.Context) HttpFrontendArrayOutput
}

type HttpFrontendArray []HttpFrontendInput

func (HttpFrontendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpFrontend)(nil)).Elem()
}

func (i HttpFrontendArray) ToHttpFrontendArrayOutput() HttpFrontendArrayOutput {
	return i.ToHttpFrontendArrayOutputWithContext(context.Background())
}

func (i HttpFrontendArray) ToHttpFrontendArrayOutputWithContext(ctx context.Context) HttpFrontendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFrontendArrayOutput)
}

// HttpFrontendMapInput is an input type that accepts HttpFrontendMap and HttpFrontendMapOutput values.
// You can construct a concrete instance of `HttpFrontendMapInput` via:
//
//	HttpFrontendMap{ "key": HttpFrontendArgs{...} }
type HttpFrontendMapInput interface {
	pulumi.Input

	ToHttpFrontendMapOutput() HttpFrontendMapOutput
	ToHttpFrontendMapOutputWithContext(context.Context) HttpFrontendMapOutput
}

type HttpFrontendMap map[string]HttpFrontendInput

func (HttpFrontendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpFrontend)(nil)).Elem()
}

func (i HttpFrontendMap) ToHttpFrontendMapOutput() HttpFrontendMapOutput {
	return i.ToHttpFrontendMapOutputWithContext(context.Background())
}

func (i HttpFrontendMap) ToHttpFrontendMapOutputWithContext(ctx context.Context) HttpFrontendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFrontendMapOutput)
}

type HttpFrontendOutput struct{ *pulumi.OutputState }

func (HttpFrontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpFrontend)(nil)).Elem()
}

func (o HttpFrontendOutput) ToHttpFrontendOutput() HttpFrontendOutput {
	return o
}

func (o HttpFrontendOutput) ToHttpFrontendOutputWithContext(ctx context.Context) HttpFrontendOutput {
	return o
}

// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
func (o HttpFrontendOutput) AllowedSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.StringArrayOutput { return v.AllowedSources }).(pulumi.StringArrayOutput)
}

// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
func (o HttpFrontendOutput) DedicatedIpfos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.StringArrayOutput { return v.DedicatedIpfos }).(pulumi.StringArrayOutput)
}

// Default TCP Farm of your frontend
func (o HttpFrontendOutput) DefaultFarmId() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.IntOutput { return v.DefaultFarmId }).(pulumi.IntOutput)
}

// Default ssl served to your customer
func (o HttpFrontendOutput) DefaultSslId() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.IntOutput { return v.DefaultSslId }).(pulumi.IntOutput)
}

// Disable your frontend. Default: 'false'
func (o HttpFrontendOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Human readable name for your frontend, this field is for you
func (o HttpFrontendOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Port(s) attached to your frontend. Supports single port (numerical value),
// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
// and/or 'range'. Each port must be in the [1;49151] range
func (o HttpFrontendOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// Redirection HTTP'
func (o HttpFrontendOutput) RedirectLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.StringPtrOutput { return v.RedirectLocation }).(pulumi.StringPtrOutput)
}

// The internal name of your IP load balancing
func (o HttpFrontendOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// SSL deciphering. Default: 'false'
func (o HttpFrontendOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
func (o HttpFrontendOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFrontend) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type HttpFrontendArrayOutput struct{ *pulumi.OutputState }

func (HttpFrontendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpFrontend)(nil)).Elem()
}

func (o HttpFrontendArrayOutput) ToHttpFrontendArrayOutput() HttpFrontendArrayOutput {
	return o
}

func (o HttpFrontendArrayOutput) ToHttpFrontendArrayOutputWithContext(ctx context.Context) HttpFrontendArrayOutput {
	return o
}

func (o HttpFrontendArrayOutput) Index(i pulumi.IntInput) HttpFrontendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HttpFrontend {
		return vs[0].([]*HttpFrontend)[vs[1].(int)]
	}).(HttpFrontendOutput)
}

type HttpFrontendMapOutput struct{ *pulumi.OutputState }

func (HttpFrontendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpFrontend)(nil)).Elem()
}

func (o HttpFrontendMapOutput) ToHttpFrontendMapOutput() HttpFrontendMapOutput {
	return o
}

func (o HttpFrontendMapOutput) ToHttpFrontendMapOutputWithContext(ctx context.Context) HttpFrontendMapOutput {
	return o
}

func (o HttpFrontendMapOutput) MapIndex(k pulumi.StringInput) HttpFrontendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HttpFrontend {
		return vs[0].(map[string]*HttpFrontend)[vs[1].(string)]
	}).(HttpFrontendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFrontendInput)(nil)).Elem(), &HttpFrontend{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFrontendArrayInput)(nil)).Elem(), HttpFrontendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFrontendMapInput)(nil)).Elem(), HttpFrontendMap{})
	pulumi.RegisterOutputType(HttpFrontendOutput{})
	pulumi.RegisterOutputType(HttpFrontendArrayOutput{})
	pulumi.RegisterOutputType(HttpFrontendMapOutput{})
}