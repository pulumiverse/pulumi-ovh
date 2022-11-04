// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpRestrictionState extends com.pulumi.resources.ResourceArgs {

    public static final IpRestrictionState Empty = new IpRestrictionState();

    /**
     * Cluster ID.
     * 
     */
    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * Description of the IP restriction.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the IP restriction.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    @Import(name="engine")
    private @Nullable Output<String> engine;

    /**
     * @return The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * Authorized IP.
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return Authorized IP.
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Current status of the IP restriction.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Current status of the IP restriction.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private IpRestrictionState() {}

    private IpRestrictionState(IpRestrictionState $) {
        this.clusterId = $.clusterId;
        this.description = $.description;
        this.engine = $.engine;
        this.ip = $.ip;
        this.serviceName = $.serviceName;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpRestrictionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpRestrictionState $;

        public Builder() {
            $ = new IpRestrictionState();
        }

        public Builder(IpRestrictionState defaults) {
            $ = new IpRestrictionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param description Description of the IP restriction.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the IP restriction.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param engine The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param ip Authorized IP.
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip Authorized IP.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
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
         * @param status Current status of the IP restriction.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Current status of the IP restriction.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public IpRestrictionState build() {
            return $;
        }
    }

}