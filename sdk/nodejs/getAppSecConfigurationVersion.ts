// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAppSecConfigurationVersion(args: GetAppSecConfigurationVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecConfigurationVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecConfigurationVersion:getAppSecConfigurationVersion", {
        "configId": args.configId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecConfigurationVersion.
 */
export interface GetAppSecConfigurationVersionArgs {
    /**
     * . Unique identifier of the security configuration you want to return version information for.
     */
    configId: number;
    /**
     * . Version number of the security configuration you want to return information about. If not included, information about all the security configuration's versions is returned.
     */
    version?: number;
}

/**
 * A collection of values returned by getAppSecConfigurationVersion.
 */
export interface GetAppSecConfigurationVersionResult {
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly latestVersion: number;
    readonly outputText: string;
    readonly productionStatus: string;
    readonly stagingStatus: string;
    readonly version?: number;
}

export function getAppSecConfigurationVersionOutput(args: GetAppSecConfigurationVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecConfigurationVersionResult> {
    return pulumi.output(args).apply(a => getAppSecConfigurationVersion(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecConfigurationVersion.
 */
export interface GetAppSecConfigurationVersionOutputArgs {
    /**
     * . Unique identifier of the security configuration you want to return version information for.
     */
    configId: pulumi.Input<number>;
    /**
     * . Version number of the security configuration you want to return information about. If not included, information about all the security configuration's versions is returned.
     */
    version?: pulumi.Input<number>;
}
