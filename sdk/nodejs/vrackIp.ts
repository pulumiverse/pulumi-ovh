// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Attach a Ip block to a VRack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mycart = ovh.getOrderCart({
 *     ovhSubsidiary: "fr",
 *     description: "my cart",
 * });
 * const vrackOrderCartProductPlan = mycart.then(mycart => ovh.getOrderCartProductPlan({
 *     cartId: mycart.id,
 *     priceCapacity: "renew",
 *     product: "vrack",
 *     planCode: "vrack",
 * }));
 * const vrackVrack = new ovh.Vrack("vrackVrack", {
 *     description: mycart.then(mycart => mycart.description),
 *     ovhSubsidiary: mycart.then(mycart => mycart.ovhSubsidiary),
 *     paymentMean: "fidelity",
 *     plan: {
 *         duration: vrackOrderCartProductPlan.then(vrackOrderCartProductPlan => vrackOrderCartProductPlan.selectedPrices?[0]?.duration),
 *         planCode: vrackOrderCartProductPlan.then(vrackOrderCartProductPlan => vrackOrderCartProductPlan.planCode),
 *         pricingMode: vrackOrderCartProductPlan.then(vrackOrderCartProductPlan => vrackOrderCartProductPlan.selectedPrices?[0]?.pricingMode),
 *     },
 * });
 * const ipblockOrderCartProductPlan = mycart.then(mycart => ovh.getOrderCartProductPlan({
 *     cartId: mycart.id,
 *     priceCapacity: "renew",
 *     product: "ip",
 *     planCode: "ip-v4-s30-ripe",
 * }));
 * const ipblockIpService = new ovh.IpService("ipblockIpService", {
 *     ovhSubsidiary: mycart.then(mycart => mycart.ovhSubsidiary),
 *     paymentMean: "ovh-account",
 *     description: mycart.then(mycart => mycart.description),
 *     plan: {
 *         duration: ipblockOrderCartProductPlan.then(ipblockOrderCartProductPlan => ipblockOrderCartProductPlan.selectedPrices?[0]?.duration),
 *         planCode: ipblockOrderCartProductPlan.then(ipblockOrderCartProductPlan => ipblockOrderCartProductPlan.planCode),
 *         pricingMode: ipblockOrderCartProductPlan.then(ipblockOrderCartProductPlan => ipblockOrderCartProductPlan.selectedPrices?[0]?.pricingMode),
 *         configurations: [{
 *             label: "country",
 *             value: "FR",
 *         }],
 *     },
 * });
 * const vrackblock = new ovh.VrackIp("vrackblock", {
 *     serviceName: vrackVrack.serviceName,
 *     block: ipblockIpService.ip,
 * });
 * ```
 */
export class VrackIp extends pulumi.CustomResource {
    /**
     * Get an existing VrackIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VrackIpState, opts?: pulumi.CustomResourceOptions): VrackIp {
        return new VrackIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/vrackIp:VrackIp';

    /**
     * Returns true if the given object is an instance of VrackIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VrackIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VrackIp.__pulumiType;
    }

    /**
     * Your IP block.
     */
    public readonly block!: pulumi.Output<string>;
    /**
     * Your gateway
     */
    public /*out*/ readonly gateway!: pulumi.Output<string>;
    /**
     * Your IP block
     */
    public /*out*/ readonly ip!: pulumi.Output<string>;
    /**
     * The internal name of your vrack
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Where you want your block announced on the network
     */
    public /*out*/ readonly zone!: pulumi.Output<string>;

    /**
     * Create a VrackIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VrackIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VrackIpArgs | VrackIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VrackIpState | undefined;
            resourceInputs["block"] = state ? state.block : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as VrackIpArgs | undefined;
            if ((!args || args.block === undefined) && !opts.urn) {
                throw new Error("Missing required property 'block'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["block"] = args ? args.block : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["gateway"] = undefined /*out*/;
            resourceInputs["ip"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VrackIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VrackIp resources.
 */
export interface VrackIpState {
    /**
     * Your IP block.
     */
    block?: pulumi.Input<string>;
    /**
     * Your gateway
     */
    gateway?: pulumi.Input<string>;
    /**
     * Your IP block
     */
    ip?: pulumi.Input<string>;
    /**
     * The internal name of your vrack
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Where you want your block announced on the network
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VrackIp resource.
 */
export interface VrackIpArgs {
    /**
     * Your IP block.
     */
    block: pulumi.Input<string>;
    /**
     * The internal name of your vrack
     */
    serviceName: pulumi.Input<string>;
}
