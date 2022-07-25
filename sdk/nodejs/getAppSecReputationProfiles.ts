// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security configuration; reputation profile
 *
 * Returns information about your reputation profiles. Reputation profiles grade the security risk of an IP address based on previous activities associated with that address. Depending on the reputation score, and depending on how your configuration has been set up, requests from a specific IP address can trigger an alert, or even be blocked.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/reputation-profiles](https://techdocs.akamai.com/application-security/reference/get-reputation-profiles)
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
 * const reputationProfiles = configuration.then(configuration => akamai.getAppSecReputationProfiles({
 *     configId: configuration.configId,
 * }));
 * export const reputationProfilesOutput = reputationProfiles.then(reputationProfiles => reputationProfiles.outputText);
 * export const reputationProfilesJson = reputationProfiles.then(reputationProfiles => reputationProfiles.json);
 * const reputationProfile = configuration.then(configuration => akamai.getAppSecReputationProfiles({
 *     configId: configuration.configId,
 *     reputationProfileId: 12345,
 * }));
 * export const reputationProfileJson = reputationProfile.then(reputationProfile => reputationProfile.json);
 * export const reputationProfileOutput = reputationProfile.then(reputationProfile => reputationProfile.outputText);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `outputText`. Tabular report of the details about the specified reputation profile or profiles.
 * - `json`. JSON-formatted report of the details about the specified reputation profile or profiles.
 */
export function getAppSecReputationProfiles(args: GetAppSecReputationProfilesArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecReputationProfilesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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
     * . Unique identifier of the security configuration associated with the reputation profiles.
     */
    configId: number;
    /**
     * . Unique identifier of the reputation profile you want to return information for. If not included, information is returned for all your reputation profiles.
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
    readonly json: string;
    readonly outputText: string;
    readonly reputationProfileId?: number;
}

export function getAppSecReputationProfilesOutput(args: GetAppSecReputationProfilesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecReputationProfilesResult> {
    return pulumi.output(args).apply(a => getAppSecReputationProfiles(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecReputationProfiles.
 */
export interface GetAppSecReputationProfilesOutputArgs {
    /**
     * . Unique identifier of the security configuration associated with the reputation profiles.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the reputation profile you want to return information for. If not included, information is returned for all your reputation profiles.
     */
    reputationProfileId?: pulumi.Input<number>;
}
