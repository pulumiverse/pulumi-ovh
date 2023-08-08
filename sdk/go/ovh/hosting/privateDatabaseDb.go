// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new database on your private cloud database service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Hosting"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Hosting.NewPrivateDatabaseDb(ctx, "database", &Hosting.PrivateDatabaseDbArgs{
//				DatabaseName: pulumi.String("XXXXXX"),
//				ServiceName:  pulumi.String("XXXXXX"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// OVHcloud Webhosting database can be imported using the `service_name` and the `database_name`, separated by "/" E.g.,
//
// ```sh
//
//	$ pulumi import ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb database service_name/database_name
//
// ```
type PrivateDatabaseDb struct {
	pulumi.CustomResourceState

	// Name of your new database
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The internal name of your private database.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewPrivateDatabaseDb registers a new resource with the given unique name, arguments, and options.
func NewPrivateDatabaseDb(ctx *pulumi.Context,
	name string, args *PrivateDatabaseDbArgs, opts ...pulumi.ResourceOption) (*PrivateDatabaseDb, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateDatabaseDb
	err := ctx.RegisterResource("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDatabaseDb gets an existing PrivateDatabaseDb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDatabaseDb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDatabaseDbState, opts ...pulumi.ResourceOption) (*PrivateDatabaseDb, error) {
	var resource PrivateDatabaseDb
	err := ctx.ReadResource("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDatabaseDb resources.
type privateDatabaseDbState struct {
	// Name of your new database
	DatabaseName *string `pulumi:"databaseName"`
	// The internal name of your private database.
	ServiceName *string `pulumi:"serviceName"`
}

type PrivateDatabaseDbState struct {
	// Name of your new database
	DatabaseName pulumi.StringPtrInput
	// The internal name of your private database.
	ServiceName pulumi.StringPtrInput
}

func (PrivateDatabaseDbState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDatabaseDbState)(nil)).Elem()
}

type privateDatabaseDbArgs struct {
	// Name of your new database
	DatabaseName string `pulumi:"databaseName"`
	// The internal name of your private database.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a PrivateDatabaseDb resource.
type PrivateDatabaseDbArgs struct {
	// Name of your new database
	DatabaseName pulumi.StringInput
	// The internal name of your private database.
	ServiceName pulumi.StringInput
}

func (PrivateDatabaseDbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDatabaseDbArgs)(nil)).Elem()
}

type PrivateDatabaseDbInput interface {
	pulumi.Input

	ToPrivateDatabaseDbOutput() PrivateDatabaseDbOutput
	ToPrivateDatabaseDbOutputWithContext(ctx context.Context) PrivateDatabaseDbOutput
}

func (*PrivateDatabaseDb) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDatabaseDb)(nil)).Elem()
}

func (i *PrivateDatabaseDb) ToPrivateDatabaseDbOutput() PrivateDatabaseDbOutput {
	return i.ToPrivateDatabaseDbOutputWithContext(context.Background())
}

func (i *PrivateDatabaseDb) ToPrivateDatabaseDbOutputWithContext(ctx context.Context) PrivateDatabaseDbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseDbOutput)
}

// PrivateDatabaseDbArrayInput is an input type that accepts PrivateDatabaseDbArray and PrivateDatabaseDbArrayOutput values.
// You can construct a concrete instance of `PrivateDatabaseDbArrayInput` via:
//
//	PrivateDatabaseDbArray{ PrivateDatabaseDbArgs{...} }
type PrivateDatabaseDbArrayInput interface {
	pulumi.Input

	ToPrivateDatabaseDbArrayOutput() PrivateDatabaseDbArrayOutput
	ToPrivateDatabaseDbArrayOutputWithContext(context.Context) PrivateDatabaseDbArrayOutput
}

type PrivateDatabaseDbArray []PrivateDatabaseDbInput

func (PrivateDatabaseDbArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDatabaseDb)(nil)).Elem()
}

func (i PrivateDatabaseDbArray) ToPrivateDatabaseDbArrayOutput() PrivateDatabaseDbArrayOutput {
	return i.ToPrivateDatabaseDbArrayOutputWithContext(context.Background())
}

func (i PrivateDatabaseDbArray) ToPrivateDatabaseDbArrayOutputWithContext(ctx context.Context) PrivateDatabaseDbArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseDbArrayOutput)
}

// PrivateDatabaseDbMapInput is an input type that accepts PrivateDatabaseDbMap and PrivateDatabaseDbMapOutput values.
// You can construct a concrete instance of `PrivateDatabaseDbMapInput` via:
//
//	PrivateDatabaseDbMap{ "key": PrivateDatabaseDbArgs{...} }
type PrivateDatabaseDbMapInput interface {
	pulumi.Input

	ToPrivateDatabaseDbMapOutput() PrivateDatabaseDbMapOutput
	ToPrivateDatabaseDbMapOutputWithContext(context.Context) PrivateDatabaseDbMapOutput
}

type PrivateDatabaseDbMap map[string]PrivateDatabaseDbInput

func (PrivateDatabaseDbMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDatabaseDb)(nil)).Elem()
}

func (i PrivateDatabaseDbMap) ToPrivateDatabaseDbMapOutput() PrivateDatabaseDbMapOutput {
	return i.ToPrivateDatabaseDbMapOutputWithContext(context.Background())
}

func (i PrivateDatabaseDbMap) ToPrivateDatabaseDbMapOutputWithContext(ctx context.Context) PrivateDatabaseDbMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseDbMapOutput)
}

type PrivateDatabaseDbOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseDbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDatabaseDb)(nil)).Elem()
}

func (o PrivateDatabaseDbOutput) ToPrivateDatabaseDbOutput() PrivateDatabaseDbOutput {
	return o
}

func (o PrivateDatabaseDbOutput) ToPrivateDatabaseDbOutputWithContext(ctx context.Context) PrivateDatabaseDbOutput {
	return o
}

// Name of your new database
func (o PrivateDatabaseDbOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseDb) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// The internal name of your private database.
func (o PrivateDatabaseDbOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseDb) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type PrivateDatabaseDbArrayOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseDbArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDatabaseDb)(nil)).Elem()
}

func (o PrivateDatabaseDbArrayOutput) ToPrivateDatabaseDbArrayOutput() PrivateDatabaseDbArrayOutput {
	return o
}

func (o PrivateDatabaseDbArrayOutput) ToPrivateDatabaseDbArrayOutputWithContext(ctx context.Context) PrivateDatabaseDbArrayOutput {
	return o
}

func (o PrivateDatabaseDbArrayOutput) Index(i pulumi.IntInput) PrivateDatabaseDbOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateDatabaseDb {
		return vs[0].([]*PrivateDatabaseDb)[vs[1].(int)]
	}).(PrivateDatabaseDbOutput)
}

type PrivateDatabaseDbMapOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseDbMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDatabaseDb)(nil)).Elem()
}

func (o PrivateDatabaseDbMapOutput) ToPrivateDatabaseDbMapOutput() PrivateDatabaseDbMapOutput {
	return o
}

func (o PrivateDatabaseDbMapOutput) ToPrivateDatabaseDbMapOutputWithContext(ctx context.Context) PrivateDatabaseDbMapOutput {
	return o
}

func (o PrivateDatabaseDbMapOutput) MapIndex(k pulumi.StringInput) PrivateDatabaseDbOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateDatabaseDb {
		return vs[0].(map[string]*PrivateDatabaseDb)[vs[1].(string)]
	}).(PrivateDatabaseDbOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseDbInput)(nil)).Elem(), &PrivateDatabaseDb{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseDbArrayInput)(nil)).Elem(), PrivateDatabaseDbArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseDbMapInput)(nil)).Elem(), PrivateDatabaseDbMap{})
	pulumi.RegisterOutputType(PrivateDatabaseDbOutput{})
	pulumi.RegisterOutputType(PrivateDatabaseDbArrayOutput{})
	pulumi.RegisterOutputType(PrivateDatabaseDbMapOutput{})
}
