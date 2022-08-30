// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.outputs.GetMeInstallationTemplatePartitionSchemeHardwareRaid;
import com.pulumi.ovh.outputs.GetMeInstallationTemplatePartitionSchemePartition;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMeInstallationTemplatePartitionScheme {
    private List<GetMeInstallationTemplatePartitionSchemeHardwareRaid> hardwareRaids;
    private String name;
    private List<GetMeInstallationTemplatePartitionSchemePartition> partitions;
    private Integer priority;

    private GetMeInstallationTemplatePartitionScheme() {}
    public List<GetMeInstallationTemplatePartitionSchemeHardwareRaid> hardwareRaids() {
        return this.hardwareRaids;
    }
    public String name() {
        return this.name;
    }
    public List<GetMeInstallationTemplatePartitionSchemePartition> partitions() {
        return this.partitions;
    }
    public Integer priority() {
        return this.priority;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMeInstallationTemplatePartitionScheme defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetMeInstallationTemplatePartitionSchemeHardwareRaid> hardwareRaids;
        private String name;
        private List<GetMeInstallationTemplatePartitionSchemePartition> partitions;
        private Integer priority;
        public Builder() {}
        public Builder(GetMeInstallationTemplatePartitionScheme defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hardwareRaids = defaults.hardwareRaids;
    	      this.name = defaults.name;
    	      this.partitions = defaults.partitions;
    	      this.priority = defaults.priority;
        }

        @CustomType.Setter
        public Builder hardwareRaids(List<GetMeInstallationTemplatePartitionSchemeHardwareRaid> hardwareRaids) {
            this.hardwareRaids = Objects.requireNonNull(hardwareRaids);
            return this;
        }
        public Builder hardwareRaids(GetMeInstallationTemplatePartitionSchemeHardwareRaid... hardwareRaids) {
            return hardwareRaids(List.of(hardwareRaids));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder partitions(List<GetMeInstallationTemplatePartitionSchemePartition> partitions) {
            this.partitions = Objects.requireNonNull(partitions);
            return this;
        }
        public Builder partitions(GetMeInstallationTemplatePartitionSchemePartition... partitions) {
            return partitions(List.of(partitions));
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        public GetMeInstallationTemplatePartitionScheme build() {
            final var o = new GetMeInstallationTemplatePartitionScheme();
            o.hardwareRaids = hardwareRaids;
            o.name = name;
            o.partitions = partitions;
            o.priority = priority;
            return o;
        }
    }
}
