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
    public sealed class GetIpServiceRoutedToResult
    {
        /// <summary>
        /// The service name
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private GetIpServiceRoutedToResult(string serviceName)
        {
            ServiceName = serviceName;
        }
    }
}
