// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.outputs.GetMeInstallationTemplateCustomization;
import com.pulumi.ovh.outputs.GetMeInstallationTemplatePartitionScheme;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMeInstallationTemplateResult {
    private List<String> availableLanguages;
    private Boolean beta;
    private Integer bitFormat;
    private String category;
    private List<GetMeInstallationTemplateCustomization> customizations;
    private String defaultLanguage;
    private Boolean deprecated;
    private String description;
    private String distribution;
    private String family;
    private List<String> filesystems;
    private Boolean hardRaidConfiguration;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String lastModification;
    private Boolean lvmReady;
    private List<GetMeInstallationTemplatePartitionScheme> partitionSchemes;
    private Boolean supportsDistributionKernel;
    private Boolean supportsGptLabel;
    private Boolean supportsRtm;
    private Boolean supportsSqlServer;
    private String supportsUefi;
    private String templateName;

    private GetMeInstallationTemplateResult() {}
    public List<String> availableLanguages() {
        return this.availableLanguages;
    }
    public Boolean beta() {
        return this.beta;
    }
    public Integer bitFormat() {
        return this.bitFormat;
    }
    public String category() {
        return this.category;
    }
    public List<GetMeInstallationTemplateCustomization> customizations() {
        return this.customizations;
    }
    public String defaultLanguage() {
        return this.defaultLanguage;
    }
    public Boolean deprecated() {
        return this.deprecated;
    }
    public String description() {
        return this.description;
    }
    public String distribution() {
        return this.distribution;
    }
    public String family() {
        return this.family;
    }
    public List<String> filesystems() {
        return this.filesystems;
    }
    public Boolean hardRaidConfiguration() {
        return this.hardRaidConfiguration;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String lastModification() {
        return this.lastModification;
    }
    public Boolean lvmReady() {
        return this.lvmReady;
    }
    public List<GetMeInstallationTemplatePartitionScheme> partitionSchemes() {
        return this.partitionSchemes;
    }
    public Boolean supportsDistributionKernel() {
        return this.supportsDistributionKernel;
    }
    public Boolean supportsGptLabel() {
        return this.supportsGptLabel;
    }
    public Boolean supportsRtm() {
        return this.supportsRtm;
    }
    public Boolean supportsSqlServer() {
        return this.supportsSqlServer;
    }
    public String supportsUefi() {
        return this.supportsUefi;
    }
    public String templateName() {
        return this.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMeInstallationTemplateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> availableLanguages;
        private Boolean beta;
        private Integer bitFormat;
        private String category;
        private List<GetMeInstallationTemplateCustomization> customizations;
        private String defaultLanguage;
        private Boolean deprecated;
        private String description;
        private String distribution;
        private String family;
        private List<String> filesystems;
        private Boolean hardRaidConfiguration;
        private String id;
        private String lastModification;
        private Boolean lvmReady;
        private List<GetMeInstallationTemplatePartitionScheme> partitionSchemes;
        private Boolean supportsDistributionKernel;
        private Boolean supportsGptLabel;
        private Boolean supportsRtm;
        private Boolean supportsSqlServer;
        private String supportsUefi;
        private String templateName;
        public Builder() {}
        public Builder(GetMeInstallationTemplateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availableLanguages = defaults.availableLanguages;
    	      this.beta = defaults.beta;
    	      this.bitFormat = defaults.bitFormat;
    	      this.category = defaults.category;
    	      this.customizations = defaults.customizations;
    	      this.defaultLanguage = defaults.defaultLanguage;
    	      this.deprecated = defaults.deprecated;
    	      this.description = defaults.description;
    	      this.distribution = defaults.distribution;
    	      this.family = defaults.family;
    	      this.filesystems = defaults.filesystems;
    	      this.hardRaidConfiguration = defaults.hardRaidConfiguration;
    	      this.id = defaults.id;
    	      this.lastModification = defaults.lastModification;
    	      this.lvmReady = defaults.lvmReady;
    	      this.partitionSchemes = defaults.partitionSchemes;
    	      this.supportsDistributionKernel = defaults.supportsDistributionKernel;
    	      this.supportsGptLabel = defaults.supportsGptLabel;
    	      this.supportsRtm = defaults.supportsRtm;
    	      this.supportsSqlServer = defaults.supportsSqlServer;
    	      this.supportsUefi = defaults.supportsUefi;
    	      this.templateName = defaults.templateName;
        }

        @CustomType.Setter
        public Builder availableLanguages(List<String> availableLanguages) {
            this.availableLanguages = Objects.requireNonNull(availableLanguages);
            return this;
        }
        public Builder availableLanguages(String... availableLanguages) {
            return availableLanguages(List.of(availableLanguages));
        }
        @CustomType.Setter
        public Builder beta(Boolean beta) {
            this.beta = Objects.requireNonNull(beta);
            return this;
        }
        @CustomType.Setter
        public Builder bitFormat(Integer bitFormat) {
            this.bitFormat = Objects.requireNonNull(bitFormat);
            return this;
        }
        @CustomType.Setter
        public Builder category(String category) {
            this.category = Objects.requireNonNull(category);
            return this;
        }
        @CustomType.Setter
        public Builder customizations(List<GetMeInstallationTemplateCustomization> customizations) {
            this.customizations = Objects.requireNonNull(customizations);
            return this;
        }
        public Builder customizations(GetMeInstallationTemplateCustomization... customizations) {
            return customizations(List.of(customizations));
        }
        @CustomType.Setter
        public Builder defaultLanguage(String defaultLanguage) {
            this.defaultLanguage = Objects.requireNonNull(defaultLanguage);
            return this;
        }
        @CustomType.Setter
        public Builder deprecated(Boolean deprecated) {
            this.deprecated = Objects.requireNonNull(deprecated);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder distribution(String distribution) {
            this.distribution = Objects.requireNonNull(distribution);
            return this;
        }
        @CustomType.Setter
        public Builder family(String family) {
            this.family = Objects.requireNonNull(family);
            return this;
        }
        @CustomType.Setter
        public Builder filesystems(List<String> filesystems) {
            this.filesystems = Objects.requireNonNull(filesystems);
            return this;
        }
        public Builder filesystems(String... filesystems) {
            return filesystems(List.of(filesystems));
        }
        @CustomType.Setter
        public Builder hardRaidConfiguration(Boolean hardRaidConfiguration) {
            this.hardRaidConfiguration = Objects.requireNonNull(hardRaidConfiguration);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lastModification(String lastModification) {
            this.lastModification = Objects.requireNonNull(lastModification);
            return this;
        }
        @CustomType.Setter
        public Builder lvmReady(Boolean lvmReady) {
            this.lvmReady = Objects.requireNonNull(lvmReady);
            return this;
        }
        @CustomType.Setter
        public Builder partitionSchemes(List<GetMeInstallationTemplatePartitionScheme> partitionSchemes) {
            this.partitionSchemes = Objects.requireNonNull(partitionSchemes);
            return this;
        }
        public Builder partitionSchemes(GetMeInstallationTemplatePartitionScheme... partitionSchemes) {
            return partitionSchemes(List.of(partitionSchemes));
        }
        @CustomType.Setter
        public Builder supportsDistributionKernel(Boolean supportsDistributionKernel) {
            this.supportsDistributionKernel = Objects.requireNonNull(supportsDistributionKernel);
            return this;
        }
        @CustomType.Setter
        public Builder supportsGptLabel(Boolean supportsGptLabel) {
            this.supportsGptLabel = Objects.requireNonNull(supportsGptLabel);
            return this;
        }
        @CustomType.Setter
        public Builder supportsRtm(Boolean supportsRtm) {
            this.supportsRtm = Objects.requireNonNull(supportsRtm);
            return this;
        }
        @CustomType.Setter
        public Builder supportsSqlServer(Boolean supportsSqlServer) {
            this.supportsSqlServer = Objects.requireNonNull(supportsSqlServer);
            return this;
        }
        @CustomType.Setter
        public Builder supportsUefi(String supportsUefi) {
            this.supportsUefi = Objects.requireNonNull(supportsUefi);
            return this;
        }
        @CustomType.Setter
        public Builder templateName(String templateName) {
            this.templateName = Objects.requireNonNull(templateName);
            return this;
        }
        public GetMeInstallationTemplateResult build() {
            final var o = new GetMeInstallationTemplateResult();
            o.availableLanguages = availableLanguages;
            o.beta = beta;
            o.bitFormat = bitFormat;
            o.category = category;
            o.customizations = customizations;
            o.defaultLanguage = defaultLanguage;
            o.deprecated = deprecated;
            o.description = description;
            o.distribution = distribution;
            o.family = family;
            o.filesystems = filesystems;
            o.hardRaidConfiguration = hardRaidConfiguration;
            o.id = id;
            o.lastModification = lastModification;
            o.lvmReady = lvmReady;
            o.partitionSchemes = partitionSchemes;
            o.supportsDistributionKernel = supportsDistributionKernel;
            o.supportsGptLabel = supportsGptLabel;
            o.supportsRtm = supportsRtm;
            o.supportsSqlServer = supportsSqlServer;
            o.supportsUefi = supportsUefi;
            o.templateName = templateName;
            return o;
        }
    }
}
