// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpLoadBalancingTcpRouteRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpLoadBalancingTcpRouteRuleArgs Empty = new IpLoadBalancingTcpRouteRuleArgs();

    /**
     * Human readable name for your rule, this field is for you
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Human readable name for your rule, this field is for you
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Name of the field to match like &#34;protocol&#34; or &#34;host&#34;. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34; for a list of available rules
     * 
     */
    @Import(name="field", required=true)
    private Output<String> field;

    /**
     * @return Name of the field to match like &#34;protocol&#34; or &#34;host&#34;. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34; for a list of available rules
     * 
     */
    public Output<String> field() {
        return this.field;
    }

    /**
     * Matching operator. Not all operators are available for all fields. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34;
     * 
     */
    @Import(name="match", required=true)
    private Output<String> match;

    /**
     * @return Matching operator. Not all operators are available for all fields. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34;
     * 
     */
    public Output<String> match() {
        return this.match;
    }

    /**
     * Invert the matching operator effect
     * 
     */
    @Import(name="negate")
    private @Nullable Output<Boolean> negate;

    /**
     * @return Invert the matching operator effect
     * 
     */
    public Optional<Output<Boolean>> negate() {
        return Optional.ofNullable(this.negate);
    }

    /**
     * Value to match against this match. Interpretation if this field depends on the match and field
     * 
     */
    @Import(name="pattern")
    private @Nullable Output<String> pattern;

    /**
     * @return Value to match against this match. Interpretation if this field depends on the match and field
     * 
     */
    public Optional<Output<String>> pattern() {
        return Optional.ofNullable(this.pattern);
    }

    /**
     * The route to apply this rule
     * 
     */
    @Import(name="routeId", required=true)
    private Output<String> routeId;

    /**
     * @return The route to apply this rule
     * 
     */
    public Output<String> routeId() {
        return this.routeId;
    }

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
     * Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     * 
     */
    @Import(name="subField")
    private @Nullable Output<String> subField;

    /**
     * @return Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     * 
     */
    public Optional<Output<String>> subField() {
        return Optional.ofNullable(this.subField);
    }

    private IpLoadBalancingTcpRouteRuleArgs() {}

    private IpLoadBalancingTcpRouteRuleArgs(IpLoadBalancingTcpRouteRuleArgs $) {
        this.displayName = $.displayName;
        this.field = $.field;
        this.match = $.match;
        this.negate = $.negate;
        this.pattern = $.pattern;
        this.routeId = $.routeId;
        this.serviceName = $.serviceName;
        this.subField = $.subField;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpLoadBalancingTcpRouteRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpLoadBalancingTcpRouteRuleArgs $;

        public Builder() {
            $ = new IpLoadBalancingTcpRouteRuleArgs();
        }

        public Builder(IpLoadBalancingTcpRouteRuleArgs defaults) {
            $ = new IpLoadBalancingTcpRouteRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName Human readable name for your rule, this field is for you
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Human readable name for your rule, this field is for you
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param field Name of the field to match like &#34;protocol&#34; or &#34;host&#34;. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34; for a list of available rules
         * 
         * @return builder
         * 
         */
        public Builder field(Output<String> field) {
            $.field = field;
            return this;
        }

        /**
         * @param field Name of the field to match like &#34;protocol&#34; or &#34;host&#34;. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34; for a list of available rules
         * 
         * @return builder
         * 
         */
        public Builder field(String field) {
            return field(Output.of(field));
        }

        /**
         * @param match Matching operator. Not all operators are available for all fields. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34;
         * 
         * @return builder
         * 
         */
        public Builder match(Output<String> match) {
            $.match = match;
            return this;
        }

        /**
         * @param match Matching operator. Not all operators are available for all fields. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34;
         * 
         * @return builder
         * 
         */
        public Builder match(String match) {
            return match(Output.of(match));
        }

        /**
         * @param negate Invert the matching operator effect
         * 
         * @return builder
         * 
         */
        public Builder negate(@Nullable Output<Boolean> negate) {
            $.negate = negate;
            return this;
        }

        /**
         * @param negate Invert the matching operator effect
         * 
         * @return builder
         * 
         */
        public Builder negate(Boolean negate) {
            return negate(Output.of(negate));
        }

        /**
         * @param pattern Value to match against this match. Interpretation if this field depends on the match and field
         * 
         * @return builder
         * 
         */
        public Builder pattern(@Nullable Output<String> pattern) {
            $.pattern = pattern;
            return this;
        }

        /**
         * @param pattern Value to match against this match. Interpretation if this field depends on the match and field
         * 
         * @return builder
         * 
         */
        public Builder pattern(String pattern) {
            return pattern(Output.of(pattern));
        }

        /**
         * @param routeId The route to apply this rule
         * 
         * @return builder
         * 
         */
        public Builder routeId(Output<String> routeId) {
            $.routeId = routeId;
            return this;
        }

        /**
         * @param routeId The route to apply this rule
         * 
         * @return builder
         * 
         */
        public Builder routeId(String routeId) {
            return routeId(Output.of(routeId));
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
         * @param subField Name of sub-field, if applicable. This may be a Cookie or Header name for instance
         * 
         * @return builder
         * 
         */
        public Builder subField(@Nullable Output<String> subField) {
            $.subField = subField;
            return this;
        }

        /**
         * @param subField Name of sub-field, if applicable. This may be a Cookie or Header name for instance
         * 
         * @return builder
         * 
         */
        public Builder subField(String subField) {
            return subField(Output.of(subField));
        }

        public IpLoadBalancingTcpRouteRuleArgs build() {
            $.field = Objects.requireNonNull($.field, "expected parameter 'field' to be non-null");
            $.match = Objects.requireNonNull($.match, "expected parameter 'match' to be non-null");
            $.routeId = Objects.requireNonNull($.routeId, "expected parameter 'routeId' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
