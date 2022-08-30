// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetCloudProjectContainerRegistriesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudProjectContainerRegistriesPlainArgs Empty = new GetCloudProjectContainerRegistriesPlainArgs();

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

    private GetCloudProjectContainerRegistriesPlainArgs() {}

    private GetCloudProjectContainerRegistriesPlainArgs(GetCloudProjectContainerRegistriesPlainArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudProjectContainerRegistriesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudProjectContainerRegistriesPlainArgs $;

        public Builder() {
            $ = new GetCloudProjectContainerRegistriesPlainArgs();
        }

        public Builder(GetCloudProjectContainerRegistriesPlainArgs defaults) {
            $ = new GetCloudProjectContainerRegistriesPlainArgs(Objects.requireNonNull(defaults));
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

        public GetCloudProjectContainerRegistriesPlainArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
