// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecWafMode` data source to retrieve the mode that indicates how the WAF rules of the given security configuration and security policy will be updated.
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
 * const wafMode = configuration.then(configuration => akamai.getAppSecWafMode({
 *     configId: configuration.configId,
 *     securityPolicyId: _var.policy_id,
 * }));
 * export const wafModeMode = wafMode.then(wafMode => wafMode.mode);
 * export const wafModeCurrentRuleset = wafMode.then(wafMode => wafMode.currentRuleset);
 * export const wafModeEvalStatus = wafMode.then(wafMode => wafMode.evalStatus);
 * export const wafModeEvalRuleset = wafMode.then(wafMode => wafMode.evalRuleset);
 * export const wafModeEvalExpirationDate = wafMode.then(wafMode => wafMode.evalExpirationDate);
 * export const wafModeText = wafMode.then(wafMode => wafMode.outputText);
 * export const wafModeJson = wafMode.then(wafMode => wafMode.json);
 * ```
 */
export function getAppSecWafMode(args: GetAppSecWafModeArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecWafModeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecWafMode:getAppSecWafMode", {
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecWafMode.
 */
export interface GetAppSecWafModeArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecWafMode.
 */
export interface GetAppSecWafModeResult {
    readonly configId: number;
    /**
     * The current rule set version and the ISO 8601 date the rule set version was introduced; this date acts like a version number.
     */
    readonly currentRuleset: string;
    /**
     * The ISO 8601 time stamp when the evaluation is expiring. This value only appears when `eval` is set to "enabled".
     */
    readonly evalExpirationDate: string;
    /**
     * The evaluation rule set version and the ISO 8601 date the evaluation starts.
     */
    readonly evalRuleset: string;
    /**
     * Whether the evaluation mode is enabled or disabled."
     */
    readonly evalStatus: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted list of the mode information.
     */
    readonly json: string;
    /**
     * The security policy mode, either `KRS` (update manually) or `AAG` (update automatically),
     */
    readonly mode: string;
    /**
     * A tabular display of the mode information.
     */
    readonly outputText: string;
    readonly securityPolicyId: string;
}
