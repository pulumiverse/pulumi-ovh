// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package order

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information of order cart product options.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Order"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mycart, err := Order.GetCart(ctx, &order.GetCartArgs{
//				OvhSubsidiary: "fr",
//				Description:   pulumi.StringRef("my cart"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Order.GetCartProductOptions(ctx, &order.GetCartProductOptionsArgs{
//				CartId:   mycart.Id,
//				Product:  "cloud",
//				PlanCode: "project",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCartProductOptions(ctx *pulumi.Context, args *LookupCartProductOptionsArgs, opts ...pulumi.InvokeOption) (*LookupCartProductOptionsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCartProductOptionsResult
	err := ctx.Invoke("ovh:Order/getCartProductOptions:getCartProductOptions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCartProductOptions.
type LookupCartProductOptionsArgs struct {
	// Cart identifier
	CartId string `pulumi:"cartId"`
	// Catalog name
	CatalogName *string `pulumi:"catalogName"`
	// Product offer identifier
	PlanCode string `pulumi:"planCode"`
	// Product
	Product string `pulumi:"product"`
}

// A collection of values returned by getCartProductOptions.
type LookupCartProductOptionsResult struct {
	CartId      string  `pulumi:"cartId"`
	CatalogName *string `pulumi:"catalogName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Product offer identifier
	PlanCode string `pulumi:"planCode"`
	Product  string `pulumi:"product"`
	// products results
	Results []GetCartProductOptionsResult `pulumi:"results"`
}

func LookupCartProductOptionsOutput(ctx *pulumi.Context, args LookupCartProductOptionsOutputArgs, opts ...pulumi.InvokeOption) LookupCartProductOptionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCartProductOptionsResult, error) {
			args := v.(LookupCartProductOptionsArgs)
			r, err := LookupCartProductOptions(ctx, &args, opts...)
			var s LookupCartProductOptionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCartProductOptionsResultOutput)
}

// A collection of arguments for invoking getCartProductOptions.
type LookupCartProductOptionsOutputArgs struct {
	// Cart identifier
	CartId pulumi.StringInput `pulumi:"cartId"`
	// Catalog name
	CatalogName pulumi.StringPtrInput `pulumi:"catalogName"`
	// Product offer identifier
	PlanCode pulumi.StringInput `pulumi:"planCode"`
	// Product
	Product pulumi.StringInput `pulumi:"product"`
}

func (LookupCartProductOptionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCartProductOptionsArgs)(nil)).Elem()
}

// A collection of values returned by getCartProductOptions.
type LookupCartProductOptionsResultOutput struct{ *pulumi.OutputState }

func (LookupCartProductOptionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCartProductOptionsResult)(nil)).Elem()
}

func (o LookupCartProductOptionsResultOutput) ToLookupCartProductOptionsResultOutput() LookupCartProductOptionsResultOutput {
	return o
}

func (o LookupCartProductOptionsResultOutput) ToLookupCartProductOptionsResultOutputWithContext(ctx context.Context) LookupCartProductOptionsResultOutput {
	return o
}

func (o LookupCartProductOptionsResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCartProductOptionsResult) string { return v.CartId }).(pulumi.StringOutput)
}

func (o LookupCartProductOptionsResultOutput) CatalogName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCartProductOptionsResult) *string { return v.CatalogName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCartProductOptionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCartProductOptionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Product offer identifier
func (o LookupCartProductOptionsResultOutput) PlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCartProductOptionsResult) string { return v.PlanCode }).(pulumi.StringOutput)
}

func (o LookupCartProductOptionsResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCartProductOptionsResult) string { return v.Product }).(pulumi.StringOutput)
}

// products results
func (o LookupCartProductOptionsResultOutput) Results() GetCartProductOptionsResultArrayOutput {
	return o.ApplyT(func(v LookupCartProductOptionsResult) []GetCartProductOptionsResult { return v.Results }).(GetCartProductOptionsResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCartProductOptionsResultOutput{})
}
