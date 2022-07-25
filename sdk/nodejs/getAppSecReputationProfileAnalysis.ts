// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy
 *
 * Returns information about the following two reputation analysis settings:
 *
 * - `forwardToHTTPHeader`. When enabled, client reputation information associated with a request is forwarded to origin servers by using an HTTP header.
 * - `forwardSharedIPToHTTPHeaderAndSIEM`. When enabled, both the HTTP header and SIEM integration events include a value indicating that the IP addresses is shared address.
 *
 * The returned information is described in the [ReputationAnalysis members](https://developer.akamai.com/api/cloud_security/application_security/v1.html#f06bb20c) section of the Application Security API.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/reputation-analysis](https://techdocs.akamai.com/application-security/reference/get-reputation-analysis)
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
 * const reputationAnalysis = configuration.then(configuration => akamai.getAppSecReputationProfileAnalysis({
 *     configId: configuration.configId,
 *     securityPolicyId: "gms1_134637",
 * }));
 * export const reputationAnalysisText = reputationAnalysis.then(reputationAnalysis => reputationAnalysis.outputText);
 * export const reputationAnalysisJson = reputationAnalysis.then(reputationAnalysis => reputationAnalysis.json);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `json`. JSON-formatted list of the reputation analysis settings.
 * - `outputText`. Tabular report showing the reputation analysis settings.
 */
export function getAppSecReputationProfileAnalysis(args: GetAppSecReputationProfileAnalysisArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecReputationProfileAnalysisResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecReputationProfileAnalysis:getAppSecReputationProfileAnalysis", {
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecReputationProfileAnalysis.
 */
export interface GetAppSecReputationProfileAnalysisArgs {
    /**
     * . Unique identifier of the security configuration associated with the reputation profile analysis settings.
     */
    configId: number;
    /**
     * . Unique identifier of the security policy associated with the reputation profile analysis settings.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecReputationProfileAnalysis.
 */
export interface GetAppSecReputationProfileAnalysisResult {
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly outputText: string;
    readonly securityPolicyId: string;
}

export function getAppSecReputationProfileAnalysisOutput(args: GetAppSecReputationProfileAnalysisOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecReputationProfileAnalysisResult> {
    return pulumi.output(args).apply(a => getAppSecReputationProfileAnalysis(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecReputationProfileAnalysis.
 */
export interface GetAppSecReputationProfileAnalysisOutputArgs {
    /**
     * . Unique identifier of the security configuration associated with the reputation profile analysis settings.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the reputation profile analysis settings.
     */
    securityPolicyId: pulumi.Input<string>;
}
