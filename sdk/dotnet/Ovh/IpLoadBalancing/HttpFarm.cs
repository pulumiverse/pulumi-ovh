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
    /// Creates a HTTP backend server group (farm) to be used by loadbalancing frontend(s)
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
    ///     var farmname = new Ovh.IpLoadBalancing.HttpFarm("farmname", new()
    ///     {
    ///         DisplayName = "ingress-8080-gra",
    ///         ServiceName = lb.Apply(getIpLoadBalancingResult =&gt; getIpLoadBalancingResult.ServiceName),
    ///         Zone = "GRA",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:IpLoadBalancing/httpFarm:HttpFarm")]
    public partial class HttpFarm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        /// </summary>
        [Output("balance")]
        public Output<string?> Balance { get; private set; } = null!;

        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Port for backends to receive traffic on.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// define a backend healthcheck probe
        /// </summary>
        [Output("probe")]
        public Output<Outputs.HttpFarmProbe?> Probe { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
        /// </summary>
        [Output("stickiness")]
        public Output<string?> Stickiness { get; private set; } = null!;

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Output("vrackNetworkId")]
        public Output<int?> VrackNetworkId { get; private set; } = null!;

        /// <summary>
        /// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a HttpFarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HttpFarm(string name, HttpFarmArgs args, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/httpFarm:HttpFarm", name, args ?? new HttpFarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HttpFarm(string name, Input<string> id, HttpFarmState? state = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/httpFarm:HttpFarm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HttpFarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HttpFarm Get(string name, Input<string> id, HttpFarmState? state = null, CustomResourceOptions? options = null)
        {
            return new HttpFarm(name, id, state, options);
        }
    }

    public sealed class HttpFarmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        /// </summary>
        [Input("balance")]
        public Input<string>? Balance { get; set; }

        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Port for backends to receive traffic on.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// define a backend healthcheck probe
        /// </summary>
        [Input("probe")]
        public Input<Inputs.HttpFarmProbeArgs>? Probe { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
        /// </summary>
        [Input("stickiness")]
        public Input<string>? Stickiness { get; set; }

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Input("vrackNetworkId")]
        public Input<int>? VrackNetworkId { get; set; }

        /// <summary>
        /// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public HttpFarmArgs()
        {
        }
        public static new HttpFarmArgs Empty => new HttpFarmArgs();
    }

    public sealed class HttpFarmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        /// </summary>
        [Input("balance")]
        public Input<string>? Balance { get; set; }

        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Port for backends to receive traffic on.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// define a backend healthcheck probe
        /// </summary>
        [Input("probe")]
        public Input<Inputs.HttpFarmProbeGetArgs>? Probe { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
        /// </summary>
        [Input("stickiness")]
        public Input<string>? Stickiness { get; set; }

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Input("vrackNetworkId")]
        public Input<int>? VrackNetworkId { get; set; }

        /// <summary>
        /// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public HttpFarmState()
        {
        }
        public static new HttpFarmState Empty => new HttpFarmState();
    }
}
