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
    public static class GetDedicatedInstallationTemplates
    {
        /// <summary>
        /// Use this data source to get the list of installation templates available for dedicated servers.
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
        ///     var templates = Ovh.GetDedicatedInstallationTemplates.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDedicatedInstallationTemplatesResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDedicatedInstallationTemplatesResult>("ovh:index/getDedicatedInstallationTemplates:getDedicatedInstallationTemplates", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetDedicatedInstallationTemplatesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of installation templates IDs available for dedicated servers.
        /// </summary>
        public readonly ImmutableArray<string> Results;

        [OutputConstructor]
        private GetDedicatedInstallationTemplatesResult(
            string id,

            ImmutableArray<string> results)
        {
            Id = id;
            Results = results;
        }
    }
}
