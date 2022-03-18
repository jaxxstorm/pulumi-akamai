// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
 *
 * Use the `akamai.getCloudletsApiPrioritizationMatchRule` data source to build a match rule JSON object for the API Prioritization Cloudlet.
 *
 * ## Basic usage
 *
 * This example returns the JSON-encoded rules for the API Prioritization Cloudlet:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const example = pulumi.output(akamai.getCloudletsApiPrioritizationMatchRule({
 *     matchRules: [{
 *         disabled: false,
 *         end: 1645037845,
 *         matchUrl: "example.com",
 *         matches: [{
 *             caseSensitive: true,
 *             matchOperator: "equals",
 *             matchType: "method",
 *             negate: false,
 *             objectMatchValues: [{
 *                 type: "simple",
 *                 values: ["POST"],
 *             }],
 *         }],
 *         name: "rule",
 *         passThroughPercent: 10,
 *         start: 1644865045,
 *     }],
 * }));
 * ```
 *
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `type` - The type of Cloudlet the rule is for.
 * * `json` - A `matchRules` JSON structure generated from the API schema that defines the rules for this policy.
 */
export function getCloudletsApiPrioritizationMatchRule(args?: GetCloudletsApiPrioritizationMatchRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudletsApiPrioritizationMatchRuleResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getCloudletsApiPrioritizationMatchRule:getCloudletsApiPrioritizationMatchRule", {
        "matchRules": args.matchRules,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudletsApiPrioritizationMatchRule.
 */
export interface GetCloudletsApiPrioritizationMatchRuleArgs {
    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     */
    matchRules?: inputs.GetCloudletsApiPrioritizationMatchRuleMatchRule[];
}

/**
 * A collection of values returned by getCloudletsApiPrioritizationMatchRule.
 */
export interface GetCloudletsApiPrioritizationMatchRuleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly matchRules?: outputs.GetCloudletsApiPrioritizationMatchRuleMatchRule[];
}

export function getCloudletsApiPrioritizationMatchRuleOutput(args?: GetCloudletsApiPrioritizationMatchRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudletsApiPrioritizationMatchRuleResult> {
    return pulumi.output(args).apply(a => getCloudletsApiPrioritizationMatchRule(a, opts))
}

/**
 * A collection of arguments for invoking getCloudletsApiPrioritizationMatchRule.
 */
export interface GetCloudletsApiPrioritizationMatchRuleOutputArgs {
    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     */
    matchRules?: pulumi.Input<pulumi.Input<inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleArgs>[]>;
}
