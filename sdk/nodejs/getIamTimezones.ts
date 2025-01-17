// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use `akamai.getIamTimezones` to list all time zones Akamai supports. Time zones are in ISO 8601 format. Use the values from this data source to set the time zone for a user. Administrators use this data source to set a user's time zone. The default time zone is GMT.
 *
 * ## Attributes reference
 *
 * These attributes are returned:
 *
 * * `timezones` — Supported timezones.
 *   * `timezone` - The time zone ID.
 *   * `description` - The description of a time zone, including the GMT +/-.
 *   * `offset` - The time zone offset from GMT.
 *   * `posix` - The time zone posix.
 *
 * [API Reference](https://techdocs.akamai.com/iam-api/reference/get-common-timezones)
 */
export function getIamTimezones(opts?: pulumi.InvokeOptions): Promise<GetIamTimezonesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getIamTimezones:getIamTimezones", {
    }, opts);
}

/**
 * A collection of values returned by getIamTimezones.
 */
export interface GetIamTimezonesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly timezones: outputs.GetIamTimezonesTimezone[];
}
