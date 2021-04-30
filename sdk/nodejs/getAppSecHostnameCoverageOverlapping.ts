// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.getAppSecHostnameCoverageOverlapping` data source to retrieve information about the configuration versions that contain a hostname also included in the current configuration version. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoverageoverlapping).
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const test = pulumi.output(akamai.getAppSecHostnameCoverageOverlapping({
 *     configId: 43253,
 *     hostname: "example.com",
 *     version: 7,
 * }, { async: true }));
 * ```
 */
export function getAppSecHostnameCoverageOverlapping(args: GetAppSecHostnameCoverageOverlappingArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecHostnameCoverageOverlappingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecHostnameCoverageOverlapping:getAppSecHostnameCoverageOverlapping", {
        "configId": args.configId,
        "hostname": args.hostname,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecHostnameCoverageOverlapping.
 */
export interface GetAppSecHostnameCoverageOverlappingArgs {
    /**
     * The configuration ID.
     */
    readonly configId: number;
    /**
     * The hostname for which to retrieve information.
     */
    readonly hostname: string;
    /**
     * The version number of the configuration.
     */
    readonly version: number;
}

/**
 * A collection of values returned by getAppSecHostnameCoverageOverlapping.
 */
export interface GetAppSecHostnameCoverageOverlappingResult {
    readonly configId: number;
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A JSON-formatted list of the overlap information.
     */
    readonly json: string;
    /**
     * A tabular display of the overlap information.
     */
    readonly outputText: string;
    readonly version: number;
}
