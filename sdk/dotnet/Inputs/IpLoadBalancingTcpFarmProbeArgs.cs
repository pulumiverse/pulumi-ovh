// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace lbrlabs.Ovh.Inputs
{

    public sealed class IpLoadBalancingTcpFarmProbeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Force use of SSL (TLS)
        /// </summary>
        [Input("forceSsl")]
        public Input<bool>? ForceSsl { get; set; }

        /// <summary>
        /// probe interval, Value between 30 and 3600 seconds, default 30
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// What to mach `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
        /// </summary>
        [Input("match")]
        public Input<string>? Match { get; set; }

        /// <summary>
        /// HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Negate probe result
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        /// <summary>
        /// Pattern to match against `match`
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Port for backends to recieve traffic on.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// URL for HTTP probe type.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public IpLoadBalancingTcpFarmProbeArgs()
        {
        }
        public static new IpLoadBalancingTcpFarmProbeArgs Empty => new IpLoadBalancingTcpFarmProbeArgs();
    }
}
