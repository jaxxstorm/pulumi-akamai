// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getPropertyRules` data source to query and retrieve the rule tree of
 * an existing property version. This data source lets you search across the contracts
 * and groups you have access to.
 *
 * ## Basic usage
 *
 * This example returns the rule tree for version 3 of a property based on the selected contract and group:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 *
 * export const propertyMatch = data.akamai_property_rules["my-example"];
 * ```
 *
 * ## Argument reference
 *
 * This data source supports these arguments:
 *
 * * `contractId` - (Required) A contract's unique ID, including the `ctr_` prefix.
 * * `groupId` - (Required) A group's unique ID, including the `grp_` prefix.
 * * `propertyId` - (Required) A property's unique ID, including the `prp_` prefix.
 * * `version` - (Optional) The version to return. Returns the latest version by default.
 *
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `rules` - A JSON-encoded rule tree for the property.
 * * `errors` - A list of validation errors for the rule tree object returned. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the Property Manager API documentation.
 */
export function getPropertyRules(args: GetPropertyRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetPropertyRulesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getPropertyRules:getPropertyRules", {
        "contractId": args.contractId,
        "groupId": args.groupId,
        "propertyId": args.propertyId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getPropertyRules.
 */
export interface GetPropertyRulesArgs {
    contractId?: string;
    groupId?: string;
    propertyId: string;
    version?: number;
}

/**
 * A collection of values returned by getPropertyRules.
 */
export interface GetPropertyRulesResult {
    readonly contractId: string;
    readonly errors: string;
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly propertyId: string;
    readonly rules: string;
    readonly version: number;
}
