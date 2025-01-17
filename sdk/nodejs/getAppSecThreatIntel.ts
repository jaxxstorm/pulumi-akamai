// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy
 *
 * Returns threat intelligence settings for a security policy Note that this data source is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information on ASE, please contact your Akamai representative.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/threat-intel](https://techdocs.akamai.com/application-security/reference/get-rules-threat-intel)
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
 * const threatIntelAppSecThreatIntel = configuration.then(configuration => akamai.getAppSecThreatIntel({
 *     configId: configuration.configId,
 *     securityPolicyId: "gms1_134637",
 * }));
 * export const threatIntel = threatIntelAppSecThreatIntel.then(threatIntelAppSecThreatIntel => threatIntelAppSecThreatIntel.threatIntel);
 * export const json = threatIntelAppSecThreatIntel.then(threatIntelAppSecThreatIntel => threatIntelAppSecThreatIntel.json);
 * export const outputText = threatIntelAppSecThreatIntel.then(threatIntelAppSecThreatIntel => threatIntelAppSecThreatIntel.outputText);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `threatIntel`. Reports the threat Intelligence setting, either **on** or **off**.
 * - `json`. JSON-formatted threat intelligence report.
 * - `outputText`. Tabular report of the threat intelligence information.
 */
export function getAppSecThreatIntel(args: GetAppSecThreatIntelArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecThreatIntelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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
     * . Unique identifier of the security configuration associated with the threat intelligence settings.
     */
    configId: number;
    /**
     * . Unique identifier of the security policy associated with the threat intelligence settings.
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
    readonly json: string;
    readonly outputText: string;
    readonly securityPolicyId: string;
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
     * . Unique identifier of the security configuration associated with the threat intelligence settings.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the threat intelligence settings.
     */
    securityPolicyId: pulumi.Input<string>;
}
