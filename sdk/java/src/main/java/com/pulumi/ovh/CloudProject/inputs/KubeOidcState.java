// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeOidcState extends com.pulumi.resources.ResourceArgs {

    public static final KubeOidcState Empty = new KubeOidcState();

    /**
     * The OIDC client ID.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The OIDC client ID.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The OIDC issuer url.
     * 
     */
    @Import(name="issuerUrl")
    private @Nullable Output<String> issuerUrl;

    /**
     * @return The OIDC issuer url.
     * 
     */
    public Optional<Output<String>> issuerUrl() {
        return Optional.ofNullable(this.issuerUrl);
    }

    /**
     * The ID of the managed kubernetes cluster.
     * 
     */
    @Import(name="kubeId")
    private @Nullable Output<String> kubeId;

    /**
     * @return The ID of the managed kubernetes cluster.
     * 
     */
    public Optional<Output<String>> kubeId() {
        return Optional.ofNullable(this.kubeId);
    }

    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private KubeOidcState() {}

    private KubeOidcState(KubeOidcState $) {
        this.clientId = $.clientId;
        this.issuerUrl = $.issuerUrl;
        this.kubeId = $.kubeId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeOidcState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeOidcState $;

        public Builder() {
            $ = new KubeOidcState();
        }

        public Builder(KubeOidcState defaults) {
            $ = new KubeOidcState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId The OIDC client ID.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The OIDC client ID.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param issuerUrl The OIDC issuer url.
         * 
         * @return builder
         * 
         */
        public Builder issuerUrl(@Nullable Output<String> issuerUrl) {
            $.issuerUrl = issuerUrl;
            return this;
        }

        /**
         * @param issuerUrl The OIDC issuer url.
         * 
         * @return builder
         * 
         */
        public Builder issuerUrl(String issuerUrl) {
            return issuerUrl(Output.of(issuerUrl));
        }

        /**
         * @param kubeId The ID of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeId(@Nullable Output<String> kubeId) {
            $.kubeId = kubeId;
            return this;
        }

        /**
         * @param kubeId The ID of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeId(String kubeId) {
            return kubeId(Output.of(kubeId));
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted,
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
         * @param serviceName The ID of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public KubeOidcState build() {
            return $;
        }
    }

}