// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Me
{
    public static class GetIdentityUsers
    {
        /// <summary>
        /// Use this data source to retrieve list of user logins of the account's identity users.
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
        ///     var users = Ovh.Me.GetIdentityUsers.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIdentityUsersResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentityUsersResult>("ovh:Me/getIdentityUsers:getIdentityUsers", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetIdentityUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of the user's logins of all the identity users.
        /// </summary>
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private GetIdentityUsersResult(
            string id,

            ImmutableArray<string> users)
        {
            Id = id;
            Users = users;
        }
    }
}
