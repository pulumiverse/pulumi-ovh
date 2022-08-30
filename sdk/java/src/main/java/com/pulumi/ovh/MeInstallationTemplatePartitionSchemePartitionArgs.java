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


public final class MeInstallationTemplatePartitionSchemePartitionArgs extends com.pulumi.resources.ResourceArgs {

    public static final MeInstallationTemplatePartitionSchemePartitionArgs Empty = new MeInstallationTemplatePartitionSchemePartitionArgs();

    /**
     * Partition filesystem
     * 
     */
    @Import(name="filesystem", required=true)
    private Output<String> filesystem;

    /**
     * @return Partition filesystem
     * 
     */
    public Output<String> filesystem() {
        return this.filesystem;
    }

    /**
     * partition mount point
     * 
     */
    @Import(name="mountpoint", required=true)
    private Output<String> mountpoint;

    /**
     * @return partition mount point
     * 
     */
    public Output<String> mountpoint() {
        return this.mountpoint;
    }

    /**
     * step or order. specifies the creation order of the partition on the disk
     * 
     */
    @Import(name="order", required=true)
    private Output<Integer> order;

    /**
     * @return step or order. specifies the creation order of the partition on the disk
     * 
     */
    public Output<Integer> order() {
        return this.order;
    }

    /**
     * raid partition type
     * 
     */
    @Import(name="raid")
    private @Nullable Output<String> raid;

    /**
     * @return raid partition type
     * 
     */
    public Optional<Output<String>> raid() {
        return Optional.ofNullable(this.raid);
    }

    /**
     * name of this partitioning scheme
     * 
     */
    @Import(name="schemeName", required=true)
    private Output<String> schemeName;

    /**
     * @return name of this partitioning scheme
     * 
     */
    public Output<String> schemeName() {
        return this.schemeName;
    }

    /**
     * size of partition in MB, 0 =&gt; rest of the space
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return size of partition in MB, 0 =&gt; rest of the space
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    /**
     * Template name
     * 
     */
    @Import(name="templateName", required=true)
    private Output<String> templateName;

    /**
     * @return Template name
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }

    /**
     * partition type
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return partition type
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * The volume name needed for proxmox distribution
     * 
     */
    @Import(name="volumeName")
    private @Nullable Output<String> volumeName;

    /**
     * @return The volume name needed for proxmox distribution
     * 
     */
    public Optional<Output<String>> volumeName() {
        return Optional.ofNullable(this.volumeName);
    }

    private MeInstallationTemplatePartitionSchemePartitionArgs() {}

    private MeInstallationTemplatePartitionSchemePartitionArgs(MeInstallationTemplatePartitionSchemePartitionArgs $) {
        this.filesystem = $.filesystem;
        this.mountpoint = $.mountpoint;
        this.order = $.order;
        this.raid = $.raid;
        this.schemeName = $.schemeName;
        this.size = $.size;
        this.templateName = $.templateName;
        this.type = $.type;
        this.volumeName = $.volumeName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MeInstallationTemplatePartitionSchemePartitionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MeInstallationTemplatePartitionSchemePartitionArgs $;

        public Builder() {
            $ = new MeInstallationTemplatePartitionSchemePartitionArgs();
        }

        public Builder(MeInstallationTemplatePartitionSchemePartitionArgs defaults) {
            $ = new MeInstallationTemplatePartitionSchemePartitionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filesystem Partition filesystem
         * 
         * @return builder
         * 
         */
        public Builder filesystem(Output<String> filesystem) {
            $.filesystem = filesystem;
            return this;
        }

        /**
         * @param filesystem Partition filesystem
         * 
         * @return builder
         * 
         */
        public Builder filesystem(String filesystem) {
            return filesystem(Output.of(filesystem));
        }

        /**
         * @param mountpoint partition mount point
         * 
         * @return builder
         * 
         */
        public Builder mountpoint(Output<String> mountpoint) {
            $.mountpoint = mountpoint;
            return this;
        }

        /**
         * @param mountpoint partition mount point
         * 
         * @return builder
         * 
         */
        public Builder mountpoint(String mountpoint) {
            return mountpoint(Output.of(mountpoint));
        }

        /**
         * @param order step or order. specifies the creation order of the partition on the disk
         * 
         * @return builder
         * 
         */
        public Builder order(Output<Integer> order) {
            $.order = order;
            return this;
        }

        /**
         * @param order step or order. specifies the creation order of the partition on the disk
         * 
         * @return builder
         * 
         */
        public Builder order(Integer order) {
            return order(Output.of(order));
        }

        /**
         * @param raid raid partition type
         * 
         * @return builder
         * 
         */
        public Builder raid(@Nullable Output<String> raid) {
            $.raid = raid;
            return this;
        }

        /**
         * @param raid raid partition type
         * 
         * @return builder
         * 
         */
        public Builder raid(String raid) {
            return raid(Output.of(raid));
        }

        /**
         * @param schemeName name of this partitioning scheme
         * 
         * @return builder
         * 
         */
        public Builder schemeName(Output<String> schemeName) {
            $.schemeName = schemeName;
            return this;
        }

        /**
         * @param schemeName name of this partitioning scheme
         * 
         * @return builder
         * 
         */
        public Builder schemeName(String schemeName) {
            return schemeName(Output.of(schemeName));
        }

        /**
         * @param size size of partition in MB, 0 =&gt; rest of the space
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size size of partition in MB, 0 =&gt; rest of the space
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param templateName Template name
         * 
         * @return builder
         * 
         */
        public Builder templateName(Output<String> templateName) {
            $.templateName = templateName;
            return this;
        }

        /**
         * @param templateName Template name
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            return templateName(Output.of(templateName));
        }

        /**
         * @param type partition type
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type partition type
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param volumeName The volume name needed for proxmox distribution
         * 
         * @return builder
         * 
         */
        public Builder volumeName(@Nullable Output<String> volumeName) {
            $.volumeName = volumeName;
            return this;
        }

        /**
         * @param volumeName The volume name needed for proxmox distribution
         * 
         * @return builder
         * 
         */
        public Builder volumeName(String volumeName) {
            return volumeName(Output.of(volumeName));
        }

        public MeInstallationTemplatePartitionSchemePartitionArgs build() {
            $.filesystem = Objects.requireNonNull($.filesystem, "expected parameter 'filesystem' to be non-null");
            $.mountpoint = Objects.requireNonNull($.mountpoint, "expected parameter 'mountpoint' to be non-null");
            $.order = Objects.requireNonNull($.order, "expected parameter 'order' to be non-null");
            $.schemeName = Objects.requireNonNull($.schemeName, "expected parameter 'schemeName' to be non-null");
            $.size = Objects.requireNonNull($.size, "expected parameter 'size' to be non-null");
            $.templateName = Objects.requireNonNull($.templateName, "expected parameter 'templateName' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
