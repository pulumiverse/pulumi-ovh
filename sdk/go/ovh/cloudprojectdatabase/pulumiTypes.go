// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OpensearchUserAcl struct {
	// Pattern of the ACL.
	Pattern string `pulumi:"pattern"`
	// Permission of the ACL
	// Available permission:
	// * `admin`
	// * `read`
	// * `write`
	// * `readwrite`
	// * `deny`
	Permission string `pulumi:"permission"`
}

// OpensearchUserAclInput is an input type that accepts OpensearchUserAclArgs and OpensearchUserAclOutput values.
// You can construct a concrete instance of `OpensearchUserAclInput` via:
//
//	OpensearchUserAclArgs{...}
type OpensearchUserAclInput interface {
	pulumi.Input

	ToOpensearchUserAclOutput() OpensearchUserAclOutput
	ToOpensearchUserAclOutputWithContext(context.Context) OpensearchUserAclOutput
}

type OpensearchUserAclArgs struct {
	// Pattern of the ACL.
	Pattern pulumi.StringInput `pulumi:"pattern"`
	// Permission of the ACL
	// Available permission:
	// * `admin`
	// * `read`
	// * `write`
	// * `readwrite`
	// * `deny`
	Permission pulumi.StringInput `pulumi:"permission"`
}

func (OpensearchUserAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpensearchUserAcl)(nil)).Elem()
}

func (i OpensearchUserAclArgs) ToOpensearchUserAclOutput() OpensearchUserAclOutput {
	return i.ToOpensearchUserAclOutputWithContext(context.Background())
}

func (i OpensearchUserAclArgs) ToOpensearchUserAclOutputWithContext(ctx context.Context) OpensearchUserAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpensearchUserAclOutput)
}

// OpensearchUserAclArrayInput is an input type that accepts OpensearchUserAclArray and OpensearchUserAclArrayOutput values.
// You can construct a concrete instance of `OpensearchUserAclArrayInput` via:
//
//	OpensearchUserAclArray{ OpensearchUserAclArgs{...} }
type OpensearchUserAclArrayInput interface {
	pulumi.Input

	ToOpensearchUserAclArrayOutput() OpensearchUserAclArrayOutput
	ToOpensearchUserAclArrayOutputWithContext(context.Context) OpensearchUserAclArrayOutput
}

type OpensearchUserAclArray []OpensearchUserAclInput

func (OpensearchUserAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpensearchUserAcl)(nil)).Elem()
}

func (i OpensearchUserAclArray) ToOpensearchUserAclArrayOutput() OpensearchUserAclArrayOutput {
	return i.ToOpensearchUserAclArrayOutputWithContext(context.Background())
}

func (i OpensearchUserAclArray) ToOpensearchUserAclArrayOutputWithContext(ctx context.Context) OpensearchUserAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpensearchUserAclArrayOutput)
}

type OpensearchUserAclOutput struct{ *pulumi.OutputState }

func (OpensearchUserAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpensearchUserAcl)(nil)).Elem()
}

func (o OpensearchUserAclOutput) ToOpensearchUserAclOutput() OpensearchUserAclOutput {
	return o
}

func (o OpensearchUserAclOutput) ToOpensearchUserAclOutputWithContext(ctx context.Context) OpensearchUserAclOutput {
	return o
}

// Pattern of the ACL.
func (o OpensearchUserAclOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v OpensearchUserAcl) string { return v.Pattern }).(pulumi.StringOutput)
}

// Permission of the ACL
// Available permission:
// * `admin`
// * `read`
// * `write`
// * `readwrite`
// * `deny`
func (o OpensearchUserAclOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v OpensearchUserAcl) string { return v.Permission }).(pulumi.StringOutput)
}

type OpensearchUserAclArrayOutput struct{ *pulumi.OutputState }

func (OpensearchUserAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpensearchUserAcl)(nil)).Elem()
}

func (o OpensearchUserAclArrayOutput) ToOpensearchUserAclArrayOutput() OpensearchUserAclArrayOutput {
	return o
}

func (o OpensearchUserAclArrayOutput) ToOpensearchUserAclArrayOutputWithContext(ctx context.Context) OpensearchUserAclArrayOutput {
	return o
}

func (o OpensearchUserAclArrayOutput) Index(i pulumi.IntInput) OpensearchUserAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpensearchUserAcl {
		return vs[0].([]OpensearchUserAcl)[vs[1].(int)]
	}).(OpensearchUserAclOutput)
}

type GetCapabilitiesEngine struct {
	// Default version used for the engine.
	DefaultVersion string `pulumi:"defaultVersion"`
	// Description of the plan.
	Description string `pulumi:"description"`
	// Name of the plan.
	Name string `pulumi:"name"`
	// SSL modes for this engine.
	SslModes []string `pulumi:"sslModes"`
	// Versions available for this engine.
	Versions []string `pulumi:"versions"`
}

// GetCapabilitiesEngineInput is an input type that accepts GetCapabilitiesEngineArgs and GetCapabilitiesEngineOutput values.
// You can construct a concrete instance of `GetCapabilitiesEngineInput` via:
//
//	GetCapabilitiesEngineArgs{...}
type GetCapabilitiesEngineInput interface {
	pulumi.Input

	ToGetCapabilitiesEngineOutput() GetCapabilitiesEngineOutput
	ToGetCapabilitiesEngineOutputWithContext(context.Context) GetCapabilitiesEngineOutput
}

type GetCapabilitiesEngineArgs struct {
	// Default version used for the engine.
	DefaultVersion pulumi.StringInput `pulumi:"defaultVersion"`
	// Description of the plan.
	Description pulumi.StringInput `pulumi:"description"`
	// Name of the plan.
	Name pulumi.StringInput `pulumi:"name"`
	// SSL modes for this engine.
	SslModes pulumi.StringArrayInput `pulumi:"sslModes"`
	// Versions available for this engine.
	Versions pulumi.StringArrayInput `pulumi:"versions"`
}

func (GetCapabilitiesEngineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesEngine)(nil)).Elem()
}

func (i GetCapabilitiesEngineArgs) ToGetCapabilitiesEngineOutput() GetCapabilitiesEngineOutput {
	return i.ToGetCapabilitiesEngineOutputWithContext(context.Background())
}

func (i GetCapabilitiesEngineArgs) ToGetCapabilitiesEngineOutputWithContext(ctx context.Context) GetCapabilitiesEngineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCapabilitiesEngineOutput)
}

// GetCapabilitiesEngineArrayInput is an input type that accepts GetCapabilitiesEngineArray and GetCapabilitiesEngineArrayOutput values.
// You can construct a concrete instance of `GetCapabilitiesEngineArrayInput` via:
//
//	GetCapabilitiesEngineArray{ GetCapabilitiesEngineArgs{...} }
type GetCapabilitiesEngineArrayInput interface {
	pulumi.Input

	ToGetCapabilitiesEngineArrayOutput() GetCapabilitiesEngineArrayOutput
	ToGetCapabilitiesEngineArrayOutputWithContext(context.Context) GetCapabilitiesEngineArrayOutput
}

type GetCapabilitiesEngineArray []GetCapabilitiesEngineInput

func (GetCapabilitiesEngineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCapabilitiesEngine)(nil)).Elem()
}

func (i GetCapabilitiesEngineArray) ToGetCapabilitiesEngineArrayOutput() GetCapabilitiesEngineArrayOutput {
	return i.ToGetCapabilitiesEngineArrayOutputWithContext(context.Background())
}

func (i GetCapabilitiesEngineArray) ToGetCapabilitiesEngineArrayOutputWithContext(ctx context.Context) GetCapabilitiesEngineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCapabilitiesEngineArrayOutput)
}

type GetCapabilitiesEngineOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesEngineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesEngine)(nil)).Elem()
}

func (o GetCapabilitiesEngineOutput) ToGetCapabilitiesEngineOutput() GetCapabilitiesEngineOutput {
	return o
}

func (o GetCapabilitiesEngineOutput) ToGetCapabilitiesEngineOutputWithContext(ctx context.Context) GetCapabilitiesEngineOutput {
	return o
}

// Default version used for the engine.
func (o GetCapabilitiesEngineOutput) DefaultVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesEngine) string { return v.DefaultVersion }).(pulumi.StringOutput)
}

// Description of the plan.
func (o GetCapabilitiesEngineOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesEngine) string { return v.Description }).(pulumi.StringOutput)
}

// Name of the plan.
func (o GetCapabilitiesEngineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesEngine) string { return v.Name }).(pulumi.StringOutput)
}

// SSL modes for this engine.
func (o GetCapabilitiesEngineOutput) SslModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCapabilitiesEngine) []string { return v.SslModes }).(pulumi.StringArrayOutput)
}

// Versions available for this engine.
func (o GetCapabilitiesEngineOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCapabilitiesEngine) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

type GetCapabilitiesEngineArrayOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesEngineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCapabilitiesEngine)(nil)).Elem()
}

func (o GetCapabilitiesEngineArrayOutput) ToGetCapabilitiesEngineArrayOutput() GetCapabilitiesEngineArrayOutput {
	return o
}

func (o GetCapabilitiesEngineArrayOutput) ToGetCapabilitiesEngineArrayOutputWithContext(ctx context.Context) GetCapabilitiesEngineArrayOutput {
	return o
}

func (o GetCapabilitiesEngineArrayOutput) Index(i pulumi.IntInput) GetCapabilitiesEngineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCapabilitiesEngine {
		return vs[0].([]GetCapabilitiesEngine)[vs[1].(int)]
	}).(GetCapabilitiesEngineOutput)
}

type GetCapabilitiesFlavor struct {
	// Flavor core number.
	Core int `pulumi:"core"`
	// Flavor ram size in GB.
	Memory int `pulumi:"memory"`
	// Name of the plan.
	Name string `pulumi:"name"`
	// Flavor disk size in GB.
	Storage int `pulumi:"storage"`
}

// GetCapabilitiesFlavorInput is an input type that accepts GetCapabilitiesFlavorArgs and GetCapabilitiesFlavorOutput values.
// You can construct a concrete instance of `GetCapabilitiesFlavorInput` via:
//
//	GetCapabilitiesFlavorArgs{...}
type GetCapabilitiesFlavorInput interface {
	pulumi.Input

	ToGetCapabilitiesFlavorOutput() GetCapabilitiesFlavorOutput
	ToGetCapabilitiesFlavorOutputWithContext(context.Context) GetCapabilitiesFlavorOutput
}

type GetCapabilitiesFlavorArgs struct {
	// Flavor core number.
	Core pulumi.IntInput `pulumi:"core"`
	// Flavor ram size in GB.
	Memory pulumi.IntInput `pulumi:"memory"`
	// Name of the plan.
	Name pulumi.StringInput `pulumi:"name"`
	// Flavor disk size in GB.
	Storage pulumi.IntInput `pulumi:"storage"`
}

func (GetCapabilitiesFlavorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesFlavor)(nil)).Elem()
}

func (i GetCapabilitiesFlavorArgs) ToGetCapabilitiesFlavorOutput() GetCapabilitiesFlavorOutput {
	return i.ToGetCapabilitiesFlavorOutputWithContext(context.Background())
}

func (i GetCapabilitiesFlavorArgs) ToGetCapabilitiesFlavorOutputWithContext(ctx context.Context) GetCapabilitiesFlavorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCapabilitiesFlavorOutput)
}

// GetCapabilitiesFlavorArrayInput is an input type that accepts GetCapabilitiesFlavorArray and GetCapabilitiesFlavorArrayOutput values.
// You can construct a concrete instance of `GetCapabilitiesFlavorArrayInput` via:
//
//	GetCapabilitiesFlavorArray{ GetCapabilitiesFlavorArgs{...} }
type GetCapabilitiesFlavorArrayInput interface {
	pulumi.Input

	ToGetCapabilitiesFlavorArrayOutput() GetCapabilitiesFlavorArrayOutput
	ToGetCapabilitiesFlavorArrayOutputWithContext(context.Context) GetCapabilitiesFlavorArrayOutput
}

type GetCapabilitiesFlavorArray []GetCapabilitiesFlavorInput

func (GetCapabilitiesFlavorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCapabilitiesFlavor)(nil)).Elem()
}

func (i GetCapabilitiesFlavorArray) ToGetCapabilitiesFlavorArrayOutput() GetCapabilitiesFlavorArrayOutput {
	return i.ToGetCapabilitiesFlavorArrayOutputWithContext(context.Background())
}

func (i GetCapabilitiesFlavorArray) ToGetCapabilitiesFlavorArrayOutputWithContext(ctx context.Context) GetCapabilitiesFlavorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCapabilitiesFlavorArrayOutput)
}

type GetCapabilitiesFlavorOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesFlavorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesFlavor)(nil)).Elem()
}

func (o GetCapabilitiesFlavorOutput) ToGetCapabilitiesFlavorOutput() GetCapabilitiesFlavorOutput {
	return o
}

func (o GetCapabilitiesFlavorOutput) ToGetCapabilitiesFlavorOutputWithContext(ctx context.Context) GetCapabilitiesFlavorOutput {
	return o
}

// Flavor core number.
func (o GetCapabilitiesFlavorOutput) Core() pulumi.IntOutput {
	return o.ApplyT(func(v GetCapabilitiesFlavor) int { return v.Core }).(pulumi.IntOutput)
}

// Flavor ram size in GB.
func (o GetCapabilitiesFlavorOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v GetCapabilitiesFlavor) int { return v.Memory }).(pulumi.IntOutput)
}

// Name of the plan.
func (o GetCapabilitiesFlavorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesFlavor) string { return v.Name }).(pulumi.StringOutput)
}

// Flavor disk size in GB.
func (o GetCapabilitiesFlavorOutput) Storage() pulumi.IntOutput {
	return o.ApplyT(func(v GetCapabilitiesFlavor) int { return v.Storage }).(pulumi.IntOutput)
}

type GetCapabilitiesFlavorArrayOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesFlavorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCapabilitiesFlavor)(nil)).Elem()
}

func (o GetCapabilitiesFlavorArrayOutput) ToGetCapabilitiesFlavorArrayOutput() GetCapabilitiesFlavorArrayOutput {
	return o
}

func (o GetCapabilitiesFlavorArrayOutput) ToGetCapabilitiesFlavorArrayOutputWithContext(ctx context.Context) GetCapabilitiesFlavorArrayOutput {
	return o
}

func (o GetCapabilitiesFlavorArrayOutput) Index(i pulumi.IntInput) GetCapabilitiesFlavorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCapabilitiesFlavor {
		return vs[0].([]GetCapabilitiesFlavor)[vs[1].(int)]
	}).(GetCapabilitiesFlavorOutput)
}

type GetCapabilitiesOption struct {
	// Name of the plan.
	Name string `pulumi:"name"`
	// Type of the option.
	Type string `pulumi:"type"`
}

// GetCapabilitiesOptionInput is an input type that accepts GetCapabilitiesOptionArgs and GetCapabilitiesOptionOutput values.
// You can construct a concrete instance of `GetCapabilitiesOptionInput` via:
//
//	GetCapabilitiesOptionArgs{...}
type GetCapabilitiesOptionInput interface {
	pulumi.Input

	ToGetCapabilitiesOptionOutput() GetCapabilitiesOptionOutput
	ToGetCapabilitiesOptionOutputWithContext(context.Context) GetCapabilitiesOptionOutput
}

type GetCapabilitiesOptionArgs struct {
	// Name of the plan.
	Name pulumi.StringInput `pulumi:"name"`
	// Type of the option.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetCapabilitiesOptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesOption)(nil)).Elem()
}

func (i GetCapabilitiesOptionArgs) ToGetCapabilitiesOptionOutput() GetCapabilitiesOptionOutput {
	return i.ToGetCapabilitiesOptionOutputWithContext(context.Background())
}

func (i GetCapabilitiesOptionArgs) ToGetCapabilitiesOptionOutputWithContext(ctx context.Context) GetCapabilitiesOptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCapabilitiesOptionOutput)
}

// GetCapabilitiesOptionArrayInput is an input type that accepts GetCapabilitiesOptionArray and GetCapabilitiesOptionArrayOutput values.
// You can construct a concrete instance of `GetCapabilitiesOptionArrayInput` via:
//
//	GetCapabilitiesOptionArray{ GetCapabilitiesOptionArgs{...} }
type GetCapabilitiesOptionArrayInput interface {
	pulumi.Input

	ToGetCapabilitiesOptionArrayOutput() GetCapabilitiesOptionArrayOutput
	ToGetCapabilitiesOptionArrayOutputWithContext(context.Context) GetCapabilitiesOptionArrayOutput
}

type GetCapabilitiesOptionArray []GetCapabilitiesOptionInput

func (GetCapabilitiesOptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCapabilitiesOption)(nil)).Elem()
}

func (i GetCapabilitiesOptionArray) ToGetCapabilitiesOptionArrayOutput() GetCapabilitiesOptionArrayOutput {
	return i.ToGetCapabilitiesOptionArrayOutputWithContext(context.Background())
}

func (i GetCapabilitiesOptionArray) ToGetCapabilitiesOptionArrayOutputWithContext(ctx context.Context) GetCapabilitiesOptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCapabilitiesOptionArrayOutput)
}

type GetCapabilitiesOptionOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesOption)(nil)).Elem()
}

func (o GetCapabilitiesOptionOutput) ToGetCapabilitiesOptionOutput() GetCapabilitiesOptionOutput {
	return o
}

func (o GetCapabilitiesOptionOutput) ToGetCapabilitiesOptionOutputWithContext(ctx context.Context) GetCapabilitiesOptionOutput {
	return o
}

// Name of the plan.
func (o GetCapabilitiesOptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesOption) string { return v.Name }).(pulumi.StringOutput)
}

// Type of the option.
func (o GetCapabilitiesOptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesOption) string { return v.Type }).(pulumi.StringOutput)
}

type GetCapabilitiesOptionArrayOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesOptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCapabilitiesOption)(nil)).Elem()
}

func (o GetCapabilitiesOptionArrayOutput) ToGetCapabilitiesOptionArrayOutput() GetCapabilitiesOptionArrayOutput {
	return o
}

func (o GetCapabilitiesOptionArrayOutput) ToGetCapabilitiesOptionArrayOutputWithContext(ctx context.Context) GetCapabilitiesOptionArrayOutput {
	return o
}

func (o GetCapabilitiesOptionArrayOutput) Index(i pulumi.IntInput) GetCapabilitiesOptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCapabilitiesOption {
		return vs[0].([]GetCapabilitiesOption)[vs[1].(int)]
	}).(GetCapabilitiesOptionOutput)
}

type GetCapabilitiesPlan struct {
	// Automatic backup retention duration.
	BackupRetention string `pulumi:"backupRetention"`
	// Description of the plan.
	Description string `pulumi:"description"`
	// Name of the plan.
	Name string `pulumi:"name"`
}

// GetCapabilitiesPlanInput is an input type that accepts GetCapabilitiesPlanArgs and GetCapabilitiesPlanOutput values.
// You can construct a concrete instance of `GetCapabilitiesPlanInput` via:
//
//	GetCapabilitiesPlanArgs{...}
type GetCapabilitiesPlanInput interface {
	pulumi.Input

	ToGetCapabilitiesPlanOutput() GetCapabilitiesPlanOutput
	ToGetCapabilitiesPlanOutputWithContext(context.Context) GetCapabilitiesPlanOutput
}

type GetCapabilitiesPlanArgs struct {
	// Automatic backup retention duration.
	BackupRetention pulumi.StringInput `pulumi:"backupRetention"`
	// Description of the plan.
	Description pulumi.StringInput `pulumi:"description"`
	// Name of the plan.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetCapabilitiesPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesPlan)(nil)).Elem()
}

func (i GetCapabilitiesPlanArgs) ToGetCapabilitiesPlanOutput() GetCapabilitiesPlanOutput {
	return i.ToGetCapabilitiesPlanOutputWithContext(context.Background())
}

func (i GetCapabilitiesPlanArgs) ToGetCapabilitiesPlanOutputWithContext(ctx context.Context) GetCapabilitiesPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCapabilitiesPlanOutput)
}

// GetCapabilitiesPlanArrayInput is an input type that accepts GetCapabilitiesPlanArray and GetCapabilitiesPlanArrayOutput values.
// You can construct a concrete instance of `GetCapabilitiesPlanArrayInput` via:
//
//	GetCapabilitiesPlanArray{ GetCapabilitiesPlanArgs{...} }
type GetCapabilitiesPlanArrayInput interface {
	pulumi.Input

	ToGetCapabilitiesPlanArrayOutput() GetCapabilitiesPlanArrayOutput
	ToGetCapabilitiesPlanArrayOutputWithContext(context.Context) GetCapabilitiesPlanArrayOutput
}

type GetCapabilitiesPlanArray []GetCapabilitiesPlanInput

func (GetCapabilitiesPlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCapabilitiesPlan)(nil)).Elem()
}

func (i GetCapabilitiesPlanArray) ToGetCapabilitiesPlanArrayOutput() GetCapabilitiesPlanArrayOutput {
	return i.ToGetCapabilitiesPlanArrayOutputWithContext(context.Background())
}

func (i GetCapabilitiesPlanArray) ToGetCapabilitiesPlanArrayOutputWithContext(ctx context.Context) GetCapabilitiesPlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCapabilitiesPlanArrayOutput)
}

type GetCapabilitiesPlanOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesPlan)(nil)).Elem()
}

func (o GetCapabilitiesPlanOutput) ToGetCapabilitiesPlanOutput() GetCapabilitiesPlanOutput {
	return o
}

func (o GetCapabilitiesPlanOutput) ToGetCapabilitiesPlanOutputWithContext(ctx context.Context) GetCapabilitiesPlanOutput {
	return o
}

// Automatic backup retention duration.
func (o GetCapabilitiesPlanOutput) BackupRetention() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesPlan) string { return v.BackupRetention }).(pulumi.StringOutput)
}

// Description of the plan.
func (o GetCapabilitiesPlanOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesPlan) string { return v.Description }).(pulumi.StringOutput)
}

// Name of the plan.
func (o GetCapabilitiesPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesPlan) string { return v.Name }).(pulumi.StringOutput)
}

type GetCapabilitiesPlanArrayOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesPlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCapabilitiesPlan)(nil)).Elem()
}

func (o GetCapabilitiesPlanArrayOutput) ToGetCapabilitiesPlanArrayOutput() GetCapabilitiesPlanArrayOutput {
	return o
}

func (o GetCapabilitiesPlanArrayOutput) ToGetCapabilitiesPlanArrayOutputWithContext(ctx context.Context) GetCapabilitiesPlanArrayOutput {
	return o
}

func (o GetCapabilitiesPlanArrayOutput) Index(i pulumi.IntInput) GetCapabilitiesPlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCapabilitiesPlan {
		return vs[0].([]GetCapabilitiesPlan)[vs[1].(int)]
	}).(GetCapabilitiesPlanOutput)
}

type GetDatabaseEndpoint struct {
	// Type of component the URI relates to.
	Component string `pulumi:"component"`
	// Domain of the cluster.
	Domain string `pulumi:"domain"`
	// Path of the endpoint.
	Path string `pulumi:"path"`
	// Connection port for the endpoint.
	Port int `pulumi:"port"`
	// Scheme used to generate the URI.
	Scheme string `pulumi:"scheme"`
	// Defines whether the endpoint uses SSL.
	Ssl bool `pulumi:"ssl"`
	// SSL mode used to connect to the service if the SSL is enabled.
	SslMode string `pulumi:"sslMode"`
	// URI of the endpoint.
	Uri string `pulumi:"uri"`
}

// GetDatabaseEndpointInput is an input type that accepts GetDatabaseEndpointArgs and GetDatabaseEndpointOutput values.
// You can construct a concrete instance of `GetDatabaseEndpointInput` via:
//
//	GetDatabaseEndpointArgs{...}
type GetDatabaseEndpointInput interface {
	pulumi.Input

	ToGetDatabaseEndpointOutput() GetDatabaseEndpointOutput
	ToGetDatabaseEndpointOutputWithContext(context.Context) GetDatabaseEndpointOutput
}

type GetDatabaseEndpointArgs struct {
	// Type of component the URI relates to.
	Component pulumi.StringInput `pulumi:"component"`
	// Domain of the cluster.
	Domain pulumi.StringInput `pulumi:"domain"`
	// Path of the endpoint.
	Path pulumi.StringInput `pulumi:"path"`
	// Connection port for the endpoint.
	Port pulumi.IntInput `pulumi:"port"`
	// Scheme used to generate the URI.
	Scheme pulumi.StringInput `pulumi:"scheme"`
	// Defines whether the endpoint uses SSL.
	Ssl pulumi.BoolInput `pulumi:"ssl"`
	// SSL mode used to connect to the service if the SSL is enabled.
	SslMode pulumi.StringInput `pulumi:"sslMode"`
	// URI of the endpoint.
	Uri pulumi.StringInput `pulumi:"uri"`
}

func (GetDatabaseEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseEndpoint)(nil)).Elem()
}

func (i GetDatabaseEndpointArgs) ToGetDatabaseEndpointOutput() GetDatabaseEndpointOutput {
	return i.ToGetDatabaseEndpointOutputWithContext(context.Background())
}

func (i GetDatabaseEndpointArgs) ToGetDatabaseEndpointOutputWithContext(ctx context.Context) GetDatabaseEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDatabaseEndpointOutput)
}

// GetDatabaseEndpointArrayInput is an input type that accepts GetDatabaseEndpointArray and GetDatabaseEndpointArrayOutput values.
// You can construct a concrete instance of `GetDatabaseEndpointArrayInput` via:
//
//	GetDatabaseEndpointArray{ GetDatabaseEndpointArgs{...} }
type GetDatabaseEndpointArrayInput interface {
	pulumi.Input

	ToGetDatabaseEndpointArrayOutput() GetDatabaseEndpointArrayOutput
	ToGetDatabaseEndpointArrayOutputWithContext(context.Context) GetDatabaseEndpointArrayOutput
}

type GetDatabaseEndpointArray []GetDatabaseEndpointInput

func (GetDatabaseEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDatabaseEndpoint)(nil)).Elem()
}

func (i GetDatabaseEndpointArray) ToGetDatabaseEndpointArrayOutput() GetDatabaseEndpointArrayOutput {
	return i.ToGetDatabaseEndpointArrayOutputWithContext(context.Background())
}

func (i GetDatabaseEndpointArray) ToGetDatabaseEndpointArrayOutputWithContext(ctx context.Context) GetDatabaseEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDatabaseEndpointArrayOutput)
}

type GetDatabaseEndpointOutput struct{ *pulumi.OutputState }

func (GetDatabaseEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseEndpoint)(nil)).Elem()
}

func (o GetDatabaseEndpointOutput) ToGetDatabaseEndpointOutput() GetDatabaseEndpointOutput {
	return o
}

func (o GetDatabaseEndpointOutput) ToGetDatabaseEndpointOutputWithContext(ctx context.Context) GetDatabaseEndpointOutput {
	return o
}

// Type of component the URI relates to.
func (o GetDatabaseEndpointOutput) Component() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseEndpoint) string { return v.Component }).(pulumi.StringOutput)
}

// Domain of the cluster.
func (o GetDatabaseEndpointOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseEndpoint) string { return v.Domain }).(pulumi.StringOutput)
}

// Path of the endpoint.
func (o GetDatabaseEndpointOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseEndpoint) string { return v.Path }).(pulumi.StringOutput)
}

// Connection port for the endpoint.
func (o GetDatabaseEndpointOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetDatabaseEndpoint) int { return v.Port }).(pulumi.IntOutput)
}

// Scheme used to generate the URI.
func (o GetDatabaseEndpointOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseEndpoint) string { return v.Scheme }).(pulumi.StringOutput)
}

// Defines whether the endpoint uses SSL.
func (o GetDatabaseEndpointOutput) Ssl() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDatabaseEndpoint) bool { return v.Ssl }).(pulumi.BoolOutput)
}

// SSL mode used to connect to the service if the SSL is enabled.
func (o GetDatabaseEndpointOutput) SslMode() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseEndpoint) string { return v.SslMode }).(pulumi.StringOutput)
}

// URI of the endpoint.
func (o GetDatabaseEndpointOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseEndpoint) string { return v.Uri }).(pulumi.StringOutput)
}

type GetDatabaseEndpointArrayOutput struct{ *pulumi.OutputState }

func (GetDatabaseEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDatabaseEndpoint)(nil)).Elem()
}

func (o GetDatabaseEndpointArrayOutput) ToGetDatabaseEndpointArrayOutput() GetDatabaseEndpointArrayOutput {
	return o
}

func (o GetDatabaseEndpointArrayOutput) ToGetDatabaseEndpointArrayOutputWithContext(ctx context.Context) GetDatabaseEndpointArrayOutput {
	return o
}

func (o GetDatabaseEndpointArrayOutput) Index(i pulumi.IntInput) GetDatabaseEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDatabaseEndpoint {
		return vs[0].([]GetDatabaseEndpoint)[vs[1].(int)]
	}).(GetDatabaseEndpointOutput)
}

type GetDatabaseNode struct {
	// Private network id in which the node should be deployed. It's the regional openstackId of the private network
	NetworkId string `pulumi:"networkId"`
	// Public cloud region in which the node should be deployed.
	Region string `pulumi:"region"`
	// Private subnet ID in which the node is.
	SubnetId string `pulumi:"subnetId"`
}

// GetDatabaseNodeInput is an input type that accepts GetDatabaseNodeArgs and GetDatabaseNodeOutput values.
// You can construct a concrete instance of `GetDatabaseNodeInput` via:
//
//	GetDatabaseNodeArgs{...}
type GetDatabaseNodeInput interface {
	pulumi.Input

	ToGetDatabaseNodeOutput() GetDatabaseNodeOutput
	ToGetDatabaseNodeOutputWithContext(context.Context) GetDatabaseNodeOutput
}

type GetDatabaseNodeArgs struct {
	// Private network id in which the node should be deployed. It's the regional openstackId of the private network
	NetworkId pulumi.StringInput `pulumi:"networkId"`
	// Public cloud region in which the node should be deployed.
	Region pulumi.StringInput `pulumi:"region"`
	// Private subnet ID in which the node is.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}

func (GetDatabaseNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseNode)(nil)).Elem()
}

func (i GetDatabaseNodeArgs) ToGetDatabaseNodeOutput() GetDatabaseNodeOutput {
	return i.ToGetDatabaseNodeOutputWithContext(context.Background())
}

func (i GetDatabaseNodeArgs) ToGetDatabaseNodeOutputWithContext(ctx context.Context) GetDatabaseNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDatabaseNodeOutput)
}

// GetDatabaseNodeArrayInput is an input type that accepts GetDatabaseNodeArray and GetDatabaseNodeArrayOutput values.
// You can construct a concrete instance of `GetDatabaseNodeArrayInput` via:
//
//	GetDatabaseNodeArray{ GetDatabaseNodeArgs{...} }
type GetDatabaseNodeArrayInput interface {
	pulumi.Input

	ToGetDatabaseNodeArrayOutput() GetDatabaseNodeArrayOutput
	ToGetDatabaseNodeArrayOutputWithContext(context.Context) GetDatabaseNodeArrayOutput
}

type GetDatabaseNodeArray []GetDatabaseNodeInput

func (GetDatabaseNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDatabaseNode)(nil)).Elem()
}

func (i GetDatabaseNodeArray) ToGetDatabaseNodeArrayOutput() GetDatabaseNodeArrayOutput {
	return i.ToGetDatabaseNodeArrayOutputWithContext(context.Background())
}

func (i GetDatabaseNodeArray) ToGetDatabaseNodeArrayOutputWithContext(ctx context.Context) GetDatabaseNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDatabaseNodeArrayOutput)
}

type GetDatabaseNodeOutput struct{ *pulumi.OutputState }

func (GetDatabaseNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseNode)(nil)).Elem()
}

func (o GetDatabaseNodeOutput) ToGetDatabaseNodeOutput() GetDatabaseNodeOutput {
	return o
}

func (o GetDatabaseNodeOutput) ToGetDatabaseNodeOutputWithContext(ctx context.Context) GetDatabaseNodeOutput {
	return o
}

// Private network id in which the node should be deployed. It's the regional openstackId of the private network
func (o GetDatabaseNodeOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseNode) string { return v.NetworkId }).(pulumi.StringOutput)
}

// Public cloud region in which the node should be deployed.
func (o GetDatabaseNodeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseNode) string { return v.Region }).(pulumi.StringOutput)
}

// Private subnet ID in which the node is.
func (o GetDatabaseNodeOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseNode) string { return v.SubnetId }).(pulumi.StringOutput)
}

type GetDatabaseNodeArrayOutput struct{ *pulumi.OutputState }

func (GetDatabaseNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDatabaseNode)(nil)).Elem()
}

func (o GetDatabaseNodeArrayOutput) ToGetDatabaseNodeArrayOutput() GetDatabaseNodeArrayOutput {
	return o
}

func (o GetDatabaseNodeArrayOutput) ToGetDatabaseNodeArrayOutputWithContext(ctx context.Context) GetDatabaseNodeArrayOutput {
	return o
}

func (o GetDatabaseNodeArrayOutput) Index(i pulumi.IntInput) GetDatabaseNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDatabaseNode {
		return vs[0].([]GetDatabaseNode)[vs[1].(int)]
	}).(GetDatabaseNodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OpensearchUserAclInput)(nil)).Elem(), OpensearchUserAclArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OpensearchUserAclArrayInput)(nil)).Elem(), OpensearchUserAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCapabilitiesEngineInput)(nil)).Elem(), GetCapabilitiesEngineArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCapabilitiesEngineArrayInput)(nil)).Elem(), GetCapabilitiesEngineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCapabilitiesFlavorInput)(nil)).Elem(), GetCapabilitiesFlavorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCapabilitiesFlavorArrayInput)(nil)).Elem(), GetCapabilitiesFlavorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCapabilitiesOptionInput)(nil)).Elem(), GetCapabilitiesOptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCapabilitiesOptionArrayInput)(nil)).Elem(), GetCapabilitiesOptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCapabilitiesPlanInput)(nil)).Elem(), GetCapabilitiesPlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCapabilitiesPlanArrayInput)(nil)).Elem(), GetCapabilitiesPlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDatabaseEndpointInput)(nil)).Elem(), GetDatabaseEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDatabaseEndpointArrayInput)(nil)).Elem(), GetDatabaseEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDatabaseNodeInput)(nil)).Elem(), GetDatabaseNodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDatabaseNodeArrayInput)(nil)).Elem(), GetDatabaseNodeArray{})
	pulumi.RegisterOutputType(OpensearchUserAclOutput{})
	pulumi.RegisterOutputType(OpensearchUserAclArrayOutput{})
	pulumi.RegisterOutputType(GetCapabilitiesEngineOutput{})
	pulumi.RegisterOutputType(GetCapabilitiesEngineArrayOutput{})
	pulumi.RegisterOutputType(GetCapabilitiesFlavorOutput{})
	pulumi.RegisterOutputType(GetCapabilitiesFlavorArrayOutput{})
	pulumi.RegisterOutputType(GetCapabilitiesOptionOutput{})
	pulumi.RegisterOutputType(GetCapabilitiesOptionArrayOutput{})
	pulumi.RegisterOutputType(GetCapabilitiesPlanOutput{})
	pulumi.RegisterOutputType(GetCapabilitiesPlanArrayOutput{})
	pulumi.RegisterOutputType(GetDatabaseEndpointOutput{})
	pulumi.RegisterOutputType(GetDatabaseEndpointArrayOutput{})
	pulumi.RegisterOutputType(GetDatabaseNodeOutput{})
	pulumi.RegisterOutputType(GetDatabaseNodeArrayOutput{})
}
