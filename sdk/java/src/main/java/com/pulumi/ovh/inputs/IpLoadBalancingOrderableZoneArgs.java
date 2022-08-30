// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpLoadBalancingOrderableZoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpLoadBalancingOrderableZoneArgs Empty = new IpLoadBalancingOrderableZoneArgs();

    /**
     * The zone three letter code
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The zone three letter code
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Plan code
     * 
     */
    @Import(name="planCode")
    private @Nullable Output<String> planCode;

    /**
     * @return Plan code
     * 
     */
    public Optional<Output<String>> planCode() {
        return Optional.ofNullable(this.planCode);
    }

    private IpLoadBalancingOrderableZoneArgs() {}

    private IpLoadBalancingOrderableZoneArgs(IpLoadBalancingOrderableZoneArgs $) {
        this.name = $.name;
        this.planCode = $.planCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpLoadBalancingOrderableZoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpLoadBalancingOrderableZoneArgs $;

        public Builder() {
            $ = new IpLoadBalancingOrderableZoneArgs();
        }

        public Builder(IpLoadBalancingOrderableZoneArgs defaults) {
            $ = new IpLoadBalancingOrderableZoneArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The zone three letter code
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The zone three letter code
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param planCode Plan code
         * 
         * @return builder
         * 
         */
        public Builder planCode(@Nullable Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        /**
         * @param planCode Plan code
         * 
         * @return builder
         * 
         */
        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        public IpLoadBalancingOrderableZoneArgs build() {
            return $;
        }
    }

}
