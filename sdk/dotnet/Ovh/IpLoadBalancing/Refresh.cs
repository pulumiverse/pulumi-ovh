// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.IpLoadBalancing
{
    /// <summary>
    /// Applies changes from other `ovh_iploadbalancing_*` resources to the production configuration of loadbalancers.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Lbrlabs.PulumiPackage.Ovh;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var lb = Ovh.IpLoadBalancing.GetIpLoadBalancing.Invoke(new()
    ///     {
    ///         ServiceName = "ip-1.2.3.4",
    ///         State = "ok",
    ///     });
    /// 
    ///     var farmname = new Ovh.IpLoadBalancing.TcpFarm("farmname", new()
    ///     {
    ///         ServiceName = lb.Apply(getIpLoadBalancingResult =&gt; getIpLoadBalancingResult.Id),
    ///         Port = 8080,
    ///         Zone = "all",
    ///     });
    /// 
    ///     var backend = new Ovh.IpLoadBalancing.TcpFarmServer("backend", new()
    ///     {
    ///         ServiceName = lb.Apply(getIpLoadBalancingResult =&gt; getIpLoadBalancingResult.Id),
    ///         FarmId = farmname.Id,
    ///         DisplayName = "mybackend",
    ///         Address = "4.5.6.7",
    ///         Status = "active",
    ///         Port = 80,
    ///         ProxyProtocolVersion = v2,
    ///         Weight = 2,
    ///         Probe = true,
    ///         Ssl = false,
    ///         Backup = true,
    ///     });
    /// 
    ///     var mylb = new Ovh.IpLoadBalancing.Refresh("mylb", new()
    ///     {
    ///         ServiceName = lb.Apply(getIpLoadBalancingResult =&gt; getIpLoadBalancingResult.Id),
    ///         Keepers = new[]
    ///         {
    ///             new[]
    ///             {
    ///                 backend,
    ///             }.Select(__item =&gt; __item.Address).ToList(),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:IpLoadBalancing/refresh:Refresh")]
    public partial class Refresh : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of values tracked to trigger refresh, used also to form implicit dependencies
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableArray<string>> Keepers { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a Refresh resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Refresh(string name, RefreshArgs args, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/refresh:Refresh", name, args ?? new RefreshArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Refresh(string name, Input<string> id, RefreshState? state = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/refresh:Refresh", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Refresh resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Refresh Get(string name, Input<string> id, RefreshState? state = null, CustomResourceOptions? options = null)
        {
            return new Refresh(name, id, state, options);
        }
    }

    public sealed class RefreshArgs : global::Pulumi.ResourceArgs
    {
        [Input("keepers", required: true)]
        private InputList<string>? _keepers;

        /// <summary>
        /// List of values tracked to trigger refresh, used also to form implicit dependencies
        /// </summary>
        public InputList<string> Keepers
        {
            get => _keepers ?? (_keepers = new InputList<string>());
            set => _keepers = value;
        }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public RefreshArgs()
        {
        }
        public static new RefreshArgs Empty => new RefreshArgs();
    }

    public sealed class RefreshState : global::Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputList<string>? _keepers;

        /// <summary>
        /// List of values tracked to trigger refresh, used also to form implicit dependencies
        /// </summary>
        public InputList<string> Keepers
        {
            get => _keepers ?? (_keepers = new InputList<string>());
            set => _keepers = value;
        }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public RefreshState()
        {
        }
        public static new RefreshState Empty => new RefreshState();
    }
}