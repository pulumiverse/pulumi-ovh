// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an S3 Credential for a user in a public cloud project.
 *
 * ## Import
 *
 * OVHcloud User S3 Credentials can be imported using the `service_name`, `user_id` and `access_key_id` of the credential, separated by "/" E.g., bash
 *
 * ```sh
 *  $ pulumi import ovh:CloudProject/s3Credential:S3Credential s3_credential service_name/user_id/access_key_id
 * ```
 */
export class S3Credential extends pulumi.CustomResource {
    /**
     * Get an existing S3Credential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: S3CredentialState, opts?: pulumi.CustomResourceOptions): S3Credential {
        return new S3Credential(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/s3Credential:S3Credential';

    /**
     * Returns true if the given object is an instance of S3Credential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is S3Credential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === S3Credential.__pulumiType;
    }

    /**
     * the Access Key ID
     */
    public /*out*/ readonly accessKeyId!: pulumi.Output<string>;
    public /*out*/ readonly internalUserId!: pulumi.Output<string>;
    /**
     * (Sensitive) the Secret Access Key
     */
    public /*out*/ readonly secretAccessKey!: pulumi.Output<string>;
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * The ID of a public cloud project's user.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a S3Credential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: S3CredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: S3CredentialArgs | S3CredentialState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as S3CredentialState | undefined;
            resourceInputs["accessKeyId"] = state ? state.accessKeyId : undefined;
            resourceInputs["internalUserId"] = state ? state.internalUserId : undefined;
            resourceInputs["secretAccessKey"] = state ? state.secretAccessKey : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as S3CredentialArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["accessKeyId"] = undefined /*out*/;
            resourceInputs["internalUserId"] = undefined /*out*/;
            resourceInputs["secretAccessKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secretAccessKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(S3Credential.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering S3Credential resources.
 */
export interface S3CredentialState {
    /**
     * the Access Key ID
     */
    accessKeyId?: pulumi.Input<string>;
    internalUserId?: pulumi.Input<string>;
    /**
     * (Sensitive) the Secret Access Key
     */
    secretAccessKey?: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The ID of a public cloud project's user.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a S3Credential resource.
 */
export interface S3CredentialArgs {
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
    /**
     * The ID of a public cloud project's user.
     */
    userId: pulumi.Input<string>;
}
