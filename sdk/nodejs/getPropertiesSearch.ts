// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getPropertiesSearch` data source to retrieve the list of properties matching a specific hostname, edge hostname or property name based on the [EdgeGrid API client token](https://techdocs.akamai.com/developer/docs/authenticate-with-edgegrid) you're using.
 *
 * ## Example Usage
 *
 * Return properties associated with the EdgeGrid API client token currently used for authentication:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 *
 * export const myPropertyList = data.akamai_properties_search.example;
 * ```
 * ## Attributes reference
 *
 * This data source returns this attribute:
 *
 * * `properties` - A list of property version matching the given criteria.
 */
export function getPropertiesSearch(args: GetPropertiesSearchArgs, opts?: pulumi.InvokeOptions): Promise<GetPropertiesSearchResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getPropertiesSearch:getPropertiesSearch", {
        "key": args.key,
        "value": args.value,
    }, opts);
}

/**
 * A collection of arguments for invoking getPropertiesSearch.
 */
export interface GetPropertiesSearchArgs {
    /**
     * Key used for search. Valid values are:
     * * **hostname**
     * * **edgeHostname**
     * * **propertyName**
     */
    key: string;
    /**
     * - (Required) Value to search for.
     */
    value: string;
}

/**
 * A collection of values returned by getPropertiesSearch.
 */
export interface GetPropertiesSearchResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly key: string;
    readonly properties: outputs.GetPropertiesSearchProperty[];
    readonly value: string;
}

export function getPropertiesSearchOutput(args: GetPropertiesSearchOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPropertiesSearchResult> {
    return pulumi.output(args).apply(a => getPropertiesSearch(a, opts))
}

/**
 * A collection of arguments for invoking getPropertiesSearch.
 */
export interface GetPropertiesSearchOutputArgs {
    /**
     * Key used for search. Valid values are:
     * * **hostname**
     * * **edgeHostname**
     * * **propertyName**
     */
    key: pulumi.Input<string>;
    /**
     * - (Required) Value to search for.
     */
    value: pulumi.Input<string>;
}
