// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.CloudProject.inputs.ProjectPlanArgs;
import com.pulumi.ovh.CloudProject.inputs.ProjectPlanOptionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectArgs Empty = new ProjectArgs();

    /**
     * A description associated with the user.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description associated with the user.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * OVHcloud Subsidiary
     * 
     */
    @Import(name="ovhSubsidiary", required=true)
    private Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary
     * 
     */
    public Output<String> ovhSubsidiary() {
        return this.ovhSubsidiary;
    }

    /**
     * Ovh payment mode
     * 
     * @deprecated
     * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     * 
     */
    @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
    @Import(name="paymentMean")
    private @Nullable Output<String> paymentMean;

    /**
     * @return Ovh payment mode
     * 
     * @deprecated
     * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     * 
     */
    @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
    public Optional<Output<String>> paymentMean() {
        return Optional.ofNullable(this.paymentMean);
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="plan", required=true)
    private Output<ProjectPlanArgs> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<ProjectPlanArgs> plan() {
        return this.plan;
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="planOptions")
    private @Nullable Output<List<ProjectPlanOptionArgs>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<List<ProjectPlanOptionArgs>>> planOptions() {
        return Optional.ofNullable(this.planOptions);
    }

    private ProjectArgs() {}

    private ProjectArgs(ProjectArgs $) {
        this.description = $.description;
        this.ovhSubsidiary = $.ovhSubsidiary;
        this.paymentMean = $.paymentMean;
        this.plan = $.plan;
        this.planOptions = $.planOptions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectArgs $;

        public Builder() {
            $ = new ProjectArgs();
        }

        public Builder(ProjectArgs defaults) {
            $ = new ProjectArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A description associated with the user.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description associated with the user.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(Output<String> ovhSubsidiary) {
            $.ovhSubsidiary = ovhSubsidiary;
            return this;
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            return ovhSubsidiary(Output.of(ovhSubsidiary));
        }

        /**
         * @param paymentMean Ovh payment mode
         * 
         * @return builder
         * 
         * @deprecated
         * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
         * 
         */
        @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
        public Builder paymentMean(@Nullable Output<String> paymentMean) {
            $.paymentMean = paymentMean;
            return this;
        }

        /**
         * @param paymentMean Ovh payment mode
         * 
         * @return builder
         * 
         * @deprecated
         * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
         * 
         */
        @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
        public Builder paymentMean(String paymentMean) {
            return paymentMean(Output.of(paymentMean));
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(Output<ProjectPlanArgs> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(ProjectPlanArgs plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(@Nullable Output<List<ProjectPlanOptionArgs>> planOptions) {
            $.planOptions = planOptions;
            return this;
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(List<ProjectPlanOptionArgs> planOptions) {
            return planOptions(Output.of(planOptions));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(ProjectPlanOptionArgs... planOptions) {
            return planOptions(List.of(planOptions));
        }

        public ProjectArgs build() {
            $.ovhSubsidiary = Objects.requireNonNull($.ovhSubsidiary, "expected parameter 'ovhSubsidiary' to be non-null");
            $.plan = Objects.requireNonNull($.plan, "expected parameter 'plan' to be non-null");
            return $;
        }
    }

}
