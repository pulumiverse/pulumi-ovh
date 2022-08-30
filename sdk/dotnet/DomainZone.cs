// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    /// <summary>
    /// Creates a domain zone.
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
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mycart = Ovh.GetOrderCart.Invoke(new()
    ///     {
    ///         OvhSubsidiary = "fr",
    ///     });
    /// 
    ///     var zoneOrderCartProductPlan = Ovh.GetOrderCartProductPlan.Invoke(new()
    ///     {
    ///         CartId = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.Id),
    ///         PriceCapacity = "renew",
    ///         Product = "dns",
    ///         PlanCode = "zone",
    ///     });
    /// 
    ///     var zoneDomainZone = new Ovh.DomainZone("zoneDomainZone", new()
    ///     {
    ///         OvhSubsidiary = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.OvhSubsidiary),
    ///         PaymentMean = "fidelity",
    ///         Plan = new Ovh.Inputs.DomainZonePlanArgs
    ///         {
    ///             Duration = zoneOrderCartProductPlan.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.SelectedPrices[0]?.Duration),
    ///             PlanCode = zoneOrderCartProductPlan.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.PlanCode),
    ///             PricingMode = zoneOrderCartProductPlan.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.SelectedPrices[0]?.PricingMode),
    ///             Configurations = new[]
    ///             {
    ///                 new Ovh.Inputs.DomainZonePlanConfigurationArgs
    ///                 {
    ///                     Label = "zone",
    ///                     Value = "myzone.mydomain.com",
    ///                 },
    ///                 new Ovh.Inputs.DomainZonePlanConfigurationArgs
    ///                 {
    ///                     Label = "template",
    ///                     Value = "minimized",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:index/domainZone:DomainZone")]
    public partial class DomainZone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Is DNSSEC supported by this zone
        /// </summary>
        [Output("dnssecSupported")]
        public Output<bool> DnssecSupported { get; private set; } = null!;

        /// <summary>
        /// hasDnsAnycast flag of the DNS zone
        /// </summary>
        [Output("hasDnsAnycast")]
        public Output<bool> HasDnsAnycast { get; private set; } = null!;

        /// <summary>
        /// Last update date of the DNS zone
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// Zone name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name servers that host the DNS zone
        /// </summary>
        [Output("nameServers")]
        public Output<ImmutableArray<string>> NameServers { get; private set; } = null!;

        /// <summary>
        /// Details about an Order
        /// </summary>
        [Output("orders")]
        public Output<ImmutableArray<Outputs.DomainZoneOrder>> Orders { get; private set; } = null!;

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
        public Output<Outputs.DomainZonePlan> Plan { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("planOptions")]
        public Output<ImmutableArray<Outputs.DomainZonePlanOption>> PlanOptions { get; private set; } = null!;


        /// <summary>
        /// Create a DomainZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainZone(string name, DomainZoneArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/domainZone:DomainZone", name, args ?? new DomainZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainZone(string name, Input<string> id, DomainZoneState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/domainZone:DomainZone", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainZone Get(string name, Input<string> id, DomainZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainZone(name, id, state, options);
        }
    }

    public sealed class DomainZoneArgs : global::Pulumi.ResourceArgs
    {
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
        public Input<Inputs.DomainZonePlanArgs> Plan { get; set; } = null!;

        [Input("planOptions")]
        private InputList<Inputs.DomainZonePlanOptionArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.DomainZonePlanOptionArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.DomainZonePlanOptionArgs>());
            set => _planOptions = value;
        }

        public DomainZoneArgs()
        {
        }
        public static new DomainZoneArgs Empty => new DomainZoneArgs();
    }

    public sealed class DomainZoneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is DNSSEC supported by this zone
        /// </summary>
        [Input("dnssecSupported")]
        public Input<bool>? DnssecSupported { get; set; }

        /// <summary>
        /// hasDnsAnycast flag of the DNS zone
        /// </summary>
        [Input("hasDnsAnycast")]
        public Input<bool>? HasDnsAnycast { get; set; }

        /// <summary>
        /// Last update date of the DNS zone
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// Zone name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameServers")]
        private InputList<string>? _nameServers;

        /// <summary>
        /// Name servers that host the DNS zone
        /// </summary>
        public InputList<string> NameServers
        {
            get => _nameServers ?? (_nameServers = new InputList<string>());
            set => _nameServers = value;
        }

        [Input("orders")]
        private InputList<Inputs.DomainZoneOrderGetArgs>? _orders;

        /// <summary>
        /// Details about an Order
        /// </summary>
        public InputList<Inputs.DomainZoneOrderGetArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.DomainZoneOrderGetArgs>());
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
        public Input<Inputs.DomainZonePlanGetArgs>? Plan { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.DomainZonePlanOptionGetArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.DomainZonePlanOptionGetArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.DomainZonePlanOptionGetArgs>());
            set => _planOptions = value;
        }

        public DomainZoneState()
        {
        }
        public static new DomainZoneState Empty => new DomainZoneState();
    }
}
