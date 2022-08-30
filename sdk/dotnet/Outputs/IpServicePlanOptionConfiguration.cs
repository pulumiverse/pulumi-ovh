// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class IpServicePlanOptionConfiguration
    {
        /// <summary>
        /// Identifier of the resource
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Path to the resource in API.OVH.COM
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private IpServicePlanOptionConfiguration(
            string label,

            string value)
        {
            Label = label;
            Value = value;
        }
    }
}
