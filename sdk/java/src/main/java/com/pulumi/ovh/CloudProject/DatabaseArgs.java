// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.CloudProject.inputs.DatabaseNodeArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseArgs Empty = new DatabaseArgs();

    /**
     * Advanced configuration key / value.
     * 
     */
    @Import(name="advancedConfiguration")
    private @Nullable Output<Map<String,String>> advancedConfiguration;

    /**
     * @return Advanced configuration key / value.
     * 
     */
    public Optional<Output<Map<String,String>>> advancedConfiguration() {
        return Optional.ofNullable(this.advancedConfiguration);
    }

    /**
     * Small description of the database service.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Small description of the database service.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The disk size (in GB) of the database service.
     * 
     */
    @Import(name="diskSize")
    private @Nullable Output<Integer> diskSize;

    /**
     * @return The disk size (in GB) of the database service.
     * 
     */
    public Optional<Output<Integer>> diskSize() {
        return Optional.ofNullable(this.diskSize);
    }

    /**
     * The database engine you want to deploy. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    @Import(name="engine", required=true)
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
     * A valid OVHcloud public cloud database flavor name in which the nodes will be started.
     * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
     * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
     * 
     */
    @Import(name="flavor", required=true)
    private Output<String> flavor;

    /**
     * @return A valid OVHcloud public cloud database flavor name in which the nodes will be started.
     * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
     * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
     * 
     */
    public Output<String> flavor() {
        return this.flavor;
    }

    /**
     * Defines whether the REST API is enabled on a kafka cluster
     * 
     */
    @Import(name="kafkaRestApi")
    private @Nullable Output<Boolean> kafkaRestApi;

    /**
     * @return Defines whether the REST API is enabled on a kafka cluster
     * 
     */
    public Optional<Output<Boolean>> kafkaRestApi() {
        return Optional.ofNullable(this.kafkaRestApi);
    }

    /**
     * List of nodes object.
     * Multi region cluster are not yet available, all node should be identical.
     * 
     */
    @Import(name="nodes", required=true)
    private Output<List<DatabaseNodeArgs>> nodes;

    /**
     * @return List of nodes object.
     * Multi region cluster are not yet available, all node should be identical.
     * 
     */
    public Output<List<DatabaseNodeArgs>> nodes() {
        return this.nodes;
    }

    /**
     * Defines whether the ACLs are enabled on an OpenSearch cluster
     * 
     */
    @Import(name="opensearchAclsEnabled")
    private @Nullable Output<Boolean> opensearchAclsEnabled;

    /**
     * @return Defines whether the ACLs are enabled on an OpenSearch cluster
     * 
     */
    public Optional<Output<Boolean>> opensearchAclsEnabled() {
        return Optional.ofNullable(this.opensearchAclsEnabled);
    }

    /**
     * Plan of the cluster.
     * Enum: &#34;essential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
     * 
     */
    @Import(name="plan", required=true)
    private Output<String> plan;

    /**
     * @return Plan of the cluster.
     * Enum: &#34;essential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
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
    @Import(name="serviceName", required=true)
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
     * The version of the engine in which the service should be deployed
     * 
     */
    @Import(name="version", required=true)
    private Output<String> version;

    /**
     * @return The version of the engine in which the service should be deployed
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    private DatabaseArgs() {}

    private DatabaseArgs(DatabaseArgs $) {
        this.advancedConfiguration = $.advancedConfiguration;
        this.description = $.description;
        this.diskSize = $.diskSize;
        this.engine = $.engine;
        this.flavor = $.flavor;
        this.kafkaRestApi = $.kafkaRestApi;
        this.nodes = $.nodes;
        this.opensearchAclsEnabled = $.opensearchAclsEnabled;
        this.plan = $.plan;
        this.serviceName = $.serviceName;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseArgs $;

        public Builder() {
            $ = new DatabaseArgs();
        }

        public Builder(DatabaseArgs defaults) {
            $ = new DatabaseArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param advancedConfiguration Advanced configuration key / value.
         * 
         * @return builder
         * 
         */
        public Builder advancedConfiguration(@Nullable Output<Map<String,String>> advancedConfiguration) {
            $.advancedConfiguration = advancedConfiguration;
            return this;
        }

        /**
         * @param advancedConfiguration Advanced configuration key / value.
         * 
         * @return builder
         * 
         */
        public Builder advancedConfiguration(Map<String,String> advancedConfiguration) {
            return advancedConfiguration(Output.of(advancedConfiguration));
        }

        /**
         * @param description Small description of the database service.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Small description of the database service.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param diskSize The disk size (in GB) of the database service.
         * 
         * @return builder
         * 
         */
        public Builder diskSize(@Nullable Output<Integer> diskSize) {
            $.diskSize = diskSize;
            return this;
        }

        /**
         * @param diskSize The disk size (in GB) of the database service.
         * 
         * @return builder
         * 
         */
        public Builder diskSize(Integer diskSize) {
            return diskSize(Output.of(diskSize));
        }

        /**
         * @param engine The database engine you want to deploy. To get a full list of available engine visit.
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine The database engine you want to deploy. To get a full list of available engine visit.
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param flavor A valid OVHcloud public cloud database flavor name in which the nodes will be started.
         * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
         * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
         * 
         * @return builder
         * 
         */
        public Builder flavor(Output<String> flavor) {
            $.flavor = flavor;
            return this;
        }

        /**
         * @param flavor A valid OVHcloud public cloud database flavor name in which the nodes will be started.
         * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
         * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
         * 
         * @return builder
         * 
         */
        public Builder flavor(String flavor) {
            return flavor(Output.of(flavor));
        }

        /**
         * @param kafkaRestApi Defines whether the REST API is enabled on a kafka cluster
         * 
         * @return builder
         * 
         */
        public Builder kafkaRestApi(@Nullable Output<Boolean> kafkaRestApi) {
            $.kafkaRestApi = kafkaRestApi;
            return this;
        }

        /**
         * @param kafkaRestApi Defines whether the REST API is enabled on a kafka cluster
         * 
         * @return builder
         * 
         */
        public Builder kafkaRestApi(Boolean kafkaRestApi) {
            return kafkaRestApi(Output.of(kafkaRestApi));
        }

        /**
         * @param nodes List of nodes object.
         * Multi region cluster are not yet available, all node should be identical.
         * 
         * @return builder
         * 
         */
        public Builder nodes(Output<List<DatabaseNodeArgs>> nodes) {
            $.nodes = nodes;
            return this;
        }

        /**
         * @param nodes List of nodes object.
         * Multi region cluster are not yet available, all node should be identical.
         * 
         * @return builder
         * 
         */
        public Builder nodes(List<DatabaseNodeArgs> nodes) {
            return nodes(Output.of(nodes));
        }

        /**
         * @param nodes List of nodes object.
         * Multi region cluster are not yet available, all node should be identical.
         * 
         * @return builder
         * 
         */
        public Builder nodes(DatabaseNodeArgs... nodes) {
            return nodes(List.of(nodes));
        }

        /**
         * @param opensearchAclsEnabled Defines whether the ACLs are enabled on an OpenSearch cluster
         * 
         * @return builder
         * 
         */
        public Builder opensearchAclsEnabled(@Nullable Output<Boolean> opensearchAclsEnabled) {
            $.opensearchAclsEnabled = opensearchAclsEnabled;
            return this;
        }

        /**
         * @param opensearchAclsEnabled Defines whether the ACLs are enabled on an OpenSearch cluster
         * 
         * @return builder
         * 
         */
        public Builder opensearchAclsEnabled(Boolean opensearchAclsEnabled) {
            return opensearchAclsEnabled(Output.of(opensearchAclsEnabled));
        }

        /**
         * @param plan Plan of the cluster.
         * Enum: &#34;essential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
         * 
         * @return builder
         * 
         */
        public Builder plan(Output<String> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Plan of the cluster.
         * Enum: &#34;essential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
         * 
         * @return builder
         * 
         */
        public Builder plan(String plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param version The version of the engine in which the service should be deployed
         * 
         * @return builder
         * 
         */
        public Builder version(Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of the engine in which the service should be deployed
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public DatabaseArgs build() {
            $.engine = Objects.requireNonNull($.engine, "expected parameter 'engine' to be non-null");
            $.flavor = Objects.requireNonNull($.flavor, "expected parameter 'flavor' to be non-null");
            $.nodes = Objects.requireNonNull($.nodes, "expected parameter 'nodes' to be non-null");
            $.plan = Objects.requireNonNull($.plan, "expected parameter 'plan' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            $.version = Objects.requireNonNull($.version, "expected parameter 'version' to be non-null");
            return $;
        }
    }

}
