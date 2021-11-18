// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecRuleUpgradeDetails` data source to retrieve information on changes to the KRS rule sets.
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
 * const upgradeDetails = configuration.then(configuration => akamai.getAppSecRuleUpgradeDetails({
 *     configId: configuration.configId,
 *     securityPolicyId: _var.security_policy_id,
 * }));
 * export const upgradeDetailsText = upgradeDetails.then(upgradeDetails => upgradeDetails.outputText);
 * export const upgradeDetailsJson = upgradeDetails.then(upgradeDetails => upgradeDetails.json);
 * ```
 */
export function getAppSecRuleUpgradeDetails(args: GetAppSecRuleUpgradeDetailsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecRuleUpgradeDetailsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecRuleUpgradeDetails:getAppSecRuleUpgradeDetails", {
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecRuleUpgradeDetails.
 */
export interface GetAppSecRuleUpgradeDetailsArgs {
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
 * A collection of values returned by getAppSecRuleUpgradeDetails.
 */
export interface GetAppSecRuleUpgradeDetailsResult {
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted list of the changes (additions and deletions) to the rules for the specified security policy.
     */
    readonly json: string;
    /**
     * A tabular display showing changes (additions and deletions) to the rules for the specified security policy.
     */
    readonly outputText: string;
    readonly securityPolicyId: string;
}

export function getAppSecRuleUpgradeDetailsOutput(args: GetAppSecRuleUpgradeDetailsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecRuleUpgradeDetailsResult> {
    return pulumi.output(args).apply(a => getAppSecRuleUpgradeDetails(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecRuleUpgradeDetails.
 */
export interface GetAppSecRuleUpgradeDetailsOutputArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: pulumi.Input<number>;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId: pulumi.Input<string>;
}
