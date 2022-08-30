// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudProjectFailoverIpAttachPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudProjectFailoverIpAttachPlainArgs Empty = new GetCloudProjectFailoverIpAttachPlainArgs();

    /**
     * The IP block
     * 
     */
    @Import(name="block")
    private @Nullable String block;

    /**
     * @return The IP block
     * 
     */
    public Optional<String> block() {
        return Optional.ofNullable(this.block);
    }

    @Import(name="continentCode")
    private @Nullable String continentCode;

    public Optional<String> continentCode() {
        return Optional.ofNullable(this.continentCode);
    }

    @Import(name="geoLoc")
    private @Nullable String geoLoc;

    public Optional<String> geoLoc() {
        return Optional.ofNullable(this.geoLoc);
    }

    /**
     * The failover ip address to query
     * 
     */
    @Import(name="ip")
    private @Nullable String ip;

    /**
     * @return The failover ip address to query
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }

    @Import(name="routedTo")
    private @Nullable String routedTo;

    public Optional<String> routedTo() {
        return Optional.ofNullable(this.routedTo);
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

    private GetCloudProjectFailoverIpAttachPlainArgs() {}

    private GetCloudProjectFailoverIpAttachPlainArgs(GetCloudProjectFailoverIpAttachPlainArgs $) {
        this.block = $.block;
        this.continentCode = $.continentCode;
        this.geoLoc = $.geoLoc;
        this.ip = $.ip;
        this.routedTo = $.routedTo;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudProjectFailoverIpAttachPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudProjectFailoverIpAttachPlainArgs $;

        public Builder() {
            $ = new GetCloudProjectFailoverIpAttachPlainArgs();
        }

        public Builder(GetCloudProjectFailoverIpAttachPlainArgs defaults) {
            $ = new GetCloudProjectFailoverIpAttachPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param block The IP block
         * 
         * @return builder
         * 
         */
        public Builder block(@Nullable String block) {
            $.block = block;
            return this;
        }

        public Builder continentCode(@Nullable String continentCode) {
            $.continentCode = continentCode;
            return this;
        }

        public Builder geoLoc(@Nullable String geoLoc) {
            $.geoLoc = geoLoc;
            return this;
        }

        /**
         * @param ip The failover ip address to query
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable String ip) {
            $.ip = ip;
            return this;
        }

        public Builder routedTo(@Nullable String routedTo) {
            $.routedTo = routedTo;
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

        public GetCloudProjectFailoverIpAttachPlainArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
