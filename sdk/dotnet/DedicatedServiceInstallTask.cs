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
    [OvhResourceType("ovh:index/dedicatedServiceInstallTask:DedicatedServiceInstallTask")]
    public partial class DedicatedServiceInstallTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase.
        /// </summary>
        [Output("bootidOnDestroy")]
        public Output<int?> BootidOnDestroy { get; private set; } = null!;

        /// <summary>
        /// Details of this task. (should be `Install asked`)
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// see `details` block below.
        /// </summary>
        [Output("details")]
        public Output<Outputs.DedicatedServiceInstallTaskDetails?> Details { get; private set; } = null!;

        /// <summary>
        /// Completion date in RFC3339 format.
        /// </summary>
        [Output("doneDate")]
        public Output<string> DoneDate { get; private set; } = null!;

        /// <summary>
        /// Function name (should be `hardInstall`).
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// Last update in RFC3339 format.
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Output("partitionSchemeName")]
        public Output<string?> PartitionSchemeName { get; private set; } = null!;

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Task creation date in RFC3339 format.
        /// </summary>
        [Output("startDate")]
        public Output<string> StartDate { get; private set; } = null!;

        /// <summary>
        /// Task status (should be `done`)
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Template name.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedServiceInstallTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedServiceInstallTask(string name, DedicatedServiceInstallTaskArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/dedicatedServiceInstallTask:DedicatedServiceInstallTask", name, args ?? new DedicatedServiceInstallTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedServiceInstallTask(string name, Input<string> id, DedicatedServiceInstallTaskState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/dedicatedServiceInstallTask:DedicatedServiceInstallTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DedicatedServiceInstallTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedServiceInstallTask Get(string name, Input<string> id, DedicatedServiceInstallTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new DedicatedServiceInstallTask(name, id, state, options);
        }
    }

    public sealed class DedicatedServiceInstallTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase.
        /// </summary>
        [Input("bootidOnDestroy")]
        public Input<int>? BootidOnDestroy { get; set; }

        /// <summary>
        /// see `details` block below.
        /// </summary>
        [Input("details")]
        public Input<Inputs.DedicatedServiceInstallTaskDetailsArgs>? Details { get; set; }

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Input("partitionSchemeName")]
        public Input<string>? PartitionSchemeName { get; set; }

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Template name.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public DedicatedServiceInstallTaskArgs()
        {
        }
        public static new DedicatedServiceInstallTaskArgs Empty => new DedicatedServiceInstallTaskArgs();
    }

    public sealed class DedicatedServiceInstallTaskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase.
        /// </summary>
        [Input("bootidOnDestroy")]
        public Input<int>? BootidOnDestroy { get; set; }

        /// <summary>
        /// Details of this task. (should be `Install asked`)
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// see `details` block below.
        /// </summary>
        [Input("details")]
        public Input<Inputs.DedicatedServiceInstallTaskDetailsGetArgs>? Details { get; set; }

        /// <summary>
        /// Completion date in RFC3339 format.
        /// </summary>
        [Input("doneDate")]
        public Input<string>? DoneDate { get; set; }

        /// <summary>
        /// Function name (should be `hardInstall`).
        /// </summary>
        [Input("function")]
        public Input<string>? Function { get; set; }

        /// <summary>
        /// Last update in RFC3339 format.
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Input("partitionSchemeName")]
        public Input<string>? PartitionSchemeName { get; set; }

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Task creation date in RFC3339 format.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// Task status (should be `done`)
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Template name.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        public DedicatedServiceInstallTaskState()
        {
        }
        public static new DedicatedServiceInstallTaskState Empty => new DedicatedServiceInstallTaskState();
    }
}
