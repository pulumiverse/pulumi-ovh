// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a nodepool in a OVH Managed Kubernetes Service cluster.
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
//			_, err := ovh.NewCloudProjectKubeNodePool(ctx, "pool", &ovh.CloudProjectKubeNodePoolArgs{
//				DesiredNodes: pulumi.Int(3),
//				FlavorName:   pulumi.String("b2-7"),
//				KubeId:       pulumi.String("xxxxxxxx-2bf9-xxxx-xxxx-xxxxxxxxxxxx"),
//				MaxNodes:     pulumi.Int(3),
//				MinNodes:     pulumi.Int(3),
//				ServiceName:  pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Template: &CloudProjectKubeNodePoolTemplateArgs{
//					Metadata: &CloudProjectKubeNodePoolTemplateMetadataArgs{
//						Annotations: pulumi.StringMap{
//							"k1": pulumi.String("v1"),
//							"k2": pulumi.String("v2"),
//						},
//						Finalizers: pulumi.StringArray{
//							pulumi.String("F1"),
//							pulumi.String("F2"),
//						},
//						Labels: pulumi.StringMap{
//							"k3": pulumi.String("v3"),
//							"k4": pulumi.String("v4"),
//						},
//					},
//					Spec: &CloudProjectKubeNodePoolTemplateSpecArgs{
//						Taints: pulumi.AnyMapArray{
//							pulumi.AnyMap{
//								"effect": pulumi.Any("PreferNoSchedule"),
//								"key":    pulumi.Any("k"),
//								"value":  pulumi.Any("v"),
//							},
//						},
//						Unschedulable: pulumi.Bool(false),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CloudProjectKubeNodePool struct {
	pulumi.CustomResourceState

	// should the pool use the anti-affinity feature. Default to `false`.
	AntiAffinity pulumi.BoolPtrOutput `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool. Default to `false`.
	Autoscale pulumi.BoolPtrOutput `pulumi:"autoscale"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes pulumi.IntOutput `pulumi:"availableNodes"`
	// Creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes pulumi.IntOutput `pulumi:"currentNodes"`
	// number of nodes to start.
	DesiredNodes pulumi.IntOutput `pulumi:"desiredNodes"`
	// Flavor name
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// a valid OVH public cloud flavor ID in which the nodes will be started.
	// Ex: "b2-7". Changing this value recreates the resource.
	// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
	FlavorName pulumi.StringOutput `pulumi:"flavorName"`
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringOutput `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes pulumi.IntOutput `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes pulumi.IntOutput `pulumi:"minNodes"`
	// should the nodes be billed on a monthly basis. Default to `false`.
	MonthlyBilled pulumi.BoolPtrOutput `pulumi:"monthlyBilled"`
	// The name of the nodepool.
	// Changing this value recreates the resource.
	// Warning: "_" char is not allowed!
	Name pulumi.StringOutput `pulumi:"name"`
	// Project id
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus pulumi.StringOutput `pulumi:"sizeStatus"`
	// Current status
	Status pulumi.StringOutput `pulumi:"status"`
	// Node pool template
	Template CloudProjectKubeNodePoolTemplatePtrOutput `pulumi:"template"`
	// Number of nodes with latest version installed in the pool
	UpToDateNodes pulumi.IntOutput `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewCloudProjectKubeNodePool registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectKubeNodePool(ctx *pulumi.Context,
	name string, args *CloudProjectKubeNodePoolArgs, opts ...pulumi.ResourceOption) (*CloudProjectKubeNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlavorName == nil {
		return nil, errors.New("invalid value for required argument 'FlavorName'")
	}
	if args.KubeId == nil {
		return nil, errors.New("invalid value for required argument 'KubeId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudProjectKubeNodePool
	err := ctx.RegisterResource("ovh:index/cloudProjectKubeNodePool:CloudProjectKubeNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectKubeNodePool gets an existing CloudProjectKubeNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectKubeNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectKubeNodePoolState, opts ...pulumi.ResourceOption) (*CloudProjectKubeNodePool, error) {
	var resource CloudProjectKubeNodePool
	err := ctx.ReadResource("ovh:index/cloudProjectKubeNodePool:CloudProjectKubeNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectKubeNodePool resources.
type cloudProjectKubeNodePoolState struct {
	// should the pool use the anti-affinity feature. Default to `false`.
	AntiAffinity *bool `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool. Default to `false`.
	Autoscale *bool `pulumi:"autoscale"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes *int `pulumi:"availableNodes"`
	// Creation date
	CreatedAt *string `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes *int `pulumi:"currentNodes"`
	// number of nodes to start.
	DesiredNodes *int `pulumi:"desiredNodes"`
	// Flavor name
	Flavor *string `pulumi:"flavor"`
	// a valid OVH public cloud flavor ID in which the nodes will be started.
	// Ex: "b2-7". Changing this value recreates the resource.
	// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
	FlavorName *string `pulumi:"flavorName"`
	// The id of the managed kubernetes cluster.
	KubeId *string `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes *int `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes *int `pulumi:"minNodes"`
	// should the nodes be billed on a monthly basis. Default to `false`.
	MonthlyBilled *bool `pulumi:"monthlyBilled"`
	// The name of the nodepool.
	// Changing this value recreates the resource.
	// Warning: "_" char is not allowed!
	Name *string `pulumi:"name"`
	// Project id
	ProjectId *string `pulumi:"projectId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus *string `pulumi:"sizeStatus"`
	// Current status
	Status *string `pulumi:"status"`
	// Node pool template
	Template *CloudProjectKubeNodePoolTemplate `pulumi:"template"`
	// Number of nodes with latest version installed in the pool
	UpToDateNodes *int `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt *string `pulumi:"updatedAt"`
}

type CloudProjectKubeNodePoolState struct {
	// should the pool use the anti-affinity feature. Default to `false`.
	AntiAffinity pulumi.BoolPtrInput
	// Enable auto-scaling for the pool. Default to `false`.
	Autoscale pulumi.BoolPtrInput
	// Number of nodes which are actually ready in the pool
	AvailableNodes pulumi.IntPtrInput
	// Creation date
	CreatedAt pulumi.StringPtrInput
	// Number of nodes present in the pool
	CurrentNodes pulumi.IntPtrInput
	// number of nodes to start.
	DesiredNodes pulumi.IntPtrInput
	// Flavor name
	Flavor pulumi.StringPtrInput
	// a valid OVH public cloud flavor ID in which the nodes will be started.
	// Ex: "b2-7". Changing this value recreates the resource.
	// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
	FlavorName pulumi.StringPtrInput
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringPtrInput
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes pulumi.IntPtrInput
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes pulumi.IntPtrInput
	// should the nodes be billed on a monthly basis. Default to `false`.
	MonthlyBilled pulumi.BoolPtrInput
	// The name of the nodepool.
	// Changing this value recreates the resource.
	// Warning: "_" char is not allowed!
	Name pulumi.StringPtrInput
	// Project id
	ProjectId pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus pulumi.StringPtrInput
	// Current status
	Status pulumi.StringPtrInput
	// Node pool template
	Template CloudProjectKubeNodePoolTemplatePtrInput
	// Number of nodes with latest version installed in the pool
	UpToDateNodes pulumi.IntPtrInput
	// Last update date
	UpdatedAt pulumi.StringPtrInput
}

func (CloudProjectKubeNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectKubeNodePoolState)(nil)).Elem()
}

type cloudProjectKubeNodePoolArgs struct {
	// should the pool use the anti-affinity feature. Default to `false`.
	AntiAffinity *bool `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool. Default to `false`.
	Autoscale *bool `pulumi:"autoscale"`
	// number of nodes to start.
	DesiredNodes *int `pulumi:"desiredNodes"`
	// a valid OVH public cloud flavor ID in which the nodes will be started.
	// Ex: "b2-7". Changing this value recreates the resource.
	// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
	FlavorName string `pulumi:"flavorName"`
	// The id of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes *int `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes *int `pulumi:"minNodes"`
	// should the nodes be billed on a monthly basis. Default to `false`.
	MonthlyBilled *bool `pulumi:"monthlyBilled"`
	// The name of the nodepool.
	// Changing this value recreates the resource.
	// Warning: "_" char is not allowed!
	Name *string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// Node pool template
	Template *CloudProjectKubeNodePoolTemplate `pulumi:"template"`
}

// The set of arguments for constructing a CloudProjectKubeNodePool resource.
type CloudProjectKubeNodePoolArgs struct {
	// should the pool use the anti-affinity feature. Default to `false`.
	AntiAffinity pulumi.BoolPtrInput
	// Enable auto-scaling for the pool. Default to `false`.
	Autoscale pulumi.BoolPtrInput
	// number of nodes to start.
	DesiredNodes pulumi.IntPtrInput
	// a valid OVH public cloud flavor ID in which the nodes will be started.
	// Ex: "b2-7". Changing this value recreates the resource.
	// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
	FlavorName pulumi.StringInput
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringInput
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes pulumi.IntPtrInput
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes pulumi.IntPtrInput
	// should the nodes be billed on a monthly basis. Default to `false`.
	MonthlyBilled pulumi.BoolPtrInput
	// The name of the nodepool.
	// Changing this value recreates the resource.
	// Warning: "_" char is not allowed!
	Name pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// Node pool template
	Template CloudProjectKubeNodePoolTemplatePtrInput
}

func (CloudProjectKubeNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectKubeNodePoolArgs)(nil)).Elem()
}

type CloudProjectKubeNodePoolInput interface {
	pulumi.Input

	ToCloudProjectKubeNodePoolOutput() CloudProjectKubeNodePoolOutput
	ToCloudProjectKubeNodePoolOutputWithContext(ctx context.Context) CloudProjectKubeNodePoolOutput
}

func (*CloudProjectKubeNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectKubeNodePool)(nil)).Elem()
}

func (i *CloudProjectKubeNodePool) ToCloudProjectKubeNodePoolOutput() CloudProjectKubeNodePoolOutput {
	return i.ToCloudProjectKubeNodePoolOutputWithContext(context.Background())
}

func (i *CloudProjectKubeNodePool) ToCloudProjectKubeNodePoolOutputWithContext(ctx context.Context) CloudProjectKubeNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodePoolOutput)
}

// CloudProjectKubeNodePoolArrayInput is an input type that accepts CloudProjectKubeNodePoolArray and CloudProjectKubeNodePoolArrayOutput values.
// You can construct a concrete instance of `CloudProjectKubeNodePoolArrayInput` via:
//
//	CloudProjectKubeNodePoolArray{ CloudProjectKubeNodePoolArgs{...} }
type CloudProjectKubeNodePoolArrayInput interface {
	pulumi.Input

	ToCloudProjectKubeNodePoolArrayOutput() CloudProjectKubeNodePoolArrayOutput
	ToCloudProjectKubeNodePoolArrayOutputWithContext(context.Context) CloudProjectKubeNodePoolArrayOutput
}

type CloudProjectKubeNodePoolArray []CloudProjectKubeNodePoolInput

func (CloudProjectKubeNodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectKubeNodePool)(nil)).Elem()
}

func (i CloudProjectKubeNodePoolArray) ToCloudProjectKubeNodePoolArrayOutput() CloudProjectKubeNodePoolArrayOutput {
	return i.ToCloudProjectKubeNodePoolArrayOutputWithContext(context.Background())
}

func (i CloudProjectKubeNodePoolArray) ToCloudProjectKubeNodePoolArrayOutputWithContext(ctx context.Context) CloudProjectKubeNodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodePoolArrayOutput)
}

// CloudProjectKubeNodePoolMapInput is an input type that accepts CloudProjectKubeNodePoolMap and CloudProjectKubeNodePoolMapOutput values.
// You can construct a concrete instance of `CloudProjectKubeNodePoolMapInput` via:
//
//	CloudProjectKubeNodePoolMap{ "key": CloudProjectKubeNodePoolArgs{...} }
type CloudProjectKubeNodePoolMapInput interface {
	pulumi.Input

	ToCloudProjectKubeNodePoolMapOutput() CloudProjectKubeNodePoolMapOutput
	ToCloudProjectKubeNodePoolMapOutputWithContext(context.Context) CloudProjectKubeNodePoolMapOutput
}

type CloudProjectKubeNodePoolMap map[string]CloudProjectKubeNodePoolInput

func (CloudProjectKubeNodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectKubeNodePool)(nil)).Elem()
}

func (i CloudProjectKubeNodePoolMap) ToCloudProjectKubeNodePoolMapOutput() CloudProjectKubeNodePoolMapOutput {
	return i.ToCloudProjectKubeNodePoolMapOutputWithContext(context.Background())
}

func (i CloudProjectKubeNodePoolMap) ToCloudProjectKubeNodePoolMapOutputWithContext(ctx context.Context) CloudProjectKubeNodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodePoolMapOutput)
}

type CloudProjectKubeNodePoolOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectKubeNodePool)(nil)).Elem()
}

func (o CloudProjectKubeNodePoolOutput) ToCloudProjectKubeNodePoolOutput() CloudProjectKubeNodePoolOutput {
	return o
}

func (o CloudProjectKubeNodePoolOutput) ToCloudProjectKubeNodePoolOutputWithContext(ctx context.Context) CloudProjectKubeNodePoolOutput {
	return o
}

// should the pool use the anti-affinity feature. Default to `false`.
func (o CloudProjectKubeNodePoolOutput) AntiAffinity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.BoolPtrOutput { return v.AntiAffinity }).(pulumi.BoolPtrOutput)
}

// Enable auto-scaling for the pool. Default to `false`.
func (o CloudProjectKubeNodePoolOutput) Autoscale() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.BoolPtrOutput { return v.Autoscale }).(pulumi.BoolPtrOutput)
}

// Number of nodes which are actually ready in the pool
func (o CloudProjectKubeNodePoolOutput) AvailableNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.IntOutput { return v.AvailableNodes }).(pulumi.IntOutput)
}

// Creation date
func (o CloudProjectKubeNodePoolOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Number of nodes present in the pool
func (o CloudProjectKubeNodePoolOutput) CurrentNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.IntOutput { return v.CurrentNodes }).(pulumi.IntOutput)
}

// number of nodes to start.
func (o CloudProjectKubeNodePoolOutput) DesiredNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.IntOutput { return v.DesiredNodes }).(pulumi.IntOutput)
}

// Flavor name
func (o CloudProjectKubeNodePoolOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.Flavor }).(pulumi.StringOutput)
}

// a valid OVH public cloud flavor ID in which the nodes will be started.
// Ex: "b2-7". Changing this value recreates the resource.
// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
func (o CloudProjectKubeNodePoolOutput) FlavorName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.FlavorName }).(pulumi.StringOutput)
}

// The id of the managed kubernetes cluster.
func (o CloudProjectKubeNodePoolOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.KubeId }).(pulumi.StringOutput)
}

// maximum number of nodes allowed in the pool.
// Setting `desiredNodes` over this value will raise an error.
func (o CloudProjectKubeNodePoolOutput) MaxNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.IntOutput { return v.MaxNodes }).(pulumi.IntOutput)
}

// minimum number of nodes allowed in the pool.
// Setting `desiredNodes` under this value will raise an error.
func (o CloudProjectKubeNodePoolOutput) MinNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.IntOutput { return v.MinNodes }).(pulumi.IntOutput)
}

// should the nodes be billed on a monthly basis. Default to `false`.
func (o CloudProjectKubeNodePoolOutput) MonthlyBilled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.BoolPtrOutput { return v.MonthlyBilled }).(pulumi.BoolPtrOutput)
}

// The name of the nodepool.
// Changing this value recreates the resource.
// Warning: "_" char is not allowed!
func (o CloudProjectKubeNodePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Project id
func (o CloudProjectKubeNodePoolOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectKubeNodePoolOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Status describing the state between number of nodes wanted and available ones
func (o CloudProjectKubeNodePoolOutput) SizeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.SizeStatus }).(pulumi.StringOutput)
}

// Current status
func (o CloudProjectKubeNodePoolOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Node pool template
func (o CloudProjectKubeNodePoolOutput) Template() CloudProjectKubeNodePoolTemplatePtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) CloudProjectKubeNodePoolTemplatePtrOutput { return v.Template }).(CloudProjectKubeNodePoolTemplatePtrOutput)
}

// Number of nodes with latest version installed in the pool
func (o CloudProjectKubeNodePoolOutput) UpToDateNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.IntOutput { return v.UpToDateNodes }).(pulumi.IntOutput)
}

// Last update date
func (o CloudProjectKubeNodePoolOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodePool) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type CloudProjectKubeNodePoolArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeNodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectKubeNodePool)(nil)).Elem()
}

func (o CloudProjectKubeNodePoolArrayOutput) ToCloudProjectKubeNodePoolArrayOutput() CloudProjectKubeNodePoolArrayOutput {
	return o
}

func (o CloudProjectKubeNodePoolArrayOutput) ToCloudProjectKubeNodePoolArrayOutputWithContext(ctx context.Context) CloudProjectKubeNodePoolArrayOutput {
	return o
}

func (o CloudProjectKubeNodePoolArrayOutput) Index(i pulumi.IntInput) CloudProjectKubeNodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProjectKubeNodePool {
		return vs[0].([]*CloudProjectKubeNodePool)[vs[1].(int)]
	}).(CloudProjectKubeNodePoolOutput)
}

type CloudProjectKubeNodePoolMapOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeNodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectKubeNodePool)(nil)).Elem()
}

func (o CloudProjectKubeNodePoolMapOutput) ToCloudProjectKubeNodePoolMapOutput() CloudProjectKubeNodePoolMapOutput {
	return o
}

func (o CloudProjectKubeNodePoolMapOutput) ToCloudProjectKubeNodePoolMapOutputWithContext(ctx context.Context) CloudProjectKubeNodePoolMapOutput {
	return o
}

func (o CloudProjectKubeNodePoolMapOutput) MapIndex(k pulumi.StringInput) CloudProjectKubeNodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProjectKubeNodePool {
		return vs[0].(map[string]*CloudProjectKubeNodePool)[vs[1].(string)]
	}).(CloudProjectKubeNodePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectKubeNodePoolInput)(nil)).Elem(), &CloudProjectKubeNodePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectKubeNodePoolArrayInput)(nil)).Elem(), CloudProjectKubeNodePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectKubeNodePoolMapInput)(nil)).Elem(), CloudProjectKubeNodePoolMap{})
	pulumi.RegisterOutputType(CloudProjectKubeNodePoolOutput{})
	pulumi.RegisterOutputType(CloudProjectKubeNodePoolArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectKubeNodePoolMapOutput{})
}
