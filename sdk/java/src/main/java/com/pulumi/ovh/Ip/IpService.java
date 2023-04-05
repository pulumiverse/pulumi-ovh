// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Ip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Ip.IpServiceArgs;
import com.pulumi.ovh.Ip.inputs.IpServiceState;
import com.pulumi.ovh.Ip.outputs.IpServiceOrder;
import com.pulumi.ovh.Ip.outputs.IpServicePlan;
import com.pulumi.ovh.Ip.outputs.IpServicePlanOption;
import com.pulumi.ovh.Ip.outputs.IpServiceRoutedTo;
import com.pulumi.ovh.Utilities;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Order.OrderFunctions;
 * import com.pulumi.ovh.Order.inputs.GetCartArgs;
 * import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
 * import com.pulumi.ovh.Ip.IpService;
 * import com.pulumi.ovh.Ip.IpServiceArgs;
 * import com.pulumi.ovh.Ip.inputs.IpServicePlanArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
 *             .ovhSubsidiary(&#34;fr&#34;)
 *             .description(&#34;order ip block&#34;)
 *             .build());
 * 
 *         final var ipblockCartProductPlan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
 *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
 *             .priceCapacity(&#34;renew&#34;)
 *             .product(&#34;ip&#34;)
 *             .planCode(&#34;ip-v4-s30-ripe&#34;)
 *             .build());
 * 
 *         var ipblockIpService = new IpService(&#34;ipblockIpService&#34;, IpServiceArgs.builder()        
 *             .ovhSubsidiary(mycart.applyValue(getCartResult -&gt; getCartResult.ovhSubsidiary()))
 *             .description(&#34;my ip block&#34;)
 *             .plan(IpServicePlanArgs.builder()
 *                 .duration(ipblockCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].duration()))
 *                 .planCode(ipblockCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.planCode()))
 *                 .pricingMode(ipblockCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].pricingMode()))
 *                 .configurations(IpServicePlanConfigurationArgs.builder()
 *                     .label(&#34;country&#34;)
 *                     .value(&#34;FR&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="ovh:Ip/ipService:IpService")
public class IpService extends com.pulumi.resources.CustomResource {
    /**
     * can be terminated
     * 
     */
    @Export(name="canBeTerminated", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> canBeTerminated;

    /**
     * @return can be terminated
     * 
     */
    public Output<Boolean> canBeTerminated() {
        return this.canBeTerminated;
    }
    /**
     * country
     * 
     */
    @Export(name="country", refs={String.class}, tree="[0]")
    private Output<String> country;

    /**
     * @return country
     * 
     */
    public Output<String> country() {
        return this.country;
    }
    /**
     * Custom description on your ip.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Custom description on your ip.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * ip block
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return ip block
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * Details about an Order
     * 
     */
    @Export(name="orders", refs={List.class,IpServiceOrder.class}, tree="[0,1]")
    private Output<List<IpServiceOrder>> orders;

    /**
     * @return Details about an Order
     * 
     */
    public Output<List<IpServiceOrder>> orders() {
        return this.orders;
    }
    /**
     * IP block organisation Id
     * 
     */
    @Export(name="organisationId", refs={String.class}, tree="[0]")
    private Output<String> organisationId;

    /**
     * @return IP block organisation Id
     * 
     */
    public Output<String> organisationId() {
        return this.organisationId;
    }
    /**
     * OVHcloud Subsidiary
     * 
     */
    @Export(name="ovhSubsidiary", refs={String.class}, tree="[0]")
    private Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary
     * 
     */
    public Output<String> ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * Ovh payment mode
     * 
     * @deprecated
     * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     * 
     */
    @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
    @Export(name="paymentMean", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> paymentMean;

    /**
     * @return Ovh payment mode
     * 
     */
    public Output<Optional<String>> paymentMean() {
        return Codegen.optional(this.paymentMean);
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="plan", refs={IpServicePlan.class}, tree="[0]")
    private Output<IpServicePlan> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<IpServicePlan> plan() {
        return this.plan;
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="planOptions", refs={List.class,IpServicePlanOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<IpServicePlanOption>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<Optional<List<IpServicePlanOption>>> planOptions() {
        return Codegen.optional(this.planOptions);
    }
    /**
     * Routage information
     * 
     */
    @Export(name="routedTos", refs={List.class,IpServiceRoutedTo.class}, tree="[0,1]")
    private Output<List<IpServiceRoutedTo>> routedTos;

    /**
     * @return Routage information
     * 
     */
    public Output<List<IpServiceRoutedTo>> routedTos() {
        return this.routedTos;
    }
    /**
     * service name
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Possible values for ip type
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Possible values for ip type
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IpService(String name) {
        this(name, IpServiceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpService(String name, IpServiceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpService(String name, IpServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Ip/ipService:IpService", name, args == null ? IpServiceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IpService(String name, Output<String> id, @Nullable IpServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Ip/ipService:IpService", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static IpService get(String name, Output<String> id, @Nullable IpServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpService(name, id, state, options);
    }
}
