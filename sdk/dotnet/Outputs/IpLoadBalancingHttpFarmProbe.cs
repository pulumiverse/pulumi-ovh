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
    public sealed class IpLoadBalancingHttpFarmProbe
    {
        /// <summary>
        /// Force use of SSL (TLS)
        /// </summary>
        public readonly bool? ForceSsl;
        /// <summary>
        /// probe interval, Value between 30 and 3600 seconds, default 30
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// What to mach `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
        /// </summary>
        public readonly string? Match;
        /// <summary>
        /// HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
        /// </summary>
        public readonly string? Method;
        /// <summary>
        /// Negate probe result
        /// </summary>
        public readonly bool? Negate;
        /// <summary>
        /// Pattern to match against `match`
        /// </summary>
        public readonly string? Pattern;
        /// <summary>
        /// Port for backends to recieve traffic on.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// URL for HTTP probe type.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private IpLoadBalancingHttpFarmProbe(
            bool? forceSsl,

            int? interval,

            string? match,

            string? method,

            bool? negate,

            string? pattern,

            int? port,

            string type,

            string? url)
        {
            ForceSsl = forceSsl;
            Interval = interval;
            Match = match;
            Method = method;
            Negate = negate;
            Pattern = pattern;
            Port = port;
            Type = type;
            Url = url;
        }
    }
}
