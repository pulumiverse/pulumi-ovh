// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetOpenSearchPatternPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOpenSearchPatternPlainArgs Empty = new GetOpenSearchPatternPlainArgs();

    /**
     * Cluster ID
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return Cluster ID
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * Pattern ID.
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return Pattern ID.
     * 
     */
    public String id() {
        return this.id;
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetOpenSearchPatternPlainArgs() {}

    private GetOpenSearchPatternPlainArgs(GetOpenSearchPatternPlainArgs $) {
        this.clusterId = $.clusterId;
        this.id = $.id;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOpenSearchPatternPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOpenSearchPatternPlainArgs $;

        public Builder() {
            $ = new GetOpenSearchPatternPlainArgs();
        }

        public Builder(GetOpenSearchPatternPlainArgs defaults) {
            $ = new GetOpenSearchPatternPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param id Pattern ID.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
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
            $.serviceName = serviceName;
            return this;
        }

        public GetOpenSearchPatternPlainArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
