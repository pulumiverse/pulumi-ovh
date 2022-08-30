// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProjectDatabaseIpRestrictionArgs;
import com.pulumi.ovh.Utilities;
import com.pulumi.ovh.inputs.CloudProjectDatabaseIpRestrictionState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Apply IP restrictions to an OVHcloud Managed Database cluster.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.OvhFunctions;
 * import com.pulumi.ovh.inputs.GetCloudProjectDatabaseArgs;
 * import com.pulumi.ovh.CloudProjectDatabaseIpRestriction;
 * import com.pulumi.ovh.CloudProjectDatabaseIpRestrictionArgs;
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
 *         final var db = OvhFunctions.getCloudProjectDatabase(GetCloudProjectDatabaseArgs.builder()
 *             .serviceName(&#34;XXXX&#34;)
 *             .engine(&#34;YYYY&#34;)
 *             .clusterId(&#34;ZZZZ&#34;)
 *             .build());
 * 
 *         var iprestriction = new CloudProjectDatabaseIpRestriction(&#34;iprestriction&#34;, CloudProjectDatabaseIpRestrictionArgs.builder()        
 *             .serviceName(ovh_cloud_project_database.db().service_name())
 *             .engine(ovh_cloud_project_database.db().engine())
 *             .clusterId(ovh_cloud_project_database.db().cluster_id())
 *             .ip(&#34;178.97.6.0/24&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OVHcloud Managed database cluster IP restrictions can be imported using the `service_name`, `engine`, `cluster_id` and the `ip`, separated by &#34;/&#34; E.g.,
 * 
 * ```sh
 *  $ pulumi import ovh:index/cloudProjectDatabaseIpRestriction:CloudProjectDatabaseIpRestriction my_ip_restriction &lt;service_name&gt;/&lt;engine&gt;/&lt;cluster_id&gt;/178.97.6.0/24
 * ```
 * 
 */
@ResourceType(type="ovh:index/cloudProjectDatabaseIpRestriction:CloudProjectDatabaseIpRestriction")
public class CloudProjectDatabaseIpRestriction extends com.pulumi.resources.CustomResource {
    /**
     * Cluster ID.
     * 
     */
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Description of the IP restriction.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the IP restriction.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    @Export(name="engine", type=String.class, parameters={})
    private Output<String> engine;

    /**
     * @return The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * Authorized IP.
     * 
     */
    @Export(name="ip", type=String.class, parameters={})
    private Output<String> ip;

    /**
     * @return Authorized IP.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", type=String.class, parameters={})
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Current status of the IP restriction.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Current status of the IP restriction.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudProjectDatabaseIpRestriction(String name) {
        this(name, CloudProjectDatabaseIpRestrictionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudProjectDatabaseIpRestriction(String name, CloudProjectDatabaseIpRestrictionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudProjectDatabaseIpRestriction(String name, CloudProjectDatabaseIpRestrictionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:index/cloudProjectDatabaseIpRestriction:CloudProjectDatabaseIpRestriction", name, args == null ? CloudProjectDatabaseIpRestrictionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CloudProjectDatabaseIpRestriction(String name, Output<String> id, @Nullable CloudProjectDatabaseIpRestrictionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:index/cloudProjectDatabaseIpRestriction:CloudProjectDatabaseIpRestriction", name, state, makeResourceOptions(options, id));
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
    public static CloudProjectDatabaseIpRestriction get(String name, Output<String> id, @Nullable CloudProjectDatabaseIpRestrictionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudProjectDatabaseIpRestriction(name, id, state, options);
    }
}
