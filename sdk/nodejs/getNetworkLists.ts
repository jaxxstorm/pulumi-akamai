// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getNetworkLists` data source to retrieve information about the available network lists,
 * optionally filtered by list type or based on a search string. The information available is described
 * [here](https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlists).
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const networkLists = akamai.getNetworkLists({});
 * export const networkListsText = networkLists.then(networkLists => networkLists.outputText);
 * export const networkListsJson = networkLists.then(networkLists => networkLists.json);
 * export const networkListsList = networkLists.then(networkLists => networkLists.lists);
 * const networkListsFilter = akamai.getNetworkLists({
 *     name: "Test Whitelist",
 *     type: "IP",
 * });
 * export const networkListsFilterText = networkListsFilter.then(networkListsFilter => networkListsFilter.outputText);
 * export const networkListsFilterJson = networkListsFilter.then(networkListsFilter => networkListsFilter.json);
 * export const networkListsFilterList = networkListsFilter.then(networkListsFilter => networkListsFilter.lists);
 * ```
 */
export function getNetworkLists(args?: GetNetworkListsArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkListsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getNetworkLists:getNetworkLists", {
        "name": args.name,
        "networkListId": args.networkListId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkLists.
 */
export interface GetNetworkListsArgs {
    /**
     * The name of a specific network list to retrieve. If not supplied, information about all network
     * lists will be returned.
     */
    name?: string;
    /**
     * The ID of a specific network list to retrieve.
     * If not supplied, information about all network lists will be returned.
     */
    networkListId?: string;
    /**
     * The type of network lists to be retrieved; must be either "IP" or "GEO". If not supplied,
     * information about both types will be returned.
     */
    type?: string;
}

/**
 * A collection of values returned by getNetworkLists.
 */
export interface GetNetworkListsResult {
    readonly contractId: string;
    readonly groupId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted list of information about the specified network list(s).
     */
    readonly json: string;
    /**
     * A list containing the IDs of the specified network lists(s).
     */
    readonly lists: string[];
    readonly name?: string;
    readonly networkListId: string;
    /**
     * A tabular display showing the network list information.
     */
    readonly outputText: string;
    readonly type?: string;
}

export function getNetworkListsOutput(args?: GetNetworkListsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkListsResult> {
    return pulumi.output(args).apply(a => getNetworkLists(a, opts))
}

/**
 * A collection of arguments for invoking getNetworkLists.
 */
export interface GetNetworkListsOutputArgs {
    /**
     * The name of a specific network list to retrieve. If not supplied, information about all network
     * lists will be returned.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of a specific network list to retrieve.
     * If not supplied, information about all network lists will be returned.
     */
    networkListId?: pulumi.Input<string>;
    /**
     * The type of network lists to be retrieved; must be either "IP" or "GEO". If not supplied,
     * information about both types will be returned.
     */
    type?: pulumi.Input<string>;
}
