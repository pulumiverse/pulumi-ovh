// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.CloudProject.outputs.GetKubeCustomizationApiserverAdmissionplugins;
import java.util.Objects;

@CustomType
public final class GetKubeCustomizationApiserver {
    private GetKubeCustomizationApiserverAdmissionplugins admissionplugins;

    private GetKubeCustomizationApiserver() {}
    public GetKubeCustomizationApiserverAdmissionplugins admissionplugins() {
        return this.admissionplugins;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubeCustomizationApiserver defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private GetKubeCustomizationApiserverAdmissionplugins admissionplugins;
        public Builder() {}
        public Builder(GetKubeCustomizationApiserver defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.admissionplugins = defaults.admissionplugins;
        }

        @CustomType.Setter
        public Builder admissionplugins(GetKubeCustomizationApiserverAdmissionplugins admissionplugins) {
            this.admissionplugins = Objects.requireNonNull(admissionplugins);
            return this;
        }
        public GetKubeCustomizationApiserver build() {
            final var o = new GetKubeCustomizationApiserver();
            o.admissionplugins = admissionplugins;
            return o;
        }
    }
}
