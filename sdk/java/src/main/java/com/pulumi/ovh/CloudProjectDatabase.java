// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProjectDatabaseArgs;
import com.pulumi.ovh.Utilities;
import com.pulumi.ovh.inputs.CloudProjectDatabaseState;
import com.pulumi.ovh.outputs.CloudProjectDatabaseEndpoint;
import com.pulumi.ovh.outputs.CloudProjectDatabaseNode;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * To deploy a business PostgreSQL service with two nodes on public network:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProjectDatabase;
 * import com.pulumi.ovh.CloudProjectDatabaseArgs;
 * import com.pulumi.ovh.inputs.CloudProjectDatabaseNodeArgs;
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
 *         var postgresql = new CloudProjectDatabase(&#34;postgresql&#34;, CloudProjectDatabaseArgs.builder()        
 *             .description(&#34;my-first-postgresql&#34;)
 *             .engine(&#34;postgresql&#34;)
 *             .flavor(&#34;db1-15&#34;)
 *             .nodes(            
 *                 CloudProjectDatabaseNodeArgs.builder()
 *                     .region(&#34;GRA&#34;)
 *                     .build(),
 *                 CloudProjectDatabaseNodeArgs.builder()
 *                     .region(&#34;GRA&#34;)
 *                     .build())
 *             .plan(&#34;business&#34;)
 *             .serviceName(&#34;xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx&#34;)
 *             .version(&#34;14&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * To deploy an enterprise MongoDB service with three nodes on private network:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProjectDatabase;
 * import com.pulumi.ovh.CloudProjectDatabaseArgs;
 * import com.pulumi.ovh.inputs.CloudProjectDatabaseNodeArgs;
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
 *         var mongodb = new CloudProjectDatabase(&#34;mongodb&#34;, CloudProjectDatabaseArgs.builder()        
 *             .serviceName(var_.openstack_infos().project_id())
 *             .description(&#34;my-first-mongodb&#34;)
 *             .engine(&#34;mongodb&#34;)
 *             .version(&#34;5.0&#34;)
 *             .plan(&#34;enterprise&#34;)
 *             .nodes(            
 *                 CloudProjectDatabaseNodeArgs.builder()
 *                     .region(&#34;SBG&#34;)
 *                     .subnetId(&#34;XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX&#34;)
 *                     .networkId(&#34;XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX&#34;)
 *                     .build(),
 *                 CloudProjectDatabaseNodeArgs.builder()
 *                     .region(&#34;SBG&#34;)
 *                     .subnetId(&#34;XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX&#34;)
 *                     .networkId(&#34;XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX&#34;)
 *                     .build(),
 *                 CloudProjectDatabaseNodeArgs.builder()
 *                     .region(&#34;SBG&#34;)
 *                     .subnetId(&#34;XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX&#34;)
 *                     .networkId(&#34;XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX&#34;)
 *                     .build())
 *             .flavor(&#34;db1-30&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OVHcloud Managed database clusters can be imported using the `service_name`, `engine`, `id` of the cluster, separated by &#34;/&#34; E.g.,
 * 
 * ```sh
 *  $ pulumi import ovh:index/cloudProjectDatabase:CloudProjectDatabase my_database_cluster &lt;service_name&gt;/&lt;engine&gt;/&lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="ovh:index/cloudProjectDatabase:CloudProjectDatabase")
public class CloudProjectDatabase extends com.pulumi.resources.CustomResource {
    /**
     * Time on which backups start every day.
     * 
     */
    @Export(name="backupTime", type=String.class, parameters={})
    private Output<String> backupTime;

    /**
     * @return Time on which backups start every day.
     * 
     */
    public Output<String> backupTime() {
        return this.backupTime;
    }
    /**
     * Date of the creation of the cluster.
     * 
     */
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    /**
     * @return Date of the creation of the cluster.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Small description of the database service.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Small description of the database service.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * List of all endpoints objects of the service.
     * 
     */
    @Export(name="endpoints", type=List.class, parameters={CloudProjectDatabaseEndpoint.class})
    private Output<List<CloudProjectDatabaseEndpoint>> endpoints;

    /**
     * @return List of all endpoints objects of the service.
     * 
     */
    public Output<List<CloudProjectDatabaseEndpoint>> endpoints() {
        return this.endpoints;
    }
    /**
     * The database engine you want to deploy. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    @Export(name="engine", type=String.class, parameters={})
    private Output<String> engine;

    /**
     * @return The database engine you want to deploy. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * a valid OVH public cloud database flavor name in which the nodes will be started.
     * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
     * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
     * 
     */
    @Export(name="flavor", type=String.class, parameters={})
    private Output<String> flavor;

    /**
     * @return a valid OVH public cloud database flavor name in which the nodes will be started.
     * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
     * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
     * 
     */
    public Output<String> flavor() {
        return this.flavor;
    }
    /**
     * Time on which maintenances can start every day.
     * 
     */
    @Export(name="maintenanceTime", type=String.class, parameters={})
    private Output<String> maintenanceTime;

    /**
     * @return Time on which maintenances can start every day.
     * 
     */
    public Output<String> maintenanceTime() {
        return this.maintenanceTime;
    }
    /**
     * Type of network of the cluster.
     * 
     */
    @Export(name="networkType", type=String.class, parameters={})
    private Output<String> networkType;

    /**
     * @return Type of network of the cluster.
     * 
     */
    public Output<String> networkType() {
        return this.networkType;
    }
    /**
     * List of nodes object.
     * Multi region cluster are not yet available, all node should be identical.
     * 
     */
    @Export(name="nodes", type=List.class, parameters={CloudProjectDatabaseNode.class})
    private Output<List<CloudProjectDatabaseNode>> nodes;

    /**
     * @return List of nodes object.
     * Multi region cluster are not yet available, all node should be identical.
     * 
     */
    public Output<List<CloudProjectDatabaseNode>> nodes() {
        return this.nodes;
    }
    /**
     * List of nodes object.
     * Enum: &#34;esential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
     * 
     */
    @Export(name="plan", type=String.class, parameters={})
    private Output<String> plan;

    /**
     * @return List of nodes object.
     * Enum: &#34;esential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
     * 
     */
    public Output<String> plan() {
        return this.plan;
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
     * Current status of the cluster.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Current status of the cluster.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The version of the engine in which the service should be deployed
     * 
     */
    @Export(name="version", type=String.class, parameters={})
    private Output<String> version;

    /**
     * @return The version of the engine in which the service should be deployed
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudProjectDatabase(String name) {
        this(name, CloudProjectDatabaseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudProjectDatabase(String name, CloudProjectDatabaseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudProjectDatabase(String name, CloudProjectDatabaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:index/cloudProjectDatabase:CloudProjectDatabase", name, args == null ? CloudProjectDatabaseArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CloudProjectDatabase(String name, Output<String> id, @Nullable CloudProjectDatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:index/cloudProjectDatabase:CloudProjectDatabase", name, state, makeResourceOptions(options, id));
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
    public static CloudProjectDatabase get(String name, Output<String> id, @Nullable CloudProjectDatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudProjectDatabase(name, id, state, options);
    }
}
