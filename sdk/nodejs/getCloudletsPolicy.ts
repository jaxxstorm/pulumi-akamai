// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.CloudletsPolicy` data source to list details about a policy with and its specified version, or latest if not specified.
 *
 * ## Basic usage
 *
 * This example returns the policy details based on the policy ID and optionally, a version:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const example = pulumi.output(akamai.getCloudletsPolicy({
 *     policyId: 1234,
 *     version: 1,
 * }));
 * ```
 *
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `groupId` - Defines the group association for the policy. You must have edit privileges for the group.
 * * `name` - The unique name of the policy.
 * * `apiVersion` - The specific version of the Cloudlets API.
 * * `cloudletId` - A unique identifier that corresponds to a Cloudlets policy type. Enter `0` for Edge Redirector, `1` for Visitor Prioritization, `3` for Forward Rewrite, `5` for API Prioritization, `7` for Phased Release, or `9` for Application Load Balancer.
 * * `cloudletCode` - The two- or three- character code for the type of Cloudlet, `ALB` for Application Load Balancer, `AP` for API Prioritization, `CD` for Phased Release, `ER` for Edge Redirector, `FR` for Forward Rewrite, and `VP` for Visitor Prioritization.
 * * `revisionId` - A unique identifier given to every policy version update.
 * * `description` - The description of this specific policy.
 * * `versionDescription` - The description of this specific policy version.
 * * `rulesLocked` - Whether editing `matchRules` for the Cloudlet policy version is blocked.
 * * `matchRules`- A JSON structure that defines the rules for this policy.
 * * `matchRuleFormat` - The format of the Cloudlet-specific `matchRules`.
 * * `warnings` - A JSON encoded list of warnings.
 * * `activations` - A list of of current policy activation information, including:
 *   * `apiVersion` - The specific version of the Cloudlets API.
 *   * `network` - The network, either `staging` or `prod` on which a property or a Cloudlets policy has been activated.
 *   * `policyInfo` - A list of Cloudlet policy information, including:
 *       * `policyId` - An integer identifier that is associated with all versions of a policy.
 *       * `name` - The name of the policy.
 *       * `version` - The version number of the policy.
 *       * `status` - The activation status for the policy. Values include the following: `inactive` where the policy version has not been activated. No active property versions reference this policy. `active` where the policy version is currently active (published) and its associated property version is also active. `deactivated` where the policy version was previously activated but it has been superseded by a more recent activation of another policy version. `pending` where the policy version is proceeding through the activation workflow. `failed` where the policy version activation workflow has failed.
 *       * `statusDetail` - Information about the status of an activation operation. This field is not returned when it has no value.
 *       * `activatedBy` - The name of the user who activated the policy.
 *       * `activationDate` - The date on which the policy was activated in milliseconds since epoch.
 *   * `propertyInfo` A list of Cloudlet property information, including:
 *       * `name` - The name of the property.
 *       * `version` - The version number of the activated property.
 *       * `groupId` - Defines the group association for the policy or property. If returns `0`, the policy is not tied to a group and in effect appears in all groups for the account. You must have edit privileges for the group.
 *       * `status` - The activation status for the property. Values include the following: `inactive` where the policy version has not been activated. No active property versions reference this policy. `active` where the policy version is currently active (published) and its associated property version is also active. `deactivated` where the policy version was previously activated but it has been superseded by a more recent activation of another policy version. `pending` where the policy version is proceeding through the activation workflow. `failed` where the policy version activation workflow has failed.
 *       * `activatedBy` - The name of the user who activated the property.
 *       * `activationDate` - The date on which the property was activated in milliseconds since epoch.
 */
export function getCloudletsPolicy(args: GetCloudletsPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudletsPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getCloudletsPolicy:getCloudletsPolicy", {
        "policyId": args.policyId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudletsPolicy.
 */
export interface GetCloudletsPolicyArgs {
    /**
     * - (Required) An integer identifier that is associated with all versions of a policy.
     */
    policyId: number;
    /**
     * - (Optional) The version number of a policy.
     */
    version?: number;
}

/**
 * A collection of values returned by getCloudletsPolicy.
 */
export interface GetCloudletsPolicyResult {
    readonly activations: outputs.GetCloudletsPolicyActivation[];
    readonly apiVersion: string;
    readonly cloudletCode: string;
    readonly cloudletId: number;
    readonly description: string;
    readonly groupId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly matchRuleFormat: string;
    readonly matchRules: string;
    readonly name: string;
    readonly policyId: number;
    readonly revisionId: number;
    readonly rulesLocked: boolean;
    readonly version?: number;
    readonly versionDescription: string;
    readonly warnings: string;
}

export function getCloudletsPolicyOutput(args: GetCloudletsPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudletsPolicyResult> {
    return pulumi.output(args).apply(a => getCloudletsPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getCloudletsPolicy.
 */
export interface GetCloudletsPolicyOutputArgs {
    /**
     * - (Required) An integer identifier that is associated with all versions of a policy.
     */
    policyId: pulumi.Input<number>;
    /**
     * - (Optional) The version number of a policy.
     */
    version?: pulumi.Input<number>;
}
