// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetSshKeyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSshKeyArgs Empty = new GetSshKeyArgs();

    /**
     * The name of the SSH key.
     * 
     */
    @Import(name="keyName", required=true)
    private Output<String> keyName;

    /**
     * @return The name of the SSH key.
     * 
     */
    public Output<String> keyName() {
        return this.keyName;
    }

    private GetSshKeyArgs() {}

    private GetSshKeyArgs(GetSshKeyArgs $) {
        this.keyName = $.keyName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSshKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSshKeyArgs $;

        public Builder() {
            $ = new GetSshKeyArgs();
        }

        public Builder(GetSshKeyArgs defaults) {
            $ = new GetSshKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param keyName The name of the SSH key.
         * 
         * @return builder
         * 
         */
        public Builder keyName(Output<String> keyName) {
            $.keyName = keyName;
            return this;
        }

        /**
         * @param keyName The name of the SSH key.
         * 
         * @return builder
         * 
         */
        public Builder keyName(String keyName) {
            return keyName(Output.of(keyName));
        }

        public GetSshKeyArgs build() {
            $.keyName = Objects.requireNonNull($.keyName, "expected parameter 'keyName' to be non-null");
            return $;
        }
    }

}