// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecSecurityPolicyProtections` data source to retrieve the protections in effect for a given security policy.
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
 * const protections = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecSecurityPolicyProtections({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 *     securityPolicyId: _var.security_policy_id,
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
 */
export function getAppSecSecurityPolicyProtections(args: GetAppSecSecurityPolicyProtectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecSecurityPolicyProtectionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecSecurityPolicyProtections:getAppSecSecurityPolicyProtections", {
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecSecurityPolicyProtections.
 */
export interface GetAppSecSecurityPolicyProtectionsArgs {
    /**
     * The ID of the security configuration to use.
     */
    readonly configId: number;
    /**
     * The ID of the security policy to use.
     */
    readonly securityPolicyId: string;
    /**
     * The version number of the security configuration to use.
     */
    readonly version: number;
}

/**
 * A collection of values returned by getAppSecSecurityPolicyProtections.
 */
export interface GetAppSecSecurityPolicyProtectionsResult {
    /**
     * `true` or `false`, indicating whether API constraints are in effect.
     */
    readonly applyApiConstraints: boolean;
    /**
     * `true` or `false`, indicating whether application layer controls are in effect.
     */
    readonly applyApplicationLayerControls: boolean;
    /**
     * `true` or `false`, indicating whether botman controls are in effect.
     */
    readonly applyBotmanControls: boolean;
    /**
     * `true` or `false`, indicating whether network layer controls are in effect.
     */
    readonly applyNetworkLayerControls: boolean;
    /**
     * `true` or `false`, indicating whether rate controls are in effect.
     */
    readonly applyRateControls: boolean;
    /**
     * `true` or `false`, indicating whether reputation controls are in effect.
     */
    readonly applyReputationControls: boolean;
    /**
     * `true` or `false`, indicating whether slow post controls are in effect.
     */
    readonly applySlowPostControls: boolean;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * a JSON-formatted list showing the status of the protection settings
     */
    readonly json: string;
    /**
     * a tabular display showing the status of the protection settings
     */
    readonly outputText: string;
    readonly securityPolicyId: string;
    readonly version: number;
}
