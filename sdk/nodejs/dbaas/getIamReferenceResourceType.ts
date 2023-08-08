// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to list all the IAM resource types.
 *
 * ## Important
 *
 * > Using this resource requires that the account is enrolled in the OVHcloud [IAM beta](https://labs.ovhcloud.com/en/iam/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const types = ovh.Dbaas.getIamReferenceResourceType({});
 * ```
 */
export function getIamReferenceResourceType(opts?: pulumi.InvokeOptions): Promise<GetIamReferenceResourceTypeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Dbaas/getIamReferenceResourceType:getIamReferenceResourceType", {
    }, opts);
}

/**
 * A collection of values returned by getIamReferenceResourceType.
 */
export interface GetIamReferenceResourceTypeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of resource types
     */
    readonly types: string[];
}
