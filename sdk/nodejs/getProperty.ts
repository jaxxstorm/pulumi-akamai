// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.Property` data source to query and list the property ID and rule tree based on the property name.
 *
 * ## Example Usage
 *
 * This example returns the property ID and rule tree based on the property name and optional version argument:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const example = akamai.getProperty({
 *     name: "terraform-demo",
 *     version: 1,
 * });
 * export const myPropertyID = example;
 * ```
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `property_ID` - A property's unique identifier, including the `prp_` prefix.
 * * `rules` - A JSON-encoded rule tree for a given property.
 */
export function getProperty(args: GetPropertyArgs, opts?: pulumi.InvokeOptions): Promise<GetPropertyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getProperty:getProperty", {
        "name": args.name,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getProperty.
 */
export interface GetPropertyArgs {
    /**
     * - (Required) The property name.
     */
    name: string;
    /**
     * - (Optional) The version of the property whose ID you want to list.
     */
    version?: number;
}

/**
 * A collection of values returned by getProperty.
 */
export interface GetPropertyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly rules: string;
    readonly version?: number;
}

export function getPropertyOutput(args: GetPropertyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPropertyResult> {
    return pulumi.output(args).apply(a => getProperty(a, opts))
}

/**
 * A collection of arguments for invoking getProperty.
 */
export interface GetPropertyOutputArgs {
    /**
     * - (Required) The property name.
     */
    name: pulumi.Input<string>;
    /**
     * - (Optional) The version of the property whose ID you want to list.
     */
    version?: pulumi.Input<number>;
}
