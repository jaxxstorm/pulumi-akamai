// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecByPassNetworkList` data source to retrieve information about which network
 * lists are used in the bypass network lists settings.  The information available is described
 * [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getbypassnetworklistsforawapconfigversion).
 * Note: this data source is only applicable to WAP (Web Application Protector) configurations.
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
 * const bypassNetworkLists = configuration.then(configuration => akamai.getAppSecBypassNetworkLists({
 *     configId: configuration.configId,
 * }));
 * export const bypassNetworkListsOutput = bypassNetworkLists.then(bypassNetworkLists => bypassNetworkLists.outputText);
 * export const bypassNetworkListsJson = bypassNetworkLists.then(bypassNetworkLists => bypassNetworkLists.json);
 * export const bypassNetworkListsIdList = bypassNetworkLists.then(bypassNetworkLists => bypassNetworkLists.bypassNetworkLists);
 * ```
 */
export function getAppSecBypassNetworkLists(args: GetAppSecBypassNetworkListsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecBypassNetworkListsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecBypassNetworkLists:getAppSecBypassNetworkLists", {
        "configId": args.configId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecBypassNetworkLists.
 */
export interface GetAppSecBypassNetworkListsArgs {
    /**
     * The configuration ID to use.
     */
    configId: number;
}

/**
 * A collection of values returned by getAppSecBypassNetworkLists.
 */
export interface GetAppSecBypassNetworkListsResult {
    /**
     * A list of strings containing the network list IDs.
     */
    readonly bypassNetworkLists: string[];
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted list of information about the bypass network lists.
     */
    readonly json: string;
    /**
     * A tabular display showing the bypass network list information.
     */
    readonly outputText: string;
}

export function getAppSecBypassNetworkListsOutput(args: GetAppSecBypassNetworkListsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecBypassNetworkListsResult> {
    return pulumi.output(args).apply(a => getAppSecBypassNetworkLists(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecBypassNetworkLists.
 */
export interface GetAppSecBypassNetworkListsOutputArgs {
    /**
     * The configuration ID to use.
     */
    configId: pulumi.Input<number>;
}
