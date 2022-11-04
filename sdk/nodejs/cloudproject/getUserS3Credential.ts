// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve the Secret Access Key of an Access Key ID associated with a public cloud project's user.
 */
export function getUserS3Credential(args: GetUserS3CredentialArgs, opts?: pulumi.InvokeOptions): Promise<GetUserS3CredentialResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("ovh:CloudProject/getUserS3Credential:getUserS3Credential", {
        "accessKeyId": args.accessKeyId,
        "serviceName": args.serviceName,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserS3Credential.
 */
export interface GetUserS3CredentialArgs {
    /**
     * the Access Key ID
     */
    accessKeyId: string;
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
    /**
     * The ID of a public cloud project's user.
     */
    userId: string;
}

/**
 * A collection of values returned by getUserS3Credential.
 */
export interface GetUserS3CredentialResult {
    readonly accessKeyId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Sensitive) the Secret Access Key
     */
    readonly secretAccessKey: string;
    readonly serviceName: string;
    readonly userId: string;
}

export function getUserS3CredentialOutput(args: GetUserS3CredentialOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserS3CredentialResult> {
    return pulumi.output(args).apply(a => getUserS3Credential(a, opts))
}

/**
 * A collection of arguments for invoking getUserS3Credential.
 */
export interface GetUserS3CredentialOutputArgs {
    /**
     * the Access Key ID
     */
    accessKeyId: pulumi.Input<string>;
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