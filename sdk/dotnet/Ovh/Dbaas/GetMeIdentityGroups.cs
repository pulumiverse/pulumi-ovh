// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Dbaas
{
    public static class GetMeIdentityGroups
    {
        /// <summary>
        /// Use this data source to retrieve the list of the account's identity groups
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
        ///     var groups = Ovh.Dbaas.GetMeIdentityGroups.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMeIdentityGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMeIdentityGroupsResult>("ovh:Dbaas/getMeIdentityGroups:getMeIdentityGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetMeIdentityGroupsResult
    {
        /// <summary>
        /// The list of the group names of all the identity groups.
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetMeIdentityGroupsResult(
            ImmutableArray<string> groups,

            string id)
        {
            Groups = groups;
            Id = id;
        }
    }
}
