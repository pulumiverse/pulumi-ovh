// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class GetKubeCustomizationKubeProxyResult
    {
        /// <summary>
        /// Kubernetes cluster kube-proxy customization of iptables specific config.
        /// </summary>
        public readonly Outputs.GetKubeCustomizationKubeProxyIptablesResult? Iptables;
        /// <summary>
        /// Kubernetes cluster kube-proxy customization of IPVS specific config (durations format is [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        public readonly Outputs.GetKubeCustomizationKubeProxyIpvsResult? Ipvs;

        [OutputConstructor]
        private GetKubeCustomizationKubeProxyResult(
            Outputs.GetKubeCustomizationKubeProxyIptablesResult? iptables,

            Outputs.GetKubeCustomizationKubeProxyIpvsResult? ipvs)
        {
            Iptables = iptables;
            Ipvs = ipvs;
        }
    }
}
