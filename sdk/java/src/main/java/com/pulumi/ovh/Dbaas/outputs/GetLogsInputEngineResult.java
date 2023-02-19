// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLogsInputEngineResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Boolean isDeprecated;
    private String name;
    private String serviceName;
    private String version;

    private GetLogsInputEngineResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Boolean isDeprecated() {
        return this.isDeprecated;
    }
    public String name() {
        return this.name;
    }
    public String serviceName() {
        return this.serviceName;
    }
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLogsInputEngineResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Boolean isDeprecated;
        private String name;
        private String serviceName;
        private String version;
        public Builder() {}
        public Builder(GetLogsInputEngineResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.isDeprecated = defaults.isDeprecated;
    	      this.name = defaults.name;
    	      this.serviceName = defaults.serviceName;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder isDeprecated(Boolean isDeprecated) {
            this.isDeprecated = Objects.requireNonNull(isDeprecated);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetLogsInputEngineResult build() {
            final var o = new GetLogsInputEngineResult();
            o.id = id;
            o.isDeprecated = isDeprecated;
            o.name = name;
            o.serviceName = serviceName;
            o.version = version;
            return o;
        }
    }
}
