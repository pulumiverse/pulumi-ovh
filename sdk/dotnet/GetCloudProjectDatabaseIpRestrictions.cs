// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Ovh
{
    public static class GetCloudProjectDatabaseIpRestrictions
    {
        /// <summary>
        /// Use the list of IP restrictions associated with a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// To get the list of IP restriction on a database cluster service:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var iprestrictions = Ovh.GetCloudProjectDatabaseIpRestrictions.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Engine = "YYYY",
        ///         ClusterId = "ZZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ips"] = iprestrictions.Apply(getCloudProjectDatabaseIpRestrictionsResult =&gt; getCloudProjectDatabaseIpRestrictionsResult.Ips),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectDatabaseIpRestrictionsResult> InvokeAsync(GetCloudProjectDatabaseIpRestrictionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectDatabaseIpRestrictionsResult>("ovh:index/getCloudProjectDatabaseIpRestrictions:getCloudProjectDatabaseIpRestrictions", args ?? new GetCloudProjectDatabaseIpRestrictionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use the list of IP restrictions associated with a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// To get the list of IP restriction on a database cluster service:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var iprestrictions = Ovh.GetCloudProjectDatabaseIpRestrictions.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Engine = "YYYY",
        ///         ClusterId = "ZZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ips"] = iprestrictions.Apply(getCloudProjectDatabaseIpRestrictionsResult =&gt; getCloudProjectDatabaseIpRestrictionsResult.Ips),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectDatabaseIpRestrictionsResult> Invoke(GetCloudProjectDatabaseIpRestrictionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectDatabaseIpRestrictionsResult>("ovh:index/getCloudProjectDatabaseIpRestrictions:getCloudProjectDatabaseIpRestrictions", args ?? new GetCloudProjectDatabaseIpRestrictionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectDatabaseIpRestrictionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want to list IP restrictions. To get a full list of available engine visit.
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public string Engine { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectDatabaseIpRestrictionsArgs()
        {
        }
        public static new GetCloudProjectDatabaseIpRestrictionsArgs Empty => new GetCloudProjectDatabaseIpRestrictionsArgs();
    }

    public sealed class GetCloudProjectDatabaseIpRestrictionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want to list IP restrictions. To get a full list of available engine visit.
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectDatabaseIpRestrictionsInvokeArgs()
        {
        }
        public static new GetCloudProjectDatabaseIpRestrictionsInvokeArgs Empty => new GetCloudProjectDatabaseIpRestrictionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProjectDatabaseIpRestrictionsResult
    {
        public readonly string ClusterId;
        public readonly string Engine;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of IP restriction of the database associated with the project.
        /// </summary>
        public readonly ImmutableArray<string> Ips;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetCloudProjectDatabaseIpRestrictionsResult(
            string clusterId,

            string engine,

            string id,

            ImmutableArray<string> ips,

            string serviceName)
        {
            ClusterId = clusterId;
            Engine = engine;
            Id = id;
            Ips = ips;
            ServiceName = serviceName;
        }
    }
}
