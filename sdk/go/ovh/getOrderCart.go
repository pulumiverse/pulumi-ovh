// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to create a temporary order cart to retrieve information order cart products.
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
//			_, err := ovh.GetOrderCart(ctx, &GetOrderCartArgs{
//				Description:   pulumi.StringRef("..."),
//				OvhSubsidiary: "fr",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrderCart(ctx *pulumi.Context, args *GetOrderCartArgs, opts ...pulumi.InvokeOption) (*GetOrderCartResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetOrderCartResult
	err := ctx.Invoke("ovh:index/getOrderCart:getOrderCart", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrderCart.
type GetOrderCartArgs struct {
	// Description of your cart
	Description *string `pulumi:"description"`
	// Expiration time (format: 2006-01-02T15:04:05+00:00)
	Expire *string `pulumi:"expire"`
	// Ovh Subsidiary
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
}

// A collection of values returned by getOrderCart.
type GetOrderCartResult struct {
	// Cart identifier
	CartId      string  `pulumi:"cartId"`
	Description *string `pulumi:"description"`
	Expire      string  `pulumi:"expire"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Items of your cart
	Items         []int  `pulumi:"items"`
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Indicates if the cart has already been validated
	ReadOnly bool `pulumi:"readOnly"`
}

func GetOrderCartOutput(ctx *pulumi.Context, args GetOrderCartOutputArgs, opts ...pulumi.InvokeOption) GetOrderCartResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrderCartResult, error) {
			args := v.(GetOrderCartArgs)
			r, err := GetOrderCart(ctx, &args, opts...)
			var s GetOrderCartResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrderCartResultOutput)
}

// A collection of arguments for invoking getOrderCart.
type GetOrderCartOutputArgs struct {
	// Description of your cart
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Expiration time (format: 2006-01-02T15:04:05+00:00)
	Expire pulumi.StringPtrInput `pulumi:"expire"`
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringInput `pulumi:"ovhSubsidiary"`
}

func (GetOrderCartOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrderCartArgs)(nil)).Elem()
}

// A collection of values returned by getOrderCart.
type GetOrderCartResultOutput struct{ *pulumi.OutputState }

func (GetOrderCartResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrderCartResult)(nil)).Elem()
}

func (o GetOrderCartResultOutput) ToGetOrderCartResultOutput() GetOrderCartResultOutput {
	return o
}

func (o GetOrderCartResultOutput) ToGetOrderCartResultOutputWithContext(ctx context.Context) GetOrderCartResultOutput {
	return o
}

// Cart identifier
func (o GetOrderCartResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartResult) string { return v.CartId }).(pulumi.StringOutput)
}

func (o GetOrderCartResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrderCartResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetOrderCartResultOutput) Expire() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartResult) string { return v.Expire }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrderCartResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartResult) string { return v.Id }).(pulumi.StringOutput)
}

// Items of your cart
func (o GetOrderCartResultOutput) Items() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetOrderCartResult) []int { return v.Items }).(pulumi.IntArrayOutput)
}

func (o GetOrderCartResultOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartResult) string { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Indicates if the cart has already been validated
func (o GetOrderCartResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrderCartResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrderCartResultOutput{})
}
