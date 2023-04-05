// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Dedicated
{
    public static class GetCeph
    {
        /// <summary>
        /// Use this data source to retrieve information about a dedicated CEPH. 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_ceph = Ovh.Dedicated.GetCeph.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCephResult> InvokeAsync(GetCephArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCephResult>("ovh:Dedicated/getCeph:getCeph", args ?? new GetCephArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a dedicated CEPH. 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_ceph = Ovh.Dedicated.GetCeph.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCephResult> Invoke(GetCephInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCephResult>("ovh:Dedicated/getCeph:getCeph", args ?? new GetCephInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCephArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// CEPH cluster version
        /// </summary>
        [Input("cephVersion")]
        public string? CephVersion { get; set; }

        /// <summary>
        /// The service name of the dedicated CEPH cluster.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// the status of the service
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetCephArgs()
        {
        }
        public static new GetCephArgs Empty => new GetCephArgs();
    }

    public sealed class GetCephInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// CEPH cluster version
        /// </summary>
        [Input("cephVersion")]
        public Input<string>? CephVersion { get; set; }

        /// <summary>
        /// The service name of the dedicated CEPH cluster.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// the status of the service
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetCephInvokeArgs()
        {
        }
        public static new GetCephInvokeArgs Empty => new GetCephInvokeArgs();
    }


    [OutputType]
    public sealed class GetCephResult
    {
        /// <summary>
        /// list of CEPH monitors IPs
        /// </summary>
        public readonly ImmutableArray<string> CephMons;
        /// <summary>
        /// CEPH cluster version
        /// </summary>
        public readonly string CephVersion;
        /// <summary>
        /// CRUSH algorithm settings. Possible values
        /// * OPTIMAL
        /// * DEFAULT
        /// * LEGACY
        /// * BOBTAIL
        /// * ARGONAUT
        /// * FIREFLY
        /// * HAMMER
        /// * JEWEL
        /// </summary>
        public readonly string CrushTunables;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// CEPH cluster label
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// cluster region
        /// </summary>
        public readonly string Region;
        public readonly string ServiceName;
        /// <summary>
        /// Cluster size in TB
        /// </summary>
        public readonly double Size;
        /// <summary>
        /// the state of the cluster
        /// </summary>
        public readonly string State;
        /// <summary>
        /// the status of the service
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetCephResult(
            ImmutableArray<string> cephMons,

            string cephVersion,

            string crushTunables,

            string id,

            string label,

            string region,

            string serviceName,

            double size,

            string state,

            string status)
        {
            CephMons = cephMons;
            CephVersion = cephVersion;
            CrushTunables = crushTunables;
            Id = id;
            Label = label;
            Region = region;
            ServiceName = serviceName;
            Size = size;
            State = state;
            Status = status;
        }
    }
}
