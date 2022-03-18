// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
 *
 * Use the `akamai.getCloudletsForwardRewriteMatchRule` data source to build a match rule JSON object for the Forward Rewrite Cloudlet.
 *
 * ## Basic usage
 *
 * This example returns the JSON-encoded rules for the Forward Rewrite Cloudlet:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const example = pulumi.output(akamai.getCloudletsForwardRewriteMatchRule({
 *     matchRules: [{
 *         forwardSettings: {
 *             originId: "1234",
 *             pathAndQs: "/path",
 *             useIncomingQueryString: true,
 *         },
 *         matches: [{
 *             caseSensitive: false,
 *             checkIps: "CONNECTING_IP XFF_HEADERS",
 *             matchOperator: "equals",
 *             matchType: "clientip",
 *             matchValue: "192.0.2.0",
 *             negate: false,
 *         }],
 *         name: "rule",
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
export function getCloudletsForwardRewriteMatchRule(args?: GetCloudletsForwardRewriteMatchRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudletsForwardRewriteMatchRuleResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getCloudletsForwardRewriteMatchRule:getCloudletsForwardRewriteMatchRule", {
        "matchRules": args.matchRules,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudletsForwardRewriteMatchRule.
 */
export interface GetCloudletsForwardRewriteMatchRuleArgs {
    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     */
    matchRules?: inputs.GetCloudletsForwardRewriteMatchRuleMatchRule[];
}

/**
 * A collection of values returned by getCloudletsForwardRewriteMatchRule.
 */
export interface GetCloudletsForwardRewriteMatchRuleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly matchRules?: outputs.GetCloudletsForwardRewriteMatchRuleMatchRule[];
}

export function getCloudletsForwardRewriteMatchRuleOutput(args?: GetCloudletsForwardRewriteMatchRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudletsForwardRewriteMatchRuleResult> {
    return pulumi.output(args).apply(a => getCloudletsForwardRewriteMatchRule(a, opts))
}

/**
 * A collection of arguments for invoking getCloudletsForwardRewriteMatchRule.
 */
export interface GetCloudletsForwardRewriteMatchRuleOutputArgs {
    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     */
    matchRules?: pulumi.Input<pulumi.Input<inputs.GetCloudletsForwardRewriteMatchRuleMatchRuleArgs>[]>;
}
