// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpLoadBalancingTcpRouteActionArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpLoadBalancingTcpRouteActionArgs Empty = new IpLoadBalancingTcpRouteActionArgs();

    /**
     * Farm ID for &#34;farm&#34; action type, empty for others.
     * 
     */
    @Import(name="target")
    private @Nullable Output<String> target;

    /**
     * @return Farm ID for &#34;farm&#34; action type, empty for others.
     * 
     */
    public Optional<Output<String>> target() {
        return Optional.ofNullable(this.target);
    }

    /**
     * Action to trigger if all the rules of this route matches
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Action to trigger if all the rules of this route matches
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private IpLoadBalancingTcpRouteActionArgs() {}

    private IpLoadBalancingTcpRouteActionArgs(IpLoadBalancingTcpRouteActionArgs $) {
        this.target = $.target;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpLoadBalancingTcpRouteActionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpLoadBalancingTcpRouteActionArgs $;

        public Builder() {
            $ = new IpLoadBalancingTcpRouteActionArgs();
        }

        public Builder(IpLoadBalancingTcpRouteActionArgs defaults) {
            $ = new IpLoadBalancingTcpRouteActionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param target Farm ID for &#34;farm&#34; action type, empty for others.
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target Farm ID for &#34;farm&#34; action type, empty for others.
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        /**
         * @param type Action to trigger if all the rules of this route matches
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Action to trigger if all the rules of this route matches
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public IpLoadBalancingTcpRouteActionArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
