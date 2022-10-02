// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh
{
    /// <summary>
    /// Orders a public cloud project.
    /// 
    /// ## Important
    /// 
    /// This resource is in beta state. Use with caution.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Ovh = Lbrlabs.PulumiPackage.Ovh;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mycart = Ovh.GetOrderCart.Invoke(new()
    ///     {
    ///         OvhSubsidiary = "fr",
    ///         Description = "my cloud order cart",
    ///     });
    /// 
    ///     var cloudOrderCartProductPlan = Ovh.GetOrderCartProductPlan.Invoke(new()
    ///     {
    ///         CartId = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.Id),
    ///         PriceCapacity = "renew",
    ///         Product = "cloud",
    ///         PlanCode = "project.2018",
    ///     });
    /// 
    ///     var cloudCloudProject = new Ovh.CloudProject("cloudCloudProject", new()
    ///     {
    ///         OvhSubsidiary = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.OvhSubsidiary),
    ///         Description = "my cloud project",
    ///         PaymentMean = "fidelity",
    ///         Plan = new Ovh.Inputs.CloudProjectPlanArgs
    ///         {
    ///             Duration = cloudOrderCartProductPlan.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.SelectedPrices[0]?.Duration),
    ///             PlanCode = cloudOrderCartProductPlan.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.PlanCode),
    ///             PricingMode = cloudOrderCartProductPlan.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.SelectedPrices[0]?.PricingMode),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:index/cloudProject:CloudProject")]
    public partial class CloudProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// project access
        /// </summary>
        [Output("access")]
        public Output<string> Access { get; private set; } = null!;

        /// <summary>
        /// A description associated with the user.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Details about an Order
        /// </summary>
        [Output("orders")]
        public Output<ImmutableArray<Outputs.CloudProjectOrder>> Orders { get; private set; } = null!;

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Output("ovhSubsidiary")]
        public Output<string> OvhSubsidiary { get; private set; } = null!;

        /// <summary>
        /// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
        /// </summary>
        [Output("paymentMean")]
        public Output<string> PaymentMean { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("plan")]
        public Output<Outputs.CloudProjectPlan> Plan { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("planOptions")]
        public Output<ImmutableArray<Outputs.CloudProjectPlanOption>> PlanOptions { get; private set; } = null!;

        /// <summary>
        /// openstack project id
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// openstack project name
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// project status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CloudProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudProject(string name, CloudProjectArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/cloudProject:CloudProject", name, args ?? new CloudProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudProject(string name, Input<string> id, CloudProjectState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/cloudProject:CloudProject", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudProject Get(string name, Input<string> id, CloudProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudProject(name, id, state, options);
        }
    }

    public sealed class CloudProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description associated with the user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Input("ovhSubsidiary", required: true)]
        public Input<string> OvhSubsidiary { get; set; } = null!;

        /// <summary>
        /// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
        /// </summary>
        [Input("paymentMean", required: true)]
        public Input<string> PaymentMean { get; set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Input("plan", required: true)]
        public Input<Inputs.CloudProjectPlanArgs> Plan { get; set; } = null!;

        [Input("planOptions")]
        private InputList<Inputs.CloudProjectPlanOptionArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.CloudProjectPlanOptionArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.CloudProjectPlanOptionArgs>());
            set => _planOptions = value;
        }

        public CloudProjectArgs()
        {
        }
        public static new CloudProjectArgs Empty => new CloudProjectArgs();
    }

    public sealed class CloudProjectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// project access
        /// </summary>
        [Input("access")]
        public Input<string>? Access { get; set; }

        /// <summary>
        /// A description associated with the user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("orders")]
        private InputList<Inputs.CloudProjectOrderGetArgs>? _orders;

        /// <summary>
        /// Details about an Order
        /// </summary>
        public InputList<Inputs.CloudProjectOrderGetArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.CloudProjectOrderGetArgs>());
            set => _orders = value;
        }

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Input("ovhSubsidiary")]
        public Input<string>? OvhSubsidiary { get; set; }

        /// <summary>
        /// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
        /// </summary>
        [Input("paymentMean")]
        public Input<string>? PaymentMean { get; set; }

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Input("plan")]
        public Input<Inputs.CloudProjectPlanGetArgs>? Plan { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.CloudProjectPlanOptionGetArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.CloudProjectPlanOptionGetArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.CloudProjectPlanOptionGetArgs>());
            set => _planOptions = value;
        }

        /// <summary>
        /// openstack project id
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// openstack project name
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// project status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public CloudProjectState()
        {
        }
        public static new CloudProjectState Empty => new CloudProjectState();
    }
}