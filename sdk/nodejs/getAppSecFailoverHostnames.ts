// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security configuration
 *
 * Returns a list of the failover hostnames in a configuration. The returned information is described in the [List failover hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getfailoverhostnames) section of the Application Security API.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/failover-hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getfailoverhostnames)
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
 *     name: "Documentation",
 * });
 * const failoverHostnamesAppSecFailoverHostnames = configuration.then(configuration => akamai.getAppSecFailoverHostnames({
 *     configId: configuration.configId,
 * }));
 * export const failoverHostnames = failoverHostnamesAppSecFailoverHostnames.then(failoverHostnamesAppSecFailoverHostnames => failoverHostnamesAppSecFailoverHostnames.hostnames);
 * export const failoverHostnamesOutput = failoverHostnamesAppSecFailoverHostnames.then(failoverHostnamesAppSecFailoverHostnames => failoverHostnamesAppSecFailoverHostnames.outputText);
 * export const failoverHostnamesJson = failoverHostnamesAppSecFailoverHostnames.then(failoverHostnamesAppSecFailoverHostnames => failoverHostnamesAppSecFailoverHostnames.json);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `hostnames`. List of the failover hostnames.
 * - `json`. JSON-formatted list of the failover hostnames.
 */
export function getAppSecFailoverHostnames(args: GetAppSecFailoverHostnamesArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecFailoverHostnamesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecFailoverHostnames:getAppSecFailoverHostnames", {
        "configId": args.configId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecFailoverHostnames.
 */
export interface GetAppSecFailoverHostnamesArgs {
    /**
     * . Unique identifier of the security configuration associated with the failover hosts.
     */
    configId: number;
}

/**
 * A collection of values returned by getAppSecFailoverHostnames.
 */
export interface GetAppSecFailoverHostnamesResult {
    readonly configId: number;
    readonly hostnames: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly outputText: string;
}

export function getAppSecFailoverHostnamesOutput(args: GetAppSecFailoverHostnamesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecFailoverHostnamesResult> {
    return pulumi.output(args).apply(a => getAppSecFailoverHostnames(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecFailoverHostnames.
 */
export interface GetAppSecFailoverHostnamesOutputArgs {
    /**
     * . Unique identifier of the security configuration associated with the failover hosts.
     */
    configId: pulumi.Input<number>;
}
