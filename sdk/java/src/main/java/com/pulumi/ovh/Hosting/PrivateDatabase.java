// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Hosting.PrivateDatabaseArgs;
import com.pulumi.ovh.Hosting.inputs.PrivateDatabaseState;
import com.pulumi.ovh.Hosting.outputs.PrivateDatabaseOrder;
import com.pulumi.ovh.Hosting.outputs.PrivateDatabasePlan;
import com.pulumi.ovh.Hosting.outputs.PrivateDatabasePlanOption;
import com.pulumi.ovh.Utilities;
import java.lang.Double;
import java.lang.Integer;
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
 * import com.pulumi.ovh.Hosting.PrivateDatabase;
 * import com.pulumi.ovh.Hosting.PrivateDatabaseArgs;
 * import com.pulumi.ovh.Hosting.inputs.PrivateDatabasePlanArgs;
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
 *             .description(&#34;my privatedatabase order cart&#34;)
 *             .build());
 * 
 *         final var databaseCartProductPlan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
 *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
 *             .priceCapacity(&#34;renew&#34;)
 *             .product(&#34;privateSQL&#34;)
 *             .planCode(&#34;private-sql-512-instance&#34;)
 *             .build());
 * 
 *         var databasePrivateDatabase = new PrivateDatabase(&#34;databasePrivateDatabase&#34;, PrivateDatabaseArgs.builder()        
 *             .ovhSubsidiary(mycart.applyValue(getCartResult -&gt; getCartResult.ovhSubsidiary()))
 *             .displayName(&#34;Postgresql-12&#34;)
 *             .plan(PrivateDatabasePlanArgs.builder()
 *                 .duration(databaseCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.prices()[3].duration()))
 *                 .planCode(databaseCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.planCode()))
 *                 .pricingMode(databaseCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].pricingMode()))
 *                 .configurations(                
 *                     PrivateDatabasePlanConfigurationArgs.builder()
 *                         .label(&#34;dc&#34;)
 *                         .value(&#34;gra3&#34;)
 *                         .build(),
 *                     PrivateDatabasePlanConfigurationArgs.builder()
 *                         .label(&#34;engine&#34;)
 *                         .value(&#34;postgresql_12&#34;)
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *         ctx.export(&#34;privatedatabaseServiceName&#34;, databasePrivateDatabase.serviceName());
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OVHcloud Webhosting database can be imported using the `service_name`, E.g.,
 * 
 * ```sh
 *  $ pulumi import ovh:Hosting/privateDatabase:PrivateDatabase database service_name
 * ```
 * 
 */
@ResourceType(type="ovh:Hosting/privateDatabase:PrivateDatabase")
public class PrivateDatabase extends com.pulumi.resources.CustomResource {
    /**
     * Number of CPU on your private database
     * 
     */
    @Export(name="cpu", refs={Integer.class}, tree="[0]")
    private Output<Integer> cpu;

    /**
     * @return Number of CPU on your private database
     * 
     */
    public Output<Integer> cpu() {
        return this.cpu;
    }
    /**
     * Datacenter where this private database is located
     * 
     */
    @Export(name="datacenter", refs={String.class}, tree="[0]")
    private Output<String> datacenter;

    /**
     * @return Datacenter where this private database is located
     * 
     */
    public Output<String> datacenter() {
        return this.datacenter;
    }
    /**
     * Name displayed in customer panel for your private database
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Name displayed in customer panel for your private database
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Private database hostname
     * 
     */
    @Export(name="hostname", refs={String.class}, tree="[0]")
    private Output<String> hostname;

    /**
     * @return Private database hostname
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }
    /**
     * Private database FTP hostname
     * 
     */
    @Export(name="hostnameFtp", refs={String.class}, tree="[0]")
    private Output<String> hostnameFtp;

    /**
     * @return Private database FTP hostname
     * 
     */
    public Output<String> hostnameFtp() {
        return this.hostnameFtp;
    }
    /**
     * Infrastructure where service was stored
     * 
     */
    @Export(name="infrastructure", refs={String.class}, tree="[0]")
    private Output<String> infrastructure;

    /**
     * @return Infrastructure where service was stored
     * 
     */
    public Output<String> infrastructure() {
        return this.infrastructure;
    }
    /**
     * Type of the private database offer
     * 
     */
    @Export(name="offer", refs={String.class}, tree="[0]")
    private Output<String> offer;

    /**
     * @return Type of the private database offer
     * 
     */
    public Output<String> offer() {
        return this.offer;
    }
    /**
     * Details about your Order
     * 
     */
    @Export(name="orders", refs={List.class,PrivateDatabaseOrder.class}, tree="[0,1]")
    private Output<List<PrivateDatabaseOrder>> orders;

    /**
     * @return Details about your Order
     * 
     */
    public Output<List<PrivateDatabaseOrder>> orders() {
        return this.orders;
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
    @Export(name="plan", refs={PrivateDatabasePlan.class}, tree="[0]")
    private Output<PrivateDatabasePlan> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<PrivateDatabasePlan> plan() {
        return this.plan;
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="planOptions", refs={List.class,PrivateDatabasePlanOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PrivateDatabasePlanOption>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<Optional<List<PrivateDatabasePlanOption>>> planOptions() {
        return Codegen.optional(this.planOptions);
    }
    /**
     * Private database service port
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return Private database service port
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * Private database FTP port
     * 
     */
    @Export(name="portFtp", refs={Integer.class}, tree="[0]")
    private Output<Integer> portFtp;

    /**
     * @return Private database FTP port
     * 
     */
    public Output<Integer> portFtp() {
        return this.portFtp;
    }
    /**
     * Space allowed (in MB) on your private database
     * 
     */
    @Export(name="quotaSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> quotaSize;

    /**
     * @return Space allowed (in MB) on your private database
     * 
     */
    public Output<Integer> quotaSize() {
        return this.quotaSize;
    }
    /**
     * Sapce used (in MB) on your private database
     * 
     */
    @Export(name="quotaUsed", refs={Integer.class}, tree="[0]")
    private Output<Integer> quotaUsed;

    /**
     * @return Sapce used (in MB) on your private database
     * 
     */
    public Output<Integer> quotaUsed() {
        return this.quotaUsed;
    }
    /**
     * Amount of ram (in MB) on your private database
     * 
     */
    @Export(name="ram", refs={Integer.class}, tree="[0]")
    private Output<Integer> ram;

    /**
     * @return Amount of ram (in MB) on your private database
     * 
     */
    public Output<Integer> ram() {
        return this.ram;
    }
    /**
     * Private database server name
     * 
     */
    @Export(name="server", refs={String.class}, tree="[0]")
    private Output<String> server;

    /**
     * @return Private database server name
     * 
     */
    public Output<String> server() {
        return this.server;
    }
    /**
     * Service name
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return Service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Private database state
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Private database state
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Private database type
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Private database type
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * URN of the private database, used when writing IAM policies
     * 
     */
    @Export(name="urn", refs={String.class}, tree="[0]")
    private Output<String> urn;

    /**
     * @return URN of the private database, used when writing IAM policies
     * 
     */
    public Output<String> urn_() {
        return this.urn;
    }
    /**
     * Private database available versions
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return Private database available versions
     * 
     */
    public Output<String> version() {
        return this.version;
    }
    /**
     * Private database version label
     * 
     */
    @Export(name="versionLabel", refs={String.class}, tree="[0]")
    private Output<String> versionLabel;

    /**
     * @return Private database version label
     * 
     */
    public Output<String> versionLabel() {
        return this.versionLabel;
    }
    /**
     * Private database version number
     * 
     */
    @Export(name="versionNumber", refs={Double.class}, tree="[0]")
    private Output<Double> versionNumber;

    /**
     * @return Private database version number
     * 
     */
    public Output<Double> versionNumber() {
        return this.versionNumber;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivateDatabase(String name) {
        this(name, PrivateDatabaseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivateDatabase(String name, PrivateDatabaseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivateDatabase(String name, PrivateDatabaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabase:PrivateDatabase", name, args == null ? PrivateDatabaseArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivateDatabase(String name, Output<String> id, @Nullable PrivateDatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabase:PrivateDatabase", name, state, makeResourceOptions(options, id));
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
    public static PrivateDatabase get(String name, Output<String> id, @Nullable PrivateDatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivateDatabase(name, id, state, options);
    }
}
