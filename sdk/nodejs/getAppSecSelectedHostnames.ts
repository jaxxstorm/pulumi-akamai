// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecSelectedHostnames` data source to retrieve a list of the hostnames that are currently protected under a given security configuration version.
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
 *     name: "Akamai Tools",
 * });
 * const selectedHostnamesAppSecSelectedHostnames = Promise.all([configuration, configuration]).then(([configuration, configuration1]) => akamai.getAppSecSelectedHostnames({
 *     configId: configuration.configId,
 *     version: configuration1.latestVersion,
 * }));
 * export const selectedHostnames = selectedHostnamesAppSecSelectedHostnames.then(selectedHostnamesAppSecSelectedHostnames => selectedHostnamesAppSecSelectedHostnames.hostnames);
 * export const selectedHostnamesJson = selectedHostnamesAppSecSelectedHostnames.then(selectedHostnamesAppSecSelectedHostnames => selectedHostnamesAppSecSelectedHostnames.hostnamesJson);
 * export const selectedHostnamesOutputText = selectedHostnamesAppSecSelectedHostnames.then(selectedHostnamesAppSecSelectedHostnames => selectedHostnamesAppSecSelectedHostnames.outputText);
 * ```
 */
export function getAppSecSelectedHostnames(args: GetAppSecSelectedHostnamesArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecSelectedHostnamesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecSelectedHostnames:getAppSecSelectedHostnames", {
        "configId": args.configId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecSelectedHostnames.
 */
export interface GetAppSecSelectedHostnamesArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: number;
    /**
     * The version number of the security configuration to use.
     */
    version: number;
}

/**
 * A collection of values returned by getAppSecSelectedHostnames.
 */
export interface GetAppSecSelectedHostnamesResult {
    readonly configId: number;
    /**
     * The list of selected hostnames.
     */
    readonly hostnames: string[];
    /**
     * The list of selected hostnames in JSON format.
     */
    readonly hostnamesJson: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A tabular display of the selected hostnames.
     */
    readonly outputText: string;
    readonly version: number;
}
