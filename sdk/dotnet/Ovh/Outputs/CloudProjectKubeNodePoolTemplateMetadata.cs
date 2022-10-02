// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Outputs
{

    [OutputType]
    public sealed class CloudProjectKubeNodePoolTemplateMetadata
    {
        public readonly ImmutableDictionary<string, string>? Annotations;
        public readonly ImmutableArray<string> Finalizers;
        public readonly ImmutableDictionary<string, string>? Labels;

        [OutputConstructor]
        private CloudProjectKubeNodePoolTemplateMetadata(
            ImmutableDictionary<string, string>? annotations,

            ImmutableArray<string> finalizers,

            ImmutableDictionary<string, string>? labels)
        {
            Annotations = annotations;
            Finalizers = finalizers;
            Labels = labels;
        }
    }
}