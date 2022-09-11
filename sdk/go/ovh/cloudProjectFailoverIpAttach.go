// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches a failover IP address to a compute instance
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
//			_, err := ovh.NewCloudProjectFailoverIpAttach(ctx, "myfailoverip", &ovh.CloudProjectFailoverIpAttachArgs{
//				Ip:          pulumi.String("XXXXXX"),
//				RoutedTo:    pulumi.String("XXXXXX"),
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
type CloudProjectFailoverIpAttach struct {
	pulumi.CustomResourceState

	// The IP block
	Block pulumi.StringOutput `pulumi:"block"`
	// Ip continent
	ContinentCode pulumi.StringOutput `pulumi:"continentCode"`
	// Ip location
	GeoLoc pulumi.StringOutput `pulumi:"geoLoc"`
	// The failover ip address to attach
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Current operation progress in percent
	Progress pulumi.IntOutput `pulumi:"progress"`
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo pulumi.StringOutput `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Ip status, can be `ok` or `operationPending`
	Status pulumi.StringOutput `pulumi:"status"`
	// IP sub type
	SubType pulumi.StringOutput `pulumi:"subType"`
}

// NewCloudProjectFailoverIpAttach registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectFailoverIpAttach(ctx *pulumi.Context,
	name string, args *CloudProjectFailoverIpAttachArgs, opts ...pulumi.ResourceOption) (*CloudProjectFailoverIpAttach, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudProjectFailoverIpAttach
	err := ctx.RegisterResource("ovh:index/cloudProjectFailoverIpAttach:CloudProjectFailoverIpAttach", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectFailoverIpAttach gets an existing CloudProjectFailoverIpAttach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectFailoverIpAttach(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectFailoverIpAttachState, opts ...pulumi.ResourceOption) (*CloudProjectFailoverIpAttach, error) {
	var resource CloudProjectFailoverIpAttach
	err := ctx.ReadResource("ovh:index/cloudProjectFailoverIpAttach:CloudProjectFailoverIpAttach", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectFailoverIpAttach resources.
type cloudProjectFailoverIpAttachState struct {
	// The IP block
	Block *string `pulumi:"block"`
	// Ip continent
	ContinentCode *string `pulumi:"continentCode"`
	// Ip location
	GeoLoc *string `pulumi:"geoLoc"`
	// The failover ip address to attach
	Ip *string `pulumi:"ip"`
	// Current operation progress in percent
	Progress *int `pulumi:"progress"`
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo *string `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Ip status, can be `ok` or `operationPending`
	Status *string `pulumi:"status"`
	// IP sub type
	SubType *string `pulumi:"subType"`
}

type CloudProjectFailoverIpAttachState struct {
	// The IP block
	Block pulumi.StringPtrInput
	// Ip continent
	ContinentCode pulumi.StringPtrInput
	// Ip location
	GeoLoc pulumi.StringPtrInput
	// The failover ip address to attach
	Ip pulumi.StringPtrInput
	// Current operation progress in percent
	Progress pulumi.IntPtrInput
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Ip status, can be `ok` or `operationPending`
	Status pulumi.StringPtrInput
	// IP sub type
	SubType pulumi.StringPtrInput
}

func (CloudProjectFailoverIpAttachState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectFailoverIpAttachState)(nil)).Elem()
}

type cloudProjectFailoverIpAttachArgs struct {
	// The IP block
	Block *string `pulumi:"block"`
	// Ip continent
	ContinentCode *string `pulumi:"continentCode"`
	// Ip location
	GeoLoc *string `pulumi:"geoLoc"`
	// The failover ip address to attach
	Ip *string `pulumi:"ip"`
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo *string `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectFailoverIpAttach resource.
type CloudProjectFailoverIpAttachArgs struct {
	// The IP block
	Block pulumi.StringPtrInput
	// Ip continent
	ContinentCode pulumi.StringPtrInput
	// Ip location
	GeoLoc pulumi.StringPtrInput
	// The failover ip address to attach
	Ip pulumi.StringPtrInput
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (CloudProjectFailoverIpAttachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectFailoverIpAttachArgs)(nil)).Elem()
}

type CloudProjectFailoverIpAttachInput interface {
	pulumi.Input

	ToCloudProjectFailoverIpAttachOutput() CloudProjectFailoverIpAttachOutput
	ToCloudProjectFailoverIpAttachOutputWithContext(ctx context.Context) CloudProjectFailoverIpAttachOutput
}

func (*CloudProjectFailoverIpAttach) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectFailoverIpAttach)(nil)).Elem()
}

func (i *CloudProjectFailoverIpAttach) ToCloudProjectFailoverIpAttachOutput() CloudProjectFailoverIpAttachOutput {
	return i.ToCloudProjectFailoverIpAttachOutputWithContext(context.Background())
}

func (i *CloudProjectFailoverIpAttach) ToCloudProjectFailoverIpAttachOutputWithContext(ctx context.Context) CloudProjectFailoverIpAttachOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectFailoverIpAttachOutput)
}

// CloudProjectFailoverIpAttachArrayInput is an input type that accepts CloudProjectFailoverIpAttachArray and CloudProjectFailoverIpAttachArrayOutput values.
// You can construct a concrete instance of `CloudProjectFailoverIpAttachArrayInput` via:
//
//	CloudProjectFailoverIpAttachArray{ CloudProjectFailoverIpAttachArgs{...} }
type CloudProjectFailoverIpAttachArrayInput interface {
	pulumi.Input

	ToCloudProjectFailoverIpAttachArrayOutput() CloudProjectFailoverIpAttachArrayOutput
	ToCloudProjectFailoverIpAttachArrayOutputWithContext(context.Context) CloudProjectFailoverIpAttachArrayOutput
}

type CloudProjectFailoverIpAttachArray []CloudProjectFailoverIpAttachInput

func (CloudProjectFailoverIpAttachArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectFailoverIpAttach)(nil)).Elem()
}

func (i CloudProjectFailoverIpAttachArray) ToCloudProjectFailoverIpAttachArrayOutput() CloudProjectFailoverIpAttachArrayOutput {
	return i.ToCloudProjectFailoverIpAttachArrayOutputWithContext(context.Background())
}

func (i CloudProjectFailoverIpAttachArray) ToCloudProjectFailoverIpAttachArrayOutputWithContext(ctx context.Context) CloudProjectFailoverIpAttachArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectFailoverIpAttachArrayOutput)
}

// CloudProjectFailoverIpAttachMapInput is an input type that accepts CloudProjectFailoverIpAttachMap and CloudProjectFailoverIpAttachMapOutput values.
// You can construct a concrete instance of `CloudProjectFailoverIpAttachMapInput` via:
//
//	CloudProjectFailoverIpAttachMap{ "key": CloudProjectFailoverIpAttachArgs{...} }
type CloudProjectFailoverIpAttachMapInput interface {
	pulumi.Input

	ToCloudProjectFailoverIpAttachMapOutput() CloudProjectFailoverIpAttachMapOutput
	ToCloudProjectFailoverIpAttachMapOutputWithContext(context.Context) CloudProjectFailoverIpAttachMapOutput
}

type CloudProjectFailoverIpAttachMap map[string]CloudProjectFailoverIpAttachInput

func (CloudProjectFailoverIpAttachMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectFailoverIpAttach)(nil)).Elem()
}

func (i CloudProjectFailoverIpAttachMap) ToCloudProjectFailoverIpAttachMapOutput() CloudProjectFailoverIpAttachMapOutput {
	return i.ToCloudProjectFailoverIpAttachMapOutputWithContext(context.Background())
}

func (i CloudProjectFailoverIpAttachMap) ToCloudProjectFailoverIpAttachMapOutputWithContext(ctx context.Context) CloudProjectFailoverIpAttachMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectFailoverIpAttachMapOutput)
}

type CloudProjectFailoverIpAttachOutput struct{ *pulumi.OutputState }

func (CloudProjectFailoverIpAttachOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectFailoverIpAttach)(nil)).Elem()
}

func (o CloudProjectFailoverIpAttachOutput) ToCloudProjectFailoverIpAttachOutput() CloudProjectFailoverIpAttachOutput {
	return o
}

func (o CloudProjectFailoverIpAttachOutput) ToCloudProjectFailoverIpAttachOutputWithContext(ctx context.Context) CloudProjectFailoverIpAttachOutput {
	return o
}

// The IP block
func (o CloudProjectFailoverIpAttachOutput) Block() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.StringOutput { return v.Block }).(pulumi.StringOutput)
}

// Ip continent
func (o CloudProjectFailoverIpAttachOutput) ContinentCode() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.StringOutput { return v.ContinentCode }).(pulumi.StringOutput)
}

// Ip location
func (o CloudProjectFailoverIpAttachOutput) GeoLoc() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.StringOutput { return v.GeoLoc }).(pulumi.StringOutput)
}

// The failover ip address to attach
func (o CloudProjectFailoverIpAttachOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Current operation progress in percent
func (o CloudProjectFailoverIpAttachOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.IntOutput { return v.Progress }).(pulumi.IntOutput)
}

// The GUID of an instance to which the failover IP address is be attached
func (o CloudProjectFailoverIpAttachOutput) RoutedTo() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.StringOutput { return v.RoutedTo }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectFailoverIpAttachOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Ip status, can be `ok` or `operationPending`
func (o CloudProjectFailoverIpAttachOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// IP sub type
func (o CloudProjectFailoverIpAttachOutput) SubType() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectFailoverIpAttach) pulumi.StringOutput { return v.SubType }).(pulumi.StringOutput)
}

type CloudProjectFailoverIpAttachArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectFailoverIpAttachArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectFailoverIpAttach)(nil)).Elem()
}

func (o CloudProjectFailoverIpAttachArrayOutput) ToCloudProjectFailoverIpAttachArrayOutput() CloudProjectFailoverIpAttachArrayOutput {
	return o
}

func (o CloudProjectFailoverIpAttachArrayOutput) ToCloudProjectFailoverIpAttachArrayOutputWithContext(ctx context.Context) CloudProjectFailoverIpAttachArrayOutput {
	return o
}

func (o CloudProjectFailoverIpAttachArrayOutput) Index(i pulumi.IntInput) CloudProjectFailoverIpAttachOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProjectFailoverIpAttach {
		return vs[0].([]*CloudProjectFailoverIpAttach)[vs[1].(int)]
	}).(CloudProjectFailoverIpAttachOutput)
}

type CloudProjectFailoverIpAttachMapOutput struct{ *pulumi.OutputState }

func (CloudProjectFailoverIpAttachMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectFailoverIpAttach)(nil)).Elem()
}

func (o CloudProjectFailoverIpAttachMapOutput) ToCloudProjectFailoverIpAttachMapOutput() CloudProjectFailoverIpAttachMapOutput {
	return o
}

func (o CloudProjectFailoverIpAttachMapOutput) ToCloudProjectFailoverIpAttachMapOutputWithContext(ctx context.Context) CloudProjectFailoverIpAttachMapOutput {
	return o
}

func (o CloudProjectFailoverIpAttachMapOutput) MapIndex(k pulumi.StringInput) CloudProjectFailoverIpAttachOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProjectFailoverIpAttach {
		return vs[0].(map[string]*CloudProjectFailoverIpAttach)[vs[1].(string)]
	}).(CloudProjectFailoverIpAttachOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectFailoverIpAttachInput)(nil)).Elem(), &CloudProjectFailoverIpAttach{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectFailoverIpAttachArrayInput)(nil)).Elem(), CloudProjectFailoverIpAttachArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectFailoverIpAttachMapInput)(nil)).Elem(), CloudProjectFailoverIpAttachMap{})
	pulumi.RegisterOutputType(CloudProjectFailoverIpAttachOutput{})
	pulumi.RegisterOutputType(CloudProjectFailoverIpAttachArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectFailoverIpAttachMapOutput{})
}
