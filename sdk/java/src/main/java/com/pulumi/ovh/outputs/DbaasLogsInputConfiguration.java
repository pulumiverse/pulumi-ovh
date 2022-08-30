// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.outputs.DbaasLogsInputConfigurationFlowgger;
import com.pulumi.ovh.outputs.DbaasLogsInputConfigurationLogstash;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DbaasLogsInputConfiguration {
    /**
     * @return Flowgger configuration
     * 
     */
    private @Nullable DbaasLogsInputConfigurationFlowgger flowgger;
    /**
     * @return Logstash configuration
     * 
     */
    private @Nullable DbaasLogsInputConfigurationLogstash logstash;

    private DbaasLogsInputConfiguration() {}
    /**
     * @return Flowgger configuration
     * 
     */
    public Optional<DbaasLogsInputConfigurationFlowgger> flowgger() {
        return Optional.ofNullable(this.flowgger);
    }
    /**
     * @return Logstash configuration
     * 
     */
    public Optional<DbaasLogsInputConfigurationLogstash> logstash() {
        return Optional.ofNullable(this.logstash);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DbaasLogsInputConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable DbaasLogsInputConfigurationFlowgger flowgger;
        private @Nullable DbaasLogsInputConfigurationLogstash logstash;
        public Builder() {}
        public Builder(DbaasLogsInputConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.flowgger = defaults.flowgger;
    	      this.logstash = defaults.logstash;
        }

        @CustomType.Setter
        public Builder flowgger(@Nullable DbaasLogsInputConfigurationFlowgger flowgger) {
            this.flowgger = flowgger;
            return this;
        }
        @CustomType.Setter
        public Builder logstash(@Nullable DbaasLogsInputConfigurationLogstash logstash) {
            this.logstash = logstash;
            return this;
        }
        public DbaasLogsInputConfiguration build() {
            final var o = new DbaasLogsInputConfiguration();
            o.flowgger = flowgger;
            o.logstash = logstash;
            return o;
        }
    }
}
