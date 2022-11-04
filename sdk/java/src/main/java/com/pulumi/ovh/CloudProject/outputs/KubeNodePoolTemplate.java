// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.CloudProject.outputs.KubeNodePoolTemplateMetadata;
import com.pulumi.ovh.CloudProject.outputs.KubeNodePoolTemplateSpec;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class KubeNodePoolTemplate {
    private @Nullable KubeNodePoolTemplateMetadata metadata;
    private @Nullable KubeNodePoolTemplateSpec spec;

    private KubeNodePoolTemplate() {}
    public Optional<KubeNodePoolTemplateMetadata> metadata() {
        return Optional.ofNullable(this.metadata);
    }
    public Optional<KubeNodePoolTemplateSpec> spec() {
        return Optional.ofNullable(this.spec);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubeNodePoolTemplate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable KubeNodePoolTemplateMetadata metadata;
        private @Nullable KubeNodePoolTemplateSpec spec;
        public Builder() {}
        public Builder(KubeNodePoolTemplate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.metadata = defaults.metadata;
    	      this.spec = defaults.spec;
        }

        @CustomType.Setter
        public Builder metadata(@Nullable KubeNodePoolTemplateMetadata metadata) {
            this.metadata = metadata;
            return this;
        }
        @CustomType.Setter
        public Builder spec(@Nullable KubeNodePoolTemplateSpec spec) {
            this.spec = spec;
            return this;
        }
        public KubeNodePoolTemplate build() {
            final var o = new KubeNodePoolTemplate();
            o.metadata = metadata;
            o.spec = spec;
            return o;
        }
    }
}