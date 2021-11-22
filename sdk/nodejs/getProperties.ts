// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getProperties` data source to query and retrieve the list of properties for a group and contract
 * based on the [EdgeGrid API client token](https://developer.akamai.com/getting-started/edgegrid) you're using.
 *
 * ## Example Usage
 *
 * Return properties associated with the EdgeGrid API client token currently used for authentication:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 *
 * export const myPropertyList = data.akamai_properties.example;
 * ```
 * ## Argument reference
 *
 * This data source supports these arguments:
 *
 * * `contractId` - (Required) A contract's unique ID, including the `ctr_` prefix.
 * * `groupId` - (Required) A group's unique ID, including the `grp_` prefix.
 *
 * ## Attributes reference
 *
 * This data source returns this attribute:
 *
 * * `properties` - A list of properties available for the contract and group IDs provided.
 */
export function getProperties(args: GetPropertiesArgs, opts?: pulumi.InvokeOptions): Promise<GetPropertiesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getProperties:getProperties", {
        "contractId": args.contractId,
        "groupId": args.groupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProperties.
 */
export interface GetPropertiesArgs {
    contractId: string;
    groupId: string;
}

/**
 * A collection of values returned by getProperties.
 */
export interface GetPropertiesResult {
    readonly contractId: string;
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly properties: outputs.GetPropertiesProperty[];
}

export function getPropertiesOutput(args: GetPropertiesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPropertiesResult> {
    return pulumi.output(args).apply(a => getProperties(a, opts))
}

/**
 * A collection of arguments for invoking getProperties.
 */
export interface GetPropertiesOutputArgs {
    contractId: pulumi.Input<string>;
    groupId: pulumi.Input<string>;
}
