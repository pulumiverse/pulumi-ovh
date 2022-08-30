// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IpLoadBalancingHttpFarmProbe {
    /**
     * @return Force use of SSL (TLS)
     * 
     */
    private @Nullable Boolean forceSsl;
    /**
     * @return probe interval, Value between 30 and 3600 seconds, default 30
     * 
     */
    private @Nullable Integer interval;
    /**
     * @return What to mach `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
     * 
     */
    private @Nullable String match;
    /**
     * @return HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
     * 
     */
    private @Nullable String method;
    /**
     * @return Negate probe result
     * 
     */
    private @Nullable Boolean negate;
    /**
     * @return Pattern to match against `match`
     * 
     */
    private @Nullable String pattern;
    /**
     * @return Port for backends to recieve traffic on.
     * 
     */
    private @Nullable Integer port;
    /**
     * @return Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
     * 
     */
    private String type;
    /**
     * @return URL for HTTP probe type.
     * 
     */
    private @Nullable String url;

    private IpLoadBalancingHttpFarmProbe() {}
    /**
     * @return Force use of SSL (TLS)
     * 
     */
    public Optional<Boolean> forceSsl() {
        return Optional.ofNullable(this.forceSsl);
    }
    /**
     * @return probe interval, Value between 30 and 3600 seconds, default 30
     * 
     */
    public Optional<Integer> interval() {
        return Optional.ofNullable(this.interval);
    }
    /**
     * @return What to mach `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
     * 
     */
    public Optional<String> match() {
        return Optional.ofNullable(this.match);
    }
    /**
     * @return HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
     * 
     */
    public Optional<String> method() {
        return Optional.ofNullable(this.method);
    }
    /**
     * @return Negate probe result
     * 
     */
    public Optional<Boolean> negate() {
        return Optional.ofNullable(this.negate);
    }
    /**
     * @return Pattern to match against `match`
     * 
     */
    public Optional<String> pattern() {
        return Optional.ofNullable(this.pattern);
    }
    /**
     * @return Port for backends to recieve traffic on.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return URL for HTTP probe type.
     * 
     */
    public Optional<String> url() {
        return Optional.ofNullable(this.url);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IpLoadBalancingHttpFarmProbe defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean forceSsl;
        private @Nullable Integer interval;
        private @Nullable String match;
        private @Nullable String method;
        private @Nullable Boolean negate;
        private @Nullable String pattern;
        private @Nullable Integer port;
        private String type;
        private @Nullable String url;
        public Builder() {}
        public Builder(IpLoadBalancingHttpFarmProbe defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.forceSsl = defaults.forceSsl;
    	      this.interval = defaults.interval;
    	      this.match = defaults.match;
    	      this.method = defaults.method;
    	      this.negate = defaults.negate;
    	      this.pattern = defaults.pattern;
    	      this.port = defaults.port;
    	      this.type = defaults.type;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder forceSsl(@Nullable Boolean forceSsl) {
            this.forceSsl = forceSsl;
            return this;
        }
        @CustomType.Setter
        public Builder interval(@Nullable Integer interval) {
            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder match(@Nullable String match) {
            this.match = match;
            return this;
        }
        @CustomType.Setter
        public Builder method(@Nullable String method) {
            this.method = method;
            return this;
        }
        @CustomType.Setter
        public Builder negate(@Nullable Boolean negate) {
            this.negate = negate;
            return this;
        }
        @CustomType.Setter
        public Builder pattern(@Nullable String pattern) {
            this.pattern = pattern;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder url(@Nullable String url) {
            this.url = url;
            return this;
        }
        public IpLoadBalancingHttpFarmProbe build() {
            final var o = new IpLoadBalancingHttpFarmProbe();
            o.forceSsl = forceSsl;
            o.interval = interval;
            o.match = match;
            o.method = method;
            o.negate = negate;
            o.pattern = pattern;
            o.port = port;
            o.type = type;
            o.url = url;
            return o;
        }
    }
}
