// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecEvalRuleActions` data source to retrieve the rules available for evaluation and their actions, or the action for a specific rule available for evaluation.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const configuration = akamai.getAppSecConfiguration({
 *     name: _var.security_configuration,
 * });
 * const ruleActions = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecEvalRuleActions({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 *     securityPolicyId: _var.security_policy_id,
 * }));
 * export const ruleActionsText = ruleActions.then(ruleActions => ruleActions.outputText);
 * export const ruleActionsJson = ruleActions.then(ruleActions => ruleActions.json);
 * const ruleActionAppSecEvalRuleActions = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecEvalRuleActions({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 *     securityPolicyId: _var.security_policy_id,
 *     ruleId: _var.rule_id,
 * }));
 * export const ruleAction = akamai_appsec_eval_rule_actions.rule_action.action;
 * export const ruleId = akamai_appsec_eval_rule_actions.rule_action.rule_id;
 * ```
 */
export function getAppSecEvalRuleActions(args: GetAppSecEvalRuleActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecEvalRuleActionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecEvalRuleActions:getAppSecEvalRuleActions", {
        "configId": args.configId,
        "ruleId": args.ruleId,
        "securityPolicyId": args.securityPolicyId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecEvalRuleActions.
 */
export interface GetAppSecEvalRuleActionsArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * The ID of a specific rule. If not supplied, information about all eval rules will be returned.
     */
    ruleId?: number;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId: string;
    /**
     * The version number of the security configuration to use.
     */
    version: number;
}

/**
 * A collection of values returned by getAppSecEvalRuleActions.
 */
export interface GetAppSecEvalRuleActionsResult {
    /**
     * The action configured for the given rule if a `ruleId` was specified.
     */
    readonly action: string;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted display of the ID and action for all rules in the security policy.
     */
    readonly json: string;
    /**
     * A tabular display of the ID and action for all rules in the security policy.
     */
    readonly outputText: string;
    readonly ruleId?: number;
    readonly securityPolicyId: string;
    readonly version: number;
}
