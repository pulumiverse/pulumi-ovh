// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh
{
    public static class GetCloudProjectFailoverIpAttach
    {
        /// <summary>
        /// Use this data source to get the details of a failover ip address of a service in a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myfailoverip = Ovh.GetCloudProjectFailoverIpAttach.Invoke(new()
        ///     {
        ///         Ip = "XXXXXX",
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectFailoverIpAttachResult> InvokeAsync(GetCloudProjectFailoverIpAttachArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectFailoverIpAttachResult>("ovh:index/getCloudProjectFailoverIpAttach:getCloudProjectFailoverIpAttach", args ?? new GetCloudProjectFailoverIpAttachArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the details of a failover ip address of a service in a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myfailoverip = Ovh.GetCloudProjectFailoverIpAttach.Invoke(new()
        ///     {
        ///         Ip = "XXXXXX",
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectFailoverIpAttachResult> Invoke(GetCloudProjectFailoverIpAttachInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectFailoverIpAttachResult>("ovh:index/getCloudProjectFailoverIpAttach:getCloudProjectFailoverIpAttach", args ?? new GetCloudProjectFailoverIpAttachInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectFailoverIpAttachArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP block
        /// </summary>
        [Input("block")]
        public string? Block { get; set; }

        [Input("continentCode")]
        public string? ContinentCode { get; set; }

        [Input("geoLoc")]
        public string? GeoLoc { get; set; }

        /// <summary>
        /// The failover ip address to query
        /// </summary>
        [Input("ip")]
        public string? Ip { get; set; }

        [Input("routedTo")]
        public string? RoutedTo { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectFailoverIpAttachArgs()
        {
        }
        public static new GetCloudProjectFailoverIpAttachArgs Empty => new GetCloudProjectFailoverIpAttachArgs();
    }

    public sealed class GetCloudProjectFailoverIpAttachInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP block
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        [Input("continentCode")]
        public Input<string>? ContinentCode { get; set; }

        [Input("geoLoc")]
        public Input<string>? GeoLoc { get; set; }

        /// <summary>
        /// The failover ip address to query
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("routedTo")]
        public Input<string>? RoutedTo { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectFailoverIpAttachInvokeArgs()
        {
        }
        public static new GetCloudProjectFailoverIpAttachInvokeArgs Empty => new GetCloudProjectFailoverIpAttachInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProjectFailoverIpAttachResult
    {
        /// <summary>
        /// The IP block
        /// </summary>
        public readonly string Block;
        public readonly string ContinentCode;
        public readonly string GeoLoc;
        /// <summary>
        /// The Ip id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Ip Address
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Current operation progress in percent
        /// </summary>
        public readonly int Progress;
        public readonly string RoutedTo;
        public readonly string ServiceName;
        /// <summary>
        /// Ip status, can be `ok` or `operationPending`
        /// </summary>
        public readonly string Status;
        public readonly string SubType;

        [OutputConstructor]
        private GetCloudProjectFailoverIpAttachResult(
            string block,

            string continentCode,

            string geoLoc,

            string id,

            string ip,

            int progress,

            string routedTo,

            string serviceName,

            string status,

            string subType)
        {
            Block = block;
            ContinentCode = continentCode;
            GeoLoc = geoLoc;
            Id = id;
            Ip = ip;
            Progress = progress;
            RoutedTo = routedTo;
            ServiceName = serviceName;
            Status = status;
            SubType = subType;
        }
    }
}