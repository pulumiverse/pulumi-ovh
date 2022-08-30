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
    public static class GetCloudProjectDatabase
    {
        /// <summary>
        /// Use this data source to get the managed database of a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// To get information of a database cluster service:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var db = Ovh.GetCloudProjectDatabase.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Engine = "YYYY",
        ///         ClusterId = "ZZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["clusterId"] = db.Apply(getCloudProjectDatabaseResult =&gt; getCloudProjectDatabaseResult.ClusterId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectDatabaseResult> InvokeAsync(GetCloudProjectDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectDatabaseResult>("ovh:index/getCloudProjectDatabase:getCloudProjectDatabase", args ?? new GetCloudProjectDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the managed database of a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// To get information of a database cluster service:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var db = Ovh.GetCloudProjectDatabase.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Engine = "YYYY",
        ///         ClusterId = "ZZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["clusterId"] = db.Apply(getCloudProjectDatabaseResult =&gt; getCloudProjectDatabaseResult.ClusterId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectDatabaseResult> Invoke(GetCloudProjectDatabaseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectDatabaseResult>("ovh:index/getCloudProjectDatabase:getCloudProjectDatabase", args ?? new GetCloudProjectDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectDatabaseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The database engine you want to get information. To get a full list of available engine visit.
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

        public GetCloudProjectDatabaseArgs()
        {
        }
        public static new GetCloudProjectDatabaseArgs Empty => new GetCloudProjectDatabaseArgs();
    }

    public sealed class GetCloudProjectDatabaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The database engine you want to get information. To get a full list of available engine visit.
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

        public GetCloudProjectDatabaseInvokeArgs()
        {
        }
        public static new GetCloudProjectDatabaseInvokeArgs Empty => new GetCloudProjectDatabaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProjectDatabaseResult
    {
        /// <summary>
        /// Time on which backups start every day.
        /// </summary>
        public readonly string BackupTime;
        public readonly string ClusterId;
        /// <summary>
        /// Date of the creation of the cluster.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// List of all endpoints objects of the service.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudProjectDatabaseEndpointResult> Endpoints;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Flavor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Time on which maintenances can start every day.
        /// </summary>
        public readonly string MaintenanceTime;
        /// <summary>
        /// Type of network of the cluster.
        /// </summary>
        public readonly string NetworkType;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudProjectDatabaseNodeResult> Nodes;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Current status of the cluster.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetCloudProjectDatabaseResult(
            string backupTime,

            string clusterId,

            string createdAt,

            string description,

            ImmutableArray<Outputs.GetCloudProjectDatabaseEndpointResult> endpoints,

            string engine,

            string flavor,

            string id,

            string maintenanceTime,

            string networkType,

            ImmutableArray<Outputs.GetCloudProjectDatabaseNodeResult> nodes,

            string plan,

            string serviceName,

            string status,

            string version)
        {
            BackupTime = backupTime;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Description = description;
            Endpoints = endpoints;
            Engine = engine;
            Flavor = flavor;
            Id = id;
            MaintenanceTime = maintenanceTime;
            NetworkType = networkType;
            Nodes = nodes;
            Plan = plan;
            ServiceName = serviceName;
            Status = status;
            Version = version;
        }
    }
}
