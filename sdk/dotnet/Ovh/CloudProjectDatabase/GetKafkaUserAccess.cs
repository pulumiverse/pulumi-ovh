// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProjectDatabase
{
    public static class GetKafkaUserAccess
    {
        /// <summary>
        /// Use this data source to get information about user acces of a kafka cluster associated with a public cloud project.
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
        ///     var access = Ovh.CloudProjectDatabase.GetKafkaUserAccess.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         UserId = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["accessCert"] = access.Apply(getKafkaUserAccessResult =&gt; getKafkaUserAccessResult.Cert),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKafkaUserAccessResult> InvokeAsync(GetKafkaUserAccessArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKafkaUserAccessResult>("ovh:CloudProjectDatabase/getKafkaUserAccess:getKafkaUserAccess", args ?? new GetKafkaUserAccessArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about user acces of a kafka cluster associated with a public cloud project.
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
        ///     var access = Ovh.CloudProjectDatabase.GetKafkaUserAccess.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         UserId = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["accessCert"] = access.Apply(getKafkaUserAccessResult =&gt; getKafkaUserAccessResult.Cert),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetKafkaUserAccessResult> Invoke(GetKafkaUserAccessInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKafkaUserAccessResult>("ovh:CloudProjectDatabase/getKafkaUserAccess:getKafkaUserAccess", args ?? new GetKafkaUserAccessInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKafkaUserAccessArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// User ID
        /// </summary>
        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetKafkaUserAccessArgs()
        {
        }
        public static new GetKafkaUserAccessArgs Empty => new GetKafkaUserAccessArgs();
    }

    public sealed class GetKafkaUserAccessInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// User ID
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GetKafkaUserAccessInvokeArgs()
        {
        }
        public static new GetKafkaUserAccessInvokeArgs Empty => new GetKafkaUserAccessInvokeArgs();
    }


    [OutputType]
    public sealed class GetKafkaUserAccessResult
    {
        /// <summary>
        /// User cert.
        /// </summary>
        public readonly string Cert;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Sensitive) User key for the cert.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private GetKafkaUserAccessResult(
            string cert,

            string clusterId,

            string id,

            string key,

            string serviceName,

            string userId)
        {
            Cert = cert;
            ClusterId = clusterId;
            Id = id;
            Key = key;
            ServiceName = serviceName;
            UserId = userId;
        }
    }
}
