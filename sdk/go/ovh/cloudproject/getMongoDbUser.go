// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a user of a mongodb cluster associated with a public cloud project.
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
//			mongouser, err := CloudProject.GetMongoDbUser(ctx, &cloudproject.GetMongoDbUserArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//				Name:        "ZZZ@admin",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("mongouserRoles", mongouser.Roles)
//			return nil
//		})
//	}
//
// ```
func GetMongoDbUser(ctx *pulumi.Context, args *GetMongoDbUserArgs, opts ...pulumi.InvokeOption) (*GetMongoDbUserResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMongoDbUserResult
	err := ctx.Invoke("ovh:CloudProject/getMongoDbUser:getMongoDbUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMongoDbUser.
type GetMongoDbUserArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// Name of the user with the authentication database in the format name@authDB, for example: johndoe@admin
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getMongoDbUser.
type GetMongoDbUserResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// Roles the user belongs to
	Roles []string `pulumi:"roles"`
	// Current status of the user.
	ServiceName string `pulumi:"serviceName"`
	// Current status of the user.
	Status string `pulumi:"status"`
}

func GetMongoDbUserOutput(ctx *pulumi.Context, args GetMongoDbUserOutputArgs, opts ...pulumi.InvokeOption) GetMongoDbUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMongoDbUserResult, error) {
			args := v.(GetMongoDbUserArgs)
			r, err := GetMongoDbUser(ctx, &args, opts...)
			var s GetMongoDbUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMongoDbUserResultOutput)
}

// A collection of arguments for invoking getMongoDbUser.
type GetMongoDbUserOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Name of the user with the authentication database in the format name@authDB, for example: johndoe@admin
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetMongoDbUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMongoDbUserArgs)(nil)).Elem()
}

// A collection of values returned by getMongoDbUser.
type GetMongoDbUserResultOutput struct{ *pulumi.OutputState }

func (GetMongoDbUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMongoDbUserResult)(nil)).Elem()
}

func (o GetMongoDbUserResultOutput) ToGetMongoDbUserResultOutput() GetMongoDbUserResultOutput {
	return o
}

func (o GetMongoDbUserResultOutput) ToGetMongoDbUserResultOutputWithContext(ctx context.Context) GetMongoDbUserResultOutput {
	return o
}

// See Argument Reference above.
func (o GetMongoDbUserResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoDbUserResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Date of the creation of the user.
func (o GetMongoDbUserResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoDbUserResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMongoDbUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoDbUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetMongoDbUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoDbUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Roles the user belongs to
func (o GetMongoDbUserResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMongoDbUserResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// Current status of the user.
func (o GetMongoDbUserResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoDbUserResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o GetMongoDbUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoDbUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMongoDbUserResultOutput{})
}