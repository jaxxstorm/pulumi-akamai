// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecCustomRuleActions` data source to retrieve information about the actions defined for the custom rules, or a specific custom rule, associated with a specific security configuration and security policy.
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
 *     name: "Akamai Tools",
 * });
 * const customRuleActionsAppSecCustomRuleActions = configuration.then(configuration => akamai.getAppSecCustomRuleActions({
 *     configId: configuration.configId,
 *     securityPolicyId: "crAP_75829",
 * }));
 * export const customRuleActions = customRuleActionsAppSecCustomRuleActions.then(customRuleActionsAppSecCustomRuleActions => customRuleActionsAppSecCustomRuleActions.outputText);
 * ```
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
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * A specific custom rule for which to retrieve information. If not supplied, information about all custom rules will be returned.
     */
    customRuleId?: number;
    /**
     * The ID of the security policy to use
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
    /**
     * A tabular display showing the ID, name, and action of all custom rules, or of the specific custom rule, associated with the specified security configuration, version and security policy.
     */
    readonly outputText: string;
    readonly securityPolicyId: string;
}
