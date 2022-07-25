// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security configuration
 *
 * Returns the most recent version notes for a security configuration.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/version-notes](https://techdocs.akamai.com/application-security/reference/get-version-notes)
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
 * const versionNotes = configuration.then(configuration => akamai.getAppSecVersionNotes({
 *     configId: configuration.configId,
 * }));
 * export const versionNotesText = versionNotes.then(versionNotes => versionNotes.outputText);
 * export const versionNotesJson = versionNotes.then(versionNotes => versionNotes.json);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `json`. JSON-formatted list showing the version notes.
 * - `outputText`. Tabular report showing the version notes.
 */
export function getAppSecVersionNotes(args: GetAppSecVersionNotesArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecVersionNotesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecVersionNotes:getAppSecVersionNotes", {
        "configId": args.configId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecVersionNotes.
 */
export interface GetAppSecVersionNotesArgs {
    /**
     * . Unique identifier of the security configuration you want to return information for.
     */
    configId: number;
}

/**
 * A collection of values returned by getAppSecVersionNotes.
 */
export interface GetAppSecVersionNotesResult {
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly outputText: string;
}

export function getAppSecVersionNotesOutput(args: GetAppSecVersionNotesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecVersionNotesResult> {
    return pulumi.output(args).apply(a => getAppSecVersionNotes(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecVersionNotes.
 */
export interface GetAppSecVersionNotesOutputArgs {
    /**
     * . Unique identifier of the security configuration you want to return information for.
     */
    configId: pulumi.Input<number>;
}
