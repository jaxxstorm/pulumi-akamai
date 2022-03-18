// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getGroups` data source to list groups associated with the [EdgeGrid API client token](https://techdocs.akamai.com/developer/docs/authenticate-with-edgegrid) you're using.
 *
 * ## Basic usage
 *
 * Return groups associated with the EdgeGrid API client token you're using:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const my-example = akamai.getGroups({});
 * export const propertyMatch = my_example;
 * ```
 *
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `groups` - A list of supported groups, with the following attributes:
 *   * `groupId` - A group's unique ID, including the `grp_` prefix.
 *   * `groupName` - The name of the group.
 *   * `parentGroupId` - The ID of the parent group, if applicable.
 *   * `contractIds` - An array of strings listing the contract IDs for each group.
 */
export function getGroups(opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getGroups:getGroups", {
    }, opts);
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    readonly groups: outputs.GetGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
