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

    public sealed class IpServiceOrderDetailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom description on your ip.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// expiration date
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// order detail id
        /// </summary>
        [Input("orderDetailId")]
        public Input<int>? OrderDetailId { get; set; }

        /// <summary>
        /// quantity
        /// </summary>
        [Input("quantity")]
        public Input<string>? Quantity { get; set; }

        public IpServiceOrderDetailArgs()
        {
        }
        public static new IpServiceOrderDetailArgs Empty => new IpServiceOrderDetailArgs();
    }
}
