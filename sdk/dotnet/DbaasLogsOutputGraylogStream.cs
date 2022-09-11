// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace lbrlabs.Ovh
{
    [OvhResourceType("ovh:index/dbaasLogsOutputGraylogStream:DbaasLogsOutputGraylogStream")]
    public partial class DbaasLogsOutputGraylogStream : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if the current user can create alert on the stream
        /// </summary>
        [Output("canAlert")]
        public Output<bool> CanAlert { get; private set; } = null!;

        /// <summary>
        /// Cold storage compression method
        /// </summary>
        [Output("coldStorageCompression")]
        public Output<string> ColdStorageCompression { get; private set; } = null!;

        /// <summary>
        /// ColdStorage content
        /// </summary>
        [Output("coldStorageContent")]
        public Output<string> ColdStorageContent { get; private set; } = null!;

        /// <summary>
        /// Is Cold storage enabled?
        /// </summary>
        [Output("coldStorageEnabled")]
        public Output<bool> ColdStorageEnabled { get; private set; } = null!;

        /// <summary>
        /// Notify on new Cold storage archive
        /// </summary>
        [Output("coldStorageNotifyEnabled")]
        public Output<bool> ColdStorageNotifyEnabled { get; private set; } = null!;

        /// <summary>
        /// Cold storage retention in year
        /// </summary>
        [Output("coldStorageRetention")]
        public Output<int> ColdStorageRetention { get; private set; } = null!;

        /// <summary>
        /// ColdStorage destination
        /// </summary>
        [Output("coldStorageTarget")]
        public Output<string> ColdStorageTarget { get; private set; } = null!;

        /// <summary>
        /// Stream creation
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Stream description
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Enable ES indexing
        /// </summary>
        [Output("indexingEnabled")]
        public Output<bool> IndexingEnabled { get; private set; } = null!;

        /// <summary>
        /// Maximum indexing size (in GB)
        /// </summary>
        [Output("indexingMaxSize")]
        public Output<int> IndexingMaxSize { get; private set; } = null!;

        /// <summary>
        /// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
        /// </summary>
        [Output("indexingNotifyEnabled")]
        public Output<bool> IndexingNotifyEnabled { get; private set; } = null!;

        /// <summary>
        /// Indicates if you are allowed to edit entry
        /// </summary>
        [Output("isEditable")]
        public Output<bool> IsEditable { get; private set; } = null!;

        /// <summary>
        /// Indicates if you are allowed to share entry
        /// </summary>
        [Output("isShareable")]
        public Output<bool> IsShareable { get; private set; } = null!;

        /// <summary>
        /// Number of alert condition
        /// </summary>
        [Output("nbAlertCondition")]
        public Output<int> NbAlertCondition { get; private set; } = null!;

        /// <summary>
        /// Number of coldstored archives
        /// </summary>
        [Output("nbArchive")]
        public Output<int> NbArchive { get; private set; } = null!;

        /// <summary>
        /// Parent stream ID
        /// </summary>
        [Output("parentStreamId")]
        public Output<string?> ParentStreamId { get; private set; } = null!;

        /// <summary>
        /// If set, pause indexing when maximum size is reach
        /// </summary>
        [Output("pauseIndexingOnMaxSize")]
        public Output<bool> PauseIndexingOnMaxSize { get; private set; } = null!;

        /// <summary>
        /// Retention ID
        /// </summary>
        [Output("retentionId")]
        public Output<string> RetentionId { get; private set; } = null!;

        /// <summary>
        /// The service name
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Stream ID
        /// </summary>
        [Output("streamId")]
        public Output<string> StreamId { get; private set; } = null!;

        /// <summary>
        /// Stream description
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Stream last update
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Enable Websocket
        /// </summary>
        [Output("webSocketEnabled")]
        public Output<bool> WebSocketEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a DbaasLogsOutputGraylogStream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbaasLogsOutputGraylogStream(string name, DbaasLogsOutputGraylogStreamArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/dbaasLogsOutputGraylogStream:DbaasLogsOutputGraylogStream", name, args ?? new DbaasLogsOutputGraylogStreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbaasLogsOutputGraylogStream(string name, Input<string> id, DbaasLogsOutputGraylogStreamState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/dbaasLogsOutputGraylogStream:DbaasLogsOutputGraylogStream", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DbaasLogsOutputGraylogStream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbaasLogsOutputGraylogStream Get(string name, Input<string> id, DbaasLogsOutputGraylogStreamState? state = null, CustomResourceOptions? options = null)
        {
            return new DbaasLogsOutputGraylogStream(name, id, state, options);
        }
    }

    public sealed class DbaasLogsOutputGraylogStreamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cold storage compression method
        /// </summary>
        [Input("coldStorageCompression")]
        public Input<string>? ColdStorageCompression { get; set; }

        /// <summary>
        /// ColdStorage content
        /// </summary>
        [Input("coldStorageContent")]
        public Input<string>? ColdStorageContent { get; set; }

        /// <summary>
        /// Is Cold storage enabled?
        /// </summary>
        [Input("coldStorageEnabled")]
        public Input<bool>? ColdStorageEnabled { get; set; }

        /// <summary>
        /// Notify on new Cold storage archive
        /// </summary>
        [Input("coldStorageNotifyEnabled")]
        public Input<bool>? ColdStorageNotifyEnabled { get; set; }

        /// <summary>
        /// Cold storage retention in year
        /// </summary>
        [Input("coldStorageRetention")]
        public Input<int>? ColdStorageRetention { get; set; }

        /// <summary>
        /// ColdStorage destination
        /// </summary>
        [Input("coldStorageTarget")]
        public Input<string>? ColdStorageTarget { get; set; }

        /// <summary>
        /// Stream description
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Enable ES indexing
        /// </summary>
        [Input("indexingEnabled")]
        public Input<bool>? IndexingEnabled { get; set; }

        /// <summary>
        /// Maximum indexing size (in GB)
        /// </summary>
        [Input("indexingMaxSize")]
        public Input<int>? IndexingMaxSize { get; set; }

        /// <summary>
        /// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
        /// </summary>
        [Input("indexingNotifyEnabled")]
        public Input<bool>? IndexingNotifyEnabled { get; set; }

        /// <summary>
        /// Parent stream ID
        /// </summary>
        [Input("parentStreamId")]
        public Input<string>? ParentStreamId { get; set; }

        /// <summary>
        /// If set, pause indexing when maximum size is reach
        /// </summary>
        [Input("pauseIndexingOnMaxSize")]
        public Input<bool>? PauseIndexingOnMaxSize { get; set; }

        /// <summary>
        /// Retention ID
        /// </summary>
        [Input("retentionId")]
        public Input<string>? RetentionId { get; set; }

        /// <summary>
        /// The service name
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Stream description
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// Enable Websocket
        /// </summary>
        [Input("webSocketEnabled")]
        public Input<bool>? WebSocketEnabled { get; set; }

        public DbaasLogsOutputGraylogStreamArgs()
        {
        }
        public static new DbaasLogsOutputGraylogStreamArgs Empty => new DbaasLogsOutputGraylogStreamArgs();
    }

    public sealed class DbaasLogsOutputGraylogStreamState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the current user can create alert on the stream
        /// </summary>
        [Input("canAlert")]
        public Input<bool>? CanAlert { get; set; }

        /// <summary>
        /// Cold storage compression method
        /// </summary>
        [Input("coldStorageCompression")]
        public Input<string>? ColdStorageCompression { get; set; }

        /// <summary>
        /// ColdStorage content
        /// </summary>
        [Input("coldStorageContent")]
        public Input<string>? ColdStorageContent { get; set; }

        /// <summary>
        /// Is Cold storage enabled?
        /// </summary>
        [Input("coldStorageEnabled")]
        public Input<bool>? ColdStorageEnabled { get; set; }

        /// <summary>
        /// Notify on new Cold storage archive
        /// </summary>
        [Input("coldStorageNotifyEnabled")]
        public Input<bool>? ColdStorageNotifyEnabled { get; set; }

        /// <summary>
        /// Cold storage retention in year
        /// </summary>
        [Input("coldStorageRetention")]
        public Input<int>? ColdStorageRetention { get; set; }

        /// <summary>
        /// ColdStorage destination
        /// </summary>
        [Input("coldStorageTarget")]
        public Input<string>? ColdStorageTarget { get; set; }

        /// <summary>
        /// Stream creation
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Stream description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable ES indexing
        /// </summary>
        [Input("indexingEnabled")]
        public Input<bool>? IndexingEnabled { get; set; }

        /// <summary>
        /// Maximum indexing size (in GB)
        /// </summary>
        [Input("indexingMaxSize")]
        public Input<int>? IndexingMaxSize { get; set; }

        /// <summary>
        /// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
        /// </summary>
        [Input("indexingNotifyEnabled")]
        public Input<bool>? IndexingNotifyEnabled { get; set; }

        /// <summary>
        /// Indicates if you are allowed to edit entry
        /// </summary>
        [Input("isEditable")]
        public Input<bool>? IsEditable { get; set; }

        /// <summary>
        /// Indicates if you are allowed to share entry
        /// </summary>
        [Input("isShareable")]
        public Input<bool>? IsShareable { get; set; }

        /// <summary>
        /// Number of alert condition
        /// </summary>
        [Input("nbAlertCondition")]
        public Input<int>? NbAlertCondition { get; set; }

        /// <summary>
        /// Number of coldstored archives
        /// </summary>
        [Input("nbArchive")]
        public Input<int>? NbArchive { get; set; }

        /// <summary>
        /// Parent stream ID
        /// </summary>
        [Input("parentStreamId")]
        public Input<string>? ParentStreamId { get; set; }

        /// <summary>
        /// If set, pause indexing when maximum size is reach
        /// </summary>
        [Input("pauseIndexingOnMaxSize")]
        public Input<bool>? PauseIndexingOnMaxSize { get; set; }

        /// <summary>
        /// Retention ID
        /// </summary>
        [Input("retentionId")]
        public Input<string>? RetentionId { get; set; }

        /// <summary>
        /// The service name
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Stream ID
        /// </summary>
        [Input("streamId")]
        public Input<string>? StreamId { get; set; }

        /// <summary>
        /// Stream description
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Stream last update
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Enable Websocket
        /// </summary>
        [Input("webSocketEnabled")]
        public Input<bool>? WebSocketEnabled { get; set; }

        public DbaasLogsOutputGraylogStreamState()
        {
        }
        public static new DbaasLogsOutputGraylogStreamState Empty => new DbaasLogsOutputGraylogStreamState();
    }
}
