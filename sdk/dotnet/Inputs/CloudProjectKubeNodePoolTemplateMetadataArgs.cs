// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Ovh.Inputs
{

    public sealed class CloudProjectKubeNodePoolTemplateMetadataArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        [Input("finalizers")]
        private InputList<string>? _finalizers;
        public InputList<string> Finalizers
        {
            get => _finalizers ?? (_finalizers = new InputList<string>());
            set => _finalizers = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        public CloudProjectKubeNodePoolTemplateMetadataArgs()
        {
        }
        public static new CloudProjectKubeNodePoolTemplateMetadataArgs Empty => new CloudProjectKubeNodePoolTemplateMetadataArgs();
    }
}
