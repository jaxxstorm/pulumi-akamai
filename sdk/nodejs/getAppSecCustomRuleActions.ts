// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy; custom rule
 *
 * Retrieve information about the actions defined for your custom rules. Custom rules are rules that you create yourself: these rules aren't part of Akamai's Kona Rule Set.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/custom-rules/{ruleId}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getcustomruleactions)
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
 * const customRuleActionsAppSecCustomRuleActions = configuration.then(configuration => akamai.getAppSecCustomRuleActions({
 *     configId: configuration.configId,
 *     securityPolicyId: "gms1_134637",
 * }));
 * export const customRuleActions = customRuleActionsAppSecCustomRuleActions.then(customRuleActionsAppSecCustomRuleActions => customRuleActionsAppSecCustomRuleActions.outputText);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `outputText`. Tabular report showing the ID, name, and action of the custom rules.
 */
export function getAppSecCustomRuleActions(args: GetAppSecCustomRuleActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecCustomRuleActionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecCustomRuleActions:getAppSecCustomRuleActions", {
        "configId": args.configId,
        "customRuleId": args.customRuleId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecCustomRuleActions.
 */
export interface GetAppSecCustomRuleActionsArgs {
    /**
     * . Unique identifier of the security configuration associated with the custom rules.
     */
    configId: number;
    /**
     * . Unique identifier of the custom rule you want to return information for. If not included, action information is returned for all your custom rules.
     */
    customRuleId?: number;
    /**
     * . Unique identifier of the security policy associated with the custom rules.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecCustomRuleActions.
 */
export interface GetAppSecCustomRuleActionsResult {
    readonly configId: number;
    readonly customRuleId?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputText: string;
    readonly securityPolicyId: string;
}

export function getAppSecCustomRuleActionsOutput(args: GetAppSecCustomRuleActionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecCustomRuleActionsResult> {
    return pulumi.output(args).apply(a => getAppSecCustomRuleActions(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecCustomRuleActions.
 */
export interface GetAppSecCustomRuleActionsOutputArgs {
    /**
     * . Unique identifier of the security configuration associated with the custom rules.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the custom rule you want to return information for. If not included, action information is returned for all your custom rules.
     */
    customRuleId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the custom rules.
     */
    securityPolicyId: pulumi.Input<string>;
}
