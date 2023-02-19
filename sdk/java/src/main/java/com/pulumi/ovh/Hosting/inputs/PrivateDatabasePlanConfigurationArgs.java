// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class PrivateDatabasePlanConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final PrivateDatabasePlanConfigurationArgs Empty = new PrivateDatabasePlanConfigurationArgs();

    /**
     * Identifier of the resource
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return Identifier of the resource
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    /**
     * Path to the resource in API.OVH.COM
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Path to the resource in API.OVH.COM
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private PrivateDatabasePlanConfigurationArgs() {}

    private PrivateDatabasePlanConfigurationArgs(PrivateDatabasePlanConfigurationArgs $) {
        this.label = $.label;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateDatabasePlanConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateDatabasePlanConfigurationArgs $;

        public Builder() {
            $ = new PrivateDatabasePlanConfigurationArgs();
        }

        public Builder(PrivateDatabasePlanConfigurationArgs defaults) {
            $ = new PrivateDatabasePlanConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param label Identifier of the resource
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label Identifier of the resource
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param value Path to the resource in API.OVH.COM
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Path to the resource in API.OVH.COM
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public PrivateDatabasePlanConfigurationArgs build() {
            $.label = Objects.requireNonNull($.label, "expected parameter 'label' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}
