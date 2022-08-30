// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a private network in a public cloud project.
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
//			_, err := ovh.NewCloudProjectNetworkPrivate(ctx, "net", &ovh.CloudProjectNetworkPrivateArgs{
//				Regions: pulumi.StringArray{
//					pulumi.String("GRA1"),
//					pulumi.String("BHS1"),
//				},
//				ServiceName: pulumi.String("XXXXXX"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CloudProjectNetworkPrivate struct {
	pulumi.CustomResourceState

	// The name of the network.
	Name pulumi.StringOutput `pulumi:"name"`
	// an array of valid OVH public cloud region ID in which the network
	// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// A map representing information about the region.
	// * `regions_attributes/region` - The id of the region.
	// * `regions_attributes/status` - The status of the network in the region.
	// * `regions_attributes/openstackid` - The private network id in the region.
	RegionsAttributes CloudProjectNetworkPrivateRegionsAttributeArrayOutput `pulumi:"regionsAttributes"`
	// (Deprecated) A map representing the status of the network per region.
	// * `regions_status/region` - (Deprecated) The id of the region.
	// * `regions_status/status` - (Deprecated) The status of the network in the region.
	//
	// Deprecated: use the regions_attributes field instead
	RegionsStatuses CloudProjectNetworkPrivateRegionsStatusArrayOutput `pulumi:"regionsStatuses"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// the status of the network. should be normally set to 'ACTIVE'.
	Status pulumi.StringOutput `pulumi:"status"`
	// the type of the network. Either 'private' or 'public'.
	Type pulumi.StringOutput `pulumi:"type"`
	// a vlan id to associate with the network.
	// Changing this value recreates the resource. Defaults to 0.
	VlanId pulumi.IntPtrOutput `pulumi:"vlanId"`
}

// NewCloudProjectNetworkPrivate registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectNetworkPrivate(ctx *pulumi.Context,
	name string, args *CloudProjectNetworkPrivateArgs, opts ...pulumi.ResourceOption) (*CloudProjectNetworkPrivate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudProjectNetworkPrivate
	err := ctx.RegisterResource("ovh:index/cloudProjectNetworkPrivate:CloudProjectNetworkPrivate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectNetworkPrivate gets an existing CloudProjectNetworkPrivate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectNetworkPrivate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectNetworkPrivateState, opts ...pulumi.ResourceOption) (*CloudProjectNetworkPrivate, error) {
	var resource CloudProjectNetworkPrivate
	err := ctx.ReadResource("ovh:index/cloudProjectNetworkPrivate:CloudProjectNetworkPrivate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectNetworkPrivate resources.
type cloudProjectNetworkPrivateState struct {
	// The name of the network.
	Name *string `pulumi:"name"`
	// an array of valid OVH public cloud region ID in which the network
	// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
	Regions []string `pulumi:"regions"`
	// A map representing information about the region.
	// * `regions_attributes/region` - The id of the region.
	// * `regions_attributes/status` - The status of the network in the region.
	// * `regions_attributes/openstackid` - The private network id in the region.
	RegionsAttributes []CloudProjectNetworkPrivateRegionsAttribute `pulumi:"regionsAttributes"`
	// (Deprecated) A map representing the status of the network per region.
	// * `regions_status/region` - (Deprecated) The id of the region.
	// * `regions_status/status` - (Deprecated) The status of the network in the region.
	//
	// Deprecated: use the regions_attributes field instead
	RegionsStatuses []CloudProjectNetworkPrivateRegionsStatus `pulumi:"regionsStatuses"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// the status of the network. should be normally set to 'ACTIVE'.
	Status *string `pulumi:"status"`
	// the type of the network. Either 'private' or 'public'.
	Type *string `pulumi:"type"`
	// a vlan id to associate with the network.
	// Changing this value recreates the resource. Defaults to 0.
	VlanId *int `pulumi:"vlanId"`
}

type CloudProjectNetworkPrivateState struct {
	// The name of the network.
	Name pulumi.StringPtrInput
	// an array of valid OVH public cloud region ID in which the network
	// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
	Regions pulumi.StringArrayInput
	// A map representing information about the region.
	// * `regions_attributes/region` - The id of the region.
	// * `regions_attributes/status` - The status of the network in the region.
	// * `regions_attributes/openstackid` - The private network id in the region.
	RegionsAttributes CloudProjectNetworkPrivateRegionsAttributeArrayInput
	// (Deprecated) A map representing the status of the network per region.
	// * `regions_status/region` - (Deprecated) The id of the region.
	// * `regions_status/status` - (Deprecated) The status of the network in the region.
	//
	// Deprecated: use the regions_attributes field instead
	RegionsStatuses CloudProjectNetworkPrivateRegionsStatusArrayInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// the status of the network. should be normally set to 'ACTIVE'.
	Status pulumi.StringPtrInput
	// the type of the network. Either 'private' or 'public'.
	Type pulumi.StringPtrInput
	// a vlan id to associate with the network.
	// Changing this value recreates the resource. Defaults to 0.
	VlanId pulumi.IntPtrInput
}

func (CloudProjectNetworkPrivateState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectNetworkPrivateState)(nil)).Elem()
}

type cloudProjectNetworkPrivateArgs struct {
	// The name of the network.
	Name *string `pulumi:"name"`
	// an array of valid OVH public cloud region ID in which the network
	// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
	Regions []string `pulumi:"regions"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// a vlan id to associate with the network.
	// Changing this value recreates the resource. Defaults to 0.
	VlanId *int `pulumi:"vlanId"`
}

// The set of arguments for constructing a CloudProjectNetworkPrivate resource.
type CloudProjectNetworkPrivateArgs struct {
	// The name of the network.
	Name pulumi.StringPtrInput
	// an array of valid OVH public cloud region ID in which the network
	// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
	Regions pulumi.StringArrayInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// a vlan id to associate with the network.
	// Changing this value recreates the resource. Defaults to 0.
	VlanId pulumi.IntPtrInput
}

func (CloudProjectNetworkPrivateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectNetworkPrivateArgs)(nil)).Elem()
}

type CloudProjectNetworkPrivateInput interface {
	pulumi.Input

	ToCloudProjectNetworkPrivateOutput() CloudProjectNetworkPrivateOutput
	ToCloudProjectNetworkPrivateOutputWithContext(ctx context.Context) CloudProjectNetworkPrivateOutput
}

func (*CloudProjectNetworkPrivate) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectNetworkPrivate)(nil)).Elem()
}

func (i *CloudProjectNetworkPrivate) ToCloudProjectNetworkPrivateOutput() CloudProjectNetworkPrivateOutput {
	return i.ToCloudProjectNetworkPrivateOutputWithContext(context.Background())
}

func (i *CloudProjectNetworkPrivate) ToCloudProjectNetworkPrivateOutputWithContext(ctx context.Context) CloudProjectNetworkPrivateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectNetworkPrivateOutput)
}

// CloudProjectNetworkPrivateArrayInput is an input type that accepts CloudProjectNetworkPrivateArray and CloudProjectNetworkPrivateArrayOutput values.
// You can construct a concrete instance of `CloudProjectNetworkPrivateArrayInput` via:
//
//	CloudProjectNetworkPrivateArray{ CloudProjectNetworkPrivateArgs{...} }
type CloudProjectNetworkPrivateArrayInput interface {
	pulumi.Input

	ToCloudProjectNetworkPrivateArrayOutput() CloudProjectNetworkPrivateArrayOutput
	ToCloudProjectNetworkPrivateArrayOutputWithContext(context.Context) CloudProjectNetworkPrivateArrayOutput
}

type CloudProjectNetworkPrivateArray []CloudProjectNetworkPrivateInput

func (CloudProjectNetworkPrivateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectNetworkPrivate)(nil)).Elem()
}

func (i CloudProjectNetworkPrivateArray) ToCloudProjectNetworkPrivateArrayOutput() CloudProjectNetworkPrivateArrayOutput {
	return i.ToCloudProjectNetworkPrivateArrayOutputWithContext(context.Background())
}

func (i CloudProjectNetworkPrivateArray) ToCloudProjectNetworkPrivateArrayOutputWithContext(ctx context.Context) CloudProjectNetworkPrivateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectNetworkPrivateArrayOutput)
}

// CloudProjectNetworkPrivateMapInput is an input type that accepts CloudProjectNetworkPrivateMap and CloudProjectNetworkPrivateMapOutput values.
// You can construct a concrete instance of `CloudProjectNetworkPrivateMapInput` via:
//
//	CloudProjectNetworkPrivateMap{ "key": CloudProjectNetworkPrivateArgs{...} }
type CloudProjectNetworkPrivateMapInput interface {
	pulumi.Input

	ToCloudProjectNetworkPrivateMapOutput() CloudProjectNetworkPrivateMapOutput
	ToCloudProjectNetworkPrivateMapOutputWithContext(context.Context) CloudProjectNetworkPrivateMapOutput
}

type CloudProjectNetworkPrivateMap map[string]CloudProjectNetworkPrivateInput

func (CloudProjectNetworkPrivateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectNetworkPrivate)(nil)).Elem()
}

func (i CloudProjectNetworkPrivateMap) ToCloudProjectNetworkPrivateMapOutput() CloudProjectNetworkPrivateMapOutput {
	return i.ToCloudProjectNetworkPrivateMapOutputWithContext(context.Background())
}

func (i CloudProjectNetworkPrivateMap) ToCloudProjectNetworkPrivateMapOutputWithContext(ctx context.Context) CloudProjectNetworkPrivateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectNetworkPrivateMapOutput)
}

type CloudProjectNetworkPrivateOutput struct{ *pulumi.OutputState }

func (CloudProjectNetworkPrivateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectNetworkPrivate)(nil)).Elem()
}

func (o CloudProjectNetworkPrivateOutput) ToCloudProjectNetworkPrivateOutput() CloudProjectNetworkPrivateOutput {
	return o
}

func (o CloudProjectNetworkPrivateOutput) ToCloudProjectNetworkPrivateOutputWithContext(ctx context.Context) CloudProjectNetworkPrivateOutput {
	return o
}

// The name of the network.
func (o CloudProjectNetworkPrivateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectNetworkPrivate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// an array of valid OVH public cloud region ID in which the network
// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
func (o CloudProjectNetworkPrivateOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProjectNetworkPrivate) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// A map representing information about the region.
// * `regions_attributes/region` - The id of the region.
// * `regions_attributes/status` - The status of the network in the region.
// * `regions_attributes/openstackid` - The private network id in the region.
func (o CloudProjectNetworkPrivateOutput) RegionsAttributes() CloudProjectNetworkPrivateRegionsAttributeArrayOutput {
	return o.ApplyT(func(v *CloudProjectNetworkPrivate) CloudProjectNetworkPrivateRegionsAttributeArrayOutput {
		return v.RegionsAttributes
	}).(CloudProjectNetworkPrivateRegionsAttributeArrayOutput)
}

// (Deprecated) A map representing the status of the network per region.
// * `regions_status/region` - (Deprecated) The id of the region.
// * `regions_status/status` - (Deprecated) The status of the network in the region.
//
// Deprecated: use the regions_attributes field instead
func (o CloudProjectNetworkPrivateOutput) RegionsStatuses() CloudProjectNetworkPrivateRegionsStatusArrayOutput {
	return o.ApplyT(func(v *CloudProjectNetworkPrivate) CloudProjectNetworkPrivateRegionsStatusArrayOutput {
		return v.RegionsStatuses
	}).(CloudProjectNetworkPrivateRegionsStatusArrayOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectNetworkPrivateOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectNetworkPrivate) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// the status of the network. should be normally set to 'ACTIVE'.
func (o CloudProjectNetworkPrivateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectNetworkPrivate) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// the type of the network. Either 'private' or 'public'.
func (o CloudProjectNetworkPrivateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectNetworkPrivate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// a vlan id to associate with the network.
// Changing this value recreates the resource. Defaults to 0.
func (o CloudProjectNetworkPrivateOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudProjectNetworkPrivate) pulumi.IntPtrOutput { return v.VlanId }).(pulumi.IntPtrOutput)
}

type CloudProjectNetworkPrivateArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectNetworkPrivateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectNetworkPrivate)(nil)).Elem()
}

func (o CloudProjectNetworkPrivateArrayOutput) ToCloudProjectNetworkPrivateArrayOutput() CloudProjectNetworkPrivateArrayOutput {
	return o
}

func (o CloudProjectNetworkPrivateArrayOutput) ToCloudProjectNetworkPrivateArrayOutputWithContext(ctx context.Context) CloudProjectNetworkPrivateArrayOutput {
	return o
}

func (o CloudProjectNetworkPrivateArrayOutput) Index(i pulumi.IntInput) CloudProjectNetworkPrivateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProjectNetworkPrivate {
		return vs[0].([]*CloudProjectNetworkPrivate)[vs[1].(int)]
	}).(CloudProjectNetworkPrivateOutput)
}

type CloudProjectNetworkPrivateMapOutput struct{ *pulumi.OutputState }

func (CloudProjectNetworkPrivateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectNetworkPrivate)(nil)).Elem()
}

func (o CloudProjectNetworkPrivateMapOutput) ToCloudProjectNetworkPrivateMapOutput() CloudProjectNetworkPrivateMapOutput {
	return o
}

func (o CloudProjectNetworkPrivateMapOutput) ToCloudProjectNetworkPrivateMapOutputWithContext(ctx context.Context) CloudProjectNetworkPrivateMapOutput {
	return o
}

func (o CloudProjectNetworkPrivateMapOutput) MapIndex(k pulumi.StringInput) CloudProjectNetworkPrivateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProjectNetworkPrivate {
		return vs[0].(map[string]*CloudProjectNetworkPrivate)[vs[1].(string)]
	}).(CloudProjectNetworkPrivateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectNetworkPrivateInput)(nil)).Elem(), &CloudProjectNetworkPrivate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectNetworkPrivateArrayInput)(nil)).Elem(), CloudProjectNetworkPrivateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectNetworkPrivateMapInput)(nil)).Elem(), CloudProjectNetworkPrivateMap{})
	pulumi.RegisterOutputType(CloudProjectNetworkPrivateOutput{})
	pulumi.RegisterOutputType(CloudProjectNetworkPrivateArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectNetworkPrivateMapOutput{})
}
