// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetIpLoadbalancingVrackNetworkArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpLoadbalancingVrackNetworkArgs Empty = new GetIpLoadbalancingVrackNetworkArgs();

    /**
     * The internal name of your IP load balancing
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * Internal Load Balancer identifier of the vRack private network
     * 
     */
    @Import(name="vrackNetworkId", required=true)
    private Output<Integer> vrackNetworkId;

    /**
     * @return Internal Load Balancer identifier of the vRack private network
     * 
     */
    public Output<Integer> vrackNetworkId() {
        return this.vrackNetworkId;
    }

    private GetIpLoadbalancingVrackNetworkArgs() {}

    private GetIpLoadbalancingVrackNetworkArgs(GetIpLoadbalancingVrackNetworkArgs $) {
        this.serviceName = $.serviceName;
        this.vrackNetworkId = $.vrackNetworkId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpLoadbalancingVrackNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpLoadbalancingVrackNetworkArgs $;

        public Builder() {
            $ = new GetIpLoadbalancingVrackNetworkArgs();
        }

        public Builder(GetIpLoadbalancingVrackNetworkArgs defaults) {
            $ = new GetIpLoadbalancingVrackNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param vrackNetworkId Internal Load Balancer identifier of the vRack private network
         * 
         * @return builder
         * 
         */
        public Builder vrackNetworkId(Output<Integer> vrackNetworkId) {
            $.vrackNetworkId = vrackNetworkId;
            return this;
        }

        /**
         * @param vrackNetworkId Internal Load Balancer identifier of the vRack private network
         * 
         * @return builder
         * 
         */
        public Builder vrackNetworkId(Integer vrackNetworkId) {
            return vrackNetworkId(Output.of(vrackNetworkId));
        }

        public GetIpLoadbalancingVrackNetworkArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            $.vrackNetworkId = Objects.requireNonNull($.vrackNetworkId, "expected parameter 'vrackNetworkId' to be non-null");
            return $;
        }
    }

}
