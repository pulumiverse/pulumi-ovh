// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpLoadBalancingHttpFarmProbeArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpLoadBalancingHttpFarmProbeArgs Empty = new IpLoadBalancingHttpFarmProbeArgs();

    /**
     * Force use of SSL (TLS)
     * 
     */
    @Import(name="forceSsl")
    private @Nullable Output<Boolean> forceSsl;

    /**
     * @return Force use of SSL (TLS)
     * 
     */
    public Optional<Output<Boolean>> forceSsl() {
        return Optional.ofNullable(this.forceSsl);
    }

    /**
     * probe interval, Value between 30 and 3600 seconds, default 30
     * 
     */
    @Import(name="interval")
    private @Nullable Output<Integer> interval;

    /**
     * @return probe interval, Value between 30 and 3600 seconds, default 30
     * 
     */
    public Optional<Output<Integer>> interval() {
        return Optional.ofNullable(this.interval);
    }

    /**
     * What to mach `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
     * 
     */
    @Import(name="match")
    private @Nullable Output<String> match;

    /**
     * @return What to mach `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
     * 
     */
    public Optional<Output<String>> match() {
        return Optional.ofNullable(this.match);
    }

    /**
     * HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
     * 
     */
    @Import(name="method")
    private @Nullable Output<String> method;

    /**
     * @return HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
     * 
     */
    public Optional<Output<String>> method() {
        return Optional.ofNullable(this.method);
    }

    /**
     * Negate probe result
     * 
     */
    @Import(name="negate")
    private @Nullable Output<Boolean> negate;

    /**
     * @return Negate probe result
     * 
     */
    public Optional<Output<Boolean>> negate() {
        return Optional.ofNullable(this.negate);
    }

    /**
     * Pattern to match against `match`
     * 
     */
    @Import(name="pattern")
    private @Nullable Output<String> pattern;

    /**
     * @return Pattern to match against `match`
     * 
     */
    public Optional<Output<String>> pattern() {
        return Optional.ofNullable(this.pattern);
    }

    /**
     * Port for backends to recieve traffic on.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port for backends to recieve traffic on.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * URL for HTTP probe type.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return URL for HTTP probe type.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private IpLoadBalancingHttpFarmProbeArgs() {}

    private IpLoadBalancingHttpFarmProbeArgs(IpLoadBalancingHttpFarmProbeArgs $) {
        this.forceSsl = $.forceSsl;
        this.interval = $.interval;
        this.match = $.match;
        this.method = $.method;
        this.negate = $.negate;
        this.pattern = $.pattern;
        this.port = $.port;
        this.type = $.type;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpLoadBalancingHttpFarmProbeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpLoadBalancingHttpFarmProbeArgs $;

        public Builder() {
            $ = new IpLoadBalancingHttpFarmProbeArgs();
        }

        public Builder(IpLoadBalancingHttpFarmProbeArgs defaults) {
            $ = new IpLoadBalancingHttpFarmProbeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param forceSsl Force use of SSL (TLS)
         * 
         * @return builder
         * 
         */
        public Builder forceSsl(@Nullable Output<Boolean> forceSsl) {
            $.forceSsl = forceSsl;
            return this;
        }

        /**
         * @param forceSsl Force use of SSL (TLS)
         * 
         * @return builder
         * 
         */
        public Builder forceSsl(Boolean forceSsl) {
            return forceSsl(Output.of(forceSsl));
        }

        /**
         * @param interval probe interval, Value between 30 and 3600 seconds, default 30
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<Integer> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval probe interval, Value between 30 and 3600 seconds, default 30
         * 
         * @return builder
         * 
         */
        public Builder interval(Integer interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param match What to mach `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
         * 
         * @return builder
         * 
         */
        public Builder match(@Nullable Output<String> match) {
            $.match = match;
            return this;
        }

        /**
         * @param match What to mach `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
         * 
         * @return builder
         * 
         */
        public Builder match(String match) {
            return match(Output.of(match));
        }

        /**
         * @param method HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
         * 
         * @return builder
         * 
         */
        public Builder method(@Nullable Output<String> method) {
            $.method = method;
            return this;
        }

        /**
         * @param method HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
         * 
         * @return builder
         * 
         */
        public Builder method(String method) {
            return method(Output.of(method));
        }

        /**
         * @param negate Negate probe result
         * 
         * @return builder
         * 
         */
        public Builder negate(@Nullable Output<Boolean> negate) {
            $.negate = negate;
            return this;
        }

        /**
         * @param negate Negate probe result
         * 
         * @return builder
         * 
         */
        public Builder negate(Boolean negate) {
            return negate(Output.of(negate));
        }

        /**
         * @param pattern Pattern to match against `match`
         * 
         * @return builder
         * 
         */
        public Builder pattern(@Nullable Output<String> pattern) {
            $.pattern = pattern;
            return this;
        }

        /**
         * @param pattern Pattern to match against `match`
         * 
         * @return builder
         * 
         */
        public Builder pattern(String pattern) {
            return pattern(Output.of(pattern));
        }

        /**
         * @param port Port for backends to recieve traffic on.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port for backends to recieve traffic on.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param type Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param url URL for HTTP probe type.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url URL for HTTP probe type.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public IpLoadBalancingHttpFarmProbeArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
