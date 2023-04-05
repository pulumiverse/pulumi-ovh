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
    public static class GetServer
    {
        /// <summary>
        /// Use this data source to retrieve information about a dedicated server associated with your OVHcloud Account.
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
        ///     var server = Ovh.GetServer.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("ovh:index/getServer:getServer", args ?? new GetServerArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a dedicated server associated with your OVHcloud Account.
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
        ///     var server = Ovh.GetServer.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServerResult> Invoke(GetServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerResult>("ovh:index/getServer:getServer", args ?? new GetServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetServerArgs()
        {
        }
        public static new GetServerArgs Empty => new GetServerArgs();
    }

    public sealed class GetServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetServerInvokeArgs()
        {
        }
        public static new GetServerInvokeArgs Empty => new GetServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// boot id of the server
        /// </summary>
        public readonly int BootId;
        /// <summary>
        /// dedicater server commercial range
        /// </summary>
        public readonly string CommercialRange;
        /// <summary>
        /// dedicated datacenter localisation (bhs1,bhs2,...)
        /// </summary>
        public readonly string Datacenter;
        /// <summary>
        /// List of enabled public VNI uuids
        /// </summary>
        public readonly ImmutableArray<string> EnabledPublicVnis;
        /// <summary>
        /// List of enabled vrack_aggregation VNI uuids
        /// </summary>
        public readonly ImmutableArray<string> EnabledVrackAggregationVnis;
        /// <summary>
        /// List of enabled vrack VNI uuids
        /// </summary>
        public readonly ImmutableArray<string> EnabledVrackVnis;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// dedicated server ip (IPv4)
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// dedicated server ip blocks
        /// </summary>
        public readonly ImmutableArray<string> Ips;
        /// <summary>
        /// link speed of the server
        /// </summary>
        public readonly int LinkSpeed;
        /// <summary>
        /// Icmp monitoring state
        /// </summary>
        public readonly bool Monitoring;
        /// <summary>
        /// User defined VirtualNetworkInterface name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Operating system
        /// </summary>
        public readonly string Os;
        /// <summary>
        /// Does this server have professional use option
        /// </summary>
        public readonly bool ProfessionalUse;
        /// <summary>
        /// rack id of the server
        /// </summary>
        public readonly string Rack;
        /// <summary>
        /// rescue mail of the server
        /// </summary>
        public readonly string RescueMail;
        /// <summary>
        /// dedicated server reverse
        /// </summary>
        public readonly string Reverse;
        /// <summary>
        /// root device of the server
        /// </summary>
        public readonly string RootDevice;
        /// <summary>
        /// your server id
        /// </summary>
        public readonly int ServerId;
        public readonly string ServiceName;
        /// <summary>
        /// error, hacked, hackedBlocked, ok
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Dedicated server support level (critical, fastpath, gs, pro)
        /// </summary>
        public readonly string SupportLevel;
        /// <summary>
        /// the list of Virtualnetworkinterface assiociated with this server
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerVniResult> Vnis;

        [OutputConstructor]
        private GetServerResult(
            int bootId,

            string commercialRange,

            string datacenter,

            ImmutableArray<string> enabledPublicVnis,

            ImmutableArray<string> enabledVrackAggregationVnis,

            ImmutableArray<string> enabledVrackVnis,

            string id,

            string ip,

            ImmutableArray<string> ips,

            int linkSpeed,

            bool monitoring,

            string name,

            string os,

            bool professionalUse,

            string rack,

            string rescueMail,

            string reverse,

            string rootDevice,

            int serverId,

            string serviceName,

            string state,

            string supportLevel,

            ImmutableArray<Outputs.GetServerVniResult> vnis)
        {
            BootId = bootId;
            CommercialRange = commercialRange;
            Datacenter = datacenter;
            EnabledPublicVnis = enabledPublicVnis;
            EnabledVrackAggregationVnis = enabledVrackAggregationVnis;
            EnabledVrackVnis = enabledVrackVnis;
            Id = id;
            Ip = ip;
            Ips = ips;
            LinkSpeed = linkSpeed;
            Monitoring = monitoring;
            Name = name;
            Os = os;
            ProfessionalUse = professionalUse;
            Rack = rack;
            RescueMail = rescueMail;
            Reverse = reverse;
            RootDevice = rootDevice;
            ServerId = serverId;
            ServiceName = serviceName;
            State = state;
            SupportLevel = supportLevel;
            Vnis = vnis;
        }
    }
}
