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
    /// <summary>
    /// Use this resource to create a partition in the partition scheme of a custom installation template available for dedicated servers.
    /// 
    /// ## Import
    /// 
    /// Use the fake id format to import the resource `template_name/scheme_name/mountpoint` (example"mytemplate/myscheme//").
    /// </summary>
    [OvhResourceType("ovh:index/meInstallationTemplatePartitionSchemePartition:MeInstallationTemplatePartitionSchemePartition")]
    public partial class MeInstallationTemplatePartitionSchemePartition : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Partition filesystem
        /// </summary>
        [Output("filesystem")]
        public Output<string> Filesystem { get; private set; } = null!;

        /// <summary>
        /// partition mount point
        /// </summary>
        [Output("mountpoint")]
        public Output<string> Mountpoint { get; private set; } = null!;

        /// <summary>
        /// step or order. specifies the creation order of the partition on the disk
        /// </summary>
        [Output("order")]
        public Output<int> Order { get; private set; } = null!;

        /// <summary>
        /// raid partition type
        /// </summary>
        [Output("raid")]
        public Output<string> Raid { get; private set; } = null!;

        /// <summary>
        /// name of this partitioning scheme
        /// </summary>
        [Output("schemeName")]
        public Output<string> SchemeName { get; private set; } = null!;

        /// <summary>
        /// size of partition in MB, 0 =&gt; rest of the space
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Template name
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// partition type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The volume name needed for proxmox distribution
        /// </summary>
        [Output("volumeName")]
        public Output<string> VolumeName { get; private set; } = null!;


        /// <summary>
        /// Create a MeInstallationTemplatePartitionSchemePartition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MeInstallationTemplatePartitionSchemePartition(string name, MeInstallationTemplatePartitionSchemePartitionArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/meInstallationTemplatePartitionSchemePartition:MeInstallationTemplatePartitionSchemePartition", name, args ?? new MeInstallationTemplatePartitionSchemePartitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MeInstallationTemplatePartitionSchemePartition(string name, Input<string> id, MeInstallationTemplatePartitionSchemePartitionState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/meInstallationTemplatePartitionSchemePartition:MeInstallationTemplatePartitionSchemePartition", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MeInstallationTemplatePartitionSchemePartition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MeInstallationTemplatePartitionSchemePartition Get(string name, Input<string> id, MeInstallationTemplatePartitionSchemePartitionState? state = null, CustomResourceOptions? options = null)
        {
            return new MeInstallationTemplatePartitionSchemePartition(name, id, state, options);
        }
    }

    public sealed class MeInstallationTemplatePartitionSchemePartitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Partition filesystem
        /// </summary>
        [Input("filesystem", required: true)]
        public Input<string> Filesystem { get; set; } = null!;

        /// <summary>
        /// partition mount point
        /// </summary>
        [Input("mountpoint", required: true)]
        public Input<string> Mountpoint { get; set; } = null!;

        /// <summary>
        /// step or order. specifies the creation order of the partition on the disk
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// raid partition type
        /// </summary>
        [Input("raid")]
        public Input<string>? Raid { get; set; }

        /// <summary>
        /// name of this partitioning scheme
        /// </summary>
        [Input("schemeName", required: true)]
        public Input<string> SchemeName { get; set; } = null!;

        /// <summary>
        /// size of partition in MB, 0 =&gt; rest of the space
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// Template name
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        /// <summary>
        /// partition type
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The volume name needed for proxmox distribution
        /// </summary>
        [Input("volumeName")]
        public Input<string>? VolumeName { get; set; }

        public MeInstallationTemplatePartitionSchemePartitionArgs()
        {
        }
        public static new MeInstallationTemplatePartitionSchemePartitionArgs Empty => new MeInstallationTemplatePartitionSchemePartitionArgs();
    }

    public sealed class MeInstallationTemplatePartitionSchemePartitionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Partition filesystem
        /// </summary>
        [Input("filesystem")]
        public Input<string>? Filesystem { get; set; }

        /// <summary>
        /// partition mount point
        /// </summary>
        [Input("mountpoint")]
        public Input<string>? Mountpoint { get; set; }

        /// <summary>
        /// step or order. specifies the creation order of the partition on the disk
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// raid partition type
        /// </summary>
        [Input("raid")]
        public Input<string>? Raid { get; set; }

        /// <summary>
        /// name of this partitioning scheme
        /// </summary>
        [Input("schemeName")]
        public Input<string>? SchemeName { get; set; }

        /// <summary>
        /// size of partition in MB, 0 =&gt; rest of the space
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Template name
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// partition type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The volume name needed for proxmox distribution
        /// </summary>
        [Input("volumeName")]
        public Input<string>? VolumeName { get; set; }

        public MeInstallationTemplatePartitionSchemePartitionState()
        {
        }
        public static new MeInstallationTemplatePartitionSchemePartitionState Empty => new MeInstallationTemplatePartitionSchemePartitionState();
    }
}
