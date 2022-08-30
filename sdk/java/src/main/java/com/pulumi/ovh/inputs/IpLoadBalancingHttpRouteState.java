// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.inputs.IpLoadBalancingHttpRouteActionArgs;
import com.pulumi.ovh.inputs.IpLoadBalancingHttpRouteRuleArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpLoadBalancingHttpRouteState extends com.pulumi.resources.ResourceArgs {

    public static final IpLoadBalancingHttpRouteState Empty = new IpLoadBalancingHttpRouteState();

    /**
     * Action triggered when all rules match
     * 
     */
    @Import(name="action")
    private @Nullable Output<IpLoadBalancingHttpRouteActionArgs> action;

    /**
     * @return Action triggered when all rules match
     * 
     */
    public Optional<Output<IpLoadBalancingHttpRouteActionArgs>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * Human readable name for your route, this field is for you
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Human readable name for your route, this field is for you
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Route traffic for this frontend
     * 
     */
    @Import(name="frontendId")
    private @Nullable Output<Integer> frontendId;

    /**
     * @return Route traffic for this frontend
     * 
     */
    public Optional<Output<Integer>> frontendId() {
        return Optional.ofNullable(this.frontendId);
    }

    /**
     * List of rules to match to trigger action
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<IpLoadBalancingHttpRouteRuleArgs>> rules;

    /**
     * @return List of rules to match to trigger action
     * 
     */
    public Optional<Output<List<IpLoadBalancingHttpRouteRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * The internal name of your IP load balancing
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * HTTP status code for &#34;redirect&#34; and &#34;reject&#34; actions
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return HTTP status code for &#34;redirect&#34; and &#34;reject&#34; actions
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    /**
     * @return Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
     * 
     */
    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private IpLoadBalancingHttpRouteState() {}

    private IpLoadBalancingHttpRouteState(IpLoadBalancingHttpRouteState $) {
        this.action = $.action;
        this.displayName = $.displayName;
        this.frontendId = $.frontendId;
        this.rules = $.rules;
        this.serviceName = $.serviceName;
        this.status = $.status;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpLoadBalancingHttpRouteState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpLoadBalancingHttpRouteState $;

        public Builder() {
            $ = new IpLoadBalancingHttpRouteState();
        }

        public Builder(IpLoadBalancingHttpRouteState defaults) {
            $ = new IpLoadBalancingHttpRouteState(Objects.requireNonNull(defaults));
        }

        /**
         * @param action Action triggered when all rules match
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<IpLoadBalancingHttpRouteActionArgs> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action Action triggered when all rules match
         * 
         * @return builder
         * 
         */
        public Builder action(IpLoadBalancingHttpRouteActionArgs action) {
            return action(Output.of(action));
        }

        /**
         * @param displayName Human readable name for your route, this field is for you
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Human readable name for your route, this field is for you
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param frontendId Route traffic for this frontend
         * 
         * @return builder
         * 
         */
        public Builder frontendId(@Nullable Output<Integer> frontendId) {
            $.frontendId = frontendId;
            return this;
        }

        /**
         * @param frontendId Route traffic for this frontend
         * 
         * @return builder
         * 
         */
        public Builder frontendId(Integer frontendId) {
            return frontendId(Output.of(frontendId));
        }

        /**
         * @param rules List of rules to match to trigger action
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<IpLoadBalancingHttpRouteRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules List of rules to match to trigger action
         * 
         * @return builder
         * 
         */
        public Builder rules(List<IpLoadBalancingHttpRouteRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules List of rules to match to trigger action
         * 
         * @return builder
         * 
         */
        public Builder rules(IpLoadBalancingHttpRouteRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
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
         * @param status HTTP status code for &#34;redirect&#34; and &#34;reject&#34; actions
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status HTTP status code for &#34;redirect&#34; and &#34;reject&#34; actions
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param weight Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public IpLoadBalancingHttpRouteState build() {
            return $;
        }
    }

}
