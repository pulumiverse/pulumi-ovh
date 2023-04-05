// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject
{
    /// <summary>
    /// Manage a worflow that schedules backups of public cloud instance.
    /// Note that upon deletion, the workflow is deleted but any backups that have been created by this workflow are not.
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
    ///     var myBackup = new Ovh.CloudProject.WorkflowBackup("myBackup", new()
    ///     {
    ///         Cron = "50 4 * * *",
    ///         InstanceId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx",
    ///         MaxExecutionCount = 0,
    ///         RegionName = "GRA11",
    ///         Rotation = 7,
    ///         ServiceName = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProject/workflowBackup:WorkflowBackup")]
    public partial class WorkflowBackup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the backup files that are created. If empty, the `name` attribute is used.
        /// </summary>
        [Output("backupName")]
        public Output<string> BackupName { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The cron periodicity at which the backup workflow is scheduled
        /// </summary>
        [Output("cron")]
        public Output<string> Cron { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        /// </summary>
        [Output("maxExecutionCount")]
        public Output<int?> MaxExecutionCount { get; private set; } = null!;

        /// <summary>
        /// The worflow name that is used in the UI
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the openstack region.
        /// </summary>
        [Output("regionName")]
        public Output<string> RegionName { get; private set; } = null!;

        /// <summary>
        /// The number of backup that are retained.
        /// </summary>
        [Output("rotation")]
        public Output<int> Rotation { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a WorkflowBackup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkflowBackup(string name, WorkflowBackupArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/workflowBackup:WorkflowBackup", name, args ?? new WorkflowBackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkflowBackup(string name, Input<string> id, WorkflowBackupState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/workflowBackup:WorkflowBackup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkflowBackup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkflowBackup Get(string name, Input<string> id, WorkflowBackupState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkflowBackup(name, id, state, options);
        }
    }

    public sealed class WorkflowBackupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the backup files that are created. If empty, the `name` attribute is used.
        /// </summary>
        [Input("backupName")]
        public Input<string>? BackupName { get; set; }

        /// <summary>
        /// The cron periodicity at which the backup workflow is scheduled
        /// </summary>
        [Input("cron", required: true)]
        public Input<string> Cron { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        /// </summary>
        [Input("maxExecutionCount")]
        public Input<int>? MaxExecutionCount { get; set; }

        /// <summary>
        /// The worflow name that is used in the UI
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the openstack region.
        /// </summary>
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        /// <summary>
        /// The number of backup that are retained.
        /// </summary>
        [Input("rotation", required: true)]
        public Input<int> Rotation { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public WorkflowBackupArgs()
        {
        }
        public static new WorkflowBackupArgs Empty => new WorkflowBackupArgs();
    }

    public sealed class WorkflowBackupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the backup files that are created. If empty, the `name` attribute is used.
        /// </summary>
        [Input("backupName")]
        public Input<string>? BackupName { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The cron periodicity at which the backup workflow is scheduled
        /// </summary>
        [Input("cron")]
        public Input<string>? Cron { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        /// </summary>
        [Input("maxExecutionCount")]
        public Input<int>? MaxExecutionCount { get; set; }

        /// <summary>
        /// The worflow name that is used in the UI
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the openstack region.
        /// </summary>
        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        /// <summary>
        /// The number of backup that are retained.
        /// </summary>
        [Input("rotation")]
        public Input<int>? Rotation { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public WorkflowBackupState()
        {
        }
        public static new WorkflowBackupState Empty => new WorkflowBackupState();
    }
}
