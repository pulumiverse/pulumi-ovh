// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an IPXE Script.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Me.GetIpxeScript(ctx, &me.GetIpxeScriptArgs{
//				Name: "myscript",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIpxeScript(ctx *pulumi.Context, args *LookupIpxeScriptArgs, opts ...pulumi.InvokeOption) (*LookupIpxeScriptResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpxeScriptResult
	err := ctx.Invoke("ovh:Me/getIpxeScript:getIpxeScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpxeScript.
type LookupIpxeScriptArgs struct {
	// The name of the IPXE Script.
	Name string `pulumi:"name"`
}

// A collection of values returned by getIpxeScript.
type LookupIpxeScriptResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// The content of the script.
	Script string `pulumi:"script"`
}

func LookupIpxeScriptOutput(ctx *pulumi.Context, args LookupIpxeScriptOutputArgs, opts ...pulumi.InvokeOption) LookupIpxeScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpxeScriptResult, error) {
			args := v.(LookupIpxeScriptArgs)
			r, err := LookupIpxeScript(ctx, &args, opts...)
			var s LookupIpxeScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpxeScriptResultOutput)
}

// A collection of arguments for invoking getIpxeScript.
type LookupIpxeScriptOutputArgs struct {
	// The name of the IPXE Script.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupIpxeScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpxeScriptArgs)(nil)).Elem()
}

// A collection of values returned by getIpxeScript.
type LookupIpxeScriptResultOutput struct{ *pulumi.OutputState }

func (LookupIpxeScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpxeScriptResult)(nil)).Elem()
}

func (o LookupIpxeScriptResultOutput) ToLookupIpxeScriptResultOutput() LookupIpxeScriptResultOutput {
	return o
}

func (o LookupIpxeScriptResultOutput) ToLookupIpxeScriptResultOutputWithContext(ctx context.Context) LookupIpxeScriptResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIpxeScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpxeScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupIpxeScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpxeScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

// The content of the script.
func (o LookupIpxeScriptResultOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpxeScriptResult) string { return v.Script }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpxeScriptResultOutput{})
}
