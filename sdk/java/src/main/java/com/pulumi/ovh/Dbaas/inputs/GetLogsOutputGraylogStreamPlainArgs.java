// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetLogsOutputGraylogStreamPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLogsOutputGraylogStreamPlainArgs Empty = new GetLogsOutputGraylogStreamPlainArgs();

    /**
     * The service name
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The service name
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    /**
     * Stream description
     * 
     */
    @Import(name="title", required=true)
    private String title;

    /**
     * @return Stream description
     * 
     */
    public String title() {
        return this.title;
    }

    private GetLogsOutputGraylogStreamPlainArgs() {}

    private GetLogsOutputGraylogStreamPlainArgs(GetLogsOutputGraylogStreamPlainArgs $) {
        this.serviceName = $.serviceName;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLogsOutputGraylogStreamPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLogsOutputGraylogStreamPlainArgs $;

        public Builder() {
            $ = new GetLogsOutputGraylogStreamPlainArgs();
        }

        public Builder(GetLogsOutputGraylogStreamPlainArgs defaults) {
            $ = new GetLogsOutputGraylogStreamPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param title Stream description
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            $.title = title;
            return this;
        }

        public GetLogsOutputGraylogStreamPlainArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            $.title = Objects.requireNonNull($.title, "expected parameter 'title' to be non-null");
            return $;
        }
    }

}