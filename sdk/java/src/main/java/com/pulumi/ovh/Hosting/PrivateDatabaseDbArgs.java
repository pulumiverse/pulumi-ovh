// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class PrivateDatabaseDbArgs extends com.pulumi.resources.ResourceArgs {

    public static final PrivateDatabaseDbArgs Empty = new PrivateDatabaseDbArgs();

    /**
     * Name of your new database
     * 
     */
    @Import(name="databaseName", required=true)
    private Output<String> databaseName;

    /**
     * @return Name of your new database
     * 
     */
    public Output<String> databaseName() {
        return this.databaseName;
    }

    /**
     * The internal name of your private database.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your private database.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private PrivateDatabaseDbArgs() {}

    private PrivateDatabaseDbArgs(PrivateDatabaseDbArgs $) {
        this.databaseName = $.databaseName;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateDatabaseDbArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateDatabaseDbArgs $;

        public Builder() {
            $ = new PrivateDatabaseDbArgs();
        }

        public Builder(PrivateDatabaseDbArgs defaults) {
            $ = new PrivateDatabaseDbArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param databaseName Name of your new database
         * 
         * @return builder
         * 
         */
        public Builder databaseName(Output<String> databaseName) {
            $.databaseName = databaseName;
            return this;
        }

        /**
         * @param databaseName Name of your new database
         * 
         * @return builder
         * 
         */
        public Builder databaseName(String databaseName) {
            return databaseName(Output.of(databaseName));
        }

        /**
         * @param serviceName The internal name of your private database.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your private database.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public PrivateDatabaseDbArgs build() {
            $.databaseName = Objects.requireNonNull($.databaseName, "expected parameter 'databaseName' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
