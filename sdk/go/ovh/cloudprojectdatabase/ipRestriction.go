// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Apply IP restrictions to an OVHcloud Managed Database cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			db, err := CloudProjectDatabase.GetDatabase(ctx, &cloudprojectdatabase.GetDatabaseArgs{
//				ServiceName: "XXXX",
//				Engine:      "YYYY",
//				Id:          "ZZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = CloudProjectDatabase.NewIpRestriction(ctx, "iprestriction", &CloudProjectDatabase.IpRestrictionArgs{
//				ServiceName: pulumi.String(db.ServiceName),
//				Engine:      pulumi.String(db.Engine),
//				ClusterId:   pulumi.String(db.Id),
//				Ip:          pulumi.String("178.97.6.0/24"),
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
// OVHcloud Managed database cluster IP restrictions can be imported using the `service_name`, `engine`, `cluster_id` and the `ip`, separated by "/" E.g., bash
//
// ```sh
//
//	$ pulumi import ovh:CloudProjectDatabase/ipRestriction:IpRestriction my_ip_restriction service_name/engine/cluster_id/178.97.6.0/24
//
// ```
type IpRestriction struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Description of the IP restriction.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Authorized IP.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current status of the IP restriction.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewIpRestriction registers a new resource with the given unique name, arguments, and options.
func NewIpRestriction(ctx *pulumi.Context,
	name string, args *IpRestrictionArgs, opts ...pulumi.ResourceOption) (*IpRestriction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpRestriction
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/ipRestriction:IpRestriction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpRestriction gets an existing IpRestriction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpRestriction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpRestrictionState, opts ...pulumi.ResourceOption) (*IpRestriction, error) {
	var resource IpRestriction
	err := ctx.ReadResource("ovh:CloudProjectDatabase/ipRestriction:IpRestriction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpRestriction resources.
type ipRestrictionState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Description of the IP restriction.
	Description *string `pulumi:"description"`
	// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine *string `pulumi:"engine"`
	// Authorized IP.
	Ip *string `pulumi:"ip"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current status of the IP restriction.
	Status *string `pulumi:"status"`
}

type IpRestrictionState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Description of the IP restriction.
	Description pulumi.StringPtrInput
	// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringPtrInput
	// Authorized IP.
	Ip pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current status of the IP restriction.
	Status pulumi.StringPtrInput
}

func (IpRestrictionState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipRestrictionState)(nil)).Elem()
}

type ipRestrictionArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Description of the IP restriction.
	Description *string `pulumi:"description"`
	// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// Authorized IP.
	Ip string `pulumi:"ip"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a IpRestriction resource.
type IpRestrictionArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Description of the IP restriction.
	Description pulumi.StringPtrInput
	// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput
	// Authorized IP.
	Ip pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (IpRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipRestrictionArgs)(nil)).Elem()
}

type IpRestrictionInput interface {
	pulumi.Input

	ToIpRestrictionOutput() IpRestrictionOutput
	ToIpRestrictionOutputWithContext(ctx context.Context) IpRestrictionOutput
}

func (*IpRestriction) ElementType() reflect.Type {
	return reflect.TypeOf((**IpRestriction)(nil)).Elem()
}

func (i *IpRestriction) ToIpRestrictionOutput() IpRestrictionOutput {
	return i.ToIpRestrictionOutputWithContext(context.Background())
}

func (i *IpRestriction) ToIpRestrictionOutputWithContext(ctx context.Context) IpRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRestrictionOutput)
}

// IpRestrictionArrayInput is an input type that accepts IpRestrictionArray and IpRestrictionArrayOutput values.
// You can construct a concrete instance of `IpRestrictionArrayInput` via:
//
//	IpRestrictionArray{ IpRestrictionArgs{...} }
type IpRestrictionArrayInput interface {
	pulumi.Input

	ToIpRestrictionArrayOutput() IpRestrictionArrayOutput
	ToIpRestrictionArrayOutputWithContext(context.Context) IpRestrictionArrayOutput
}

type IpRestrictionArray []IpRestrictionInput

func (IpRestrictionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpRestriction)(nil)).Elem()
}

func (i IpRestrictionArray) ToIpRestrictionArrayOutput() IpRestrictionArrayOutput {
	return i.ToIpRestrictionArrayOutputWithContext(context.Background())
}

func (i IpRestrictionArray) ToIpRestrictionArrayOutputWithContext(ctx context.Context) IpRestrictionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRestrictionArrayOutput)
}

// IpRestrictionMapInput is an input type that accepts IpRestrictionMap and IpRestrictionMapOutput values.
// You can construct a concrete instance of `IpRestrictionMapInput` via:
//
//	IpRestrictionMap{ "key": IpRestrictionArgs{...} }
type IpRestrictionMapInput interface {
	pulumi.Input

	ToIpRestrictionMapOutput() IpRestrictionMapOutput
	ToIpRestrictionMapOutputWithContext(context.Context) IpRestrictionMapOutput
}

type IpRestrictionMap map[string]IpRestrictionInput

func (IpRestrictionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpRestriction)(nil)).Elem()
}

func (i IpRestrictionMap) ToIpRestrictionMapOutput() IpRestrictionMapOutput {
	return i.ToIpRestrictionMapOutputWithContext(context.Background())
}

func (i IpRestrictionMap) ToIpRestrictionMapOutputWithContext(ctx context.Context) IpRestrictionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRestrictionMapOutput)
}

type IpRestrictionOutput struct{ *pulumi.OutputState }

func (IpRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpRestriction)(nil)).Elem()
}

func (o IpRestrictionOutput) ToIpRestrictionOutput() IpRestrictionOutput {
	return o
}

func (o IpRestrictionOutput) ToIpRestrictionOutputWithContext(ctx context.Context) IpRestrictionOutput {
	return o
}

// Cluster ID.
func (o IpRestrictionOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRestriction) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Description of the IP restriction.
func (o IpRestrictionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpRestriction) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
func (o IpRestrictionOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRestriction) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Authorized IP.
func (o IpRestrictionOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRestriction) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o IpRestrictionOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRestriction) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the IP restriction.
func (o IpRestrictionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRestriction) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type IpRestrictionArrayOutput struct{ *pulumi.OutputState }

func (IpRestrictionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpRestriction)(nil)).Elem()
}

func (o IpRestrictionArrayOutput) ToIpRestrictionArrayOutput() IpRestrictionArrayOutput {
	return o
}

func (o IpRestrictionArrayOutput) ToIpRestrictionArrayOutputWithContext(ctx context.Context) IpRestrictionArrayOutput {
	return o
}

func (o IpRestrictionArrayOutput) Index(i pulumi.IntInput) IpRestrictionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpRestriction {
		return vs[0].([]*IpRestriction)[vs[1].(int)]
	}).(IpRestrictionOutput)
}

type IpRestrictionMapOutput struct{ *pulumi.OutputState }

func (IpRestrictionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpRestriction)(nil)).Elem()
}

func (o IpRestrictionMapOutput) ToIpRestrictionMapOutput() IpRestrictionMapOutput {
	return o
}

func (o IpRestrictionMapOutput) ToIpRestrictionMapOutputWithContext(ctx context.Context) IpRestrictionMapOutput {
	return o
}

func (o IpRestrictionMapOutput) MapIndex(k pulumi.StringInput) IpRestrictionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpRestriction {
		return vs[0].(map[string]*IpRestriction)[vs[1].(string)]
	}).(IpRestrictionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpRestrictionInput)(nil)).Elem(), &IpRestriction{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRestrictionArrayInput)(nil)).Elem(), IpRestrictionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRestrictionMapInput)(nil)).Elem(), IpRestrictionMap{})
	pulumi.RegisterOutputType(IpRestrictionOutput{})
	pulumi.RegisterOutputType(IpRestrictionArrayOutput{})
	pulumi.RegisterOutputType(IpRestrictionMapOutput{})
}
