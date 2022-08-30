// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMeInstallationTemplatePartitionSchemeHardwareRaid {
    private List<String> disks;
    private String mode;
    private String name;
    private Integer step;

    private GetMeInstallationTemplatePartitionSchemeHardwareRaid() {}
    public List<String> disks() {
        return this.disks;
    }
    public String mode() {
        return this.mode;
    }
    public String name() {
        return this.name;
    }
    public Integer step() {
        return this.step;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMeInstallationTemplatePartitionSchemeHardwareRaid defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> disks;
        private String mode;
        private String name;
        private Integer step;
        public Builder() {}
        public Builder(GetMeInstallationTemplatePartitionSchemeHardwareRaid defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.disks = defaults.disks;
    	      this.mode = defaults.mode;
    	      this.name = defaults.name;
    	      this.step = defaults.step;
        }

        @CustomType.Setter
        public Builder disks(List<String> disks) {
            this.disks = Objects.requireNonNull(disks);
            return this;
        }
        public Builder disks(String... disks) {
            return disks(List.of(disks));
        }
        @CustomType.Setter
        public Builder mode(String mode) {
            this.mode = Objects.requireNonNull(mode);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder step(Integer step) {
            this.step = Objects.requireNonNull(step);
            return this;
        }
        public GetMeInstallationTemplatePartitionSchemeHardwareRaid build() {
            final var o = new GetMeInstallationTemplatePartitionSchemeHardwareRaid();
            o.disks = disks;
            o.mode = mode;
            o.name = name;
            o.step = step;
            return o;
        }
    }
}
