// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOrderCartProductOptionsPlanSelectedPricePrice {
    /**
     * @return Currency code
     * 
     */
    private String currencyCode;
    /**
     * @return Textual representation
     * 
     */
    private String text;
    /**
     * @return The effective price
     * 
     */
    private Double value;

    private GetOrderCartProductOptionsPlanSelectedPricePrice() {}
    /**
     * @return Currency code
     * 
     */
    public String currencyCode() {
        return this.currencyCode;
    }
    /**
     * @return Textual representation
     * 
     */
    public String text() {
        return this.text;
    }
    /**
     * @return The effective price
     * 
     */
    public Double value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrderCartProductOptionsPlanSelectedPricePrice defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String currencyCode;
        private String text;
        private Double value;
        public Builder() {}
        public Builder(GetOrderCartProductOptionsPlanSelectedPricePrice defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.currencyCode = defaults.currencyCode;
    	      this.text = defaults.text;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder currencyCode(String currencyCode) {
            this.currencyCode = Objects.requireNonNull(currencyCode);
            return this;
        }
        @CustomType.Setter
        public Builder text(String text) {
            this.text = Objects.requireNonNull(text);
            return this;
        }
        @CustomType.Setter
        public Builder value(Double value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public GetOrderCartProductOptionsPlanSelectedPricePrice build() {
            final var o = new GetOrderCartProductOptionsPlanSelectedPricePrice();
            o.currencyCode = currencyCode;
            o.text = text;
            o.value = value;
            return o;
        }
    }
}
