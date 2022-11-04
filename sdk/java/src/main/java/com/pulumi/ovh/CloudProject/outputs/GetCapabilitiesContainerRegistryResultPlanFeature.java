// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;

@CustomType
public final class GetCapabilitiesContainerRegistryResultPlanFeature {
    /**
     * @return Vulnerability scanning
     * 
     */
    private Boolean vulnerability;

    private GetCapabilitiesContainerRegistryResultPlanFeature() {}
    /**
     * @return Vulnerability scanning
     * 
     */
    public Boolean vulnerability() {
        return this.vulnerability;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCapabilitiesContainerRegistryResultPlanFeature defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean vulnerability;
        public Builder() {}
        public Builder(GetCapabilitiesContainerRegistryResultPlanFeature defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.vulnerability = defaults.vulnerability;
        }

        @CustomType.Setter
        public Builder vulnerability(Boolean vulnerability) {
            this.vulnerability = Objects.requireNonNull(vulnerability);
            return this;
        }
        public GetCapabilitiesContainerRegistryResultPlanFeature build() {
            final var o = new GetCapabilitiesContainerRegistryResultPlanFeature();
            o.vulnerability = vulnerability;
            return o;
        }
    }
}