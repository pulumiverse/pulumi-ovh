// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Orders a public cloud project.
//
// ## Important
//
// This resource is in beta state. Use with caution.
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
//			mycart, err := ovh.GetOrderCart(ctx, &GetOrderCartArgs{
//				OvhSubsidiary: "fr",
//				Description:   pulumi.StringRef("my cloud order cart"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			cloudOrderCartProductPlan, err := ovh.GetOrderCartProductPlan(ctx, &GetOrderCartProductPlanArgs{
//				CartId:        mycart.Id,
//				PriceCapacity: "renew",
//				Product:       "cloud",
//				PlanCode:      "project.2018",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ovh.NewCloudProject(ctx, "cloudCloudProject", &ovh.CloudProjectArgs{
//				OvhSubsidiary: pulumi.String(mycart.OvhSubsidiary),
//				Description:   pulumi.String("my cloud project"),
//				PaymentMean:   pulumi.String("fidelity"),
//				Plan: &CloudProjectPlanArgs{
//					Duration:    pulumi.String(cloudOrderCartProductPlan.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(cloudOrderCartProductPlan.PlanCode),
//					PricingMode: pulumi.String(cloudOrderCartProductPlan.SelectedPrices[0].PricingMode),
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
type CloudProject struct {
	pulumi.CustomResourceState

	// project access
	Access pulumi.StringOutput `pulumi:"access"`
	// A description associated with the user.
	Description pulumi.StringOutput `pulumi:"description"`
	// Details about an Order
	Orders CloudProjectOrderArrayOutput `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan CloudProjectPlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions CloudProjectPlanOptionArrayOutput `pulumi:"planOptions"`
	// openstack project id
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// openstack project name
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// project status
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewCloudProject registers a new resource with the given unique name, arguments, and options.
func NewCloudProject(ctx *pulumi.Context,
	name string, args *CloudProjectArgs, opts ...pulumi.ResourceOption) (*CloudProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OvhSubsidiary == nil {
		return nil, errors.New("invalid value for required argument 'OvhSubsidiary'")
	}
	if args.PaymentMean == nil {
		return nil, errors.New("invalid value for required argument 'PaymentMean'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudProject
	err := ctx.RegisterResource("ovh:index/cloudProject:CloudProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProject gets an existing CloudProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectState, opts ...pulumi.ResourceOption) (*CloudProject, error) {
	var resource CloudProject
	err := ctx.ReadResource("ovh:index/cloudProject:CloudProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProject resources.
type cloudProjectState struct {
	// project access
	Access *string `pulumi:"access"`
	// A description associated with the user.
	Description *string `pulumi:"description"`
	// Details about an Order
	Orders []CloudProjectOrder `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *CloudProjectPlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []CloudProjectPlanOption `pulumi:"planOptions"`
	// openstack project id
	ProjectId *string `pulumi:"projectId"`
	// openstack project name
	ProjectName *string `pulumi:"projectName"`
	// project status
	Status *string `pulumi:"status"`
}

type CloudProjectState struct {
	// project access
	Access pulumi.StringPtrInput
	// A description associated with the user.
	Description pulumi.StringPtrInput
	// Details about an Order
	Orders CloudProjectOrderArrayInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan CloudProjectPlanPtrInput
	// Product Plan to order
	PlanOptions CloudProjectPlanOptionArrayInput
	// openstack project id
	ProjectId pulumi.StringPtrInput
	// openstack project name
	ProjectName pulumi.StringPtrInput
	// project status
	Status pulumi.StringPtrInput
}

func (CloudProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectState)(nil)).Elem()
}

type cloudProjectArgs struct {
	// A description associated with the user.
	Description *string `pulumi:"description"`
	// Ovh Subsidiary
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan CloudProjectPlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []CloudProjectPlanOption `pulumi:"planOptions"`
}

// The set of arguments for constructing a CloudProject resource.
type CloudProjectArgs struct {
	// A description associated with the user.
	Description pulumi.StringPtrInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringInput
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringInput
	// Product Plan to order
	Plan CloudProjectPlanInput
	// Product Plan to order
	PlanOptions CloudProjectPlanOptionArrayInput
}

func (CloudProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectArgs)(nil)).Elem()
}

type CloudProjectInput interface {
	pulumi.Input

	ToCloudProjectOutput() CloudProjectOutput
	ToCloudProjectOutputWithContext(ctx context.Context) CloudProjectOutput
}

func (*CloudProject) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProject)(nil)).Elem()
}

func (i *CloudProject) ToCloudProjectOutput() CloudProjectOutput {
	return i.ToCloudProjectOutputWithContext(context.Background())
}

func (i *CloudProject) ToCloudProjectOutputWithContext(ctx context.Context) CloudProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectOutput)
}

// CloudProjectArrayInput is an input type that accepts CloudProjectArray and CloudProjectArrayOutput values.
// You can construct a concrete instance of `CloudProjectArrayInput` via:
//
//	CloudProjectArray{ CloudProjectArgs{...} }
type CloudProjectArrayInput interface {
	pulumi.Input

	ToCloudProjectArrayOutput() CloudProjectArrayOutput
	ToCloudProjectArrayOutputWithContext(context.Context) CloudProjectArrayOutput
}

type CloudProjectArray []CloudProjectInput

func (CloudProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProject)(nil)).Elem()
}

func (i CloudProjectArray) ToCloudProjectArrayOutput() CloudProjectArrayOutput {
	return i.ToCloudProjectArrayOutputWithContext(context.Background())
}

func (i CloudProjectArray) ToCloudProjectArrayOutputWithContext(ctx context.Context) CloudProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectArrayOutput)
}

// CloudProjectMapInput is an input type that accepts CloudProjectMap and CloudProjectMapOutput values.
// You can construct a concrete instance of `CloudProjectMapInput` via:
//
//	CloudProjectMap{ "key": CloudProjectArgs{...} }
type CloudProjectMapInput interface {
	pulumi.Input

	ToCloudProjectMapOutput() CloudProjectMapOutput
	ToCloudProjectMapOutputWithContext(context.Context) CloudProjectMapOutput
}

type CloudProjectMap map[string]CloudProjectInput

func (CloudProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProject)(nil)).Elem()
}

func (i CloudProjectMap) ToCloudProjectMapOutput() CloudProjectMapOutput {
	return i.ToCloudProjectMapOutputWithContext(context.Background())
}

func (i CloudProjectMap) ToCloudProjectMapOutputWithContext(ctx context.Context) CloudProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectMapOutput)
}

type CloudProjectOutput struct{ *pulumi.OutputState }

func (CloudProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProject)(nil)).Elem()
}

func (o CloudProjectOutput) ToCloudProjectOutput() CloudProjectOutput {
	return o
}

func (o CloudProjectOutput) ToCloudProjectOutputWithContext(ctx context.Context) CloudProjectOutput {
	return o
}

// project access
func (o CloudProjectOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.Access }).(pulumi.StringOutput)
}

// A description associated with the user.
func (o CloudProjectOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Details about an Order
func (o CloudProjectOutput) Orders() CloudProjectOrderArrayOutput {
	return o.ApplyT(func(v *CloudProject) CloudProjectOrderArrayOutput { return v.Orders }).(CloudProjectOrderArrayOutput)
}

// Ovh Subsidiary
func (o CloudProjectOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
func (o CloudProjectOutput) PaymentMean() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.PaymentMean }).(pulumi.StringOutput)
}

// Product Plan to order
func (o CloudProjectOutput) Plan() CloudProjectPlanOutput {
	return o.ApplyT(func(v *CloudProject) CloudProjectPlanOutput { return v.Plan }).(CloudProjectPlanOutput)
}

// Product Plan to order
func (o CloudProjectOutput) PlanOptions() CloudProjectPlanOptionArrayOutput {
	return o.ApplyT(func(v *CloudProject) CloudProjectPlanOptionArrayOutput { return v.PlanOptions }).(CloudProjectPlanOptionArrayOutput)
}

// openstack project id
func (o CloudProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// openstack project name
func (o CloudProjectOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// project status
func (o CloudProjectOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type CloudProjectArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProject)(nil)).Elem()
}

func (o CloudProjectArrayOutput) ToCloudProjectArrayOutput() CloudProjectArrayOutput {
	return o
}

func (o CloudProjectArrayOutput) ToCloudProjectArrayOutputWithContext(ctx context.Context) CloudProjectArrayOutput {
	return o
}

func (o CloudProjectArrayOutput) Index(i pulumi.IntInput) CloudProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProject {
		return vs[0].([]*CloudProject)[vs[1].(int)]
	}).(CloudProjectOutput)
}

type CloudProjectMapOutput struct{ *pulumi.OutputState }

func (CloudProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProject)(nil)).Elem()
}

func (o CloudProjectMapOutput) ToCloudProjectMapOutput() CloudProjectMapOutput {
	return o
}

func (o CloudProjectMapOutput) ToCloudProjectMapOutputWithContext(ctx context.Context) CloudProjectMapOutput {
	return o
}

func (o CloudProjectMapOutput) MapIndex(k pulumi.StringInput) CloudProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProject {
		return vs[0].(map[string]*CloudProject)[vs[1].(string)]
	}).(CloudProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectInput)(nil)).Elem(), &CloudProject{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectArrayInput)(nil)).Elem(), CloudProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectMapInput)(nil)).Elem(), CloudProjectMap{})
	pulumi.RegisterOutputType(CloudProjectOutput{})
	pulumi.RegisterOutputType(CloudProjectArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectMapOutput{})
}
