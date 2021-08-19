// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecEvalRules` data source to list the action and condition-exception information
 * for a rule or rules you want to evaluate.
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
 * const evalRule = configuration.then(configuration => akamai.getAppSecEvalRules({
 *     configId: configuration.configId,
 *     securityPolicyId: _var.security_policy_id,
 *     ruleId: _var.rule_id,
 * }));
 * export const evalRuleAction = akamai_appsec_eval_rules.eval_rule.eval_rule_action;
 * export const conditionException = akamai_appsec_eval_rules.eval_rule.condition_exception;
 * export const json = akamai_appsec_eval_rules.eval_rule.json;
 * export const outputText = akamai_appsec_eval_rules.eval_rule.output_text;
 * ```
 */
export function getAppSecEvalRules(args: GetAppSecEvalRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecEvalRulesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecEvalRules:getAppSecEvalRules", {
        "configId": args.configId,
        "ruleId": args.ruleId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecEvalRules.
 */
export interface GetAppSecEvalRulesArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * The ID of the rule to use. If not specified, information about all rules will be returned.
     */
    ruleId?: number;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecEvalRules.
 */
export interface GetAppSecEvalRulesResult {
    /**
     * The eval rule's conditions and exceptions.
     */
    readonly conditionException: string;
    readonly configId: number;
    /**
     * The eval rule's action, either `alert`, `deny`, or `none`.
     */
    readonly evalRuleAction: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted list of the action and condition-exception information for the specified eval rule.
     * This output is only generated if an eval rule is specified.
     */
    readonly json: string;
    /**
     * A tabular display showing, for the specified eval rule or rules, the rule action and boolean values
     * indicating whether conditions and exceptions are present.
     */
    readonly outputText: string;
    readonly ruleId?: number;
    readonly securityPolicyId: string;
}