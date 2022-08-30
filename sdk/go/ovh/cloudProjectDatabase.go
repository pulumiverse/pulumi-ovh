// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// To deploy a business PostgreSQL service with two nodes on public network:
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
//			_, err := ovh.NewCloudProjectDatabase(ctx, "postgresql", &ovh.CloudProjectDatabaseArgs{
//				Description: pulumi.String("my-first-postgresql"),
//				Engine:      pulumi.String("postgresql"),
//				Flavor:      pulumi.String("db1-15"),
//				Nodes: CloudProjectDatabaseNodeArray{
//					&CloudProjectDatabaseNodeArgs{
//						Region: pulumi.String("GRA"),
//					},
//					&CloudProjectDatabaseNodeArgs{
//						Region: pulumi.String("GRA"),
//					},
//				},
//				Plan:        pulumi.String("business"),
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Version:     pulumi.String("14"),
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
// To deploy an enterprise MongoDB service with three nodes on private network:
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
//			_, err := ovh.NewCloudProjectDatabase(ctx, "mongodb", &ovh.CloudProjectDatabaseArgs{
//				ServiceName: pulumi.Any(_var.Openstack_infos.Project_id),
//				Description: pulumi.String("my-first-mongodb"),
//				Engine:      pulumi.String("mongodb"),
//				Version:     pulumi.String("5.0"),
//				Plan:        pulumi.String("enterprise"),
//				Nodes: CloudProjectDatabaseNodeArray{
//					&CloudProjectDatabaseNodeArgs{
//						Region:    pulumi.String("SBG"),
//						SubnetId:  pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//						NetworkId: pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//					},
//					&CloudProjectDatabaseNodeArgs{
//						Region:    pulumi.String("SBG"),
//						SubnetId:  pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//						NetworkId: pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//					},
//					&CloudProjectDatabaseNodeArgs{
//						Region:    pulumi.String("SBG"),
//						SubnetId:  pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//						NetworkId: pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//					},
//				},
//				Flavor: pulumi.String("db1-30"),
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
// OVHcloud Managed database clusters can be imported using the `service_name`, `engine`, `id` of the cluster, separated by "/" E.g.,
//
// ```sh
//
//	$ pulumi import ovh:index/cloudProjectDatabase:CloudProjectDatabase my_database_cluster <service_name>/<engine>/<id>
//
// ```
type CloudProjectDatabase struct {
	pulumi.CustomResourceState

	// Time on which backups start every day.
	BackupTime pulumi.StringOutput `pulumi:"backupTime"`
	// Date of the creation of the cluster.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Small description of the database service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of all endpoints objects of the service.
	Endpoints CloudProjectDatabaseEndpointArrayOutput `pulumi:"endpoints"`
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringOutput `pulumi:"engine"`
	// a valid OVH public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// Time on which maintenances can start every day.
	MaintenanceTime pulumi.StringOutput `pulumi:"maintenanceTime"`
	// Type of network of the cluster.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes CloudProjectDatabaseNodeArrayOutput `pulumi:"nodes"`
	// List of nodes object.
	// Enum: "esential", "business", "enterprise".
	Plan pulumi.StringOutput `pulumi:"plan"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current status of the cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// The version of the engine in which the service should be deployed
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewCloudProjectDatabase registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectDatabase(ctx *pulumi.Context,
	name string, args *CloudProjectDatabaseArgs, opts ...pulumi.ResourceOption) (*CloudProjectDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.Flavor == nil {
		return nil, errors.New("invalid value for required argument 'Flavor'")
	}
	if args.Nodes == nil {
		return nil, errors.New("invalid value for required argument 'Nodes'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudProjectDatabase
	err := ctx.RegisterResource("ovh:index/cloudProjectDatabase:CloudProjectDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectDatabase gets an existing CloudProjectDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectDatabaseState, opts ...pulumi.ResourceOption) (*CloudProjectDatabase, error) {
	var resource CloudProjectDatabase
	err := ctx.ReadResource("ovh:index/cloudProjectDatabase:CloudProjectDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectDatabase resources.
type cloudProjectDatabaseState struct {
	// Time on which backups start every day.
	BackupTime *string `pulumi:"backupTime"`
	// Date of the creation of the cluster.
	CreatedAt *string `pulumi:"createdAt"`
	// Small description of the database service.
	Description *string `pulumi:"description"`
	// List of all endpoints objects of the service.
	Endpoints []CloudProjectDatabaseEndpoint `pulumi:"endpoints"`
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine *string `pulumi:"engine"`
	// a valid OVH public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor *string `pulumi:"flavor"`
	// Time on which maintenances can start every day.
	MaintenanceTime *string `pulumi:"maintenanceTime"`
	// Type of network of the cluster.
	NetworkType *string `pulumi:"networkType"`
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes []CloudProjectDatabaseNode `pulumi:"nodes"`
	// List of nodes object.
	// Enum: "esential", "business", "enterprise".
	Plan *string `pulumi:"plan"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current status of the cluster.
	Status *string `pulumi:"status"`
	// The version of the engine in which the service should be deployed
	Version *string `pulumi:"version"`
}

type CloudProjectDatabaseState struct {
	// Time on which backups start every day.
	BackupTime pulumi.StringPtrInput
	// Date of the creation of the cluster.
	CreatedAt pulumi.StringPtrInput
	// Small description of the database service.
	Description pulumi.StringPtrInput
	// List of all endpoints objects of the service.
	Endpoints CloudProjectDatabaseEndpointArrayInput
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringPtrInput
	// a valid OVH public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor pulumi.StringPtrInput
	// Time on which maintenances can start every day.
	MaintenanceTime pulumi.StringPtrInput
	// Type of network of the cluster.
	NetworkType pulumi.StringPtrInput
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes CloudProjectDatabaseNodeArrayInput
	// List of nodes object.
	// Enum: "esential", "business", "enterprise".
	Plan pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current status of the cluster.
	Status pulumi.StringPtrInput
	// The version of the engine in which the service should be deployed
	Version pulumi.StringPtrInput
}

func (CloudProjectDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectDatabaseState)(nil)).Elem()
}

type cloudProjectDatabaseArgs struct {
	// Small description of the database service.
	Description *string `pulumi:"description"`
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// a valid OVH public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor string `pulumi:"flavor"`
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes []CloudProjectDatabaseNode `pulumi:"nodes"`
	// List of nodes object.
	// Enum: "esential", "business", "enterprise".
	Plan string `pulumi:"plan"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// The version of the engine in which the service should be deployed
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a CloudProjectDatabase resource.
type CloudProjectDatabaseArgs struct {
	// Small description of the database service.
	Description pulumi.StringPtrInput
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput
	// a valid OVH public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor pulumi.StringInput
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes CloudProjectDatabaseNodeArrayInput
	// List of nodes object.
	// Enum: "esential", "business", "enterprise".
	Plan pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// The version of the engine in which the service should be deployed
	Version pulumi.StringInput
}

func (CloudProjectDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectDatabaseArgs)(nil)).Elem()
}

type CloudProjectDatabaseInput interface {
	pulumi.Input

	ToCloudProjectDatabaseOutput() CloudProjectDatabaseOutput
	ToCloudProjectDatabaseOutputWithContext(ctx context.Context) CloudProjectDatabaseOutput
}

func (*CloudProjectDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectDatabase)(nil)).Elem()
}

func (i *CloudProjectDatabase) ToCloudProjectDatabaseOutput() CloudProjectDatabaseOutput {
	return i.ToCloudProjectDatabaseOutputWithContext(context.Background())
}

func (i *CloudProjectDatabase) ToCloudProjectDatabaseOutputWithContext(ctx context.Context) CloudProjectDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectDatabaseOutput)
}

// CloudProjectDatabaseArrayInput is an input type that accepts CloudProjectDatabaseArray and CloudProjectDatabaseArrayOutput values.
// You can construct a concrete instance of `CloudProjectDatabaseArrayInput` via:
//
//	CloudProjectDatabaseArray{ CloudProjectDatabaseArgs{...} }
type CloudProjectDatabaseArrayInput interface {
	pulumi.Input

	ToCloudProjectDatabaseArrayOutput() CloudProjectDatabaseArrayOutput
	ToCloudProjectDatabaseArrayOutputWithContext(context.Context) CloudProjectDatabaseArrayOutput
}

type CloudProjectDatabaseArray []CloudProjectDatabaseInput

func (CloudProjectDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectDatabase)(nil)).Elem()
}

func (i CloudProjectDatabaseArray) ToCloudProjectDatabaseArrayOutput() CloudProjectDatabaseArrayOutput {
	return i.ToCloudProjectDatabaseArrayOutputWithContext(context.Background())
}

func (i CloudProjectDatabaseArray) ToCloudProjectDatabaseArrayOutputWithContext(ctx context.Context) CloudProjectDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectDatabaseArrayOutput)
}

// CloudProjectDatabaseMapInput is an input type that accepts CloudProjectDatabaseMap and CloudProjectDatabaseMapOutput values.
// You can construct a concrete instance of `CloudProjectDatabaseMapInput` via:
//
//	CloudProjectDatabaseMap{ "key": CloudProjectDatabaseArgs{...} }
type CloudProjectDatabaseMapInput interface {
	pulumi.Input

	ToCloudProjectDatabaseMapOutput() CloudProjectDatabaseMapOutput
	ToCloudProjectDatabaseMapOutputWithContext(context.Context) CloudProjectDatabaseMapOutput
}

type CloudProjectDatabaseMap map[string]CloudProjectDatabaseInput

func (CloudProjectDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectDatabase)(nil)).Elem()
}

func (i CloudProjectDatabaseMap) ToCloudProjectDatabaseMapOutput() CloudProjectDatabaseMapOutput {
	return i.ToCloudProjectDatabaseMapOutputWithContext(context.Background())
}

func (i CloudProjectDatabaseMap) ToCloudProjectDatabaseMapOutputWithContext(ctx context.Context) CloudProjectDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectDatabaseMapOutput)
}

type CloudProjectDatabaseOutput struct{ *pulumi.OutputState }

func (CloudProjectDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectDatabase)(nil)).Elem()
}

func (o CloudProjectDatabaseOutput) ToCloudProjectDatabaseOutput() CloudProjectDatabaseOutput {
	return o
}

func (o CloudProjectDatabaseOutput) ToCloudProjectDatabaseOutputWithContext(ctx context.Context) CloudProjectDatabaseOutput {
	return o
}

// Time on which backups start every day.
func (o CloudProjectDatabaseOutput) BackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.BackupTime }).(pulumi.StringOutput)
}

// Date of the creation of the cluster.
func (o CloudProjectDatabaseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Small description of the database service.
func (o CloudProjectDatabaseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of all endpoints objects of the service.
func (o CloudProjectDatabaseOutput) Endpoints() CloudProjectDatabaseEndpointArrayOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) CloudProjectDatabaseEndpointArrayOutput { return v.Endpoints }).(CloudProjectDatabaseEndpointArrayOutput)
}

// The database engine you want to deploy. To get a full list of available engine visit.
// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
func (o CloudProjectDatabaseOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// a valid OVH public cloud database flavor name in which the nodes will be started.
// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
func (o CloudProjectDatabaseOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.Flavor }).(pulumi.StringOutput)
}

// Time on which maintenances can start every day.
func (o CloudProjectDatabaseOutput) MaintenanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.MaintenanceTime }).(pulumi.StringOutput)
}

// Type of network of the cluster.
func (o CloudProjectDatabaseOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// List of nodes object.
// Multi region cluster are not yet available, all node should be identical.
func (o CloudProjectDatabaseOutput) Nodes() CloudProjectDatabaseNodeArrayOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) CloudProjectDatabaseNodeArrayOutput { return v.Nodes }).(CloudProjectDatabaseNodeArrayOutput)
}

// List of nodes object.
// Enum: "esential", "business", "enterprise".
func (o CloudProjectDatabaseOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.Plan }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectDatabaseOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the cluster.
func (o CloudProjectDatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The version of the engine in which the service should be deployed
func (o CloudProjectDatabaseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectDatabase) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type CloudProjectDatabaseArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectDatabase)(nil)).Elem()
}

func (o CloudProjectDatabaseArrayOutput) ToCloudProjectDatabaseArrayOutput() CloudProjectDatabaseArrayOutput {
	return o
}

func (o CloudProjectDatabaseArrayOutput) ToCloudProjectDatabaseArrayOutputWithContext(ctx context.Context) CloudProjectDatabaseArrayOutput {
	return o
}

func (o CloudProjectDatabaseArrayOutput) Index(i pulumi.IntInput) CloudProjectDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProjectDatabase {
		return vs[0].([]*CloudProjectDatabase)[vs[1].(int)]
	}).(CloudProjectDatabaseOutput)
}

type CloudProjectDatabaseMapOutput struct{ *pulumi.OutputState }

func (CloudProjectDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectDatabase)(nil)).Elem()
}

func (o CloudProjectDatabaseMapOutput) ToCloudProjectDatabaseMapOutput() CloudProjectDatabaseMapOutput {
	return o
}

func (o CloudProjectDatabaseMapOutput) ToCloudProjectDatabaseMapOutputWithContext(ctx context.Context) CloudProjectDatabaseMapOutput {
	return o
}

func (o CloudProjectDatabaseMapOutput) MapIndex(k pulumi.StringInput) CloudProjectDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProjectDatabase {
		return vs[0].(map[string]*CloudProjectDatabase)[vs[1].(string)]
	}).(CloudProjectDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectDatabaseInput)(nil)).Elem(), &CloudProjectDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectDatabaseArrayInput)(nil)).Elem(), CloudProjectDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectDatabaseMapInput)(nil)).Elem(), CloudProjectDatabaseMap{})
	pulumi.RegisterOutputType(CloudProjectDatabaseOutput{})
	pulumi.RegisterOutputType(CloudProjectDatabaseArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectDatabaseMapOutput{})
}
