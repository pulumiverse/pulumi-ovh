// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class S3PolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final S3PolicyArgs Empty = new S3PolicyArgs();

    /**
     * The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
     * 
     */
    @Import(name="policy", required=true)
    private Output<String> policy;

    /**
     * @return The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * The ID of a public cloud project&#39;s user.
     * 
     */
    @Import(name="userId", required=true)
    private Output<String> userId;

    /**
     * @return The ID of a public cloud project&#39;s user.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    private S3PolicyArgs() {}

    private S3PolicyArgs(S3PolicyArgs $) {
        this.policy = $.policy;
        this.serviceName = $.serviceName;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(S3PolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private S3PolicyArgs $;

        public Builder() {
            $ = new S3PolicyArgs();
        }

        public Builder(S3PolicyArgs defaults) {
            $ = new S3PolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param policy The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
         * 
         * @return builder
         * 
         */
        public Builder policy(Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param userId The ID of a public cloud project&#39;s user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of a public cloud project&#39;s user.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        public S3PolicyArgs build() {
            $.policy = Objects.requireNonNull($.policy, "expected parameter 'policy' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            $.userId = Objects.requireNonNull($.userId, "expected parameter 'userId' to be non-null");
            return $;
        }
    }

}
