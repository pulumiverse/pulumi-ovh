// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DedicatedServiceInstallTask struct {
	pulumi.CustomResourceState

	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy pulumi.IntPtrOutput `pulumi:"bootidOnDestroy"`
	// Details of this task. (should be `Install asked`)
	Comment pulumi.StringOutput `pulumi:"comment"`
	// see `details` block below.
	Details DedicatedServiceInstallTaskDetailsPtrOutput `pulumi:"details"`
	// Completion date in RFC3339 format.
	DoneDate pulumi.StringOutput `pulumi:"doneDate"`
	// Function name (should be `hardInstall`).
	Function pulumi.StringOutput `pulumi:"function"`
	// Last update in RFC3339 format.
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// Partition scheme name.
	PartitionSchemeName pulumi.StringPtrOutput `pulumi:"partitionSchemeName"`
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Task creation date in RFC3339 format.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Task status (should be `done`)
	Status pulumi.StringOutput `pulumi:"status"`
	// Template name.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewDedicatedServiceInstallTask registers a new resource with the given unique name, arguments, and options.
func NewDedicatedServiceInstallTask(ctx *pulumi.Context,
	name string, args *DedicatedServiceInstallTaskArgs, opts ...pulumi.ResourceOption) (*DedicatedServiceInstallTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DedicatedServiceInstallTask
	err := ctx.RegisterResource("ovh:index/dedicatedServiceInstallTask:DedicatedServiceInstallTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedServiceInstallTask gets an existing DedicatedServiceInstallTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedServiceInstallTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedServiceInstallTaskState, opts ...pulumi.ResourceOption) (*DedicatedServiceInstallTask, error) {
	var resource DedicatedServiceInstallTask
	err := ctx.ReadResource("ovh:index/dedicatedServiceInstallTask:DedicatedServiceInstallTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedServiceInstallTask resources.
type dedicatedServiceInstallTaskState struct {
	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy *int `pulumi:"bootidOnDestroy"`
	// Details of this task. (should be `Install asked`)
	Comment *string `pulumi:"comment"`
	// see `details` block below.
	Details *DedicatedServiceInstallTaskDetails `pulumi:"details"`
	// Completion date in RFC3339 format.
	DoneDate *string `pulumi:"doneDate"`
	// Function name (should be `hardInstall`).
	Function *string `pulumi:"function"`
	// Last update in RFC3339 format.
	LastUpdate *string `pulumi:"lastUpdate"`
	// Partition scheme name.
	PartitionSchemeName *string `pulumi:"partitionSchemeName"`
	// The serviceName of your dedicated server.
	ServiceName *string `pulumi:"serviceName"`
	// Task creation date in RFC3339 format.
	StartDate *string `pulumi:"startDate"`
	// Task status (should be `done`)
	Status *string `pulumi:"status"`
	// Template name.
	TemplateName *string `pulumi:"templateName"`
}

type DedicatedServiceInstallTaskState struct {
	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy pulumi.IntPtrInput
	// Details of this task. (should be `Install asked`)
	Comment pulumi.StringPtrInput
	// see `details` block below.
	Details DedicatedServiceInstallTaskDetailsPtrInput
	// Completion date in RFC3339 format.
	DoneDate pulumi.StringPtrInput
	// Function name (should be `hardInstall`).
	Function pulumi.StringPtrInput
	// Last update in RFC3339 format.
	LastUpdate pulumi.StringPtrInput
	// Partition scheme name.
	PartitionSchemeName pulumi.StringPtrInput
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringPtrInput
	// Task creation date in RFC3339 format.
	StartDate pulumi.StringPtrInput
	// Task status (should be `done`)
	Status pulumi.StringPtrInput
	// Template name.
	TemplateName pulumi.StringPtrInput
}

func (DedicatedServiceInstallTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServiceInstallTaskState)(nil)).Elem()
}

type dedicatedServiceInstallTaskArgs struct {
	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy *int `pulumi:"bootidOnDestroy"`
	// see `details` block below.
	Details *DedicatedServiceInstallTaskDetails `pulumi:"details"`
	// Partition scheme name.
	PartitionSchemeName *string `pulumi:"partitionSchemeName"`
	// The serviceName of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
	// Template name.
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a DedicatedServiceInstallTask resource.
type DedicatedServiceInstallTaskArgs struct {
	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy pulumi.IntPtrInput
	// see `details` block below.
	Details DedicatedServiceInstallTaskDetailsPtrInput
	// Partition scheme name.
	PartitionSchemeName pulumi.StringPtrInput
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringInput
	// Template name.
	TemplateName pulumi.StringInput
}

func (DedicatedServiceInstallTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServiceInstallTaskArgs)(nil)).Elem()
}

type DedicatedServiceInstallTaskInput interface {
	pulumi.Input

	ToDedicatedServiceInstallTaskOutput() DedicatedServiceInstallTaskOutput
	ToDedicatedServiceInstallTaskOutputWithContext(ctx context.Context) DedicatedServiceInstallTaskOutput
}

func (*DedicatedServiceInstallTask) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServiceInstallTask)(nil)).Elem()
}

func (i *DedicatedServiceInstallTask) ToDedicatedServiceInstallTaskOutput() DedicatedServiceInstallTaskOutput {
	return i.ToDedicatedServiceInstallTaskOutputWithContext(context.Background())
}

func (i *DedicatedServiceInstallTask) ToDedicatedServiceInstallTaskOutputWithContext(ctx context.Context) DedicatedServiceInstallTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServiceInstallTaskOutput)
}

// DedicatedServiceInstallTaskArrayInput is an input type that accepts DedicatedServiceInstallTaskArray and DedicatedServiceInstallTaskArrayOutput values.
// You can construct a concrete instance of `DedicatedServiceInstallTaskArrayInput` via:
//
//	DedicatedServiceInstallTaskArray{ DedicatedServiceInstallTaskArgs{...} }
type DedicatedServiceInstallTaskArrayInput interface {
	pulumi.Input

	ToDedicatedServiceInstallTaskArrayOutput() DedicatedServiceInstallTaskArrayOutput
	ToDedicatedServiceInstallTaskArrayOutputWithContext(context.Context) DedicatedServiceInstallTaskArrayOutput
}

type DedicatedServiceInstallTaskArray []DedicatedServiceInstallTaskInput

func (DedicatedServiceInstallTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedServiceInstallTask)(nil)).Elem()
}

func (i DedicatedServiceInstallTaskArray) ToDedicatedServiceInstallTaskArrayOutput() DedicatedServiceInstallTaskArrayOutput {
	return i.ToDedicatedServiceInstallTaskArrayOutputWithContext(context.Background())
}

func (i DedicatedServiceInstallTaskArray) ToDedicatedServiceInstallTaskArrayOutputWithContext(ctx context.Context) DedicatedServiceInstallTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServiceInstallTaskArrayOutput)
}

// DedicatedServiceInstallTaskMapInput is an input type that accepts DedicatedServiceInstallTaskMap and DedicatedServiceInstallTaskMapOutput values.
// You can construct a concrete instance of `DedicatedServiceInstallTaskMapInput` via:
//
//	DedicatedServiceInstallTaskMap{ "key": DedicatedServiceInstallTaskArgs{...} }
type DedicatedServiceInstallTaskMapInput interface {
	pulumi.Input

	ToDedicatedServiceInstallTaskMapOutput() DedicatedServiceInstallTaskMapOutput
	ToDedicatedServiceInstallTaskMapOutputWithContext(context.Context) DedicatedServiceInstallTaskMapOutput
}

type DedicatedServiceInstallTaskMap map[string]DedicatedServiceInstallTaskInput

func (DedicatedServiceInstallTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedServiceInstallTask)(nil)).Elem()
}

func (i DedicatedServiceInstallTaskMap) ToDedicatedServiceInstallTaskMapOutput() DedicatedServiceInstallTaskMapOutput {
	return i.ToDedicatedServiceInstallTaskMapOutputWithContext(context.Background())
}

func (i DedicatedServiceInstallTaskMap) ToDedicatedServiceInstallTaskMapOutputWithContext(ctx context.Context) DedicatedServiceInstallTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServiceInstallTaskMapOutput)
}

type DedicatedServiceInstallTaskOutput struct{ *pulumi.OutputState }

func (DedicatedServiceInstallTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServiceInstallTask)(nil)).Elem()
}

func (o DedicatedServiceInstallTaskOutput) ToDedicatedServiceInstallTaskOutput() DedicatedServiceInstallTaskOutput {
	return o
}

func (o DedicatedServiceInstallTaskOutput) ToDedicatedServiceInstallTaskOutputWithContext(ctx context.Context) DedicatedServiceInstallTaskOutput {
	return o
}

// If set, reboot the server on the specified boot id during destroy phase.
func (o DedicatedServiceInstallTaskOutput) BootidOnDestroy() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.IntPtrOutput { return v.BootidOnDestroy }).(pulumi.IntPtrOutput)
}

// Details of this task. (should be `Install asked`)
func (o DedicatedServiceInstallTaskOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

// see `details` block below.
func (o DedicatedServiceInstallTaskOutput) Details() DedicatedServiceInstallTaskDetailsPtrOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) DedicatedServiceInstallTaskDetailsPtrOutput { return v.Details }).(DedicatedServiceInstallTaskDetailsPtrOutput)
}

// Completion date in RFC3339 format.
func (o DedicatedServiceInstallTaskOutput) DoneDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringOutput { return v.DoneDate }).(pulumi.StringOutput)
}

// Function name (should be `hardInstall`).
func (o DedicatedServiceInstallTaskOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringOutput { return v.Function }).(pulumi.StringOutput)
}

// Last update in RFC3339 format.
func (o DedicatedServiceInstallTaskOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringOutput { return v.LastUpdate }).(pulumi.StringOutput)
}

// Partition scheme name.
func (o DedicatedServiceInstallTaskOutput) PartitionSchemeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringPtrOutput { return v.PartitionSchemeName }).(pulumi.StringPtrOutput)
}

// The serviceName of your dedicated server.
func (o DedicatedServiceInstallTaskOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Task creation date in RFC3339 format.
func (o DedicatedServiceInstallTaskOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// Task status (should be `done`)
func (o DedicatedServiceInstallTaskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Template name.
func (o DedicatedServiceInstallTaskOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServiceInstallTask) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

type DedicatedServiceInstallTaskArrayOutput struct{ *pulumi.OutputState }

func (DedicatedServiceInstallTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedServiceInstallTask)(nil)).Elem()
}

func (o DedicatedServiceInstallTaskArrayOutput) ToDedicatedServiceInstallTaskArrayOutput() DedicatedServiceInstallTaskArrayOutput {
	return o
}

func (o DedicatedServiceInstallTaskArrayOutput) ToDedicatedServiceInstallTaskArrayOutputWithContext(ctx context.Context) DedicatedServiceInstallTaskArrayOutput {
	return o
}

func (o DedicatedServiceInstallTaskArrayOutput) Index(i pulumi.IntInput) DedicatedServiceInstallTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedServiceInstallTask {
		return vs[0].([]*DedicatedServiceInstallTask)[vs[1].(int)]
	}).(DedicatedServiceInstallTaskOutput)
}

type DedicatedServiceInstallTaskMapOutput struct{ *pulumi.OutputState }

func (DedicatedServiceInstallTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedServiceInstallTask)(nil)).Elem()
}

func (o DedicatedServiceInstallTaskMapOutput) ToDedicatedServiceInstallTaskMapOutput() DedicatedServiceInstallTaskMapOutput {
	return o
}

func (o DedicatedServiceInstallTaskMapOutput) ToDedicatedServiceInstallTaskMapOutputWithContext(ctx context.Context) DedicatedServiceInstallTaskMapOutput {
	return o
}

func (o DedicatedServiceInstallTaskMapOutput) MapIndex(k pulumi.StringInput) DedicatedServiceInstallTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedServiceInstallTask {
		return vs[0].(map[string]*DedicatedServiceInstallTask)[vs[1].(string)]
	}).(DedicatedServiceInstallTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServiceInstallTaskInput)(nil)).Elem(), &DedicatedServiceInstallTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServiceInstallTaskArrayInput)(nil)).Elem(), DedicatedServiceInstallTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServiceInstallTaskMapInput)(nil)).Elem(), DedicatedServiceInstallTaskMap{})
	pulumi.RegisterOutputType(DedicatedServiceInstallTaskOutput{})
	pulumi.RegisterOutputType(DedicatedServiceInstallTaskArrayOutput{})
	pulumi.RegisterOutputType(DedicatedServiceInstallTaskMapOutput{})
}
