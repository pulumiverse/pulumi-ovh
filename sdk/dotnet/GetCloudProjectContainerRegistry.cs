// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace lbrlabs.Ovh
{
    public static class GetCloudProjectContainerRegistry
    {
        /// <summary>
        /// Use this data source to get information about a container registry associated with a public cloud project.
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
        ///     var registry = Ovh.GetCloudProjectContainerRegistry.Invoke(new()
        ///     {
        ///         RegistryId = "yyyy",
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectContainerRegistryResult> InvokeAsync(GetCloudProjectContainerRegistryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectContainerRegistryResult>("ovh:index/getCloudProjectContainerRegistry:getCloudProjectContainerRegistry", args ?? new GetCloudProjectContainerRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a container registry associated with a public cloud project.
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
        ///     var registry = Ovh.GetCloudProjectContainerRegistry.Invoke(new()
        ///     {
        ///         RegistryId = "yyyy",
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectContainerRegistryResult> Invoke(GetCloudProjectContainerRegistryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectContainerRegistryResult>("ovh:index/getCloudProjectContainerRegistry:getCloudProjectContainerRegistry", args ?? new GetCloudProjectContainerRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectContainerRegistryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Registry ID
        /// </summary>
        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectContainerRegistryArgs()
        {
        }
        public static new GetCloudProjectContainerRegistryArgs Empty => new GetCloudProjectContainerRegistryArgs();
    }

    public sealed class GetCloudProjectContainerRegistryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Registry ID
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectContainerRegistryInvokeArgs()
        {
        }
        public static new GetCloudProjectContainerRegistryInvokeArgs Empty => new GetCloudProjectContainerRegistryInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProjectContainerRegistryResult
    {
        /// <summary>
        /// Registry creation date
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Registry name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Project ID of your registry
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Region of the registry
        /// </summary>
        public readonly string Region;
        public readonly string RegistryId;
        public readonly string ServiceName;
        /// <summary>
        /// Current size of the registry (bytes)
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Registry status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Registry last update date
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// Access url of the registry
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Version of your registry
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetCloudProjectContainerRegistryResult(
            string createdAt,

            string id,

            string name,

            string projectId,

            string region,

            string registryId,

            string serviceName,

            int size,

            string status,

            string updatedAt,

            string url,

            string version)
        {
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            ProjectId = projectId;
            Region = region;
            RegistryId = registryId;
            ServiceName = serviceName;
            Size = size;
            Status = status;
            UpdatedAt = updatedAt;
            Url = url;
            Version = version;
        }
    }
}
