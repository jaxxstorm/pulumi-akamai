// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecMatchTargets` data source to retrieve information about the match targets associated with a given configuration, or about a specific match target.
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
 * const matchTargetsAppSecMatchTargets = configuration.then(configuration => akamai.getAppSecMatchTargets({
 *     configId: configuration.configId,
 * }));
 * export const matchTargets = matchTargetsAppSecMatchTargets.then(matchTargetsAppSecMatchTargets => matchTargetsAppSecMatchTargets.outputText);
 * const matchTarget = configuration.then(configuration => akamai.getAppSecMatchTargets({
 *     configId: configuration.configId,
 *     matchTargetId: _var.match_target_id,
 * }));
 * export const matchTargetOutput = matchTarget.then(matchTarget => matchTarget.outputText);
 * ```
 */
export function getAppSecMatchTargets(args: GetAppSecMatchTargetsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecMatchTargetsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecMatchTargets:getAppSecMatchTargets", {
        "configId": args.configId,
        "matchTargetId": args.matchTargetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecMatchTargets.
 */
export interface GetAppSecMatchTargetsArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * The ID of the match target to use. If not supplied, information about all match targets is returned.
     */
    matchTargetId?: number;
}

/**
 * A collection of values returned by getAppSecMatchTargets.
 */
export interface GetAppSecMatchTargetsResult {
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted list of the match target information.
     */
    readonly json: string;
    readonly matchTargetId?: number;
    /**
     * A tabular display showing the ID and Policy ID of all match targets associated with the specified security configuration, or of the specific match target if `matchTargetId` was supplied.
     */
    readonly outputText: string;
}

export function getAppSecMatchTargetsOutput(args: GetAppSecMatchTargetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecMatchTargetsResult> {
    return pulumi.output(args).apply(a => getAppSecMatchTargets(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecMatchTargets.
 */
export interface GetAppSecMatchTargetsOutputArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: pulumi.Input<number>;
    /**
     * The ID of the match target to use. If not supplied, information about all match targets is returned.
     */
    matchTargetId?: pulumi.Input<number>;
}
