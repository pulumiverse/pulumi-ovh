// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Add a new access ACL for the given network/mask.
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
//			my_ceph, err := ovh.GetDedicatedCeph(ctx, &GetDedicatedCephArgs{
//				ServiceName: "94d423da-0e55-45f2-9812-836460a19939",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ovh.NewDedicatedCephAcl(ctx, "my-acl", &ovh.DedicatedCephAclArgs{
//				ServiceName: pulumi.String(my_ceph.Id),
//				Network:     pulumi.String("1.2.3.4"),
//				Netmask:     pulumi.String("255.255.255.255"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DedicatedCephAcl struct {
	pulumi.CustomResourceState

	// IP family. `IPv4` or `IPv6`
	Family pulumi.StringOutput `pulumi:"family"`
	// The network mask to apply
	Netmask pulumi.StringOutput `pulumi:"netmask"`
	// The network IP to authorize
	Network pulumi.StringOutput `pulumi:"network"`
	// The internal name of your dedicated CEPH
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewDedicatedCephAcl registers a new resource with the given unique name, arguments, and options.
func NewDedicatedCephAcl(ctx *pulumi.Context,
	name string, args *DedicatedCephAclArgs, opts ...pulumi.ResourceOption) (*DedicatedCephAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Netmask == nil {
		return nil, errors.New("invalid value for required argument 'Netmask'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DedicatedCephAcl
	err := ctx.RegisterResource("ovh:index/dedicatedCephAcl:DedicatedCephAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedCephAcl gets an existing DedicatedCephAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedCephAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedCephAclState, opts ...pulumi.ResourceOption) (*DedicatedCephAcl, error) {
	var resource DedicatedCephAcl
	err := ctx.ReadResource("ovh:index/dedicatedCephAcl:DedicatedCephAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedCephAcl resources.
type dedicatedCephAclState struct {
	// IP family. `IPv4` or `IPv6`
	Family *string `pulumi:"family"`
	// The network mask to apply
	Netmask *string `pulumi:"netmask"`
	// The network IP to authorize
	Network *string `pulumi:"network"`
	// The internal name of your dedicated CEPH
	ServiceName *string `pulumi:"serviceName"`
}

type DedicatedCephAclState struct {
	// IP family. `IPv4` or `IPv6`
	Family pulumi.StringPtrInput
	// The network mask to apply
	Netmask pulumi.StringPtrInput
	// The network IP to authorize
	Network pulumi.StringPtrInput
	// The internal name of your dedicated CEPH
	ServiceName pulumi.StringPtrInput
}

func (DedicatedCephAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCephAclState)(nil)).Elem()
}

type dedicatedCephAclArgs struct {
	// The network mask to apply
	Netmask string `pulumi:"netmask"`
	// The network IP to authorize
	Network string `pulumi:"network"`
	// The internal name of your dedicated CEPH
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a DedicatedCephAcl resource.
type DedicatedCephAclArgs struct {
	// The network mask to apply
	Netmask pulumi.StringInput
	// The network IP to authorize
	Network pulumi.StringInput
	// The internal name of your dedicated CEPH
	ServiceName pulumi.StringInput
}

func (DedicatedCephAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCephAclArgs)(nil)).Elem()
}

type DedicatedCephAclInput interface {
	pulumi.Input

	ToDedicatedCephAclOutput() DedicatedCephAclOutput
	ToDedicatedCephAclOutputWithContext(ctx context.Context) DedicatedCephAclOutput
}

func (*DedicatedCephAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCephAcl)(nil)).Elem()
}

func (i *DedicatedCephAcl) ToDedicatedCephAclOutput() DedicatedCephAclOutput {
	return i.ToDedicatedCephAclOutputWithContext(context.Background())
}

func (i *DedicatedCephAcl) ToDedicatedCephAclOutputWithContext(ctx context.Context) DedicatedCephAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCephAclOutput)
}

// DedicatedCephAclArrayInput is an input type that accepts DedicatedCephAclArray and DedicatedCephAclArrayOutput values.
// You can construct a concrete instance of `DedicatedCephAclArrayInput` via:
//
//	DedicatedCephAclArray{ DedicatedCephAclArgs{...} }
type DedicatedCephAclArrayInput interface {
	pulumi.Input

	ToDedicatedCephAclArrayOutput() DedicatedCephAclArrayOutput
	ToDedicatedCephAclArrayOutputWithContext(context.Context) DedicatedCephAclArrayOutput
}

type DedicatedCephAclArray []DedicatedCephAclInput

func (DedicatedCephAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedCephAcl)(nil)).Elem()
}

func (i DedicatedCephAclArray) ToDedicatedCephAclArrayOutput() DedicatedCephAclArrayOutput {
	return i.ToDedicatedCephAclArrayOutputWithContext(context.Background())
}

func (i DedicatedCephAclArray) ToDedicatedCephAclArrayOutputWithContext(ctx context.Context) DedicatedCephAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCephAclArrayOutput)
}

// DedicatedCephAclMapInput is an input type that accepts DedicatedCephAclMap and DedicatedCephAclMapOutput values.
// You can construct a concrete instance of `DedicatedCephAclMapInput` via:
//
//	DedicatedCephAclMap{ "key": DedicatedCephAclArgs{...} }
type DedicatedCephAclMapInput interface {
	pulumi.Input

	ToDedicatedCephAclMapOutput() DedicatedCephAclMapOutput
	ToDedicatedCephAclMapOutputWithContext(context.Context) DedicatedCephAclMapOutput
}

type DedicatedCephAclMap map[string]DedicatedCephAclInput

func (DedicatedCephAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedCephAcl)(nil)).Elem()
}

func (i DedicatedCephAclMap) ToDedicatedCephAclMapOutput() DedicatedCephAclMapOutput {
	return i.ToDedicatedCephAclMapOutputWithContext(context.Background())
}

func (i DedicatedCephAclMap) ToDedicatedCephAclMapOutputWithContext(ctx context.Context) DedicatedCephAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCephAclMapOutput)
}

type DedicatedCephAclOutput struct{ *pulumi.OutputState }

func (DedicatedCephAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCephAcl)(nil)).Elem()
}

func (o DedicatedCephAclOutput) ToDedicatedCephAclOutput() DedicatedCephAclOutput {
	return o
}

func (o DedicatedCephAclOutput) ToDedicatedCephAclOutputWithContext(ctx context.Context) DedicatedCephAclOutput {
	return o
}

// IP family. `IPv4` or `IPv6`
func (o DedicatedCephAclOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCephAcl) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// The network mask to apply
func (o DedicatedCephAclOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCephAcl) pulumi.StringOutput { return v.Netmask }).(pulumi.StringOutput)
}

// The network IP to authorize
func (o DedicatedCephAclOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCephAcl) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The internal name of your dedicated CEPH
func (o DedicatedCephAclOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCephAcl) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type DedicatedCephAclArrayOutput struct{ *pulumi.OutputState }

func (DedicatedCephAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedCephAcl)(nil)).Elem()
}

func (o DedicatedCephAclArrayOutput) ToDedicatedCephAclArrayOutput() DedicatedCephAclArrayOutput {
	return o
}

func (o DedicatedCephAclArrayOutput) ToDedicatedCephAclArrayOutputWithContext(ctx context.Context) DedicatedCephAclArrayOutput {
	return o
}

func (o DedicatedCephAclArrayOutput) Index(i pulumi.IntInput) DedicatedCephAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedCephAcl {
		return vs[0].([]*DedicatedCephAcl)[vs[1].(int)]
	}).(DedicatedCephAclOutput)
}

type DedicatedCephAclMapOutput struct{ *pulumi.OutputState }

func (DedicatedCephAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedCephAcl)(nil)).Elem()
}

func (o DedicatedCephAclMapOutput) ToDedicatedCephAclMapOutput() DedicatedCephAclMapOutput {
	return o
}

func (o DedicatedCephAclMapOutput) ToDedicatedCephAclMapOutputWithContext(ctx context.Context) DedicatedCephAclMapOutput {
	return o
}

func (o DedicatedCephAclMapOutput) MapIndex(k pulumi.StringInput) DedicatedCephAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedCephAcl {
		return vs[0].(map[string]*DedicatedCephAcl)[vs[1].(string)]
	}).(DedicatedCephAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCephAclInput)(nil)).Elem(), &DedicatedCephAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCephAclArrayInput)(nil)).Elem(), DedicatedCephAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCephAclMapInput)(nil)).Elem(), DedicatedCephAclMap{})
	pulumi.RegisterOutputType(DedicatedCephAclOutput{})
	pulumi.RegisterOutputType(DedicatedCephAclArrayOutput{})
	pulumi.RegisterOutputType(DedicatedCephAclMapOutput{})
}
