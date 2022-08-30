// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Ovh.Outputs
{

    [OutputType]
    public sealed class CloudProjectNetworkPrivateRegionsStatus
    {
        public readonly string? Region;
        /// <summary>
        /// the status of the network. should be normally set to 'ACTIVE'.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private CloudProjectNetworkPrivateRegionsStatus(
            string? region,

            string status)
        {
            Region = region;
            Status = status;
        }
    }
}
