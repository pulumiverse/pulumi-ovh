// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates an OIDC configuration in an OVHcloud Managed Kubernetes cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const my_oidc = new ovh.CloudProjectKubeOidc("my-oidc", {
 *     serviceName: _var.projectid,
 *     kubeId: ovh_cloud_project_kube.mykube.id,
 *     clientId: "xxx",
 *     issuerUrl: "https://ovh.com",
 * });
 * ```
 */
export class CloudProjectKubeOidc extends pulumi.CustomResource {
    /**
     * Get an existing CloudProjectKubeOidc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudProjectKubeOidcState, opts?: pulumi.CustomResourceOptions): CloudProjectKubeOidc {
        return new CloudProjectKubeOidc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/cloudProjectKubeOidc:CloudProjectKubeOidc';

    /**
     * Returns true if the given object is an instance of CloudProjectKubeOidc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudProjectKubeOidc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudProjectKubeOidc.__pulumiType;
    }

    /**
     * The OIDC client ID.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The OIDC issuer url.
     */
    public readonly issuerUrl!: pulumi.Output<string>;
    /**
     * The ID of the managed kubernetes cluster.
     */
    public readonly kubeId!: pulumi.Output<string>;
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a CloudProjectKubeOidc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudProjectKubeOidcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudProjectKubeOidcArgs | CloudProjectKubeOidcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudProjectKubeOidcState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["issuerUrl"] = state ? state.issuerUrl : undefined;
            resourceInputs["kubeId"] = state ? state.kubeId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as CloudProjectKubeOidcArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.issuerUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuerUrl'");
            }
            if ((!args || args.kubeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubeId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["issuerUrl"] = args ? args.issuerUrl : undefined;
            resourceInputs["kubeId"] = args ? args.kubeId : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudProjectKubeOidc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProjectKubeOidc resources.
 */
export interface CloudProjectKubeOidcState {
    /**
     * The OIDC client ID.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The OIDC issuer url.
     */
    issuerUrl?: pulumi.Input<string>;
    /**
     * The ID of the managed kubernetes cluster.
     */
    kubeId?: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudProjectKubeOidc resource.
 */
export interface CloudProjectKubeOidcArgs {
    /**
     * The OIDC client ID.
     */
    clientId: pulumi.Input<string>;
    /**
     * The OIDC issuer url.
     */
    issuerUrl: pulumi.Input<string>;
    /**
     * The ID of the managed kubernetes cluster.
     */
    kubeId: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
