// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecAttackGroupActionConditionException` data source to retrieve an attack group's conditions and exceptions.
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
 * const conditionException = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecAttackGroupConditionException({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 *     securityPolicyId: _var.security_policy_id,
 *     attackGroup: _var.attack_group,
 * }));
 * export const conditionExceptionText = conditionException.then(conditionException => conditionException.outputText);
 * export const conditionExceptionJson = conditionException.then(conditionException => conditionException.json);
 * ```
 */
export function getAppSecAttackGroupConditionException(args: GetAppSecAttackGroupConditionExceptionArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecAttackGroupConditionExceptionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecAttackGroupConditionException:getAppSecAttackGroupConditionException", {
        "attackGroup": args.attackGroup,
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecAttackGroupConditionException.
 */
export interface GetAppSecAttackGroupConditionExceptionArgs {
    /**
     * The attack group to use.
     */
    attackGroup: string;
    /**
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId: string;
    /**
     * The version number of the security configuration to use.
     */
    version: number;
}

/**
 * A collection of values returned by getAppSecAttackGroupConditionException.
 */
export interface GetAppSecAttackGroupConditionExceptionResult {
    readonly attackGroup: string;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The condition and exception information in JSON format.
     */
    readonly json: string;
    /**
     * A tabular display showing the condition and exception information.
     */
    readonly outputText: string;
    readonly securityPolicyId: string;
    readonly version: number;
}
