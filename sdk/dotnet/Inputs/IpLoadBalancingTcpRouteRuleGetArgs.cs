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

    public sealed class IpLoadBalancingTcpRouteRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the field to match like "protocol" or "host" "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules
        /// </summary>
        [Input("field")]
        public Input<string>? Field { get; set; }

        /// <summary>
        /// Matching operator. Not all operators are available for all fields. See "availableRules"
        /// * `negate`- Invert the matching operator effect
        /// </summary>
        [Input("match")]
        public Input<string>? Match { get; set; }

        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        /// <summary>
        /// Value to match against this match. Interpretation if this field depends on the match and field
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Id of your rule
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        /// </summary>
        [Input("subField")]
        public Input<string>? SubField { get; set; }

        public IpLoadBalancingTcpRouteRuleGetArgs()
        {
        }
        public static new IpLoadBalancingTcpRouteRuleGetArgs Empty => new IpLoadBalancingTcpRouteRuleGetArgs();
    }
}
