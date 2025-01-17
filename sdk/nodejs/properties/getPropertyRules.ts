// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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
 * import * as akamai from "@pulumi/akamai";
 *
 * const my-example = akamai.getPropertyRules({
 *     propertyId: "prp_123",
 *     groupId: "grp_12345",
 *     contractId: "ctr_1-AB123",
 *     version: 3,
 * });
 * export const propertyMatch = my_example;
 * ```
 *
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `ruleFormat` - The rule tree version used. Property rule objects are versioned infrequently, and are known as rule formats. See [Rule format schemas](https://techdocs.akamai.com/property-mgr/reference/rule-format-schemas) to learn more.
 * * `rules` - A JSON-encoded rule tree for the property.
 * * `errors` - A list of validation errors for the rule tree object returned. For more information see [Errors](https://techdocs.akamai.com/property-mgr/reference/api-errors) in the Property Manager API documentation.
 */
/** @deprecated akamai.properties.getPropertyRules has been deprecated in favor of akamai.getPropertyRules */
export function getPropertyRules(args: GetPropertyRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetPropertyRulesResult> {
    pulumi.log.warn("getPropertyRules is deprecated: akamai.properties.getPropertyRules has been deprecated in favor of akamai.getPropertyRules")
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:properties/getPropertyRules:getPropertyRules", {
        "contractId": args.contractId,
        "groupId": args.groupId,
        "propertyId": args.propertyId,
        "ruleFormat": args.ruleFormat,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getPropertyRules.
 */
export interface GetPropertyRulesArgs {
    /**
     * - (Required) A contract's unique ID, including the `ctr_` prefix.
     */
    contractId?: string;
    /**
     * - (Required) A group's unique ID, including the `grp_` prefix.
     */
    groupId?: string;
    /**
     * - (Required) A property's unique ID, including the `prp_` prefix.
     */
    propertyId: string;
    ruleFormat?: string;
    /**
     * - (Optional) The version to return. Returns the latest version by default.
     */
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
    readonly ruleFormat?: string;
    readonly rules: string;
    readonly version: number;
}

export function getPropertyRulesOutput(args: GetPropertyRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPropertyRulesResult> {
    return pulumi.output(args).apply(a => getPropertyRules(a, opts))
}

/**
 * A collection of arguments for invoking getPropertyRules.
 */
export interface GetPropertyRulesOutputArgs {
    /**
     * - (Required) A contract's unique ID, including the `ctr_` prefix.
     */
    contractId?: pulumi.Input<string>;
    /**
     * - (Required) A group's unique ID, including the `grp_` prefix.
     */
    groupId?: pulumi.Input<string>;
    /**
     * - (Required) A property's unique ID, including the `prp_` prefix.
     */
    propertyId: pulumi.Input<string>;
    ruleFormat?: pulumi.Input<string>;
    /**
     * - (Optional) The version to return. Returns the latest version by default.
     */
    version?: pulumi.Input<number>;
}
