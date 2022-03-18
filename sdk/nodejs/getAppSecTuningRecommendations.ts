// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns tuning recommendations for the specified attack group (or, if the `attackGroup` argument is not included, returns tuning recommendations for all the attack groups in the specified security policy).
 * Tuning recommendations help minimize the number of false positives triggered by a security policy. With a false positive, a client request is marked as having violated the security policy restrictions even though it actually did not.
 * Tuning recommendations are returned as attack group exceptions: if you choose, you can copy the response and use the `akamai.AppSecAttackGroup` resource to add the recommended exception to a security policy or attack group.
 * If the data source response is empty, that means that there are no further recommendations for tuning your security policy or attack group.
 * If you need, you can manually merge a recommended exception for an attack group with the exception previously configured in the attack group resource.
 * You can find additional information in our [Application Security API v1 documentation](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getrecommendations).
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
 * const policyRecommendations = configuration.then(configuration => akamai.getAppSecTuningRecommendations({
 *     configId: configuration.configId,
 *     securityPolicyId: _var.security_policy_id,
 * }));
 * export const policyRecommendationsJson = policyRecommendations.then(policyRecommendations => policyRecommendations.json);
 * const attackGroupRecommendations = configuration.then(configuration => akamai.getAppSecTuningRecommendations({
 *     configId: configuration.configId,
 *     securityPolicyId: _var.security_policy_id,
 *     attackGroup: _var.attack_group,
 * }));
 * export const attackGroupRecommendationsJson = attackGroupRecommendations.then(attackGroupRecommendations => attackGroupRecommendations.json);
 * ```
 */
export function getAppSecTuningRecommendations(args: GetAppSecTuningRecommendationsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecTuningRecommendationsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecTuningRecommendations:getAppSecTuningRecommendations", {
        "attackGroup": args.attackGroup,
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecTuningRecommendations.
 */
export interface GetAppSecTuningRecommendationsArgs {
    /**
     * . Unique name of the attack group you want to return tuning recommendations for. If not included, recommendations are returned for all your attack groups.
     */
    attackGroup?: string;
    /**
     * . Unique identifier of the security configuration you want to return tuning recommendations for.
     */
    configId: number;
    /**
     * . Unique identifier of the security policy you want to return tuning recommendations for.
     */
    securityPolicyId?: string;
}

/**
 * A collection of values returned by getAppSecTuningRecommendations.
 */
export interface GetAppSecTuningRecommendationsResult {
    readonly attackGroup?: string;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * JSON-formatted list of the tuning recommendations for the security policy or the attack group. The exception block format in a recommendation conforms to the exception block format used in `conditionException` element of `attackGroup` resource.
     */
    readonly json: string;
    readonly securityPolicyId?: string;
}

export function getAppSecTuningRecommendationsOutput(args: GetAppSecTuningRecommendationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecTuningRecommendationsResult> {
    return pulumi.output(args).apply(a => getAppSecTuningRecommendations(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecTuningRecommendations.
 */
export interface GetAppSecTuningRecommendationsOutputArgs {
    /**
     * . Unique name of the attack group you want to return tuning recommendations for. If not included, recommendations are returned for all your attack groups.
     */
    attackGroup?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration you want to return tuning recommendations for.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy you want to return tuning recommendations for.
     */
    securityPolicyId?: pulumi.Input<string>;
}
