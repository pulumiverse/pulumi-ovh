// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get the details of Vrack network available for your IPLoadbalancer associated with your OVH account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lbNetwork = pulumi.output(ovh.getIpLoadbalancingVrackNetwork({
 *     serviceName: "XXXXXX",
 *     vrackNetworkId: Number.parseFloat("yyy"),
 * }));
 * ```
 */
export function getIpLoadbalancingVrackNetwork(args: GetIpLoadbalancingVrackNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetIpLoadbalancingVrackNetworkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("ovh:index/getIpLoadbalancingVrackNetwork:getIpLoadbalancingVrackNetwork", {
        "serviceName": args.serviceName,
        "vrackNetworkId": args.vrackNetworkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpLoadbalancingVrackNetwork.
 */
export interface GetIpLoadbalancingVrackNetworkArgs {
    /**
     * The internal name of your IP load balancing
     */
    serviceName: string;
    /**
     * Internal Load Balancer identifier of the vRack private network
     */
    vrackNetworkId: number;
}

/**
 * A collection of values returned by getIpLoadbalancingVrackNetwork.
 */
export interface GetIpLoadbalancingVrackNetworkResult {
    /**
     * Human readable name for your vrack network
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
     */
    readonly natIp: string;
    readonly serviceName: string;
    /**
     * IP block of the private network in the vRack
     */
    readonly subnet: string;
    /**
     * VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
     */
    readonly vlan: number;
    readonly vrackNetworkId: number;
}

export function getIpLoadbalancingVrackNetworkOutput(args: GetIpLoadbalancingVrackNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpLoadbalancingVrackNetworkResult> {
    return pulumi.output(args).apply(a => getIpLoadbalancingVrackNetwork(a, opts))
}

/**
 * A collection of arguments for invoking getIpLoadbalancingVrackNetwork.
 */
export interface GetIpLoadbalancingVrackNetworkOutputArgs {
    /**
     * The internal name of your IP load balancing
     */
    serviceName: pulumi.Input<string>;
    /**
     * Internal Load Balancer identifier of the vRack private network
     */
    vrackNetworkId: pulumi.Input<number>;
}
