// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an SSH Key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-ovh/sdk/go/ovh"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ovh.NewMeSshKey(ctx, "mykey", &ovh.MeSshKeyArgs{
//				Key:     pulumi.String("ssh-ed25519 AAAAC3..."),
//				KeyName: pulumi.String("mykey"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type MeSshKey struct {
	pulumi.CustomResourceState

	// True when this public SSH key is used for rescue mode and reinstallations.
	Default pulumi.BoolOutput `pulumi:"default"`
	// The content of the public key in the form "ssh-algo content", e.g. "ssh-ed25519 AAAAC3...".
	Key pulumi.StringOutput `pulumi:"key"`
	// The friendly name of this SSH key.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
}

// NewMeSshKey registers a new resource with the given unique name, arguments, and options.
func NewMeSshKey(ctx *pulumi.Context,
	name string, args *MeSshKeyArgs, opts ...pulumi.ResourceOption) (*MeSshKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MeSshKey
	err := ctx.RegisterResource("ovh:index/meSshKey:MeSshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMeSshKey gets an existing MeSshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMeSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MeSshKeyState, opts ...pulumi.ResourceOption) (*MeSshKey, error) {
	var resource MeSshKey
	err := ctx.ReadResource("ovh:index/meSshKey:MeSshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MeSshKey resources.
type meSshKeyState struct {
	// True when this public SSH key is used for rescue mode and reinstallations.
	Default *bool `pulumi:"default"`
	// The content of the public key in the form "ssh-algo content", e.g. "ssh-ed25519 AAAAC3...".
	Key *string `pulumi:"key"`
	// The friendly name of this SSH key.
	KeyName *string `pulumi:"keyName"`
}

type MeSshKeyState struct {
	// True when this public SSH key is used for rescue mode and reinstallations.
	Default pulumi.BoolPtrInput
	// The content of the public key in the form "ssh-algo content", e.g. "ssh-ed25519 AAAAC3...".
	Key pulumi.StringPtrInput
	// The friendly name of this SSH key.
	KeyName pulumi.StringPtrInput
}

func (MeSshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*meSshKeyState)(nil)).Elem()
}

type meSshKeyArgs struct {
	// True when this public SSH key is used for rescue mode and reinstallations.
	Default *bool `pulumi:"default"`
	// The content of the public key in the form "ssh-algo content", e.g. "ssh-ed25519 AAAAC3...".
	Key string `pulumi:"key"`
	// The friendly name of this SSH key.
	KeyName string `pulumi:"keyName"`
}

// The set of arguments for constructing a MeSshKey resource.
type MeSshKeyArgs struct {
	// True when this public SSH key is used for rescue mode and reinstallations.
	Default pulumi.BoolPtrInput
	// The content of the public key in the form "ssh-algo content", e.g. "ssh-ed25519 AAAAC3...".
	Key pulumi.StringInput
	// The friendly name of this SSH key.
	KeyName pulumi.StringInput
}

func (MeSshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*meSshKeyArgs)(nil)).Elem()
}

type MeSshKeyInput interface {
	pulumi.Input

	ToMeSshKeyOutput() MeSshKeyOutput
	ToMeSshKeyOutputWithContext(ctx context.Context) MeSshKeyOutput
}

func (*MeSshKey) ElementType() reflect.Type {
	return reflect.TypeOf((**MeSshKey)(nil)).Elem()
}

func (i *MeSshKey) ToMeSshKeyOutput() MeSshKeyOutput {
	return i.ToMeSshKeyOutputWithContext(context.Background())
}

func (i *MeSshKey) ToMeSshKeyOutputWithContext(ctx context.Context) MeSshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeSshKeyOutput)
}

// MeSshKeyArrayInput is an input type that accepts MeSshKeyArray and MeSshKeyArrayOutput values.
// You can construct a concrete instance of `MeSshKeyArrayInput` via:
//
//	MeSshKeyArray{ MeSshKeyArgs{...} }
type MeSshKeyArrayInput interface {
	pulumi.Input

	ToMeSshKeyArrayOutput() MeSshKeyArrayOutput
	ToMeSshKeyArrayOutputWithContext(context.Context) MeSshKeyArrayOutput
}

type MeSshKeyArray []MeSshKeyInput

func (MeSshKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MeSshKey)(nil)).Elem()
}

func (i MeSshKeyArray) ToMeSshKeyArrayOutput() MeSshKeyArrayOutput {
	return i.ToMeSshKeyArrayOutputWithContext(context.Background())
}

func (i MeSshKeyArray) ToMeSshKeyArrayOutputWithContext(ctx context.Context) MeSshKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeSshKeyArrayOutput)
}

// MeSshKeyMapInput is an input type that accepts MeSshKeyMap and MeSshKeyMapOutput values.
// You can construct a concrete instance of `MeSshKeyMapInput` via:
//
//	MeSshKeyMap{ "key": MeSshKeyArgs{...} }
type MeSshKeyMapInput interface {
	pulumi.Input

	ToMeSshKeyMapOutput() MeSshKeyMapOutput
	ToMeSshKeyMapOutputWithContext(context.Context) MeSshKeyMapOutput
}

type MeSshKeyMap map[string]MeSshKeyInput

func (MeSshKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MeSshKey)(nil)).Elem()
}

func (i MeSshKeyMap) ToMeSshKeyMapOutput() MeSshKeyMapOutput {
	return i.ToMeSshKeyMapOutputWithContext(context.Background())
}

func (i MeSshKeyMap) ToMeSshKeyMapOutputWithContext(ctx context.Context) MeSshKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeSshKeyMapOutput)
}

type MeSshKeyOutput struct{ *pulumi.OutputState }

func (MeSshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MeSshKey)(nil)).Elem()
}

func (o MeSshKeyOutput) ToMeSshKeyOutput() MeSshKeyOutput {
	return o
}

func (o MeSshKeyOutput) ToMeSshKeyOutputWithContext(ctx context.Context) MeSshKeyOutput {
	return o
}

// True when this public SSH key is used for rescue mode and reinstallations.
func (o MeSshKeyOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v *MeSshKey) pulumi.BoolOutput { return v.Default }).(pulumi.BoolOutput)
}

// The content of the public key in the form "ssh-algo content", e.g. "ssh-ed25519 AAAAC3...".
func (o MeSshKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *MeSshKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The friendly name of this SSH key.
func (o MeSshKeyOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *MeSshKey) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

type MeSshKeyArrayOutput struct{ *pulumi.OutputState }

func (MeSshKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MeSshKey)(nil)).Elem()
}

func (o MeSshKeyArrayOutput) ToMeSshKeyArrayOutput() MeSshKeyArrayOutput {
	return o
}

func (o MeSshKeyArrayOutput) ToMeSshKeyArrayOutputWithContext(ctx context.Context) MeSshKeyArrayOutput {
	return o
}

func (o MeSshKeyArrayOutput) Index(i pulumi.IntInput) MeSshKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MeSshKey {
		return vs[0].([]*MeSshKey)[vs[1].(int)]
	}).(MeSshKeyOutput)
}

type MeSshKeyMapOutput struct{ *pulumi.OutputState }

func (MeSshKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MeSshKey)(nil)).Elem()
}

func (o MeSshKeyMapOutput) ToMeSshKeyMapOutput() MeSshKeyMapOutput {
	return o
}

func (o MeSshKeyMapOutput) ToMeSshKeyMapOutputWithContext(ctx context.Context) MeSshKeyMapOutput {
	return o
}

func (o MeSshKeyMapOutput) MapIndex(k pulumi.StringInput) MeSshKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MeSshKey {
		return vs[0].(map[string]*MeSshKey)[vs[1].(string)]
	}).(MeSshKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MeSshKeyInput)(nil)).Elem(), &MeSshKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*MeSshKeyArrayInput)(nil)).Elem(), MeSshKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MeSshKeyMapInput)(nil)).Elem(), MeSshKeyMap{})
	pulumi.RegisterOutputType(MeSshKeyOutput{})
	pulumi.RegisterOutputType(MeSshKeyArrayOutput{})
	pulumi.RegisterOutputType(MeSshKeyMapOutput{})
}
