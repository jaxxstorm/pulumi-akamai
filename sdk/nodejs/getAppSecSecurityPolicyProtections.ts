// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy
 *
 * Returns information about the protections in effect for the specified security policy.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/protections](https://techdocs.akamai.com/application-security/reference/get-policy-protections)
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
 * const protections = configuration.then(configuration => akamai.getAppSecSecurityPolicyProtections({
 *     configId: configuration.configId,
 *     securityPolicyId: "gms1_134637",
 * }));
 * export const protectionsJson = protections.then(protections => protections.json);
 * export const protectionsApplyApiConstraints = protections.then(protections => protections.applyApiConstraints);
 * export const protectionsApplyApplicationLayerControls = protections.then(protections => protections.applyApplicationLayerControls);
 * export const protectionsApplyBotmanControls = protections.then(protections => protections.applyBotmanControls);
 * export const protectionsApplyNetworkLayerControls = protections.then(protections => protections.applyNetworkLayerControls);
 * export const protectionsApplyRateControls = protections.then(protections => protections.applyRateControls);
 * export const protectionsApplyReputationControls = protections.then(protections => protections.applyReputationControls);
 * export const protectionsApplySlowPostControls = protections.then(protections => protections.applySlowPostControls);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned and how that returned information is formatted:
 *
 * - `applyApplicationLayerControls`. Returns **true** if application layer controls are enabled; returns **false** if they are not.
 * - `applyNetworkLayerControls`. Returns **true** if network layer controls are enabled; returns **false** if they are not.
 * - `applyRateControls`. Returns **true** if rate controls are enabled; returns **false** if they are not.
 * - `applyReputationControls`. Returns **true** if reputation controls are enabled; returns **false** if they are not.
 * - `applyBotmanControls`. Returns **true** if Bot Manager controls are enabled; returns **false** if they are not.
 * - `applyApiConstraints`. Returns **true** if API constraints are enabled; returns **false** if they are not.
 * - `applySlowPostControls`. Returns **true** if slow POST controls are enabled; returns **false** if they are not.
 * - `json`. JSON-formatted list showing the status of the protection settings.
 * - `outputText`. Tabular report showing the status of the protection settings.
 */
export function getAppSecSecurityPolicyProtections(args: GetAppSecSecurityPolicyProtectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecSecurityPolicyProtectionsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecSecurityPolicyProtections:getAppSecSecurityPolicyProtections", {
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecSecurityPolicyProtections.
 */
export interface GetAppSecSecurityPolicyProtectionsArgs {
    /**
     * . Unique identifier of the security configuration associated with the security policy protections.
     */
    configId: number;
    /**
     * . Unique identifier of the security policy you want to return protections information for.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecSecurityPolicyProtections.
 */
export interface GetAppSecSecurityPolicyProtectionsResult {
    readonly applyApiConstraints: boolean;
    readonly applyApplicationLayerControls: boolean;
    readonly applyBotmanControls: boolean;
    readonly applyNetworkLayerControls: boolean;
    readonly applyRateControls: boolean;
    readonly applyReputationControls: boolean;
    readonly applySlowPostControls: boolean;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly outputText: string;
    readonly securityPolicyId: string;
}

export function getAppSecSecurityPolicyProtectionsOutput(args: GetAppSecSecurityPolicyProtectionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecSecurityPolicyProtectionsResult> {
    return pulumi.output(args).apply(a => getAppSecSecurityPolicyProtections(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecSecurityPolicyProtections.
 */
export interface GetAppSecSecurityPolicyProtectionsOutputArgs {
    /**
     * . Unique identifier of the security configuration associated with the security policy protections.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy you want to return protections information for.
     */
    securityPolicyId: pulumi.Input<string>;
}
