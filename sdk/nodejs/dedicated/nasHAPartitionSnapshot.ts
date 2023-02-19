// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource for managing **snapshot** to partitions on HA-NAS services
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const my_partition = new ovh.Dedicated.NasHAPartitionSnapshot("my-partition", {
 *     partitionName: "my-partition",
 *     serviceName: "zpool-12345",
 *     type: "day-3",
 * });
 * ```
 *
 * ## Import
 *
 * HA-NAS partition snapshot can be imported using the `{service_name}/{partition_name}/{type}`, e.g.
 *
 * ```sh
 *  $ pulumi import ovh:Dedicated/nasHAPartitionSnapshot:NasHAPartitionSnapshot my-partition zpool-12345/my-partition/day-3`
 * ```
 */
export class NasHAPartitionSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing NasHAPartitionSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NasHAPartitionSnapshotState, opts?: pulumi.CustomResourceOptions): NasHAPartitionSnapshot {
        return new NasHAPartitionSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dedicated/nasHAPartitionSnapshot:NasHAPartitionSnapshot';

    /**
     * Returns true if the given object is an instance of NasHAPartitionSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NasHAPartitionSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NasHAPartitionSnapshot.__pulumiType;
    }

    /**
     * name of the partition
     */
    public readonly partitionName!: pulumi.Output<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Snapshot interval, allowed : day-1, day-2, day-3, day-7, hour-1, hour-6
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a NasHAPartitionSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NasHAPartitionSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NasHAPartitionSnapshotArgs | NasHAPartitionSnapshotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NasHAPartitionSnapshotState | undefined;
            resourceInputs["partitionName"] = state ? state.partitionName : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as NasHAPartitionSnapshotArgs | undefined;
            if ((!args || args.partitionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitionName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["partitionName"] = args ? args.partitionName : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NasHAPartitionSnapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NasHAPartitionSnapshot resources.
 */
export interface NasHAPartitionSnapshotState {
    /**
     * name of the partition
     */
    partitionName?: pulumi.Input<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Snapshot interval, allowed : day-1, day-2, day-3, day-7, hour-1, hour-6
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NasHAPartitionSnapshot resource.
 */
export interface NasHAPartitionSnapshotArgs {
    /**
     * name of the partition
     */
    partitionName: pulumi.Input<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    serviceName: pulumi.Input<string>;
    /**
     * Snapshot interval, allowed : day-1, day-2, day-3, day-7, hour-1, hour-6
     */
    type: pulumi.Input<string>;
}
