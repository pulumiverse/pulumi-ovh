// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.outputs.GetOrderCartProductResultPricePrice;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOrderCartProductResultPrice {
    /**
     * @return Capacities of the pricing (type of pricing)
     * 
     */
    private List<Object> capacities;
    /**
     * @return Description of the pricing
     * 
     */
    private String description;
    /**
     * @return Duration for ordering the product
     * 
     */
    private String duration;
    /**
     * @return Interval of renewal
     * 
     */
    private Integer interval;
    /**
     * @return Maximum quantity that can be ordered
     * 
     */
    private Integer maximumQuantity;
    /**
     * @return Maximum repeat for renewal
     * 
     */
    private Integer maximumRepeat;
    /**
     * @return Minimum quantity that can be ordered
     * 
     */
    private Integer minimumQuantity;
    /**
     * @return Minimum repeat for renewal
     * 
     */
    private Integer minimumRepeat;
    /**
     * @return Price of the product in micro-centims
     * 
     */
    private Integer priceInUcents;
    /**
     * @return Price of the product (Price with its currency and textual representation)
     * 
     */
    private List<GetOrderCartProductResultPricePrice> prices;
    /**
     * @return Pricing model identifier
     * 
     */
    private String pricingMode;
    /**
     * @return Pricing type
     * 
     */
    private String pricingType;

    private GetOrderCartProductResultPrice() {}
    /**
     * @return Capacities of the pricing (type of pricing)
     * 
     */
    public List<Object> capacities() {
        return this.capacities;
    }
    /**
     * @return Description of the pricing
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Duration for ordering the product
     * 
     */
    public String duration() {
        return this.duration;
    }
    /**
     * @return Interval of renewal
     * 
     */
    public Integer interval() {
        return this.interval;
    }
    /**
     * @return Maximum quantity that can be ordered
     * 
     */
    public Integer maximumQuantity() {
        return this.maximumQuantity;
    }
    /**
     * @return Maximum repeat for renewal
     * 
     */
    public Integer maximumRepeat() {
        return this.maximumRepeat;
    }
    /**
     * @return Minimum quantity that can be ordered
     * 
     */
    public Integer minimumQuantity() {
        return this.minimumQuantity;
    }
    /**
     * @return Minimum repeat for renewal
     * 
     */
    public Integer minimumRepeat() {
        return this.minimumRepeat;
    }
    /**
     * @return Price of the product in micro-centims
     * 
     */
    public Integer priceInUcents() {
        return this.priceInUcents;
    }
    /**
     * @return Price of the product (Price with its currency and textual representation)
     * 
     */
    public List<GetOrderCartProductResultPricePrice> prices() {
        return this.prices;
    }
    /**
     * @return Pricing model identifier
     * 
     */
    public String pricingMode() {
        return this.pricingMode;
    }
    /**
     * @return Pricing type
     * 
     */
    public String pricingType() {
        return this.pricingType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrderCartProductResultPrice defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Object> capacities;
        private String description;
        private String duration;
        private Integer interval;
        private Integer maximumQuantity;
        private Integer maximumRepeat;
        private Integer minimumQuantity;
        private Integer minimumRepeat;
        private Integer priceInUcents;
        private List<GetOrderCartProductResultPricePrice> prices;
        private String pricingMode;
        private String pricingType;
        public Builder() {}
        public Builder(GetOrderCartProductResultPrice defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.capacities = defaults.capacities;
    	      this.description = defaults.description;
    	      this.duration = defaults.duration;
    	      this.interval = defaults.interval;
    	      this.maximumQuantity = defaults.maximumQuantity;
    	      this.maximumRepeat = defaults.maximumRepeat;
    	      this.minimumQuantity = defaults.minimumQuantity;
    	      this.minimumRepeat = defaults.minimumRepeat;
    	      this.priceInUcents = defaults.priceInUcents;
    	      this.prices = defaults.prices;
    	      this.pricingMode = defaults.pricingMode;
    	      this.pricingType = defaults.pricingType;
        }

        @CustomType.Setter
        public Builder capacities(List<Object> capacities) {
            this.capacities = Objects.requireNonNull(capacities);
            return this;
        }
        public Builder capacities(Object... capacities) {
            return capacities(List.of(capacities));
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder duration(String duration) {
            this.duration = Objects.requireNonNull(duration);
            return this;
        }
        @CustomType.Setter
        public Builder interval(Integer interval) {
            this.interval = Objects.requireNonNull(interval);
            return this;
        }
        @CustomType.Setter
        public Builder maximumQuantity(Integer maximumQuantity) {
            this.maximumQuantity = Objects.requireNonNull(maximumQuantity);
            return this;
        }
        @CustomType.Setter
        public Builder maximumRepeat(Integer maximumRepeat) {
            this.maximumRepeat = Objects.requireNonNull(maximumRepeat);
            return this;
        }
        @CustomType.Setter
        public Builder minimumQuantity(Integer minimumQuantity) {
            this.minimumQuantity = Objects.requireNonNull(minimumQuantity);
            return this;
        }
        @CustomType.Setter
        public Builder minimumRepeat(Integer minimumRepeat) {
            this.minimumRepeat = Objects.requireNonNull(minimumRepeat);
            return this;
        }
        @CustomType.Setter
        public Builder priceInUcents(Integer priceInUcents) {
            this.priceInUcents = Objects.requireNonNull(priceInUcents);
            return this;
        }
        @CustomType.Setter
        public Builder prices(List<GetOrderCartProductResultPricePrice> prices) {
            this.prices = Objects.requireNonNull(prices);
            return this;
        }
        public Builder prices(GetOrderCartProductResultPricePrice... prices) {
            return prices(List.of(prices));
        }
        @CustomType.Setter
        public Builder pricingMode(String pricingMode) {
            this.pricingMode = Objects.requireNonNull(pricingMode);
            return this;
        }
        @CustomType.Setter
        public Builder pricingType(String pricingType) {
            this.pricingType = Objects.requireNonNull(pricingType);
            return this;
        }
        public GetOrderCartProductResultPrice build() {
            final var o = new GetOrderCartProductResultPrice();
            o.capacities = capacities;
            o.description = description;
            o.duration = duration;
            o.interval = interval;
            o.maximumQuantity = maximumQuantity;
            o.maximumRepeat = maximumRepeat;
            o.minimumQuantity = minimumQuantity;
            o.minimumRepeat = minimumRepeat;
            o.priceInUcents = priceInUcents;
            o.prices = prices;
            o.pricingMode = pricingMode;
            o.pricingType = pricingType;
            return o;
        }
    }
}
