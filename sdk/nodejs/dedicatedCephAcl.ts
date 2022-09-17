// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Add a new access ACL for the given network/mask.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const my-ceph = ovh.getDedicatedCeph({
 *     serviceName: "94d423da-0e55-45f2-9812-836460a19939",
 * });
 * const my_acl = new ovh.DedicatedCephAcl("my-acl", {
 *     serviceName: my_ceph.then(my_ceph => my_ceph.id),
 *     network: "1.2.3.4",
 *     netmask: "255.255.255.255",
 * });
 * ```
 */
export class DedicatedCephAcl extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedCephAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DedicatedCephAclState, opts?: pulumi.CustomResourceOptions): DedicatedCephAcl {
        return new DedicatedCephAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/dedicatedCephAcl:DedicatedCephAcl';

    /**
     * Returns true if the given object is an instance of DedicatedCephAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedCephAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedCephAcl.__pulumiType;
    }

    /**
     * IP family. `IPv4` or `IPv6`
     */
    public /*out*/ readonly family!: pulumi.Output<string>;
    /**
     * The network mask to apply
     */
    public readonly netmask!: pulumi.Output<string>;
    /**
     * The network IP to authorize
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The internal name of your dedicated CEPH
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a DedicatedCephAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedCephAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedCephAclArgs | DedicatedCephAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DedicatedCephAclState | undefined;
            resourceInputs["family"] = state ? state.family : undefined;
            resourceInputs["netmask"] = state ? state.netmask : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as DedicatedCephAclArgs | undefined;
            if ((!args || args.netmask === undefined) && !opts.urn) {
                throw new Error("Missing required property 'netmask'");
            }
            if ((!args || args.network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'network'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["netmask"] = args ? args.netmask : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["family"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DedicatedCephAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DedicatedCephAcl resources.
 */
export interface DedicatedCephAclState {
    /**
     * IP family. `IPv4` or `IPv6`
     */
    family?: pulumi.Input<string>;
    /**
     * The network mask to apply
     */
    netmask?: pulumi.Input<string>;
    /**
     * The network IP to authorize
     */
    network?: pulumi.Input<string>;
    /**
     * The internal name of your dedicated CEPH
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DedicatedCephAcl resource.
 */
export interface DedicatedCephAclArgs {
    /**
     * The network mask to apply
     */
    netmask: pulumi.Input<string>;
    /**
     * The network IP to authorize
     */
    network: pulumi.Input<string>;
    /**
     * The internal name of your dedicated CEPH
     */
    serviceName: pulumi.Input<string>;
}
