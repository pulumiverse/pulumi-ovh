// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.inputs.CloudProjectKubePrivateNetworkConfigurationArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudProjectKubeState extends com.pulumi.resources.ResourceArgs {

    public static final CloudProjectKubeState Empty = new CloudProjectKubeState();

    /**
     * True if control-plane is up to date.
     * 
     */
    @Import(name="controlPlaneIsUpToDate")
    private @Nullable Output<Boolean> controlPlaneIsUpToDate;

    /**
     * @return True if control-plane is up to date.
     * 
     */
    public Optional<Output<Boolean>> controlPlaneIsUpToDate() {
        return Optional.ofNullable(this.controlPlaneIsUpToDate);
    }

    /**
     * True if all nodes and control-plane are up to date.
     * 
     */
    @Import(name="isUpToDate")
    private @Nullable Output<Boolean> isUpToDate;

    /**
     * @return True if all nodes and control-plane are up to date.
     * 
     */
    public Optional<Output<Boolean>> isUpToDate() {
        return Optional.ofNullable(this.isUpToDate);
    }

    /**
     * The kubeconfig file. Use this file to connect to your kubernetes cluster.
     * 
     */
    @Import(name="kubeconfig")
    private @Nullable Output<String> kubeconfig;

    /**
     * @return The kubeconfig file. Use this file to connect to your kubernetes cluster.
     * 
     */
    public Optional<Output<String>> kubeconfig() {
        return Optional.ofNullable(this.kubeconfig);
    }

    /**
     * The name of the kubernetes cluster.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the kubernetes cluster.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Kubernetes versions available for upgrade.
     * 
     */
    @Import(name="nextUpgradeVersions")
    private @Nullable Output<List<String>> nextUpgradeVersions;

    /**
     * @return Kubernetes versions available for upgrade.
     * 
     */
    public Optional<Output<List<String>>> nextUpgradeVersions() {
        return Optional.ofNullable(this.nextUpgradeVersions);
    }

    /**
     * Cluster nodes URL.
     * 
     */
    @Import(name="nodesUrl")
    private @Nullable Output<String> nodesUrl;

    /**
     * @return Cluster nodes URL.
     * 
     */
    public Optional<Output<String>> nodesUrl() {
        return Optional.ofNullable(this.nodesUrl);
    }

    /**
     * The private network configuration
     * * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
     * * private_network_routing_as_default - Defines whether routing should default to using the nodes&#39; private interface, instead of their public interface. Default is false.
     * 
     */
    @Import(name="privateNetworkConfiguration")
    private @Nullable Output<CloudProjectKubePrivateNetworkConfigurationArgs> privateNetworkConfiguration;

    /**
     * @return The private network configuration
     * * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
     * * private_network_routing_as_default - Defines whether routing should default to using the nodes&#39; private interface, instead of their public interface. Default is false.
     * 
     */
    public Optional<Output<CloudProjectKubePrivateNetworkConfigurationArgs>> privateNetworkConfiguration() {
        return Optional.ofNullable(this.privateNetworkConfiguration);
    }

    /**
     * OpenStack private network ID to use.
     * Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
     * 
     */
    @Import(name="privateNetworkId")
    private @Nullable Output<String> privateNetworkId;

    /**
     * @return OpenStack private network ID to use.
     * Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
     * 
     */
    public Optional<Output<String>> privateNetworkId() {
        return Optional.ofNullable(this.privateNetworkId);
    }

    /**
     * a valid OVH public cloud region ID in which the kubernetes
     * cluster will be available. Ex.: &#34;GRA1&#34;. Defaults to all public cloud regions.
     * Changing this value recreates the resource.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return a valid OVH public cloud region ID in which the kubernetes
     * cluster will be available. Ex.: &#34;GRA1&#34;. Defaults to all public cloud regions.
     * Changing this value recreates the resource.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Cluster status. Should be normally set to &#39;READY&#39;.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Cluster status. Should be normally set to &#39;READY&#39;.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]&#39;.
     * 
     */
    @Import(name="updatePolicy")
    private @Nullable Output<String> updatePolicy;

    /**
     * @return Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]&#39;.
     * 
     */
    public Optional<Output<String>> updatePolicy() {
        return Optional.ofNullable(this.updatePolicy);
    }

    /**
     * Management URL of your cluster.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return Management URL of your cluster.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * kubernetes version to use.
     * Changing this value updates the resource. Defaults to latest available.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return kubernetes version to use.
     * Changing this value updates the resource. Defaults to latest available.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private CloudProjectKubeState() {}

    private CloudProjectKubeState(CloudProjectKubeState $) {
        this.controlPlaneIsUpToDate = $.controlPlaneIsUpToDate;
        this.isUpToDate = $.isUpToDate;
        this.kubeconfig = $.kubeconfig;
        this.name = $.name;
        this.nextUpgradeVersions = $.nextUpgradeVersions;
        this.nodesUrl = $.nodesUrl;
        this.privateNetworkConfiguration = $.privateNetworkConfiguration;
        this.privateNetworkId = $.privateNetworkId;
        this.region = $.region;
        this.serviceName = $.serviceName;
        this.status = $.status;
        this.updatePolicy = $.updatePolicy;
        this.url = $.url;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudProjectKubeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudProjectKubeState $;

        public Builder() {
            $ = new CloudProjectKubeState();
        }

        public Builder(CloudProjectKubeState defaults) {
            $ = new CloudProjectKubeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param controlPlaneIsUpToDate True if control-plane is up to date.
         * 
         * @return builder
         * 
         */
        public Builder controlPlaneIsUpToDate(@Nullable Output<Boolean> controlPlaneIsUpToDate) {
            $.controlPlaneIsUpToDate = controlPlaneIsUpToDate;
            return this;
        }

        /**
         * @param controlPlaneIsUpToDate True if control-plane is up to date.
         * 
         * @return builder
         * 
         */
        public Builder controlPlaneIsUpToDate(Boolean controlPlaneIsUpToDate) {
            return controlPlaneIsUpToDate(Output.of(controlPlaneIsUpToDate));
        }

        /**
         * @param isUpToDate True if all nodes and control-plane are up to date.
         * 
         * @return builder
         * 
         */
        public Builder isUpToDate(@Nullable Output<Boolean> isUpToDate) {
            $.isUpToDate = isUpToDate;
            return this;
        }

        /**
         * @param isUpToDate True if all nodes and control-plane are up to date.
         * 
         * @return builder
         * 
         */
        public Builder isUpToDate(Boolean isUpToDate) {
            return isUpToDate(Output.of(isUpToDate));
        }

        /**
         * @param kubeconfig The kubeconfig file. Use this file to connect to your kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeconfig(@Nullable Output<String> kubeconfig) {
            $.kubeconfig = kubeconfig;
            return this;
        }

        /**
         * @param kubeconfig The kubeconfig file. Use this file to connect to your kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeconfig(String kubeconfig) {
            return kubeconfig(Output.of(kubeconfig));
        }

        /**
         * @param name The name of the kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nextUpgradeVersions Kubernetes versions available for upgrade.
         * 
         * @return builder
         * 
         */
        public Builder nextUpgradeVersions(@Nullable Output<List<String>> nextUpgradeVersions) {
            $.nextUpgradeVersions = nextUpgradeVersions;
            return this;
        }

        /**
         * @param nextUpgradeVersions Kubernetes versions available for upgrade.
         * 
         * @return builder
         * 
         */
        public Builder nextUpgradeVersions(List<String> nextUpgradeVersions) {
            return nextUpgradeVersions(Output.of(nextUpgradeVersions));
        }

        /**
         * @param nextUpgradeVersions Kubernetes versions available for upgrade.
         * 
         * @return builder
         * 
         */
        public Builder nextUpgradeVersions(String... nextUpgradeVersions) {
            return nextUpgradeVersions(List.of(nextUpgradeVersions));
        }

        /**
         * @param nodesUrl Cluster nodes URL.
         * 
         * @return builder
         * 
         */
        public Builder nodesUrl(@Nullable Output<String> nodesUrl) {
            $.nodesUrl = nodesUrl;
            return this;
        }

        /**
         * @param nodesUrl Cluster nodes URL.
         * 
         * @return builder
         * 
         */
        public Builder nodesUrl(String nodesUrl) {
            return nodesUrl(Output.of(nodesUrl));
        }

        /**
         * @param privateNetworkConfiguration The private network configuration
         * * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
         * * private_network_routing_as_default - Defines whether routing should default to using the nodes&#39; private interface, instead of their public interface. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkConfiguration(@Nullable Output<CloudProjectKubePrivateNetworkConfigurationArgs> privateNetworkConfiguration) {
            $.privateNetworkConfiguration = privateNetworkConfiguration;
            return this;
        }

        /**
         * @param privateNetworkConfiguration The private network configuration
         * * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
         * * private_network_routing_as_default - Defines whether routing should default to using the nodes&#39; private interface, instead of their public interface. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkConfiguration(CloudProjectKubePrivateNetworkConfigurationArgs privateNetworkConfiguration) {
            return privateNetworkConfiguration(Output.of(privateNetworkConfiguration));
        }

        /**
         * @param privateNetworkId OpenStack private network ID to use.
         * Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(@Nullable Output<String> privateNetworkId) {
            $.privateNetworkId = privateNetworkId;
            return this;
        }

        /**
         * @param privateNetworkId OpenStack private network ID to use.
         * Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(String privateNetworkId) {
            return privateNetworkId(Output.of(privateNetworkId));
        }

        /**
         * @param region a valid OVH public cloud region ID in which the kubernetes
         * cluster will be available. Ex.: &#34;GRA1&#34;. Defaults to all public cloud regions.
         * Changing this value recreates the resource.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region a valid OVH public cloud region ID in which the kubernetes
         * cluster will be available. Ex.: &#34;GRA1&#34;. Defaults to all public cloud regions.
         * Changing this value recreates the resource.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param status Cluster status. Should be normally set to &#39;READY&#39;.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Cluster status. Should be normally set to &#39;READY&#39;.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param updatePolicy Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]&#39;.
         * 
         * @return builder
         * 
         */
        public Builder updatePolicy(@Nullable Output<String> updatePolicy) {
            $.updatePolicy = updatePolicy;
            return this;
        }

        /**
         * @param updatePolicy Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]&#39;.
         * 
         * @return builder
         * 
         */
        public Builder updatePolicy(String updatePolicy) {
            return updatePolicy(Output.of(updatePolicy));
        }

        /**
         * @param url Management URL of your cluster.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Management URL of your cluster.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param version kubernetes version to use.
         * Changing this value updates the resource. Defaults to latest available.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version kubernetes version to use.
         * Changing this value updates the resource. Defaults to latest available.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public CloudProjectKubeState build() {
            return $;
        }
    }

}
