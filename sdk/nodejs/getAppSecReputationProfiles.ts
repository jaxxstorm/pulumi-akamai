// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecReputationProfiles` data source to retrieve details about all reputation profiles, or a specific reputation profiles.
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
 * const reputationProfiles = configuration.then(configuration => akamai.getAppSecReputationProfiles({
 *     configId: configuration.configId,
 * }));
 * export const reputationProfilesOutput = reputationProfiles.then(reputationProfiles => reputationProfiles.outputText);
 * export const reputationProfilesJson = reputationProfiles.then(reputationProfiles => reputationProfiles.json);
 * const reputationProfile = configuration.then(configuration => akamai.getAppSecReputationProfiles({
 *     configId: configuration.configId,
 *     reputationProfileId: _var.reputation_profile_id,
 * }));
 * export const reputationProfileJson = reputationProfile.then(reputationProfile => reputationProfile.json);
 * export const reputationProfileOutput = reputationProfile.then(reputationProfile => reputationProfile.outputText);
 * ```
 */
export function getAppSecReputationProfiles(args: GetAppSecReputationProfilesArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecReputationProfilesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecReputationProfiles:getAppSecReputationProfiles", {
        "configId": args.configId,
        "reputationProfileId": args.reputationProfileId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecReputationProfiles.
 */
export interface GetAppSecReputationProfilesArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * The ID of a given reputation profile. If not supplied, information about all reputation profiles is returned.
     */
    reputationProfileId?: number;
}

/**
 * A collection of values returned by getAppSecReputationProfiles.
 */
export interface GetAppSecReputationProfilesResult {
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted display of the details about the indicated reputation profile or profiles.
     */
    readonly json: string;
    /**
     * A tabular display of the details about the indicated reputation profile or profiles.
     */
    readonly outputText: string;
    readonly reputationProfileId?: number;
}