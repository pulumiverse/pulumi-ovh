// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Dedicated
{
    /// <summary>
    /// Provides a resource for managing access rights to partitions on HA-NAS services
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Lbrlabs.PulumiPackage.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_partition = new Ovh.Dedicated.NasHAPartitionAccess("my-partition", new()
    ///     {
    ///         Ip = "123.123.123.123/32",
    ///         PartitionName = "my-partition",
    ///         ServiceName = "zpool-12345",
    ///         Type = "readwrite",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// HA-NAS partition access can be imported using the `{service_name}/{partition_name}/{ip}`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess my-partition zpool-12345/my-partition/123.123.123.123%2F32`
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess")]
    public partial class NasHAPartitionAccess : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ip block in x.x.x.x/x format
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// name of the partition
        /// </summary>
        [Output("partitionName")]
        public Output<string> PartitionName { get; private set; } = null!;

        /// <summary>
        /// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// one of "readwrite", "readonly"
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a NasHAPartitionAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NasHAPartitionAccess(string name, NasHAPartitionAccessArgs args, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess", name, args ?? new NasHAPartitionAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NasHAPartitionAccess(string name, Input<string> id, NasHAPartitionAccessState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NasHAPartitionAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NasHAPartitionAccess Get(string name, Input<string> id, NasHAPartitionAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new NasHAPartitionAccess(name, id, state, options);
        }
    }

    public sealed class NasHAPartitionAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ip block in x.x.x.x/x format
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// name of the partition
        /// </summary>
        [Input("partitionName", required: true)]
        public Input<string> PartitionName { get; set; } = null!;

        /// <summary>
        /// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// one of "readwrite", "readonly"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NasHAPartitionAccessArgs()
        {
        }
        public static new NasHAPartitionAccessArgs Empty => new NasHAPartitionAccessArgs();
    }

    public sealed class NasHAPartitionAccessState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ip block in x.x.x.x/x format
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// name of the partition
        /// </summary>
        [Input("partitionName")]
        public Input<string>? PartitionName { get; set; }

        /// <summary>
        /// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// one of "readwrite", "readonly"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NasHAPartitionAccessState()
        {
        }
        public static new NasHAPartitionAccessState Empty => new NasHAPartitionAccessState();
    }
}
