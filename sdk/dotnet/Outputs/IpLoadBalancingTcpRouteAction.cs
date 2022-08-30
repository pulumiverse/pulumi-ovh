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
    public sealed class IpLoadBalancingTcpRouteAction
    {
        /// <summary>
        /// Farm ID for "farm" action type, empty for others.
        /// </summary>
        public readonly string? Target;
        /// <summary>
        /// Action to trigger if all the rules of this route matches
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private IpLoadBalancingTcpRouteAction(
            string? target,

            string type)
        {
            Target = target;
            Type = type;
        }
    }
}
