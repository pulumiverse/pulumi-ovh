// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIpxeScriptResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String name;
    /**
     * @return The content of the script.
     * 
     */
    private String script;

    private GetIpxeScriptResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The content of the script.
     * 
     */
    public String script() {
        return this.script;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpxeScriptResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String name;
        private String script;
        public Builder() {}
        public Builder(GetIpxeScriptResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.script = defaults.script;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder script(String script) {
            this.script = Objects.requireNonNull(script);
            return this;
        }
        public GetIpxeScriptResult build() {
            final var o = new GetIpxeScriptResult();
            o.id = id;
            o.name = name;
            o.script = script;
            return o;
        }
    }
}