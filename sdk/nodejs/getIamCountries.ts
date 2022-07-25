// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use `akamai.getIamCountries` to retrieve all the possible countries that Akamai supports. Use the values from this data source to add or update a user's country information.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const countries = akamai.getIamCountries({});
 * export const supportedCountries = countries;
 * ```
 * ## Attributes reference
 *
 * These attributes are returned:
 *
 * * `countries` — A list of countries.
 *
 * [API Reference](https://developer.akamai.com/api/core_features/identity_management_user_admin/v2.html#getadmincountries)
 */
export function getIamCountries(opts?: pulumi.InvokeOptions): Promise<GetIamCountriesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getIamCountries:getIamCountries", {
    }, opts);
}

/**
 * A collection of values returned by getIamCountries.
 */
export interface GetIamCountriesResult {
    readonly countries: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
