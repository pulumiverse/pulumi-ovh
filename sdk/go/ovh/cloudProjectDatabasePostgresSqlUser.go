// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an user for a postgresql cluster associated with a public cloud project.
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
//			_, err := ovh.LookupCloudProjectDatabase(ctx, &GetCloudProjectDatabaseArgs{
//				ServiceName: "XXXX",
//				Engine:      "postgresql",
//				ClusterId:   "ZZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ovh.NewCloudProjectDatabasePostgresSqlUser(ctx, "user", &ovh.CloudProjectDatabasePostgresSqlUserArgs{
//				ServiceName: pulumi.Any(ovh_cloud_project_database.Postgresql.Service_name),
//				ClusterId:   pulumi.Any(ovh_cloud_project_database.Postgresql.Id),
//				Roles: pulumi.StringArray{
//					pulumi.String("replication"),
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
//
// ## Import
//
// OVHcloud Managed postgresql clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g.,
//
// ```sh
//
//	$ pulumi import ovh:index/cloudProjectDatabasePostgresSqlUser:CloudProjectDatabasePostgresSqlUser my_user <service_name>/<cluster_id>/<id>
//
// ```
type CloudProjectDatabasePostgresSqlUser struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password of the user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Roles the user belongs to. Possible values:
	// * `["replication"]`
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current status of the user.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewCloudProjectDatabasePostgresSqlUser registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectDatabasePostgresSqlUser(ctx *pulumi.Context,
	name string, args *CloudProjectDatabasePostgresSqlUserArgs, opts ...pulumi.ResourceOption) (*CloudProjectDatabasePostgresSqlUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudProjectDatabasePostgresSqlUser
	err := ctx.RegisterResource("ovh:index/cloudProjectDatabasePostgresSqlUser:CloudProjectDatabasePostgresSqlUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectDatabasePostgresSqlUser gets an existing CloudProjectDatabasePostgresSqlUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectDatabasePostgresSqlUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectDatabasePostgresSqlUserState, opts ...pulumi.ResourceOption) (*CloudProjectDatabasePostgresSqlUser, error) {
	var resource CloudProjectDatabasePostgresSqlUser
	err := ctx.ReadResource("ovh:index/cloudProjectDatabasePostgresSqlUser:CloudProjectDatabasePostgresSqlUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectDatabasePostgresSqlUser resources.
type cloudProjectDatabasePostgresSqlUserState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt *string `pulumi:"createdAt"`
	// Name of the user.
	Name *string `pulumi:"name"`
	// Password of the user.
	Password *string `pulumi:"password"`
	// Roles the user belongs to. Possible values:
	// * `["replication"]`
	Roles []string `pulumi:"roles"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current status of the user.
	Status *string `pulumi:"status"`
}

type CloudProjectDatabasePostgresSqlUserState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Date of the creation of the user.
	CreatedAt pulumi.StringPtrInput
	// Name of the user.
	Name pulumi.StringPtrInput
	// Password of the user.
	Password pulumi.StringPtrInput
	// Roles the user belongs to. Possible values:
	// * `["replication"]`
	Roles pulumi.StringArrayInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current status of the user.
	Status pulumi.StringPtrInput
}

func (CloudProjectDatabasePostgresSqlUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectDatabasePostgresSqlUserState)(nil)).Elem()
}

type cloudProjectDatabasePostgresSqlUserArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Name of the user.
	Name *string `pulumi:"name"`
	// Roles the user belongs to. Possible values:
	// * `["replication"]`
	Roles []string `pulumi:"roles"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectDatabasePostgresSqlUser resource.
type CloudProjectDatabasePostgresSqlUserArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Name of the user.
	Name pulumi.StringPtrInput
	// Roles the user belongs to. Possible values:
	// * `["replication"]`
	Roles pulumi.StringArrayInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (CloudProjectDatabasePostgresSqlUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectDatabasePostgresSqlUserArgs)(nil)).Elem()
}

type CloudProjectDatabasePostgresSqlUserInput interface {
	pulumi.Input

	ToCloudProjectDatabasePostgresSqlUserOutput() CloudProjectDatabasePostgresSqlUserOutput
	ToCloudProjectDatabasePostgresSqlUserOutputWithContext(ctx context.Context) CloudProjectDatabasePostgresSqlUserOutput
}

func (*CloudProjectDatabasePostgresSqlUser) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectDatabasePostgresSqlUser)(nil)).Elem()
}

func (i *CloudProjectDatabasePostgresSqlUser) ToCloudProjectDatabasePostgresSqlUserOutput() CloudProjectDatabasePostgresSqlUserOutput {
	return i.ToCloudProjectDatabasePostgresSqlUserOutputWithContext(context.Background())
}

func (i *CloudProjectDatabasePostgresSqlUser) ToCloudProjectDatabasePostgresSqlUserOutputWithContext(ctx context.Context) CloudProjectDatabasePostgresSqlUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectDatabasePostgresSqlUserOutput)
}

// CloudProjectDatabasePostgresSqlUserArrayInput is an input type that accepts CloudProjectDatabasePostgresSqlUserArray and CloudProjectDatabasePostgresSqlUserArrayOutput values.
// You can construct a concrete instance of `CloudProjectDatabasePostgresSqlUserArrayInput` via:
//
//	CloudProjectDatabasePostgresSqlUserArray{ CloudProjectDatabasePostgresSqlUserArgs{...} }
type CloudProjectDatabasePostgresSqlUserArrayInput interface {
	pulumi.Input

	ToCloudProjectDatabasePostgresSqlUserArrayOutput() CloudProjectDatabasePostgresSqlUserArrayOutput
	ToCloudProjectDatabasePostgresSqlUserArrayOutputWithContext(context.Context) CloudProjectDatabasePostgresSqlUserArrayOutput
}

type CloudProjectDatabasePostgresSqlUserArray []CloudProjectDatabasePostgresSqlUserInput

func (CloudProjectDatabasePostgresSqlUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectDatabasePostgresSqlUser)(nil)).Elem()
}

func (i CloudProjectDatabasePostgresSqlUserArray) ToCloudProjectDatabasePostgresSqlUserArrayOutput() CloudProjectDatabasePostgresSqlUserArrayOutput {
	return i.ToCloudProjectDatabasePostgresSqlUserArrayOutputWithContext(context.Background())
}

func (i CloudProjectDatabasePostgresSqlUserArray) ToCloudProjectDatabasePostgresSqlUserArrayOutputWithContext(ctx context.Context) CloudProjectDatabasePostgresSqlUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectDatabasePostgresSqlUserArrayOutput)
}

// CloudProjectDatabasePostgresSqlUserMapInput is an input type that accepts CloudProjectDatabasePostgresSqlUserMap and CloudProjectDatabasePostgresSqlUserMapOutput values.
// You can construct a concrete instance of `CloudProjectDatabasePostgresSqlUserMapInput` via:
//
//	CloudProjectDatabasePostgresSqlUserMap{ "key": CloudProjectDatabasePostgresSqlUserArgs{...} }
type CloudProjectDatabasePostgresSqlUserMapInput interface {
	pulumi.Input

	ToCloudProjectDatabasePostgresSqlUserMapOutput() CloudProjectDatabasePostgresSqlUserMapOutput
	ToCloudProjectDatabasePostgresSqlUserMapOutputWithContext(context.Context) CloudProjectDatabasePostgresSqlUserMapOutput
}

type CloudProjectDatabasePostgresSqlUserMap map[string]CloudProjectDatabasePostgresSqlUserInput

func (CloudProjectDatabasePostgresSqlUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectDatabasePostgresSqlUser)(nil)).Elem()
}

func (i CloudProjectDatabasePostgresSqlUserMap) ToCloudProjectDatabasePostgresSqlUserMapOutput() CloudProjectDatabasePostgresSqlUserMapOutput {
	return i.ToCloudProjectDatabasePostgresSqlUserMapOutputWithContext(context.Background())
}

func (i CloudProjectDatabasePostgresSqlUserMap) ToCloudProjectDatabasePostgresSqlUserMapOutputWithContext(ctx context.Context) CloudProjectDatabasePostgresSqlUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectDatabasePostgresSqlUserMapOutput)
}

type CloudProjectDatabasePostgresSqlUserOutput struct{ *pulumi.OutputState }

func (CloudProjectDatabasePostgresSqlUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectDatabasePostgresSqlUser)(nil)).Elem()
}

func (o CloudProjectDatabasePostgresSqlUserOutput) ToCloudProjectDatabasePostgresSqlUserOutput() CloudProjectDatabasePostgresSqlUserOutput {
	return o
}

func (o CloudProjectDatabasePostgresSqlUserOutput) ToCloudProjectDatabasePostgresSqlUserOutputWithContext(ctx context.Context) CloudProjectDatabasePostgresSqlUserOutput {
	return o
}

// Cluster ID.
func (o CloudProjectDatabasePostgresSqlUserOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabasePostgresSqlUser) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Date of the creation of the user.
func (o CloudProjectDatabasePostgresSqlUserOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabasePostgresSqlUser) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Name of the user.
func (o CloudProjectDatabasePostgresSqlUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabasePostgresSqlUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password of the user.
func (o CloudProjectDatabasePostgresSqlUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabasePostgresSqlUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Roles the user belongs to. Possible values:
// * `["replication"]`
func (o CloudProjectDatabasePostgresSqlUserOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProjectDatabasePostgresSqlUser) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectDatabasePostgresSqlUserOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabasePostgresSqlUser) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o CloudProjectDatabasePostgresSqlUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabasePostgresSqlUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type CloudProjectDatabasePostgresSqlUserArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectDatabasePostgresSqlUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectDatabasePostgresSqlUser)(nil)).Elem()
}

func (o CloudProjectDatabasePostgresSqlUserArrayOutput) ToCloudProjectDatabasePostgresSqlUserArrayOutput() CloudProjectDatabasePostgresSqlUserArrayOutput {
	return o
}

func (o CloudProjectDatabasePostgresSqlUserArrayOutput) ToCloudProjectDatabasePostgresSqlUserArrayOutputWithContext(ctx context.Context) CloudProjectDatabasePostgresSqlUserArrayOutput {
	return o
}

func (o CloudProjectDatabasePostgresSqlUserArrayOutput) Index(i pulumi.IntInput) CloudProjectDatabasePostgresSqlUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProjectDatabasePostgresSqlUser {
		return vs[0].([]*CloudProjectDatabasePostgresSqlUser)[vs[1].(int)]
	}).(CloudProjectDatabasePostgresSqlUserOutput)
}

type CloudProjectDatabasePostgresSqlUserMapOutput struct{ *pulumi.OutputState }

func (CloudProjectDatabasePostgresSqlUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectDatabasePostgresSqlUser)(nil)).Elem()
}

func (o CloudProjectDatabasePostgresSqlUserMapOutput) ToCloudProjectDatabasePostgresSqlUserMapOutput() CloudProjectDatabasePostgresSqlUserMapOutput {
	return o
}

func (o CloudProjectDatabasePostgresSqlUserMapOutput) ToCloudProjectDatabasePostgresSqlUserMapOutputWithContext(ctx context.Context) CloudProjectDatabasePostgresSqlUserMapOutput {
	return o
}

func (o CloudProjectDatabasePostgresSqlUserMapOutput) MapIndex(k pulumi.StringInput) CloudProjectDatabasePostgresSqlUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProjectDatabasePostgresSqlUser {
		return vs[0].(map[string]*CloudProjectDatabasePostgresSqlUser)[vs[1].(string)]
	}).(CloudProjectDatabasePostgresSqlUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectDatabasePostgresSqlUserInput)(nil)).Elem(), &CloudProjectDatabasePostgresSqlUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectDatabasePostgresSqlUserArrayInput)(nil)).Elem(), CloudProjectDatabasePostgresSqlUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectDatabasePostgresSqlUserMapInput)(nil)).Elem(), CloudProjectDatabasePostgresSqlUserMap{})
	pulumi.RegisterOutputType(CloudProjectDatabasePostgresSqlUserOutput{})
	pulumi.RegisterOutputType(CloudProjectDatabasePostgresSqlUserArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectDatabasePostgresSqlUserMapOutput{})
}
