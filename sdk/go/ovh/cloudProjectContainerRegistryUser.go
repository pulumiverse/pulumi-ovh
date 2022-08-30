// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a user for a container registry associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-ovh/sdk/go/ovh"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ovh.LookupCloudProjectContainerRegistry(ctx, &GetCloudProjectContainerRegistryArgs{
//				ServiceName: "XXXXXX",
//				RegistryId:  "yyyy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ovh.NewCloudProjectContainerRegistryUser(ctx, "user", &ovh.CloudProjectContainerRegistryUserArgs{
//				ServiceName: pulumi.Any(ovh_cloud_project_containerregistry.Registry.Service_name),
//				RegistryId:  pulumi.Any(ovh_cloud_project_containerregistry.Registry.Id),
//				Email:       pulumi.String("foo@bar.com"),
//				Login:       pulumi.String("foobar"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CloudProjectContainerRegistryUser struct {
	pulumi.CustomResourceState

	// User email
	Email pulumi.StringOutput `pulumi:"email"`
	// Registry name
	Login pulumi.StringOutput `pulumi:"login"`
	// (Sensitive) User password
	Password pulumi.StringOutput `pulumi:"password"`
	// Registry ID
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// User name
	User pulumi.StringOutput `pulumi:"user"`
}

// NewCloudProjectContainerRegistryUser registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectContainerRegistryUser(ctx *pulumi.Context,
	name string, args *CloudProjectContainerRegistryUserArgs, opts ...pulumi.ResourceOption) (*CloudProjectContainerRegistryUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudProjectContainerRegistryUser
	err := ctx.RegisterResource("ovh:index/cloudProjectContainerRegistryUser:CloudProjectContainerRegistryUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectContainerRegistryUser gets an existing CloudProjectContainerRegistryUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectContainerRegistryUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectContainerRegistryUserState, opts ...pulumi.ResourceOption) (*CloudProjectContainerRegistryUser, error) {
	var resource CloudProjectContainerRegistryUser
	err := ctx.ReadResource("ovh:index/cloudProjectContainerRegistryUser:CloudProjectContainerRegistryUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectContainerRegistryUser resources.
type cloudProjectContainerRegistryUserState struct {
	// User email
	Email *string `pulumi:"email"`
	// Registry name
	Login *string `pulumi:"login"`
	// (Sensitive) User password
	Password *string `pulumi:"password"`
	// Registry ID
	RegistryId *string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// User name
	User *string `pulumi:"user"`
}

type CloudProjectContainerRegistryUserState struct {
	// User email
	Email pulumi.StringPtrInput
	// Registry name
	Login pulumi.StringPtrInput
	// (Sensitive) User password
	Password pulumi.StringPtrInput
	// Registry ID
	RegistryId pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// User name
	User pulumi.StringPtrInput
}

func (CloudProjectContainerRegistryUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerRegistryUserState)(nil)).Elem()
}

type cloudProjectContainerRegistryUserArgs struct {
	// User email
	Email string `pulumi:"email"`
	// Registry name
	Login string `pulumi:"login"`
	// Registry ID
	RegistryId string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectContainerRegistryUser resource.
type CloudProjectContainerRegistryUserArgs struct {
	// User email
	Email pulumi.StringInput
	// Registry name
	Login pulumi.StringInput
	// Registry ID
	RegistryId pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (CloudProjectContainerRegistryUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerRegistryUserArgs)(nil)).Elem()
}

type CloudProjectContainerRegistryUserInput interface {
	pulumi.Input

	ToCloudProjectContainerRegistryUserOutput() CloudProjectContainerRegistryUserOutput
	ToCloudProjectContainerRegistryUserOutputWithContext(ctx context.Context) CloudProjectContainerRegistryUserOutput
}

func (*CloudProjectContainerRegistryUser) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerRegistryUser)(nil)).Elem()
}

func (i *CloudProjectContainerRegistryUser) ToCloudProjectContainerRegistryUserOutput() CloudProjectContainerRegistryUserOutput {
	return i.ToCloudProjectContainerRegistryUserOutputWithContext(context.Background())
}

func (i *CloudProjectContainerRegistryUser) ToCloudProjectContainerRegistryUserOutputWithContext(ctx context.Context) CloudProjectContainerRegistryUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerRegistryUserOutput)
}

// CloudProjectContainerRegistryUserArrayInput is an input type that accepts CloudProjectContainerRegistryUserArray and CloudProjectContainerRegistryUserArrayOutput values.
// You can construct a concrete instance of `CloudProjectContainerRegistryUserArrayInput` via:
//
//	CloudProjectContainerRegistryUserArray{ CloudProjectContainerRegistryUserArgs{...} }
type CloudProjectContainerRegistryUserArrayInput interface {
	pulumi.Input

	ToCloudProjectContainerRegistryUserArrayOutput() CloudProjectContainerRegistryUserArrayOutput
	ToCloudProjectContainerRegistryUserArrayOutputWithContext(context.Context) CloudProjectContainerRegistryUserArrayOutput
}

type CloudProjectContainerRegistryUserArray []CloudProjectContainerRegistryUserInput

func (CloudProjectContainerRegistryUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectContainerRegistryUser)(nil)).Elem()
}

func (i CloudProjectContainerRegistryUserArray) ToCloudProjectContainerRegistryUserArrayOutput() CloudProjectContainerRegistryUserArrayOutput {
	return i.ToCloudProjectContainerRegistryUserArrayOutputWithContext(context.Background())
}

func (i CloudProjectContainerRegistryUserArray) ToCloudProjectContainerRegistryUserArrayOutputWithContext(ctx context.Context) CloudProjectContainerRegistryUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerRegistryUserArrayOutput)
}

// CloudProjectContainerRegistryUserMapInput is an input type that accepts CloudProjectContainerRegistryUserMap and CloudProjectContainerRegistryUserMapOutput values.
// You can construct a concrete instance of `CloudProjectContainerRegistryUserMapInput` via:
//
//	CloudProjectContainerRegistryUserMap{ "key": CloudProjectContainerRegistryUserArgs{...} }
type CloudProjectContainerRegistryUserMapInput interface {
	pulumi.Input

	ToCloudProjectContainerRegistryUserMapOutput() CloudProjectContainerRegistryUserMapOutput
	ToCloudProjectContainerRegistryUserMapOutputWithContext(context.Context) CloudProjectContainerRegistryUserMapOutput
}

type CloudProjectContainerRegistryUserMap map[string]CloudProjectContainerRegistryUserInput

func (CloudProjectContainerRegistryUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectContainerRegistryUser)(nil)).Elem()
}

func (i CloudProjectContainerRegistryUserMap) ToCloudProjectContainerRegistryUserMapOutput() CloudProjectContainerRegistryUserMapOutput {
	return i.ToCloudProjectContainerRegistryUserMapOutputWithContext(context.Background())
}

func (i CloudProjectContainerRegistryUserMap) ToCloudProjectContainerRegistryUserMapOutputWithContext(ctx context.Context) CloudProjectContainerRegistryUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerRegistryUserMapOutput)
}

type CloudProjectContainerRegistryUserOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerRegistryUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerRegistryUser)(nil)).Elem()
}

func (o CloudProjectContainerRegistryUserOutput) ToCloudProjectContainerRegistryUserOutput() CloudProjectContainerRegistryUserOutput {
	return o
}

func (o CloudProjectContainerRegistryUserOutput) ToCloudProjectContainerRegistryUserOutputWithContext(ctx context.Context) CloudProjectContainerRegistryUserOutput {
	return o
}

// User email
func (o CloudProjectContainerRegistryUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistryUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Registry name
func (o CloudProjectContainerRegistryUserOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistryUser) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

// (Sensitive) User password
func (o CloudProjectContainerRegistryUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistryUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Registry ID
func (o CloudProjectContainerRegistryUserOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistryUser) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectContainerRegistryUserOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistryUser) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// User name
func (o CloudProjectContainerRegistryUserOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistryUser) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type CloudProjectContainerRegistryUserArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerRegistryUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectContainerRegistryUser)(nil)).Elem()
}

func (o CloudProjectContainerRegistryUserArrayOutput) ToCloudProjectContainerRegistryUserArrayOutput() CloudProjectContainerRegistryUserArrayOutput {
	return o
}

func (o CloudProjectContainerRegistryUserArrayOutput) ToCloudProjectContainerRegistryUserArrayOutputWithContext(ctx context.Context) CloudProjectContainerRegistryUserArrayOutput {
	return o
}

func (o CloudProjectContainerRegistryUserArrayOutput) Index(i pulumi.IntInput) CloudProjectContainerRegistryUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProjectContainerRegistryUser {
		return vs[0].([]*CloudProjectContainerRegistryUser)[vs[1].(int)]
	}).(CloudProjectContainerRegistryUserOutput)
}

type CloudProjectContainerRegistryUserMapOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerRegistryUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectContainerRegistryUser)(nil)).Elem()
}

func (o CloudProjectContainerRegistryUserMapOutput) ToCloudProjectContainerRegistryUserMapOutput() CloudProjectContainerRegistryUserMapOutput {
	return o
}

func (o CloudProjectContainerRegistryUserMapOutput) ToCloudProjectContainerRegistryUserMapOutputWithContext(ctx context.Context) CloudProjectContainerRegistryUserMapOutput {
	return o
}

func (o CloudProjectContainerRegistryUserMapOutput) MapIndex(k pulumi.StringInput) CloudProjectContainerRegistryUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProjectContainerRegistryUser {
		return vs[0].(map[string]*CloudProjectContainerRegistryUser)[vs[1].(string)]
	}).(CloudProjectContainerRegistryUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectContainerRegistryUserInput)(nil)).Elem(), &CloudProjectContainerRegistryUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectContainerRegistryUserArrayInput)(nil)).Elem(), CloudProjectContainerRegistryUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectContainerRegistryUserMapInput)(nil)).Elem(), CloudProjectContainerRegistryUserMap{})
	pulumi.RegisterOutputType(CloudProjectContainerRegistryUserOutput{})
	pulumi.RegisterOutputType(CloudProjectContainerRegistryUserArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectContainerRegistryUserMapOutput{})
}
