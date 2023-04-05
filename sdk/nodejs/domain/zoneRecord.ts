// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a OVHcloud domain zone record.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 *
 * // Add a record to a sub-domain
 * const test = new ovh.domain.ZoneRecord("test", {
 *     fieldtype: "A",
 *     subdomain: "test",
 *     target: "0.0.0.0",
 *     ttl: 3600,
 *     zone: "testdemo.ovh",
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud domain zone record can be imported using the `id`, which can be retrieved by using [OVH API portal](https://api.ovh.com/console/#/domain/zone/%7BzoneName%7D/record~GET), and the `zone`, separated by "." E.g., bash
 *
 * ```sh
 *  $ pulumi import ovh:Domain/zoneRecord:ZoneRecord test id.zone
 * ```
 */
export class ZoneRecord extends pulumi.CustomResource {
    /**
     * Get an existing ZoneRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneRecordState, opts?: pulumi.CustomResourceOptions): ZoneRecord {
        return new ZoneRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Domain/zoneRecord:ZoneRecord';

    /**
     * Returns true if the given object is an instance of ZoneRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZoneRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZoneRecord.__pulumiType;
    }

    /**
     * The type of the record
     */
    public readonly fieldtype!: pulumi.Output<string>;
    /**
     * The name of the record
     */
    public readonly subdomain!: pulumi.Output<string | undefined>;
    /**
     * The value of the record
     */
    public readonly target!: pulumi.Output<string>;
    /**
     * The TTL of the record, it shall be >= to 60.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The domain to add the record to
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a ZoneRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneRecordArgs | ZoneRecordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneRecordState | undefined;
            resourceInputs["fieldtype"] = state ? state.fieldtype : undefined;
            resourceInputs["subdomain"] = state ? state.subdomain : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ZoneRecordArgs | undefined;
            if ((!args || args.fieldtype === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fieldtype'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["fieldtype"] = args ? args.fieldtype : undefined;
            resourceInputs["subdomain"] = args ? args.subdomain : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZoneRecord.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZoneRecord resources.
 */
export interface ZoneRecordState {
    /**
     * The type of the record
     */
    fieldtype?: pulumi.Input<string>;
    /**
     * The name of the record
     */
    subdomain?: pulumi.Input<string>;
    /**
     * The value of the record
     */
    target?: pulumi.Input<string>;
    /**
     * The TTL of the record, it shall be >= to 60.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The domain to add the record to
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZoneRecord resource.
 */
export interface ZoneRecordArgs {
    /**
     * The type of the record
     */
    fieldtype: pulumi.Input<string>;
    /**
     * The name of the record
     */
    subdomain?: pulumi.Input<string>;
    /**
     * The value of the record
     */
    target: pulumi.Input<string>;
    /**
     * The TTL of the record, it shall be >= to 60.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The domain to add the record to
     */
    zone: pulumi.Input<string>;
}
