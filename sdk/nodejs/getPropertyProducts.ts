// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getPropertyProducts` data source to list the products included on a contract.
 *
 * ## Example Usage
 *
 * This example returns products associated with the [EdgeGrid client token](https://developer.akamai.com/getting-started/edgegrid) for a given contract:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 *
 * export const propertyMatch = data.akamai_property_products["my-example"];
 * ```
 * ## Argument reference
 *
 * This data source supports this argument:
 *
 * * `contractId` - (Required) A contract's unique ID, including the `ctr_` prefix.
 *
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `products` - A list of supported products for the contract, including:
 *   * `productId` - The product's unique ID, including the `prd_` prefix.
 *   * `productName` - A string containing the product name.
 */
export function getPropertyProducts(args: GetPropertyProductsArgs, opts?: pulumi.InvokeOptions): Promise<GetPropertyProductsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getPropertyProducts:getPropertyProducts", {
        "contractId": args.contractId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPropertyProducts.
 */
export interface GetPropertyProductsArgs {
    readonly contractId: string;
}

/**
 * A collection of values returned by getPropertyProducts.
 */
export interface GetPropertyProductsResult {
    readonly contractId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly products: outputs.GetPropertyProductsProduct[];
}
