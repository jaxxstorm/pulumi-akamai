// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy; evaluation rule
 *
 * Returns the action and the condition-exception information for a rule or set of rules being used in evaluation mode.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval-rules](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getevalrules)
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
 *     name: "Documentation",
 * });
 * const evalRule = configuration.then(configuration => akamai.getAppSecEvalRules({
 *     configId: configuration.configId,
 *     securityPolicyId: "gms1_134637",
 *     ruleId: 60029316,
 * }));
 * export const evalRuleAction = evalRule.then(evalRule => evalRule.evalRuleAction);
 * export const conditionException = evalRule.then(evalRule => evalRule.conditionException);
 * export const json = evalRule.then(evalRule => evalRule.json);
 * export const outputText = evalRule.then(evalRule => evalRule.outputText);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `evalRuleAction`. Action taken anytime the evaluation rule is triggered. Valid values are:
 *   - **alert**. Record the event,
 *   - **deny**. Reject the request.
 *   - **deny_custom_{custom_deny_id}**. The action defined by the custom deny is taken.
 *   - **none**. Take no action.
 * - `conditionException`. Conditions and exceptions associated with the rule.
 * - `json`. JSON-formatted list of the action and the condition-exception information for the rule. This output is only generated if the `ruleId` argument is included.
 * - `outputText`. Tabular report showing the rule action as well as Boolean values indicating whether conditions and exceptions have been configured for the rule.
 */
export function getAppSecEvalRules(args: GetAppSecEvalRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecEvalRulesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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
     * . Unique identifier of the security configuration running in evaluation mode.
     */
    configId: number;
    /**
     * . Unique identifier of the evaluation rule you want to return information for. If not included, information is returned for all your evaluation rules.
     */
    ruleId?: number;
    /**
     * . Unique identifier of the security policy associated with the evaluation rule.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecEvalRules.
 */
export interface GetAppSecEvalRulesResult {
    readonly conditionException: string;
    readonly configId: number;
    readonly evalRuleAction: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly outputText: string;
    readonly ruleId?: number;
    readonly securityPolicyId: string;
}

export function getAppSecEvalRulesOutput(args: GetAppSecEvalRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecEvalRulesResult> {
    return pulumi.output(args).apply(a => getAppSecEvalRules(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecEvalRules.
 */
export interface GetAppSecEvalRulesOutputArgs {
    /**
     * . Unique identifier of the security configuration running in evaluation mode.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the evaluation rule you want to return information for. If not included, information is returned for all your evaluation rules.
     */
    ruleId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the evaluation rule.
     */
    securityPolicyId: pulumi.Input<string>;
}
