// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DedicatedServerRebootTaskState extends com.pulumi.resources.ResourceArgs {

    public static final DedicatedServerRebootTaskState Empty = new DedicatedServerRebootTaskState();

    /**
     * Details of this task. (should be `Reboot asked`)
     * 
     */
    @Import(name="comment")
    private @Nullable Output<String> comment;

    /**
     * @return Details of this task. (should be `Reboot asked`)
     * 
     */
    public Optional<Output<String>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * Completion date in RFC3339 format.
     * 
     */
    @Import(name="doneDate")
    private @Nullable Output<String> doneDate;

    /**
     * @return Completion date in RFC3339 format.
     * 
     */
    public Optional<Output<String>> doneDate() {
        return Optional.ofNullable(this.doneDate);
    }

    /**
     * Function name (should be `hardReboot`).
     * 
     */
    @Import(name="function")
    private @Nullable Output<String> function;

    /**
     * @return Function name (should be `hardReboot`).
     * 
     */
    public Optional<Output<String>> function() {
        return Optional.ofNullable(this.function);
    }

    /**
     * List of values traccked to trigger reboot, used also to form implicit dependencies
     * 
     */
    @Import(name="keepers")
    private @Nullable Output<List<String>> keepers;

    /**
     * @return List of values traccked to trigger reboot, used also to form implicit dependencies
     * 
     */
    public Optional<Output<List<String>>> keepers() {
        return Optional.ofNullable(this.keepers);
    }

    /**
     * Last update in RFC3339 format.
     * 
     */
    @Import(name="lastUpdate")
    private @Nullable Output<String> lastUpdate;

    /**
     * @return Last update in RFC3339 format.
     * 
     */
    public Optional<Output<String>> lastUpdate() {
        return Optional.ofNullable(this.lastUpdate);
    }

    /**
     * The service_name of your dedicated server.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The service_name of your dedicated server.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Task creation date in RFC3339 format.
     * 
     */
    @Import(name="startDate")
    private @Nullable Output<String> startDate;

    /**
     * @return Task creation date in RFC3339 format.
     * 
     */
    public Optional<Output<String>> startDate() {
        return Optional.ofNullable(this.startDate);
    }

    /**
     * Task status (should be `done`)
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Task status (should be `done`)
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private DedicatedServerRebootTaskState() {}

    private DedicatedServerRebootTaskState(DedicatedServerRebootTaskState $) {
        this.comment = $.comment;
        this.doneDate = $.doneDate;
        this.function = $.function;
        this.keepers = $.keepers;
        this.lastUpdate = $.lastUpdate;
        this.serviceName = $.serviceName;
        this.startDate = $.startDate;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DedicatedServerRebootTaskState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DedicatedServerRebootTaskState $;

        public Builder() {
            $ = new DedicatedServerRebootTaskState();
        }

        public Builder(DedicatedServerRebootTaskState defaults) {
            $ = new DedicatedServerRebootTaskState(Objects.requireNonNull(defaults));
        }

        /**
         * @param comment Details of this task. (should be `Reboot asked`)
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<String> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment Details of this task. (should be `Reboot asked`)
         * 
         * @return builder
         * 
         */
        public Builder comment(String comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param doneDate Completion date in RFC3339 format.
         * 
         * @return builder
         * 
         */
        public Builder doneDate(@Nullable Output<String> doneDate) {
            $.doneDate = doneDate;
            return this;
        }

        /**
         * @param doneDate Completion date in RFC3339 format.
         * 
         * @return builder
         * 
         */
        public Builder doneDate(String doneDate) {
            return doneDate(Output.of(doneDate));
        }

        /**
         * @param function Function name (should be `hardReboot`).
         * 
         * @return builder
         * 
         */
        public Builder function(@Nullable Output<String> function) {
            $.function = function;
            return this;
        }

        /**
         * @param function Function name (should be `hardReboot`).
         * 
         * @return builder
         * 
         */
        public Builder function(String function) {
            return function(Output.of(function));
        }

        /**
         * @param keepers List of values traccked to trigger reboot, used also to form implicit dependencies
         * 
         * @return builder
         * 
         */
        public Builder keepers(@Nullable Output<List<String>> keepers) {
            $.keepers = keepers;
            return this;
        }

        /**
         * @param keepers List of values traccked to trigger reboot, used also to form implicit dependencies
         * 
         * @return builder
         * 
         */
        public Builder keepers(List<String> keepers) {
            return keepers(Output.of(keepers));
        }

        /**
         * @param keepers List of values traccked to trigger reboot, used also to form implicit dependencies
         * 
         * @return builder
         * 
         */
        public Builder keepers(String... keepers) {
            return keepers(List.of(keepers));
        }

        /**
         * @param lastUpdate Last update in RFC3339 format.
         * 
         * @return builder
         * 
         */
        public Builder lastUpdate(@Nullable Output<String> lastUpdate) {
            $.lastUpdate = lastUpdate;
            return this;
        }

        /**
         * @param lastUpdate Last update in RFC3339 format.
         * 
         * @return builder
         * 
         */
        public Builder lastUpdate(String lastUpdate) {
            return lastUpdate(Output.of(lastUpdate));
        }

        /**
         * @param serviceName The service_name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service_name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param startDate Task creation date in RFC3339 format.
         * 
         * @return builder
         * 
         */
        public Builder startDate(@Nullable Output<String> startDate) {
            $.startDate = startDate;
            return this;
        }

        /**
         * @param startDate Task creation date in RFC3339 format.
         * 
         * @return builder
         * 
         */
        public Builder startDate(String startDate) {
            return startDate(Output.of(startDate));
        }

        /**
         * @param status Task status (should be `done`)
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Task status (should be `done`)
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public DedicatedServerRebootTaskState build() {
            return $;
        }
    }

}
