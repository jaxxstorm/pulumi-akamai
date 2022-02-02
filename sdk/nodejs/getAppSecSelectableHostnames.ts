// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security configuration; contract; group
 *
 * Returns the list of hostnames that can be (but aren't yet) protected by a security configuration. You can specify the set of hostnames to be retrieved either by supplying the name of a security configuration or by supplying an Akamai group ID and contract ID.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selectable-hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getavailablehostnames)
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
 * const selectableHostnamesAppSecSelectableHostnames = configuration.then(configuration => akamai.getAppSecSelectableHostnames({
 *     configId: configuration.configId,
 * }));
 * export const selectableHostnames = selectableHostnamesAppSecSelectableHostnames.then(selectableHostnamesAppSecSelectableHostnames => selectableHostnamesAppSecSelectableHostnames.hostnames);
 * export const selectableHostnamesJson = selectableHostnamesAppSecSelectableHostnames.then(selectableHostnamesAppSecSelectableHostnames => selectableHostnamesAppSecSelectableHostnames.hostnamesJson);
 * export const selectableHostnamesOutputText = selectableHostnamesAppSecSelectableHostnames.then(selectableHostnamesAppSecSelectableHostnames => selectableHostnamesAppSecSelectableHostnames.outputText);
 * const selectableHostnamesForCreateConfigurationAppSecSelectableHostnames = akamai.getAppSecSelectableHostnames({
 *     contractId: "5-2WA382",
 *     groupId: 12198,
 * });
 * export const selectableHostnamesForCreateConfiguration = selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.then(selectableHostnamesForCreateConfigurationAppSecSelectableHostnames => selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.hostnames);
 * export const selectableHostnamesForCreateConfigurationJson = selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.then(selectableHostnamesForCreateConfigurationAppSecSelectableHostnames => selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.hostnamesJson);
 * export const selectableHostnamesForCreateConfigurationOutputText = selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.then(selectableHostnamesForCreateConfigurationAppSecSelectableHostnames => selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.outputText);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `hostnames`. List of selectable hostnames.
 * - `hostnamesJson`. JSON-formatted list of selectable hostnames.
 * - `outputText`. Tabular report of the selectable hostnames showing the name and configId of the security configuration under which the host is protected in production.
 */
export function getAppSecSelectableHostnames(args?: GetAppSecSelectableHostnamesArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecSelectableHostnamesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecSelectableHostnames:getAppSecSelectableHostnames", {
        "activeInProduction": args.activeInProduction,
        "activeInStaging": args.activeInStaging,
        "configId": args.configId,
        "contractid": args.contractid,
        "groupid": args.groupid,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecSelectableHostnames.
 */
export interface GetAppSecSelectableHostnamesArgs {
    activeInProduction?: boolean;
    activeInStaging?: boolean;
    /**
     * . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can't be used with either the `contractid` or the `groupid` arguments.
     */
    configId?: number;
    /**
     * . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can't be used with the `configId` argument.
     */
    contractid?: string;
    /**
     * . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can't be used with the `configId` argument.
     */
    groupid?: number;
}

/**
 * A collection of values returned by getAppSecSelectableHostnames.
 */
export interface GetAppSecSelectableHostnamesResult {
    readonly activeInProduction?: boolean;
    readonly activeInStaging?: boolean;
    readonly configId?: number;
    readonly contractid?: string;
    readonly groupid?: number;
    readonly hostnames: string[];
    readonly hostnamesJson: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputText: string;
}

export function getAppSecSelectableHostnamesOutput(args?: GetAppSecSelectableHostnamesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecSelectableHostnamesResult> {
    return pulumi.output(args).apply(a => getAppSecSelectableHostnames(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecSelectableHostnames.
 */
export interface GetAppSecSelectableHostnamesOutputArgs {
    activeInProduction?: pulumi.Input<boolean>;
    activeInStaging?: pulumi.Input<boolean>;
    /**
     * . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can't be used with either the `contractid` or the `groupid` arguments.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can't be used with the `configId` argument.
     */
    contractid?: pulumi.Input<string>;
    /**
     * . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can't be used with the `configId` argument.
     */
    groupid?: pulumi.Input<number>;
}
