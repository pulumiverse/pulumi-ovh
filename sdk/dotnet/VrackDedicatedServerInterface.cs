// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Ovh
{
    /// <summary>
    /// Attach a Dedicated Server Network Interface to a VRack.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Ovh = Pulumiverse.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vdsi = new Ovh.VrackDedicatedServerInterface("vdsi", new()
    ///     {
    ///         InterfaceId = "67890",
    ///         ServiceName = "12345",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:index/vrackDedicatedServerInterface:VrackDedicatedServerInterface")]
    public partial class VrackDedicatedServerInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of dedicated server network interface.
        /// </summary>
        [Output("interfaceId")]
        public Output<string> InterfaceId { get; private set; } = null!;

        /// <summary>
        /// The id of the vrack. If omitted,
        /// the `OVH_VRACK_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a VrackDedicatedServerInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VrackDedicatedServerInterface(string name, VrackDedicatedServerInterfaceArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/vrackDedicatedServerInterface:VrackDedicatedServerInterface", name, args ?? new VrackDedicatedServerInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VrackDedicatedServerInterface(string name, Input<string> id, VrackDedicatedServerInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/vrackDedicatedServerInterface:VrackDedicatedServerInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VrackDedicatedServerInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VrackDedicatedServerInterface Get(string name, Input<string> id, VrackDedicatedServerInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new VrackDedicatedServerInterface(name, id, state, options);
        }
    }

    public sealed class VrackDedicatedServerInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of dedicated server network interface.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        /// <summary>
        /// The id of the vrack. If omitted,
        /// the `OVH_VRACK_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public VrackDedicatedServerInterfaceArgs()
        {
        }
        public static new VrackDedicatedServerInterfaceArgs Empty => new VrackDedicatedServerInterfaceArgs();
    }

    public sealed class VrackDedicatedServerInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of dedicated server network interface.
        /// </summary>
        [Input("interfaceId")]
        public Input<string>? InterfaceId { get; set; }

        /// <summary>
        /// The id of the vrack. If omitted,
        /// the `OVH_VRACK_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public VrackDedicatedServerInterfaceState()
        {
        }
        public static new VrackDedicatedServerInterfaceState Empty => new VrackDedicatedServerInterfaceState();
    }
}
