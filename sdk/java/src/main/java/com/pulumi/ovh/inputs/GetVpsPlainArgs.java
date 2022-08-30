// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetVpsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpsPlainArgs Empty = new GetVpsPlainArgs();

    /**
     * The service_name of your dedicated server.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The service_name of your dedicated server.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetVpsPlainArgs() {}

    private GetVpsPlainArgs(GetVpsPlainArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpsPlainArgs $;

        public Builder() {
            $ = new GetVpsPlainArgs();
        }

        public Builder(GetVpsPlainArgs defaults) {
            $ = new GetVpsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The service_name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetVpsPlainArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
