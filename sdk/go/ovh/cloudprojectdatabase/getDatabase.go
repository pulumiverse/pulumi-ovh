// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the managed database of a public cloud project.
//
// ## Example Usage
//
// To get information of a database cluster service:
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			db, err := CloudProjectDatabase.GetDatabase(ctx, &cloudprojectdatabase.GetDatabaseArgs{
//				ServiceName: "XXXXXX",
//				Engine:      "YYYY",
//				Id:          "ZZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("clusterId", db.Id)
//			return nil
//		})
//	}
//
// ```
func GetDatabase(ctx *pulumi.Context, args *GetDatabaseArgs, opts ...pulumi.InvokeOption) (*GetDatabaseResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatabaseResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabase:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabase.
type GetDatabaseArgs struct {
	// The database engine you want to get information. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// Cluster ID
	Id string `pulumi:"id"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabase.
type GetDatabaseResult struct {
	// Advanced configuration key / value.
	AdvancedConfiguration map[string]string `pulumi:"advancedConfiguration"`
	// Time on which backups start every day.
	BackupTime string `pulumi:"backupTime"`
	// Date of the creation of the cluster.
	CreatedAt string `pulumi:"createdAt"`
	// Small description of the database service.
	Description string `pulumi:"description"`
	// The disk size (in GB) of the database service.
	DiskSize int `pulumi:"diskSize"`
	// The disk type of the database service.
	DiskType string `pulumi:"diskType"`
	// List of all endpoints objects of the service.
	Endpoints []GetDatabaseEndpoint `pulumi:"endpoints"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// A valid OVHcloud public cloud database flavor name in which the nodes will be started.
	Flavor string `pulumi:"flavor"`
	// See Argument Reference above.
	Id string `pulumi:"id"`
	// Defines whether the REST API is enabled on a kafka cluster.
	KafkaRestApi bool `pulumi:"kafkaRestApi"`
	// Time on which maintenances can start every day.
	MaintenanceTime string `pulumi:"maintenanceTime"`
	// Type of network of the cluster.
	NetworkType string `pulumi:"networkType"`
	// List of nodes object.
	Nodes                 []GetDatabaseNode `pulumi:"nodes"`
	OpensearchAclsEnabled bool              `pulumi:"opensearchAclsEnabled"`
	// Plan of the cluster.
	Plan string `pulumi:"plan"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// Current status of the cluster.
	Status string `pulumi:"status"`
	// The version of the engine in which the service should be deployed
	Version string `pulumi:"version"`
}

func GetDatabaseOutput(ctx *pulumi.Context, args GetDatabaseOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseResult, error) {
			args := v.(GetDatabaseArgs)
			r, err := GetDatabase(ctx, &args, opts...)
			var s GetDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabaseResultOutput)
}

// A collection of arguments for invoking getDatabase.
type GetDatabaseOutputArgs struct {
	// The database engine you want to get information. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput `pulumi:"engine"`
	// Cluster ID
	Id pulumi.StringInput `pulumi:"id"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getDatabase.
type GetDatabaseResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseResult)(nil)).Elem()
}

func (o GetDatabaseResultOutput) ToGetDatabaseResultOutput() GetDatabaseResultOutput {
	return o
}

func (o GetDatabaseResultOutput) ToGetDatabaseResultOutputWithContext(ctx context.Context) GetDatabaseResultOutput {
	return o
}

// Advanced configuration key / value.
func (o GetDatabaseResultOutput) AdvancedConfiguration() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]string { return v.AdvancedConfiguration }).(pulumi.StringMapOutput)
}

// Time on which backups start every day.
func (o GetDatabaseResultOutput) BackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.BackupTime }).(pulumi.StringOutput)
}

// Date of the creation of the cluster.
func (o GetDatabaseResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Small description of the database service.
func (o GetDatabaseResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Description }).(pulumi.StringOutput)
}

// The disk size (in GB) of the database service.
func (o GetDatabaseResultOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetDatabaseResult) int { return v.DiskSize }).(pulumi.IntOutput)
}

// The disk type of the database service.
func (o GetDatabaseResultOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.DiskType }).(pulumi.StringOutput)
}

// List of all endpoints objects of the service.
func (o GetDatabaseResultOutput) Endpoints() GetDatabaseEndpointArrayOutput {
	return o.ApplyT(func(v GetDatabaseResult) []GetDatabaseEndpoint { return v.Endpoints }).(GetDatabaseEndpointArrayOutput)
}

// See Argument Reference above.
func (o GetDatabaseResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Engine }).(pulumi.StringOutput)
}

// A valid OVHcloud public cloud database flavor name in which the nodes will be started.
func (o GetDatabaseResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Flavor }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// Defines whether the REST API is enabled on a kafka cluster.
func (o GetDatabaseResultOutput) KafkaRestApi() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDatabaseResult) bool { return v.KafkaRestApi }).(pulumi.BoolOutput)
}

// Time on which maintenances can start every day.
func (o GetDatabaseResultOutput) MaintenanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.MaintenanceTime }).(pulumi.StringOutput)
}

// Type of network of the cluster.
func (o GetDatabaseResultOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.NetworkType }).(pulumi.StringOutput)
}

// List of nodes object.
func (o GetDatabaseResultOutput) Nodes() GetDatabaseNodeArrayOutput {
	return o.ApplyT(func(v GetDatabaseResult) []GetDatabaseNode { return v.Nodes }).(GetDatabaseNodeArrayOutput)
}

func (o GetDatabaseResultOutput) OpensearchAclsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDatabaseResult) bool { return v.OpensearchAclsEnabled }).(pulumi.BoolOutput)
}

// Plan of the cluster.
func (o GetDatabaseResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Plan }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the cluster.
func (o GetDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

// The version of the engine in which the service should be deployed
func (o GetDatabaseResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseResultOutput{})
}
