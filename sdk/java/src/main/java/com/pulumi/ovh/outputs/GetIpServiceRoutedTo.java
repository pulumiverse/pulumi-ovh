// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIpServiceRoutedTo {
    /**
     * @return The service name
     * 
     */
    private String serviceName;

    private GetIpServiceRoutedTo() {}
    /**
     * @return The service name
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpServiceRoutedTo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String serviceName;
        public Builder() {}
        public Builder(GetIpServiceRoutedTo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        public GetIpServiceRoutedTo build() {
            final var o = new GetIpServiceRoutedTo();
            o.serviceName = serviceName;
            return o;
        }
    }
}
