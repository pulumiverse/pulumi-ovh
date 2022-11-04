// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProjectDatabase;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProjectDatabase.KafkaTopicArgs;
import com.pulumi.ovh.CloudProjectDatabase.inputs.KafkaTopicState;
import com.pulumi.ovh.Utilities;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates a topic for a kafka cluster associated with a public cloud project.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProjectDatabase.CloudProjectDatabaseFunctions;
 * import com.pulumi.ovh.CloudProjectDatabase.inputs.GetDatabaseArgs;
 * import com.pulumi.ovh.CloudProjectDatabase.KafkaTopic;
 * import com.pulumi.ovh.CloudProjectDatabase.KafkaTopicArgs;
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
 *         final var kafka = CloudProjectDatabaseFunctions.getDatabase(GetDatabaseArgs.builder()
 *             .serviceName(&#34;XXX&#34;)
 *             .engine(&#34;kafka&#34;)
 *             .id(&#34;ZZZ&#34;)
 *             .build());
 * 
 *         var topic = new KafkaTopic(&#34;topic&#34;, KafkaTopicArgs.builder()        
 *             .serviceName(kafka.applyValue(getDatabaseResult -&gt; getDatabaseResult.serviceName()))
 *             .clusterId(kafka.applyValue(getDatabaseResult -&gt; getDatabaseResult.id()))
 *             .minInsyncReplicas(1)
 *             .partitions(3)
 *             .replication(2)
 *             .retentionBytes(4)
 *             .retentionHours(5)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OVHcloud Managed kafka clusters topics can be imported using the `service_name`, `cluster_id` and `id` of the topic, separated by &#34;/&#34; E.g., bash
 * 
 * ```sh
 *  $ pulumi import ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic my_topic service_name/cluster_id/id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic")
public class KafkaTopic extends com.pulumi.resources.CustomResource {
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
     * Minimum insync replica accepted for this topic. Should be superior to 0
     * 
     */
    @Export(name="minInsyncReplicas", type=Integer.class, parameters={})
    private Output<Integer> minInsyncReplicas;

    /**
     * @return Minimum insync replica accepted for this topic. Should be superior to 0
     * 
     */
    public Output<Integer> minInsyncReplicas() {
        return this.minInsyncReplicas;
    }
    /**
     * Name of the topic. No spaces allowed.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the topic. No spaces allowed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Number of partitions for this topic. Should be superior to 0
     * 
     */
    @Export(name="partitions", type=Integer.class, parameters={})
    private Output<Integer> partitions;

    /**
     * @return Number of partitions for this topic. Should be superior to 0
     * 
     */
    public Output<Integer> partitions() {
        return this.partitions;
    }
    /**
     * Number of replication for this topic. Should be superior to 1
     * 
     */
    @Export(name="replication", type=Integer.class, parameters={})
    private Output<Integer> replication;

    /**
     * @return Number of replication for this topic. Should be superior to 1
     * 
     */
    public Output<Integer> replication() {
        return this.replication;
    }
    /**
     * Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
     * 
     */
    @Export(name="retentionBytes", type=Integer.class, parameters={})
    private Output<Integer> retentionBytes;

    /**
     * @return Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
     * 
     */
    public Output<Integer> retentionBytes() {
        return this.retentionBytes;
    }
    /**
     * Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
     * 
     */
    @Export(name="retentionHours", type=Integer.class, parameters={})
    private Output<Integer> retentionHours;

    /**
     * @return Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
     * 
     */
    public Output<Integer> retentionHours() {
        return this.retentionHours;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KafkaTopic(String name) {
        this(name, KafkaTopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KafkaTopic(String name, KafkaTopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KafkaTopic(String name, KafkaTopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic", name, args == null ? KafkaTopicArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KafkaTopic(String name, Output<String> id, @Nullable KafkaTopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic", name, state, makeResourceOptions(options, id));
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
    public static KafkaTopic get(String name, Output<String> id, @Nullable KafkaTopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KafkaTopic(name, id, state, options);
    }
}
