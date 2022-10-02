// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Inputs
{

    public sealed class IpServiceRoutedToGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Service where ip is routed to
        /// * `service_name`: service name
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public IpServiceRoutedToGetArgs()
        {
        }
        public static new IpServiceRoutedToGetArgs Empty => new IpServiceRoutedToGetArgs();
    }
}