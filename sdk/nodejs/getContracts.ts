// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getContracts` data source to list contracts associated with the [EdgeGrid API client token](https://developer.akamai.com/getting-started/edgegrid) you're using.
 *
 * ## Example Usage
 *
 * Return contracts associated with the EdgeGrid API client token currently used for authentication:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const my-example = akamai.getContracts({});
 * export const propertyMatch = my_example;
 * ```
 * ## Argument reference
 *
 * There are no arguments available for this data source.
 *
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `contracts` - A list of supported contracts, with the following properties:
 *   * `contractId` - The contract's unique ID, including the `ctr_` prefix.
 *   * `contractTypeName` - The type of contract, either `DIRECT_CUSTOMER`, `INDIRECT_CUSTOMER`, `PARENT_CUSTOMER`, `REFERRAL_PARTNER`, `TIER_1_RESELLER`, `VAR_CUSTOMER`, `VALUE_ADDED_RESELLER`, `PARTNER`, `PORTAL_PARTNER`, `STREAMING_RESELLER`, `AKAMAI_INTERNAL`, or `UNKNOWN`.
 */
export function getContracts(opts?: pulumi.InvokeOptions): Promise<GetContractsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getContracts:getContracts", {
    }, opts);
}

/**
 * A collection of values returned by getContracts.
 */
export interface GetContractsResult {
    readonly contracts: outputs.GetContractsContract[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
