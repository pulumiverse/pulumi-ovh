// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace lbrlabs.Ovh.Outputs
{

    [OutputType]
    public sealed class CloudProjectNetworkPrivateSubnetIpPool
    {
        /// <summary>
        /// Enable DHCP.
        /// Changing this forces a new resource to be created. Defaults to false.
        /// _
        /// </summary>
        public readonly bool? Dhcp;
        /// <summary>
        /// Last ip for this region.
        /// Changing this value recreates the subnet.
        /// </summary>
        public readonly string? End;
        /// <summary>
        /// Global network in CIDR format.
        /// Changing this value recreates the subnet
        /// </summary>
        public readonly string? Network;
        /// <summary>
        /// The region in which the network subnet will be created.
        /// Ex.: "GRA1". Changing this value recreates the resource.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// First ip for this region.
        /// Changing this value recreates the subnet.
        /// </summary>
        public readonly string? Start;

        [OutputConstructor]
        private CloudProjectNetworkPrivateSubnetIpPool(
            bool? dhcp,

            string? end,

            string? network,

            string? region,

            string? start)
        {
            Dhcp = dhcp;
            End = end;
            Network = network;
            Region = region;
            Start = start;
        }
    }
}
