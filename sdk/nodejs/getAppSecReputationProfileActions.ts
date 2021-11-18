// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecReputationProfileActions` data source to retrieve details about reputation profiles and their associated actions, or about the actions associated with a specific reputation profile.
 */
export function getAppSecReputationProfileActions(args: GetAppSecReputationProfileActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecReputationProfileActionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecReputationProfileActions:getAppSecReputationProfileActions", {
        "configId": args.configId,
        "reputationProfileId": args.reputationProfileId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecReputationProfileActions.
 */
export interface GetAppSecReputationProfileActionsArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * The ID of a given reputation profile. If not supplied, information about all reputation profiles is returned.
     */
    reputationProfileId?: number;
    /**
     * THe ID of the security policy to use.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecReputationProfileActions.
 */
export interface GetAppSecReputationProfileActionsResult {
    /**
     * The action that the specified reputation profile or profiles take when triggered.
     */
    readonly action: string;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted display of the specified reputation profile action information.
     */
    readonly json: string;
    /**
     * A tabular display of the specified reputation profile action information.
     */
    readonly outputText: string;
    readonly reputationProfileId?: number;
    readonly securityPolicyId: string;
}

export function getAppSecReputationProfileActionsOutput(args: GetAppSecReputationProfileActionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecReputationProfileActionsResult> {
    return pulumi.output(args).apply(a => getAppSecReputationProfileActions(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecReputationProfileActions.
 */
export interface GetAppSecReputationProfileActionsOutputArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: pulumi.Input<number>;
    /**
     * The ID of a given reputation profile. If not supplied, information about all reputation profiles is returned.
     */
    reputationProfileId?: pulumi.Input<number>;
    /**
     * THe ID of the security policy to use.
     */
    securityPolicyId: pulumi.Input<string>;
}
