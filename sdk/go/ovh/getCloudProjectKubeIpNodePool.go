// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a OVH Managed Kubernetes node pool.
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
//			nodepool, err := ovh.GetCloudProjectKubeIpNodePool(ctx, &GetCloudProjectKubeIpNodePoolArgs{
//				ServiceName: "XXXXXX",
//				KubeId:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
//				Name:        "xxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("maxNodes", nodepool.MaxNodes)
//			return nil
//		})
//	}
//
// ```
func GetCloudProjectKubeIpNodePool(ctx *pulumi.Context, args *GetCloudProjectKubeIpNodePoolArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectKubeIpNodePoolResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCloudProjectKubeIpNodePoolResult
	err := ctx.Invoke("ovh:index/getCloudProjectKubeIpNodePool:getCloudProjectKubeIpNodePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectKubeIpNodePool.
type GetCloudProjectKubeIpNodePoolArgs struct {
	// The id of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes *int `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes *int `pulumi:"minNodes"`
	// The name of the node pool.
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectKubeIpNodePool.
type GetCloudProjectKubeIpNodePoolResult struct {
	// (Optional) should the pool use the anti-affinity feature. Default to `false`.
	AntiAffinity bool `pulumi:"antiAffinity"`
	// (Optional) Enable auto-scaling for the pool. Default to `false`.
	Autoscale bool `pulumi:"autoscale"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes int `pulumi:"availableNodes"`
	// Creation date
	CreatedAt string `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes int `pulumi:"currentNodes"`
	// Number of nodes you desire in the pool
	DesiredNodes int `pulumi:"desiredNodes"`
	// Flavor name
	Flavor string `pulumi:"flavor"`
	// a valid OVH public cloud flavor ID in which the nodes will be started.
	// Ex: "b2-7". Changing this value recreates the resource.
	// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
	FlavorName string `pulumi:"flavorName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	KubeId string `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes int `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes int `pulumi:"minNodes"`
	// (Optional) should the nodes be billed on a monthly basis. Default to `false`.
	MonthlyBilled bool `pulumi:"monthlyBilled"`
	// (Optional) The name of the nodepool.
	// Changing this value recreates the resource.
	// Warning: "_" char is not allowed!
	Name string `pulumi:"name"`
	// Project id
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus string `pulumi:"sizeStatus"`
	// Current status
	Status string `pulumi:"status"`
	// Number of nodes with the latest version installed in the pool
	UpToDateNodes int `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetCloudProjectKubeIpNodePoolOutput(ctx *pulumi.Context, args GetCloudProjectKubeIpNodePoolOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectKubeIpNodePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectKubeIpNodePoolResult, error) {
			args := v.(GetCloudProjectKubeIpNodePoolArgs)
			r, err := GetCloudProjectKubeIpNodePool(ctx, &args, opts...)
			var s GetCloudProjectKubeIpNodePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectKubeIpNodePoolResultOutput)
}

// A collection of arguments for invoking getCloudProjectKubeIpNodePool.
type GetCloudProjectKubeIpNodePoolOutputArgs struct {
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringInput `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes pulumi.IntPtrInput `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes pulumi.IntPtrInput `pulumi:"minNodes"`
	// The name of the node pool.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetCloudProjectKubeIpNodePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectKubeIpNodePoolArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectKubeIpNodePool.
type GetCloudProjectKubeIpNodePoolResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectKubeIpNodePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectKubeIpNodePoolResult)(nil)).Elem()
}

func (o GetCloudProjectKubeIpNodePoolResultOutput) ToGetCloudProjectKubeIpNodePoolResultOutput() GetCloudProjectKubeIpNodePoolResultOutput {
	return o
}

func (o GetCloudProjectKubeIpNodePoolResultOutput) ToGetCloudProjectKubeIpNodePoolResultOutputWithContext(ctx context.Context) GetCloudProjectKubeIpNodePoolResultOutput {
	return o
}

// (Optional) should the pool use the anti-affinity feature. Default to `false`.
func (o GetCloudProjectKubeIpNodePoolResultOutput) AntiAffinity() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) bool { return v.AntiAffinity }).(pulumi.BoolOutput)
}

// (Optional) Enable auto-scaling for the pool. Default to `false`.
func (o GetCloudProjectKubeIpNodePoolResultOutput) Autoscale() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) bool { return v.Autoscale }).(pulumi.BoolOutput)
}

// Number of nodes which are actually ready in the pool
func (o GetCloudProjectKubeIpNodePoolResultOutput) AvailableNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) int { return v.AvailableNodes }).(pulumi.IntOutput)
}

// Creation date
func (o GetCloudProjectKubeIpNodePoolResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Number of nodes present in the pool
func (o GetCloudProjectKubeIpNodePoolResultOutput) CurrentNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) int { return v.CurrentNodes }).(pulumi.IntOutput)
}

// Number of nodes you desire in the pool
func (o GetCloudProjectKubeIpNodePoolResultOutput) DesiredNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) int { return v.DesiredNodes }).(pulumi.IntOutput)
}

// Flavor name
func (o GetCloudProjectKubeIpNodePoolResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.Flavor }).(pulumi.StringOutput)
}

// a valid OVH public cloud flavor ID in which the nodes will be started.
// Ex: "b2-7". Changing this value recreates the resource.
// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
func (o GetCloudProjectKubeIpNodePoolResultOutput) FlavorName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.FlavorName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloudProjectKubeIpNodePoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetCloudProjectKubeIpNodePoolResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.KubeId }).(pulumi.StringOutput)
}

// maximum number of nodes allowed in the pool.
// Setting `desiredNodes` over this value will raise an error.
func (o GetCloudProjectKubeIpNodePoolResultOutput) MaxNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) int { return v.MaxNodes }).(pulumi.IntOutput)
}

// minimum number of nodes allowed in the pool.
// Setting `desiredNodes` under this value will raise an error.
func (o GetCloudProjectKubeIpNodePoolResultOutput) MinNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) int { return v.MinNodes }).(pulumi.IntOutput)
}

// (Optional) should the nodes be billed on a monthly basis. Default to `false`.
func (o GetCloudProjectKubeIpNodePoolResultOutput) MonthlyBilled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) bool { return v.MonthlyBilled }).(pulumi.BoolOutput)
}

// (Optional) The name of the nodepool.
// Changing this value recreates the resource.
// Warning: "_" char is not allowed!
func (o GetCloudProjectKubeIpNodePoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project id
func (o GetCloudProjectKubeIpNodePoolResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetCloudProjectKubeIpNodePoolResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Status describing the state between number of nodes wanted and available ones
func (o GetCloudProjectKubeIpNodePoolResultOutput) SizeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.SizeStatus }).(pulumi.StringOutput)
}

// Current status
func (o GetCloudProjectKubeIpNodePoolResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.Status }).(pulumi.StringOutput)
}

// Number of nodes with the latest version installed in the pool
func (o GetCloudProjectKubeIpNodePoolResultOutput) UpToDateNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) int { return v.UpToDateNodes }).(pulumi.IntOutput)
}

// Last update date
func (o GetCloudProjectKubeIpNodePoolResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeIpNodePoolResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectKubeIpNodePoolResultOutput{})
}
