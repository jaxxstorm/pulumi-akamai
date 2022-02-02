// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAppSecAttackGroups(args: GetAppSecAttackGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecAttackGroupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecAttackGroups:getAppSecAttackGroups", {
        "attackGroup": args.attackGroup,
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecAttackGroups.
 */
export interface GetAppSecAttackGroupsArgs {
    /**
     * . Unique name of the attack group you want to return information for. If not included, information is returned for all your attack groups.
     */
    attackGroup?: string;
    /**
     * . Unique identifier of the security configuration associated with the attack group.
     */
    configId: number;
    /**
     * . Unique identifier of the security policy associated with the attack group.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecAttackGroups.
 */
export interface GetAppSecAttackGroupsResult {
    readonly attackGroup?: string;
    readonly attackGroupAction: string;
    readonly conditionException: string;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly outputText: string;
    readonly securityPolicyId: string;
}

export function getAppSecAttackGroupsOutput(args: GetAppSecAttackGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecAttackGroupsResult> {
    return pulumi.output(args).apply(a => getAppSecAttackGroups(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecAttackGroups.
 */
export interface GetAppSecAttackGroupsOutputArgs {
    /**
     * . Unique name of the attack group you want to return information for. If not included, information is returned for all your attack groups.
     */
    attackGroup?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration associated with the attack group.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the attack group.
     */
    securityPolicyId: pulumi.Input<string>;
}
