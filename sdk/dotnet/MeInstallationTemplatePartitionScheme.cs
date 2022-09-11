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
    /// Use this resource to create partition scheme for a custom installation template available for dedicated servers.
    /// 
    /// ## Import
    /// 
    /// Use the fake id format to import the resource `template_name/name`
    /// </summary>
    [OvhResourceType("ovh:index/meInstallationTemplatePartitionScheme:MeInstallationTemplatePartitionScheme")]
    public partial class MeInstallationTemplatePartitionScheme : global::Pulumi.CustomResource
    {
        /// <summary>
        /// name of this partitioning scheme
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
        /// among all the compatible partitioning schemes (given the underlying hardware specifications)
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// This template name
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;


        /// <summary>
        /// Create a MeInstallationTemplatePartitionScheme resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MeInstallationTemplatePartitionScheme(string name, MeInstallationTemplatePartitionSchemeArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/meInstallationTemplatePartitionScheme:MeInstallationTemplatePartitionScheme", name, args ?? new MeInstallationTemplatePartitionSchemeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MeInstallationTemplatePartitionScheme(string name, Input<string> id, MeInstallationTemplatePartitionSchemeState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/meInstallationTemplatePartitionScheme:MeInstallationTemplatePartitionScheme", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MeInstallationTemplatePartitionScheme resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MeInstallationTemplatePartitionScheme Get(string name, Input<string> id, MeInstallationTemplatePartitionSchemeState? state = null, CustomResourceOptions? options = null)
        {
            return new MeInstallationTemplatePartitionScheme(name, id, state, options);
        }
    }

    public sealed class MeInstallationTemplatePartitionSchemeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of this partitioning scheme
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
        /// among all the compatible partitioning schemes (given the underlying hardware specifications)
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// This template name
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public MeInstallationTemplatePartitionSchemeArgs()
        {
        }
        public static new MeInstallationTemplatePartitionSchemeArgs Empty => new MeInstallationTemplatePartitionSchemeArgs();
    }

    public sealed class MeInstallationTemplatePartitionSchemeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of this partitioning scheme
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
        /// among all the compatible partitioning schemes (given the underlying hardware specifications)
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// This template name
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        public MeInstallationTemplatePartitionSchemeState()
        {
        }
        public static new MeInstallationTemplatePartitionSchemeState Empty => new MeInstallationTemplatePartitionSchemeState();
    }
}
