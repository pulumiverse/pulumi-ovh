// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject
{
    /// <summary>
    /// Creates an OIDC configuration in an OVHcloud Managed Kubernetes cluster.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Ovh = Lbrlabs.PulumiPackage.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_oidc = new Ovh.CloudProject.KubeOidc("my-oidc", new()
    ///     {
    ///         ServiceName = @var.Projectid,
    ///         KubeId = ovh_cloud_project_kube.Mykube.Id,
    ///         ClientId = "xxx",
    ///         IssuerUrl = "https://ovh.com",
    ///         OidcUsernameClaim = "an-email",
    ///         OidcUsernamePrefix = "ovh:",
    ///         OidcGroupsClaims = new[]
    ///         {
    ///             "groups",
    ///         },
    ///         OidcGroupsPrefix = "ovh:",
    ///         OidcRequiredClaims = new[]
    ///         {
    ///             "claim1=val1",
    ///         },
    ///         OidcSigningAlgs = new[]
    ///         {
    ///             "RS512",
    ///         },
    ///         OidcCaContent = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZhekNDQTFPZ0F3SUJBZ0lVYm9YRkZrL1hCQmdQUUI4UHlqbkttUGVWekNjd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1JURUxNQWtHQTFVRUJoTUNRVlV4RXpBUkJnTlZCQWdNQ2xOdmJXVXRVM1JoZEdVeElUQWZCZ05WQkFvTQpHRWx1ZEdWeWJtVjBJRmRwWkdkcGRITWdVSFI1SUV4MFpEQWVGdzB5TWpFd01UUXdOalE0TlROYUZ3MHlNekV3Ck1UUXdOalE0TlROYU1FVXhDekFKQmdOVkJBWVRBa0ZWTVJNd0VRWURWUVFJREFwVGIyMWxMVk4wWVhSbE1TRXcKSHdZRFZRUUtEQmhKYm5SbGNtNWxkQ0JYYVdSbmFYUnpJRkIwZVNCTWRHUXdnZ0lpTUEwR0NTcUdTSWIzRFFFQgpBUVVBQTRJQ0R3QXdnZ0lLQW9JQ0FRQytPMk53bGx2QTQyT05SUHMyZWlqTUp2UHhpN21RblVSS3FrOHJEV1VkCkwzZU0yM1JXeVhtS1AydDQ5Zi9LVGsweEZNVStOSTUzTEhwWmh6N3NpK3dEUFUvWWZWSS9rQmZsRm8zeVZCMSsKZWdCSnpyNGIrQ3FoaWlCUkh0Vm5LblFKUmdvOVJjVkxhRm82UEY0N1V0UWJ2bWVuNGdERnExVkYwVHhUdnFMdwpIMzRZL0U2QUJsSlZnWFBzaWQzNm54eTErNnlKV05vRXNVekFiekpWMHhzTGhxc2hOazA0TWx4YnBhcG1XcEUxCmFFMHRIZGpjUlI3Y1dTRUUwMnRSQzNYL2tSNjBKb3MxR0N0Y0ZQTTVIN3NjOFBXNFRUem1EWWhOeDRiVjV4T28KU0xYRnI5ajBzZEgxbm1wSlI1dWxJT2dPTWV3MHA2d3JOYVV2MGpxc1hzdVdqMVpxdTRLRi81aEQ3azVhRlhKNQpjYWNTUi9mRWxreW1uZis0eHZFOG8wdkRWNFR5NHo3K3lSS1U0clZvZFNBZWZIN3lqeitLV1RRck96L0lHU2NwCmV1YTdqV0hRMDdMYWxyTjV2b0tFaU1JM3MrWjhzeUdVUGVyYXQwdzJMWlc3NnhxVGl4R002clZxUldxVlQ4L1oKQTJMMEc4WGRvNTZvV2lFYVF5RkJtRDFnMXU2UEsvTmFGVDI1L2tTNWJ1dnF5L1dLVGt0UVNhNHNZc1ZLbUlQTQp0Zys0NUZ2aFErNkRuQzd0TmVnaTZDTkdTb0w0R1dPOEE5UDZRNjE5RkJJZ1VjcGpFMTgvUHpQOEJmcTAxajhnCjZmdm1jNkVPMkxHVHhDcW1DbVp0TnI3OCtQaUxkMHZIY3pqY3E3NzhiNW5WRXRpUVNRQkUyb0ozTVlIZUFIUUkKYVFJREFRQUJvMU13VVRBZEJnTlZIUTRFRmdRVUpaMUhlVmx1U3pjY0U2NEZQYWtuNkRBWnhmSXdId1lEVlIwagpCQmd3Rm9BVUpaMUhlVmx1U3pjY0U2NEZQYWtuNkRBWnhmSXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QU5CZ2txCmhraUc5dzBCQVFzRkFBT0NBZ0VBQlhNSlU2MjJZVFZVNnZ1K2svNnkwMGNaWlRmVnZtdVJMOXhTcWxVM0I1QmQKVWdyVWx1TmdjN2dhUUlrYzkvWmh2MnhNd0xxUldMWEhiTWx1NkNvdkNiVTVpeWt0NHVWMnl5UzlZYWhmVVRNVQo3TVE0WFRta2hoS0dGbWZBQ2QzTUVwRE55T3hmWXh0UVBwM1NZT2IxRGFKMmUwY01Gc081bytORGQ5aFVBVzFoCjFLMjMwQnZzYldYYVo4MStIdTU4U1BsYTM5R3FMTG85MzR6dEs4WkRWNFRGTVJxMnNVQ1cxcWFidDh5ejd2RzAKSGV3dXdxelRwR1lTSFI1U0ZvMm45R0xKVUN4SnhxcDlOWVJjMlhUdXRUdkJESzVPMXFZZEJaQzd6cmcxSnczawp2SjI4UGx2TzBQRE42ZVlUdElJdC9yU05ZbW56eVVNRTRYREt0di9KRitLZWZNSWxDTkpzZDRHYXVTdlo5M1NOClhINmcrNEZvRkp4UzNxRmZ0WEc4czNRNnppNzNLRzh5UHZVNHU0WmZNRGd2aG92L0V5YkNLWUpFdVVZSlJWNGEKbmc3cWh3NDBabXQ0eWNCRzU5a2tFSGhNYWtxTWpPaUNkV2x4MEVjZXIxcEFGT1pqN3o1NktURXIxa0ZwUHVaRApjVER5SnNwTjh6dm9CQ0l1ancvQjR6S3kyWStOQitRR1p3dXhyTk9mRGR6ek9yQUE1Ym9OS2gwUUh4c0RxNTExClFaU3hCR21EcGJzN2QzMUQvQll3WEhIUWdwb3FoVUU5dFBGSThpN0pkM2FyeXZCdHlnTWlxSmt1VlRFVk1Ta0UKNTZ0VnFsMjlXenFhRXNrbDN3VUlmczVKKzN3RzRPcWNxRDdXaGQxWUtnc0VUMjdFTWlqVXZIYzQ4TXE0bU1rPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud Managed Kubernetes Service cluster OIDC can be imported using the tenant `service_name` and cluster id `kube_id` separated by "/" E.g., bash
    /// 
    /// ```sh
    ///  $ pulumi import ovh:CloudProject/kubeOidc:KubeOidc my-oidc service_name/kube_id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProject/kubeOidc:KubeOidc")]
    public partial class KubeOidc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The OIDC client ID.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The OIDC issuer url.
        /// </summary>
        [Output("issuerUrl")]
        public Output<string> IssuerUrl { get; private set; } = null!;

        /// <summary>
        /// The ID of the managed kubernetes cluster.
        /// </summary>
        [Output("kubeId")]
        public Output<string> KubeId { get; private set; } = null!;

        [Output("oidcCaContent")]
        public Output<string?> OidcCaContent { get; private set; } = null!;

        [Output("oidcGroupsClaims")]
        public Output<ImmutableArray<string>> OidcGroupsClaims { get; private set; } = null!;

        [Output("oidcGroupsPrefix")]
        public Output<string?> OidcGroupsPrefix { get; private set; } = null!;

        [Output("oidcRequiredClaims")]
        public Output<ImmutableArray<string>> OidcRequiredClaims { get; private set; } = null!;

        [Output("oidcSigningAlgs")]
        public Output<ImmutableArray<string>> OidcSigningAlgs { get; private set; } = null!;

        [Output("oidcUsernameClaim")]
        public Output<string?> OidcUsernameClaim { get; private set; } = null!;

        [Output("oidcUsernamePrefix")]
        public Output<string?> OidcUsernamePrefix { get; private set; } = null!;

        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a KubeOidc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubeOidc(string name, KubeOidcArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/kubeOidc:KubeOidc", name, args ?? new KubeOidcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubeOidc(string name, Input<string> id, KubeOidcState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/kubeOidc:KubeOidc", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KubeOidc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubeOidc Get(string name, Input<string> id, KubeOidcState? state = null, CustomResourceOptions? options = null)
        {
            return new KubeOidc(name, id, state, options);
        }
    }

    public sealed class KubeOidcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OIDC client ID.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The OIDC issuer url.
        /// </summary>
        [Input("issuerUrl", required: true)]
        public Input<string> IssuerUrl { get; set; } = null!;

        /// <summary>
        /// The ID of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public Input<string> KubeId { get; set; } = null!;

        [Input("oidcCaContent")]
        public Input<string>? OidcCaContent { get; set; }

        [Input("oidcGroupsClaims")]
        private InputList<string>? _oidcGroupsClaims;
        public InputList<string> OidcGroupsClaims
        {
            get => _oidcGroupsClaims ?? (_oidcGroupsClaims = new InputList<string>());
            set => _oidcGroupsClaims = value;
        }

        [Input("oidcGroupsPrefix")]
        public Input<string>? OidcGroupsPrefix { get; set; }

        [Input("oidcRequiredClaims")]
        private InputList<string>? _oidcRequiredClaims;
        public InputList<string> OidcRequiredClaims
        {
            get => _oidcRequiredClaims ?? (_oidcRequiredClaims = new InputList<string>());
            set => _oidcRequiredClaims = value;
        }

        [Input("oidcSigningAlgs")]
        private InputList<string>? _oidcSigningAlgs;
        public InputList<string> OidcSigningAlgs
        {
            get => _oidcSigningAlgs ?? (_oidcSigningAlgs = new InputList<string>());
            set => _oidcSigningAlgs = value;
        }

        [Input("oidcUsernameClaim")]
        public Input<string>? OidcUsernameClaim { get; set; }

        [Input("oidcUsernamePrefix")]
        public Input<string>? OidcUsernamePrefix { get; set; }

        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public KubeOidcArgs()
        {
        }
        public static new KubeOidcArgs Empty => new KubeOidcArgs();
    }

    public sealed class KubeOidcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OIDC client ID.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The OIDC issuer url.
        /// </summary>
        [Input("issuerUrl")]
        public Input<string>? IssuerUrl { get; set; }

        /// <summary>
        /// The ID of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId")]
        public Input<string>? KubeId { get; set; }

        [Input("oidcCaContent")]
        public Input<string>? OidcCaContent { get; set; }

        [Input("oidcGroupsClaims")]
        private InputList<string>? _oidcGroupsClaims;
        public InputList<string> OidcGroupsClaims
        {
            get => _oidcGroupsClaims ?? (_oidcGroupsClaims = new InputList<string>());
            set => _oidcGroupsClaims = value;
        }

        [Input("oidcGroupsPrefix")]
        public Input<string>? OidcGroupsPrefix { get; set; }

        [Input("oidcRequiredClaims")]
        private InputList<string>? _oidcRequiredClaims;
        public InputList<string> OidcRequiredClaims
        {
            get => _oidcRequiredClaims ?? (_oidcRequiredClaims = new InputList<string>());
            set => _oidcRequiredClaims = value;
        }

        [Input("oidcSigningAlgs")]
        private InputList<string>? _oidcSigningAlgs;
        public InputList<string> OidcSigningAlgs
        {
            get => _oidcSigningAlgs ?? (_oidcSigningAlgs = new InputList<string>());
            set => _oidcSigningAlgs = value;
        }

        [Input("oidcUsernameClaim")]
        public Input<string>? OidcUsernameClaim { get; set; }

        [Input("oidcUsernamePrefix")]
        public Input<string>? OidcUsernamePrefix { get; set; }

        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public KubeOidcState()
        {
        }
        public static new KubeOidcState Empty => new KubeOidcState();
    }
}
