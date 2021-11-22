// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecThreatIntel` data source to view threat intelligence setting for a policy
 * __BETA__ This is Adaptive Security Engine(ASE) related data source. Please contact your akamai representative if you want to learn more
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
 * const threatIntelAppSecThreatIntel = configuration.then(configuration => akamai.getAppSecThreatIntel({
 *     configId: configuration.configId,
 *     securityPolicyId: _var.security_policy_id,
 * }));
 * export const threatIntel = threatIntelAppSecThreatIntel.then(threatIntelAppSecThreatIntel => threatIntelAppSecThreatIntel.threatIntel);
 * export const json = threatIntelAppSecThreatIntel.then(threatIntelAppSecThreatIntel => threatIntelAppSecThreatIntel.json);
 * export const outputText = threatIntelAppSecThreatIntel.then(threatIntelAppSecThreatIntel => threatIntelAppSecThreatIntel.outputText);
 * ```
 */
export function getAppSecThreatIntel(args: GetAppSecThreatIntelArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecThreatIntelResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecThreatIntel:getAppSecThreatIntel", {
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecThreatIntel.
 */
export interface GetAppSecThreatIntelArgs {
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
 * A collection of values returned by getAppSecThreatIntel.
 */
export interface GetAppSecThreatIntelResult {
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted threat intelligence object
     */
    readonly json: string;
    /**
     * A tabular display of the threat intel information.
     */
    readonly outputText: string;
    readonly securityPolicyId: string;
    /**
     * Threat Intelligence setting, either `on` or `off`.
     */
    readonly threatIntel: string;
}

export function getAppSecThreatIntelOutput(args: GetAppSecThreatIntelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecThreatIntelResult> {
    return pulumi.output(args).apply(a => getAppSecThreatIntel(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecThreatIntel.
 */
export interface GetAppSecThreatIntelOutputArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: pulumi.Input<number>;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId: pulumi.Input<string>;
}
