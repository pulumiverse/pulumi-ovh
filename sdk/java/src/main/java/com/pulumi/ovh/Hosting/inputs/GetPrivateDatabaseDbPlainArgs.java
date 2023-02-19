// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPrivateDatabaseDbPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPrivateDatabaseDbPlainArgs Empty = new GetPrivateDatabaseDbPlainArgs();

    /**
     * Database name
     * 
     */
    @Import(name="databaseName", required=true)
    private String databaseName;

    /**
     * @return Database name
     * 
     */
    public String databaseName() {
        return this.databaseName;
    }

    /**
     * The internal name of your private database
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The internal name of your private database
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetPrivateDatabaseDbPlainArgs() {}

    private GetPrivateDatabaseDbPlainArgs(GetPrivateDatabaseDbPlainArgs $) {
        this.databaseName = $.databaseName;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPrivateDatabaseDbPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPrivateDatabaseDbPlainArgs $;

        public Builder() {
            $ = new GetPrivateDatabaseDbPlainArgs();
        }

        public Builder(GetPrivateDatabaseDbPlainArgs defaults) {
            $ = new GetPrivateDatabaseDbPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param databaseName Database name
         * 
         * @return builder
         * 
         */
        public Builder databaseName(String databaseName) {
            $.databaseName = databaseName;
            return this;
        }

        /**
         * @param serviceName The internal name of your private database
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetPrivateDatabaseDbPlainArgs build() {
            $.databaseName = Objects.requireNonNull($.databaseName, "expected parameter 'databaseName' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
