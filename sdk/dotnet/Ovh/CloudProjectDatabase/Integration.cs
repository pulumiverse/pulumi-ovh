// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProjectDatabase
{
    /// <summary>
    /// Creates an integration for a database cluster associated with a public cloud project.
    /// 
    /// With this resource you can create an integration for all engine exept `mongodb`.
    /// 
    /// Please take a look at the list of available `types` in the `Argument references` section in order to know the list of avaulable integrations. For example, thanks to the integration feature you can have your PostgreSQL logs in your OpenSearch Database.
    /// 
    /// ## Example Usage
    /// 
    /// Push PostgreSQL logs in an OpenSearch DB:
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
    ///     var dbpostgresql = Ovh.CloudProjectDatabase.GetDatabase.Invoke(new()
    ///     {
    ///         ServiceName = "XXXX",
    ///         Engine = "postgresql",
    ///         Id = "ZZZZ",
    ///     });
    /// 
    ///     var dbopensearch = Ovh.CloudProjectDatabase.GetDatabase.Invoke(new()
    ///     {
    ///         ServiceName = "XXXX",
    ///         Engine = "opensearch",
    ///         Id = "ZZZZ",
    ///     });
    /// 
    ///     var integration = new Ovh.CloudProjectDatabase.Integration("integration", new()
    ///     {
    ///         ServiceName = dbpostgresql.Apply(getDatabaseResult =&gt; getDatabaseResult.ServiceName),
    ///         Engine = dbpostgresql.Apply(getDatabaseResult =&gt; getDatabaseResult.Engine),
    ///         ClusterId = dbpostgresql.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         SourceServiceId = dbpostgresql.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         DestinationServiceId = dbopensearch.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         Type = "opensearchLogs",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud Managed database clusters users can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the user, separated by "/" E.g., bash
    /// 
    /// ```sh
    ///  $ pulumi import ovh:CloudProjectDatabase/integration:Integration my_user service_name/engine/cluster_id/id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProjectDatabase/integration:Integration")]
    public partial class Integration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// ID of the destination service.
        /// </summary>
        [Output("destinationServiceId")]
        public Output<string> DestinationServiceId { get; private set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// All engines available exept `mongodb`.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Parameters for the integration.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// ID of the source service.
        /// </summary>
        [Output("sourceServiceId")]
        public Output<string> SourceServiceId { get; private set; } = null!;

        /// <summary>
        /// Current status of the integration.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Type of the integration.
        /// Available types:
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/integration:Integration", name, args ?? new IntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/integration:Integration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Integration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integration Get(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new Integration(name, id, state, options);
        }
    }

    public sealed class IntegrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// ID of the destination service.
        /// </summary>
        [Input("destinationServiceId", required: true)]
        public Input<string> DestinationServiceId { get; set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// All engines available exept `mongodb`.
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Parameters for the integration.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// ID of the source service.
        /// </summary>
        [Input("sourceServiceId", required: true)]
        public Input<string> SourceServiceId { get; set; } = null!;

        /// <summary>
        /// Type of the integration.
        /// Available types:
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IntegrationArgs()
        {
        }
        public static new IntegrationArgs Empty => new IntegrationArgs();
    }

    public sealed class IntegrationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// ID of the destination service.
        /// </summary>
        [Input("destinationServiceId")]
        public Input<string>? DestinationServiceId { get; set; }

        /// <summary>
        /// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// All engines available exept `mongodb`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Parameters for the integration.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// ID of the source service.
        /// </summary>
        [Input("sourceServiceId")]
        public Input<string>? SourceServiceId { get; set; }

        /// <summary>
        /// Current status of the integration.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Type of the integration.
        /// Available types:
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IntegrationState()
        {
        }
        public static new IntegrationState Empty => new IntegrationState();
    }
}
