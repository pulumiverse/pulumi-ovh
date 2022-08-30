// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetMePaymentmeanBankAccount
    {
        /// <summary>
        /// Use this data source to retrieve information about a bank account
        /// payment mean associated with an OVH account.
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
        ///     var ba = Ovh.GetMePaymentmeanBankAccount.Invoke(new()
        ///     {
        ///         UseDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMePaymentmeanBankAccountResult> InvokeAsync(GetMePaymentmeanBankAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMePaymentmeanBankAccountResult>("ovh:index/getMePaymentmeanBankAccount:getMePaymentmeanBankAccount", args ?? new GetMePaymentmeanBankAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a bank account
        /// payment mean associated with an OVH account.
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
        ///     var ba = Ovh.GetMePaymentmeanBankAccount.Invoke(new()
        ///     {
        ///         UseDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMePaymentmeanBankAccountResult> Invoke(GetMePaymentmeanBankAccountInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMePaymentmeanBankAccountResult>("ovh:index/getMePaymentmeanBankAccount:getMePaymentmeanBankAccount", args ?? new GetMePaymentmeanBankAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMePaymentmeanBankAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// a regexp used to filter bank accounts 
        /// on their `description` attributes.
        /// </summary>
        [Input("descriptionRegexp")]
        public string? DescriptionRegexp { get; set; }

        /// <summary>
        /// Filter bank accounts on their `state` attribute.
        /// Can be "blockedForIncidents", "valid", "pendingValidation"
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        /// <summary>
        /// Retrieve bank account marked as default payment mean.
        /// </summary>
        [Input("useDefault")]
        public bool? UseDefault { get; set; }

        /// <summary>
        /// Retrieve oldest bank account.
        /// project.
        /// </summary>
        [Input("useOldest")]
        public bool? UseOldest { get; set; }

        public GetMePaymentmeanBankAccountArgs()
        {
        }
        public static new GetMePaymentmeanBankAccountArgs Empty => new GetMePaymentmeanBankAccountArgs();
    }

    public sealed class GetMePaymentmeanBankAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// a regexp used to filter bank accounts 
        /// on their `description` attributes.
        /// </summary>
        [Input("descriptionRegexp")]
        public Input<string>? DescriptionRegexp { get; set; }

        /// <summary>
        /// Filter bank accounts on their `state` attribute.
        /// Can be "blockedForIncidents", "valid", "pendingValidation"
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Retrieve bank account marked as default payment mean.
        /// </summary>
        [Input("useDefault")]
        public Input<bool>? UseDefault { get; set; }

        /// <summary>
        /// Retrieve oldest bank account.
        /// project.
        /// </summary>
        [Input("useOldest")]
        public Input<bool>? UseOldest { get; set; }

        public GetMePaymentmeanBankAccountInvokeArgs()
        {
        }
        public static new GetMePaymentmeanBankAccountInvokeArgs Empty => new GetMePaymentmeanBankAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetMePaymentmeanBankAccountResult
    {
        /// <summary>
        /// a boolean which tells if the retrieved bank account
        /// is marked as the default payment mean
        /// </summary>
        public readonly bool Default;
        /// <summary>
        /// the description attribute of the bank account
        /// </summary>
        public readonly string Description;
        public readonly string? DescriptionRegexp;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string State;
        public readonly bool? UseDefault;
        public readonly bool? UseOldest;

        [OutputConstructor]
        private GetMePaymentmeanBankAccountResult(
            bool @default,

            string description,

            string? descriptionRegexp,

            string id,

            string state,

            bool? useDefault,

            bool? useOldest)
        {
            Default = @default;
            Description = description;
            DescriptionRegexp = descriptionRegexp;
            Id = id;
            State = state;
            UseDefault = useDefault;
            UseOldest = useOldest;
        }
    }
}
