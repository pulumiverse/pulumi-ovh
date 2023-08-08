// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of OVHcloud Managed Kubernetes nodes in a specific node pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			nodesKubeNodePoolNodes, err := CloudProject.GetKubeNodePoolNodes(ctx, &cloudproject.GetKubeNodePoolNodesArgs{
//				ServiceName: "XXXXXX",
//				KubeId:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
//				Name:        "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nodes", nodesKubeNodePoolNodes)
//			return nil
//		})
//	}
//
// ```
func GetKubeNodePoolNodes(ctx *pulumi.Context, args *GetKubeNodePoolNodesArgs, opts ...pulumi.InvokeOption) (*GetKubeNodePoolNodesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubeNodePoolNodesResult
	err := ctx.Invoke("ovh:CloudProject/getKubeNodePoolNodes:getKubeNodePoolNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubeNodePoolNodes.
type GetKubeNodePoolNodesArgs struct {
	// The ID of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// Name of the node pool from which we want the nodes.
	Name string `pulumi:"name"`
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKubeNodePoolNodes.
type GetKubeNodePoolNodesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	KubeId string `pulumi:"kubeId"`
	// Name of the node.
	Name string `pulumi:"name"`
	// List of all nodes composing the kubernetes cluster.
	Nodes []GetKubeNodePoolNodesNode `pulumi:"nodes"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetKubeNodePoolNodesOutput(ctx *pulumi.Context, args GetKubeNodePoolNodesOutputArgs, opts ...pulumi.InvokeOption) GetKubeNodePoolNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubeNodePoolNodesResult, error) {
			args := v.(GetKubeNodePoolNodesArgs)
			r, err := GetKubeNodePoolNodes(ctx, &args, opts...)
			var s GetKubeNodePoolNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubeNodePoolNodesResultOutput)
}

// A collection of arguments for invoking getKubeNodePoolNodes.
type GetKubeNodePoolNodesOutputArgs struct {
	// The ID of the managed kubernetes cluster.
	KubeId pulumi.StringInput `pulumi:"kubeId"`
	// Name of the node pool from which we want the nodes.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetKubeNodePoolNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeNodePoolNodesArgs)(nil)).Elem()
}

// A collection of values returned by getKubeNodePoolNodes.
type GetKubeNodePoolNodesResultOutput struct{ *pulumi.OutputState }

func (GetKubeNodePoolNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeNodePoolNodesResult)(nil)).Elem()
}

func (o GetKubeNodePoolNodesResultOutput) ToGetKubeNodePoolNodesResultOutput() GetKubeNodePoolNodesResultOutput {
	return o
}

func (o GetKubeNodePoolNodesResultOutput) ToGetKubeNodePoolNodesResultOutputWithContext(ctx context.Context) GetKubeNodePoolNodesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetKubeNodePoolNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodePoolNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetKubeNodePoolNodesResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodePoolNodesResult) string { return v.KubeId }).(pulumi.StringOutput)
}

// Name of the node.
func (o GetKubeNodePoolNodesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodePoolNodesResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of all nodes composing the kubernetes cluster.
func (o GetKubeNodePoolNodesResultOutput) Nodes() GetKubeNodePoolNodesNodeArrayOutput {
	return o.ApplyT(func(v GetKubeNodePoolNodesResult) []GetKubeNodePoolNodesNode { return v.Nodes }).(GetKubeNodePoolNodesNodeArrayOutput)
}

// See Argument Reference above.
func (o GetKubeNodePoolNodesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodePoolNodesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubeNodePoolNodesResultOutput{})
}
