// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MeInstallationTemplatePartitionSchemeArgs extends com.pulumi.resources.ResourceArgs {

    public static final MeInstallationTemplatePartitionSchemeArgs Empty = new MeInstallationTemplatePartitionSchemeArgs();

    /**
     * name of this partitioning scheme
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return name of this partitioning scheme
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
     * among all the compatible partitioning schemes (given the underlying hardware specifications)
     * 
     */
    @Import(name="priority", required=true)
    private Output<Integer> priority;

    /**
     * @return on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
     * among all the compatible partitioning schemes (given the underlying hardware specifications)
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }

    /**
     * This template name
     * 
     */
    @Import(name="templateName", required=true)
    private Output<String> templateName;

    /**
     * @return This template name
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }

    private MeInstallationTemplatePartitionSchemeArgs() {}

    private MeInstallationTemplatePartitionSchemeArgs(MeInstallationTemplatePartitionSchemeArgs $) {
        this.name = $.name;
        this.priority = $.priority;
        this.templateName = $.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MeInstallationTemplatePartitionSchemeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MeInstallationTemplatePartitionSchemeArgs $;

        public Builder() {
            $ = new MeInstallationTemplatePartitionSchemeArgs();
        }

        public Builder(MeInstallationTemplatePartitionSchemeArgs defaults) {
            $ = new MeInstallationTemplatePartitionSchemeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name name of this partitioning scheme
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name name of this partitioning scheme
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param priority on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
         * among all the compatible partitioning schemes (given the underlying hardware specifications)
         * 
         * @return builder
         * 
         */
        public Builder priority(Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
         * among all the compatible partitioning schemes (given the underlying hardware specifications)
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param templateName This template name
         * 
         * @return builder
         * 
         */
        public Builder templateName(Output<String> templateName) {
            $.templateName = templateName;
            return this;
        }

        /**
         * @param templateName This template name
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            return templateName(Output.of(templateName));
        }

        public MeInstallationTemplatePartitionSchemeArgs build() {
            $.priority = Objects.requireNonNull($.priority, "expected parameter 'priority' to be non-null");
            $.templateName = Objects.requireNonNull($.templateName, "expected parameter 'templateName' to be non-null");
            return $;
        }
    }

}
