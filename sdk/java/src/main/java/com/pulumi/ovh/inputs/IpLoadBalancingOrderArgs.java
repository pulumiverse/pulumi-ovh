// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.inputs.IpLoadBalancingOrderDetailArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpLoadBalancingOrderArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpLoadBalancingOrderArgs Empty = new IpLoadBalancingOrderArgs();

    /**
     * date
     * 
     */
    @Import(name="date")
    private @Nullable Output<String> date;

    /**
     * @return date
     * 
     */
    public Optional<Output<String>> date() {
        return Optional.ofNullable(this.date);
    }

    /**
     * Information about a Bill entry
     * 
     */
    @Import(name="details")
    private @Nullable Output<List<IpLoadBalancingOrderDetailArgs>> details;

    /**
     * @return Information about a Bill entry
     * 
     */
    public Optional<Output<List<IpLoadBalancingOrderDetailArgs>>> details() {
        return Optional.ofNullable(this.details);
    }

    /**
     * expiration date
     * 
     */
    @Import(name="expirationDate")
    private @Nullable Output<String> expirationDate;

    /**
     * @return expiration date
     * 
     */
    public Optional<Output<String>> expirationDate() {
        return Optional.ofNullable(this.expirationDate);
    }

    /**
     * order id
     * 
     */
    @Import(name="orderId")
    private @Nullable Output<Integer> orderId;

    /**
     * @return order id
     * 
     */
    public Optional<Output<Integer>> orderId() {
        return Optional.ofNullable(this.orderId);
    }

    private IpLoadBalancingOrderArgs() {}

    private IpLoadBalancingOrderArgs(IpLoadBalancingOrderArgs $) {
        this.date = $.date;
        this.details = $.details;
        this.expirationDate = $.expirationDate;
        this.orderId = $.orderId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpLoadBalancingOrderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpLoadBalancingOrderArgs $;

        public Builder() {
            $ = new IpLoadBalancingOrderArgs();
        }

        public Builder(IpLoadBalancingOrderArgs defaults) {
            $ = new IpLoadBalancingOrderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param date date
         * 
         * @return builder
         * 
         */
        public Builder date(@Nullable Output<String> date) {
            $.date = date;
            return this;
        }

        /**
         * @param date date
         * 
         * @return builder
         * 
         */
        public Builder date(String date) {
            return date(Output.of(date));
        }

        /**
         * @param details Information about a Bill entry
         * 
         * @return builder
         * 
         */
        public Builder details(@Nullable Output<List<IpLoadBalancingOrderDetailArgs>> details) {
            $.details = details;
            return this;
        }

        /**
         * @param details Information about a Bill entry
         * 
         * @return builder
         * 
         */
        public Builder details(List<IpLoadBalancingOrderDetailArgs> details) {
            return details(Output.of(details));
        }

        /**
         * @param details Information about a Bill entry
         * 
         * @return builder
         * 
         */
        public Builder details(IpLoadBalancingOrderDetailArgs... details) {
            return details(List.of(details));
        }

        /**
         * @param expirationDate expiration date
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(@Nullable Output<String> expirationDate) {
            $.expirationDate = expirationDate;
            return this;
        }

        /**
         * @param expirationDate expiration date
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(String expirationDate) {
            return expirationDate(Output.of(expirationDate));
        }

        /**
         * @param orderId order id
         * 
         * @return builder
         * 
         */
        public Builder orderId(@Nullable Output<Integer> orderId) {
            $.orderId = orderId;
            return this;
        }

        /**
         * @param orderId order id
         * 
         * @return builder
         * 
         */
        public Builder orderId(Integer orderId) {
            return orderId(Output.of(orderId));
        }

        public IpLoadBalancingOrderArgs build() {
            return $;
        }
    }

}
