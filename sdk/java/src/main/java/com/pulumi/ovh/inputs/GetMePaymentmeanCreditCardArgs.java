// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetMePaymentmeanCreditCardArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMePaymentmeanCreditCardArgs Empty = new GetMePaymentmeanCreditCardArgs();

    /**
     * a regexp used to filter credit cards
     * on their `description` attributes.
     * 
     */
    @Import(name="descriptionRegexp")
    private @Nullable Output<String> descriptionRegexp;

    /**
     * @return a regexp used to filter credit cards
     * on their `description` attributes.
     * 
     */
    public Optional<Output<String>> descriptionRegexp() {
        return Optional.ofNullable(this.descriptionRegexp);
    }

    /**
     * Filter credit cards on their `state` attribute.
     * Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
     * 
     */
    @Import(name="states")
    private @Nullable Output<List<String>> states;

    /**
     * @return Filter credit cards on their `state` attribute.
     * Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
     * 
     */
    public Optional<Output<List<String>>> states() {
        return Optional.ofNullable(this.states);
    }

    /**
     * Retrieve credit card marked as default payment mean.
     * 
     */
    @Import(name="useDefault")
    private @Nullable Output<Boolean> useDefault;

    /**
     * @return Retrieve credit card marked as default payment mean.
     * 
     */
    public Optional<Output<Boolean>> useDefault() {
        return Optional.ofNullable(this.useDefault);
    }

    /**
     * Retrieve the credit card that will be the last
     * to expire according to its expiration date.
     * 
     */
    @Import(name="useLastToExpire")
    private @Nullable Output<Boolean> useLastToExpire;

    /**
     * @return Retrieve the credit card that will be the last
     * to expire according to its expiration date.
     * 
     */
    public Optional<Output<Boolean>> useLastToExpire() {
        return Optional.ofNullable(this.useLastToExpire);
    }

    private GetMePaymentmeanCreditCardArgs() {}

    private GetMePaymentmeanCreditCardArgs(GetMePaymentmeanCreditCardArgs $) {
        this.descriptionRegexp = $.descriptionRegexp;
        this.states = $.states;
        this.useDefault = $.useDefault;
        this.useLastToExpire = $.useLastToExpire;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMePaymentmeanCreditCardArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMePaymentmeanCreditCardArgs $;

        public Builder() {
            $ = new GetMePaymentmeanCreditCardArgs();
        }

        public Builder(GetMePaymentmeanCreditCardArgs defaults) {
            $ = new GetMePaymentmeanCreditCardArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param descriptionRegexp a regexp used to filter credit cards
         * on their `description` attributes.
         * 
         * @return builder
         * 
         */
        public Builder descriptionRegexp(@Nullable Output<String> descriptionRegexp) {
            $.descriptionRegexp = descriptionRegexp;
            return this;
        }

        /**
         * @param descriptionRegexp a regexp used to filter credit cards
         * on their `description` attributes.
         * 
         * @return builder
         * 
         */
        public Builder descriptionRegexp(String descriptionRegexp) {
            return descriptionRegexp(Output.of(descriptionRegexp));
        }

        /**
         * @param states Filter credit cards on their `state` attribute.
         * Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
         * 
         * @return builder
         * 
         */
        public Builder states(@Nullable Output<List<String>> states) {
            $.states = states;
            return this;
        }

        /**
         * @param states Filter credit cards on their `state` attribute.
         * Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
         * 
         * @return builder
         * 
         */
        public Builder states(List<String> states) {
            return states(Output.of(states));
        }

        /**
         * @param states Filter credit cards on their `state` attribute.
         * Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
         * 
         * @return builder
         * 
         */
        public Builder states(String... states) {
            return states(List.of(states));
        }

        /**
         * @param useDefault Retrieve credit card marked as default payment mean.
         * 
         * @return builder
         * 
         */
        public Builder useDefault(@Nullable Output<Boolean> useDefault) {
            $.useDefault = useDefault;
            return this;
        }

        /**
         * @param useDefault Retrieve credit card marked as default payment mean.
         * 
         * @return builder
         * 
         */
        public Builder useDefault(Boolean useDefault) {
            return useDefault(Output.of(useDefault));
        }

        /**
         * @param useLastToExpire Retrieve the credit card that will be the last
         * to expire according to its expiration date.
         * 
         * @return builder
         * 
         */
        public Builder useLastToExpire(@Nullable Output<Boolean> useLastToExpire) {
            $.useLastToExpire = useLastToExpire;
            return this;
        }

        /**
         * @param useLastToExpire Retrieve the credit card that will be the last
         * to expire according to its expiration date.
         * 
         * @return builder
         * 
         */
        public Builder useLastToExpire(Boolean useLastToExpire) {
            return useLastToExpire(Output.of(useLastToExpire));
        }

        public GetMePaymentmeanCreditCardArgs build() {
            return $;
        }
    }

}
