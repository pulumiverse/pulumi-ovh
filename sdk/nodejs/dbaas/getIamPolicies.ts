// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to list the existing IAM policies of an account.
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
 * const myPolicies = ovh.Dbaas.getIamPolicies({});
 * ```
 */
export function getIamPolicies(opts?: pulumi.InvokeOptions): Promise<GetIamPoliciesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Dbaas/getIamPolicies:getIamPolicies", {
    }, opts);
}

/**
 * A collection of values returned by getIamPolicies.
 */
export interface GetIamPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of the policies IDs.
     */
    readonly policies: string[];
}
