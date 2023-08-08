// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProject.KubeArgs;
import com.pulumi.ovh.CloudProject.inputs.KubeState;
import com.pulumi.ovh.CloudProject.outputs.KubeCustomization;
import com.pulumi.ovh.CloudProject.outputs.KubeCustomizationApiserver;
import com.pulumi.ovh.CloudProject.outputs.KubeCustomizationKubeProxy;
import com.pulumi.ovh.CloudProject.outputs.KubeKubeconfigAttribute;
import com.pulumi.ovh.CloudProject.outputs.KubePrivateNetworkConfiguration;
import com.pulumi.ovh.Utilities;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * OVHcloud Managed Kubernetes Service clusters can be imported using the `service_name` and the `id` of the cluster, separated by &#34;/&#34; E.g., bash
 * 
 * ```sh
 *  $ pulumi import ovh:CloudProject/kube:Kube my_kube_cluster service_name/kube_id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/kube:Kube")
public class Kube extends com.pulumi.resources.CustomResource {
    /**
     * True if control-plane is up-to-date.
     * 
     */
    @Export(name="controlPlaneIsUpToDate", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> controlPlaneIsUpToDate;

    /**
     * @return True if control-plane is up-to-date.
     * 
     */
    public Output<Boolean> controlPlaneIsUpToDate() {
        return this.controlPlaneIsUpToDate;
    }
    /**
     * Kubernetes API server customization
     * 
     */
    @Export(name="customizationApiservers", refs={List.class,KubeCustomizationApiserver.class}, tree="[0,1]")
    private Output<List<KubeCustomizationApiserver>> customizationApiservers;

    /**
     * @return Kubernetes API server customization
     * 
     */
    public Output<List<KubeCustomizationApiserver>> customizationApiservers() {
        return this.customizationApiservers;
    }
    /**
     * Kubernetes kube-proxy customization
     * 
     */
    @Export(name="customizationKubeProxy", refs={KubeCustomizationKubeProxy.class}, tree="[0]")
    private Output</* @Nullable */ KubeCustomizationKubeProxy> customizationKubeProxy;

    /**
     * @return Kubernetes kube-proxy customization
     * 
     */
    public Output<Optional<KubeCustomizationKubeProxy>> customizationKubeProxy() {
        return Codegen.optional(this.customizationKubeProxy);
    }
    /**
     * **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    @Export(name="customizations", refs={List.class,KubeCustomization.class}, tree="[0,1]")
    private Output<List<KubeCustomization>> customizations;

    /**
     * @return **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
     * 
     */
    public Output<List<KubeCustomization>> customizations() {
        return this.customizations;
    }
    /**
     * True if all nodes and control-plane are up-to-date.
     * 
     */
    @Export(name="isUpToDate", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isUpToDate;

    /**
     * @return True if all nodes and control-plane are up-to-date.
     * 
     */
    public Output<Boolean> isUpToDate() {
        return this.isUpToDate;
    }
    /**
     * Selected mode for kube-proxy. **Changing this value recreates the resource, including ETCD user data.** Defaults to `iptables`.
     * 
     */
    @Export(name="kubeProxyMode", refs={String.class}, tree="[0]")
    private Output<String> kubeProxyMode;

    /**
     * @return Selected mode for kube-proxy. **Changing this value recreates the resource, including ETCD user data.** Defaults to `iptables`.
     * 
     */
    public Output<String> kubeProxyMode() {
        return this.kubeProxyMode;
    }
    /**
     * The kubeconfig file. Use this file to connect to your kubernetes cluster.
     * 
     */
    @Export(name="kubeconfig", refs={String.class}, tree="[0]")
    private Output<String> kubeconfig;

    /**
     * @return The kubeconfig file. Use this file to connect to your kubernetes cluster.
     * 
     */
    public Output<String> kubeconfig() {
        return this.kubeconfig;
    }
    /**
     * The kubeconfig file attributes.
     * 
     */
    @Export(name="kubeconfigAttributes", refs={List.class,KubeKubeconfigAttribute.class}, tree="[0,1]")
    private Output<List<KubeKubeconfigAttribute>> kubeconfigAttributes;

    /**
     * @return The kubeconfig file attributes.
     * 
     */
    public Output<List<KubeKubeconfigAttribute>> kubeconfigAttributes() {
        return this.kubeconfigAttributes;
    }
    /**
     * The name of the kubernetes cluster.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the kubernetes cluster.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Kubernetes versions available for upgrade.
     * 
     */
    @Export(name="nextUpgradeVersions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> nextUpgradeVersions;

    /**
     * @return Kubernetes versions available for upgrade.
     * 
     */
    public Output<List<String>> nextUpgradeVersions() {
        return this.nextUpgradeVersions;
    }
    /**
     * Cluster nodes URL.
     * 
     */
    @Export(name="nodesUrl", refs={String.class}, tree="[0]")
    private Output<String> nodesUrl;

    /**
     * @return Cluster nodes URL.
     * 
     */
    public Output<String> nodesUrl() {
        return this.nodesUrl;
    }
    /**
     * The private network configuration
     * 
     */
    @Export(name="privateNetworkConfiguration", refs={KubePrivateNetworkConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ KubePrivateNetworkConfiguration> privateNetworkConfiguration;

    /**
     * @return The private network configuration
     * 
     */
    public Output<Optional<KubePrivateNetworkConfiguration>> privateNetworkConfiguration() {
        return Codegen.optional(this.privateNetworkConfiguration);
    }
    /**
     * OpenStack private network (or vRack) ID to use. **Changing this value recreates the resource, including ETCD user data.** Defaults - not use private network.
     * 
     * &gt; __WARNING__ Updating the private network ID resets the cluster so that all user data is deleted.
     * 
     */
    @Export(name="privateNetworkId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privateNetworkId;

    /**
     * @return OpenStack private network (or vRack) ID to use. **Changing this value recreates the resource, including ETCD user data.** Defaults - not use private network.
     * 
     * &gt; __WARNING__ Updating the private network ID resets the cluster so that all user data is deleted.
     * 
     */
    public Output<Optional<String>> privateNetworkId() {
        return Codegen.optional(this.privateNetworkId);
    }
    /**
     * a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: &#34;GRA1&#34;. Defaults to all public cloud regions. **Changing this value recreates the resource.**
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: &#34;GRA1&#34;. Defaults to all public cloud regions. **Changing this value recreates the resource.**
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Cluster status. Should be normally set to &#39;READY&#39;.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Cluster status. Should be normally set to &#39;READY&#39;.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
     * 
     */
    @Export(name="updatePolicy", refs={String.class}, tree="[0]")
    private Output<String> updatePolicy;

    /**
     * @return Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
     * 
     */
    public Output<String> updatePolicy() {
        return this.updatePolicy;
    }
    /**
     * Management URL of your cluster.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return Management URL of your cluster.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Kube(String name) {
        this(name, KubeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Kube(String name, KubeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Kube(String name, KubeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/kube:Kube", name, args == null ? KubeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Kube(String name, Output<String> id, @Nullable KubeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/kube:Kube", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "kubeconfig",
                "kubeconfigAttributes"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Kube get(String name, Output<String> id, @Nullable KubeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Kube(name, id, state, options);
    }
}
