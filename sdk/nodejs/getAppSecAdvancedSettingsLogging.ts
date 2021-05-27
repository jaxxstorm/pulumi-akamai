// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecAdvancedSettingsLogging` data source to retrieve information about the HTTP header logging controls for a configuration. This operation applies at the configuration level, and therefore applies to all policies within a configuration. You may retrieve these settings for a particular policy by specifying the policy using the securityPolicyId parameter. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethttpheaderloggingforaconfiguration).
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
 * const logging = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecAdvancedSettingsLogging({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 * }));
 * export const advancedSettingsLoggingOutput = logging.then(logging => logging.outputText);
 * export const advancedSettingsLoggingJson = logging.then(logging => logging.json);
 * const policyOverride = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecAdvancedSettingsLogging({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 *     securityPolicyId: _var.security_policy_id,
 * }));
 * export const advancedSettingsPolicyLoggingOutput = policyOverride.then(policyOverride => policyOverride.outputText);
 * export const advancedSettingsPolicyLoggingJson = policyOverride.then(policyOverride => policyOverride.json);
 * ```
 */
export function getAppSecAdvancedSettingsLogging(args: GetAppSecAdvancedSettingsLoggingArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecAdvancedSettingsLoggingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecAdvancedSettingsLogging:getAppSecAdvancedSettingsLogging", {
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecAdvancedSettingsLogging.
 */
export interface GetAppSecAdvancedSettingsLoggingArgs {
    /**
     * The configuration ID.
     */
    configId: number;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId?: string;
    /**
     * The version number of the configuration.
     */
    version: number;
}

/**
 * A collection of values returned by getAppSecAdvancedSettingsLogging.
 */
export interface GetAppSecAdvancedSettingsLoggingResult {
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted list of information about the logging settings.
     */
    readonly json: string;
    /**
     * A tabular display showing the logging settings.
     */
    readonly outputText: string;
    readonly securityPolicyId?: string;
    readonly version: number;
}
