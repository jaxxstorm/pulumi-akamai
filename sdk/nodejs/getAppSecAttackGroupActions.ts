// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecAttackGroupActions` data source to retrieve a list of attack groups and actions associated with a security policy or a specific attack group and action associated with a security policy.
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
 * const attackGroupActions = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecAttackGroupActions({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 *     securityPolicyId: _var.security_policy_id,
 * }));
 * export const attackGroupActionsText = attackGroupActions.then(attackGroupActions => attackGroupActions.outputText);
 * export const attackGroupActionsJson = attackGroupActions.then(attackGroupActions => attackGroupActions.json);
 * const attackGroupActionAppSecAttackGroupActions = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecAttackGroupActions({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 *     securityPolicyId: _var.security_policy_id,
 *     attackGroup: _var.attack_group,
 * }));
 * export const attackGroupAction = attackGroupActionAppSecAttackGroupActions.then(attackGroupActionAppSecAttackGroupActions => attackGroupActionAppSecAttackGroupActions.action);
 * ```
 */
export function getAppSecAttackGroupActions(args: GetAppSecAttackGroupActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecAttackGroupActionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecAttackGroupActions:getAppSecAttackGroupActions", {
        "attackGroup": args.attackGroup,
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecAttackGroupActions.
 */
export interface GetAppSecAttackGroupActionsArgs {
    /**
     * The attack group to use. If not supplied, information about all attack groups will be returned.
     */
    readonly attackGroup?: string;
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
 * A collection of values returned by getAppSecAttackGroupActions.
 */
export interface GetAppSecAttackGroupActionsResult {
    /**
     * The attack group action for the attack group if one was specified: `alert`, `deny`, or `none`. If the action is none, the attack group is inactive in the security policy.
     */
    readonly action: string;
    readonly attackGroup?: string;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The attack group information in JSON format.
     */
    readonly json: string;
    /**
     * A tabular display showing the `action` and `group` name for each attack group.
     */
    readonly outputText: string;
    readonly securityPolicyId: string;
    readonly version: number;
}
