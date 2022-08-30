// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetCloudProjectKubeIpNodePool
    {
        /// <summary>
        /// Use this data source to get a OVH Managed Kubernetes node pool.
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
        ///     var nodepool = Ovh.GetCloudProjectKubeIpNodePool.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
        ///         Name = "xxxxxx",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["maxNodes"] = nodepool.Apply(getCloudProjectKubeIpNodePoolResult =&gt; getCloudProjectKubeIpNodePoolResult.MaxNodes),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectKubeIpNodePoolResult> InvokeAsync(GetCloudProjectKubeIpNodePoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectKubeIpNodePoolResult>("ovh:index/getCloudProjectKubeIpNodePool:getCloudProjectKubeIpNodePool", args ?? new GetCloudProjectKubeIpNodePoolArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a OVH Managed Kubernetes node pool.
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
        ///     var nodepool = Ovh.GetCloudProjectKubeIpNodePool.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
        ///         Name = "xxxxxx",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["maxNodes"] = nodepool.Apply(getCloudProjectKubeIpNodePoolResult =&gt; getCloudProjectKubeIpNodePoolResult.MaxNodes),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectKubeIpNodePoolResult> Invoke(GetCloudProjectKubeIpNodePoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectKubeIpNodePoolResult>("ovh:index/getCloudProjectKubeIpNodePool:getCloudProjectKubeIpNodePool", args ?? new GetCloudProjectKubeIpNodePoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectKubeIpNodePoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public string KubeId { get; set; } = null!;

        /// <summary>
        /// maximum number of nodes allowed in the pool.
        /// Setting `desired_nodes` over this value will raise an error.
        /// </summary>
        [Input("maxNodes")]
        public int? MaxNodes { get; set; }

        /// <summary>
        /// minimum number of nodes allowed in the pool.
        /// Setting `desired_nodes` under this value will raise an error.
        /// </summary>
        [Input("minNodes")]
        public int? MinNodes { get; set; }

        /// <summary>
        /// The name of the node pool.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectKubeIpNodePoolArgs()
        {
        }
        public static new GetCloudProjectKubeIpNodePoolArgs Empty => new GetCloudProjectKubeIpNodePoolArgs();
    }

    public sealed class GetCloudProjectKubeIpNodePoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public Input<string> KubeId { get; set; } = null!;

        /// <summary>
        /// maximum number of nodes allowed in the pool.
        /// Setting `desired_nodes` over this value will raise an error.
        /// </summary>
        [Input("maxNodes")]
        public Input<int>? MaxNodes { get; set; }

        /// <summary>
        /// minimum number of nodes allowed in the pool.
        /// Setting `desired_nodes` under this value will raise an error.
        /// </summary>
        [Input("minNodes")]
        public Input<int>? MinNodes { get; set; }

        /// <summary>
        /// The name of the node pool.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectKubeIpNodePoolInvokeArgs()
        {
        }
        public static new GetCloudProjectKubeIpNodePoolInvokeArgs Empty => new GetCloudProjectKubeIpNodePoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProjectKubeIpNodePoolResult
    {
        /// <summary>
        /// (Optional) should the pool use the anti-affinity feature. Default to `false`.
        /// </summary>
        public readonly bool AntiAffinity;
        /// <summary>
        /// (Optional) Enable auto-scaling for the pool. Default to `false`.
        /// </summary>
        public readonly bool Autoscale;
        /// <summary>
        /// Number of nodes which are actually ready in the pool
        /// </summary>
        public readonly int AvailableNodes;
        /// <summary>
        /// Creation date
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Number of nodes present in the pool
        /// </summary>
        public readonly int CurrentNodes;
        /// <summary>
        /// Number of nodes you desire in the pool
        /// </summary>
        public readonly int DesiredNodes;
        /// <summary>
        /// Flavor name
        /// </summary>
        public readonly string Flavor;
        /// <summary>
        /// a valid OVH public cloud flavor ID in which the nodes will be started.
        /// Ex: "b2-7". Changing this value recreates the resource.
        /// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
        /// </summary>
        public readonly string FlavorName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string KubeId;
        /// <summary>
        /// maximum number of nodes allowed in the pool.
        /// Setting `desired_nodes` over this value will raise an error.
        /// </summary>
        public readonly int MaxNodes;
        /// <summary>
        /// minimum number of nodes allowed in the pool.
        /// Setting `desired_nodes` under this value will raise an error.
        /// </summary>
        public readonly int MinNodes;
        /// <summary>
        /// (Optional) should the nodes be billed on a monthly basis. Default to `false`.
        /// </summary>
        public readonly bool MonthlyBilled;
        /// <summary>
        /// (Optional) The name of the nodepool.
        /// Changing this value recreates the resource.
        /// Warning: "_" char is not allowed!
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Project id
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Status describing the state between number of nodes wanted and available ones
        /// </summary>
        public readonly string SizeStatus;
        /// <summary>
        /// Current status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Number of nodes with the latest version installed in the pool
        /// </summary>
        public readonly int UpToDateNodes;
        /// <summary>
        /// Last update date
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetCloudProjectKubeIpNodePoolResult(
            bool antiAffinity,

            bool autoscale,

            int availableNodes,

            string createdAt,

            int currentNodes,

            int desiredNodes,

            string flavor,

            string flavorName,

            string id,

            string kubeId,

            int maxNodes,

            int minNodes,

            bool monthlyBilled,

            string name,

            string projectId,

            string serviceName,

            string sizeStatus,

            string status,

            int upToDateNodes,

            string updatedAt)
        {
            AntiAffinity = antiAffinity;
            Autoscale = autoscale;
            AvailableNodes = availableNodes;
            CreatedAt = createdAt;
            CurrentNodes = currentNodes;
            DesiredNodes = desiredNodes;
            Flavor = flavor;
            FlavorName = flavorName;
            Id = id;
            KubeId = kubeId;
            MaxNodes = maxNodes;
            MinNodes = minNodes;
            MonthlyBilled = monthlyBilled;
            Name = name;
            ProjectId = projectId;
            ServiceName = serviceName;
            SizeStatus = sizeStatus;
            Status = status;
            UpToDateNodes = upToDateNodes;
            UpdatedAt = updatedAt;
        }
    }
}
