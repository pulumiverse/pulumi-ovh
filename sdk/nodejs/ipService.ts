// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Orders an ip service.
 *
 * ## Important
 *
 * This resource orders an OVH product for a long period of time and may generate heavy costs !
 * Use with caution.
 *
 * __NOTE__ 1: the "default-payment-mean" will scan your registered bank accounts, credit card and paypal payment means to find your default payment mean.
 *
 * __NOTE__ 2: this resource is in beta state. Use with caution.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 * import * as ovh from "@pulumiverse/ovh";
 *
 * const mycart = ovh.getOrderCart({
 *     ovhSubsidiary: "fr",
 *     description: "order ip block",
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
 *     description: "my ip block",
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
 * ```
 */
export class IpService extends pulumi.CustomResource {
    /**
     * Get an existing IpService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpServiceState, opts?: pulumi.CustomResourceOptions): IpService {
        return new IpService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/ipService:IpService';

    /**
     * Returns true if the given object is an instance of IpService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpService.__pulumiType;
    }

    /**
     * can be terminated
     */
    public /*out*/ readonly canBeTerminated!: pulumi.Output<boolean>;
    /**
     * country
     */
    public /*out*/ readonly country!: pulumi.Output<string>;
    /**
     * Custom description on your ip.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * ip block
     */
    public /*out*/ readonly ip!: pulumi.Output<string>;
    /**
     * Details about an Order
     */
    public /*out*/ readonly orders!: pulumi.Output<outputs.IpServiceOrder[]>;
    /**
     * IP block organisation Id
     */
    public /*out*/ readonly organisationId!: pulumi.Output<string>;
    /**
     * Ovh Subsidiary
     */
    public readonly ovhSubsidiary!: pulumi.Output<string>;
    /**
     * Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    public readonly paymentMean!: pulumi.Output<string>;
    /**
     * Product Plan to order
     */
    public readonly plan!: pulumi.Output<outputs.IpServicePlan>;
    /**
     * Product Plan to order
     */
    public readonly planOptions!: pulumi.Output<outputs.IpServicePlanOption[] | undefined>;
    /**
     * Routage information
     */
    public /*out*/ readonly routedTos!: pulumi.Output<outputs.IpServiceRoutedTo[]>;
    /**
     * Service where ip is routed to
     * * `serviceName`: service name
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * Possible values for ip type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IpService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpServiceArgs | IpServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpServiceState | undefined;
            resourceInputs["canBeTerminated"] = state ? state.canBeTerminated : undefined;
            resourceInputs["country"] = state ? state.country : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["orders"] = state ? state.orders : undefined;
            resourceInputs["organisationId"] = state ? state.organisationId : undefined;
            resourceInputs["ovhSubsidiary"] = state ? state.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = state ? state.paymentMean : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["planOptions"] = state ? state.planOptions : undefined;
            resourceInputs["routedTos"] = state ? state.routedTos : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as IpServiceArgs | undefined;
            if ((!args || args.ovhSubsidiary === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ovhSubsidiary'");
            }
            if ((!args || args.paymentMean === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentMean'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ovhSubsidiary"] = args ? args.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = args ? args.paymentMean : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["planOptions"] = args ? args.planOptions : undefined;
            resourceInputs["canBeTerminated"] = undefined /*out*/;
            resourceInputs["country"] = undefined /*out*/;
            resourceInputs["ip"] = undefined /*out*/;
            resourceInputs["orders"] = undefined /*out*/;
            resourceInputs["organisationId"] = undefined /*out*/;
            resourceInputs["routedTos"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpService resources.
 */
export interface IpServiceState {
    /**
     * can be terminated
     */
    canBeTerminated?: pulumi.Input<boolean>;
    /**
     * country
     */
    country?: pulumi.Input<string>;
    /**
     * Custom description on your ip.
     */
    description?: pulumi.Input<string>;
    /**
     * ip block
     */
    ip?: pulumi.Input<string>;
    /**
     * Details about an Order
     */
    orders?: pulumi.Input<pulumi.Input<inputs.IpServiceOrder>[]>;
    /**
     * IP block organisation Id
     */
    organisationId?: pulumi.Input<string>;
    /**
     * Ovh Subsidiary
     */
    ovhSubsidiary?: pulumi.Input<string>;
    /**
     * Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    paymentMean?: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan?: pulumi.Input<inputs.IpServicePlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.IpServicePlanOption>[]>;
    /**
     * Routage information
     */
    routedTos?: pulumi.Input<pulumi.Input<inputs.IpServiceRoutedTo>[]>;
    /**
     * Service where ip is routed to
     * * `serviceName`: service name
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Possible values for ip type
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpService resource.
 */
export interface IpServiceArgs {
    /**
     * Custom description on your ip.
     */
    description?: pulumi.Input<string>;
    /**
     * Ovh Subsidiary
     */
    ovhSubsidiary: pulumi.Input<string>;
    /**
     * Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    paymentMean: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan: pulumi.Input<inputs.IpServicePlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.IpServicePlanOption>[]>;
}
