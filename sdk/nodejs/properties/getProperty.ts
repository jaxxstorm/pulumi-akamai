// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getProperty(args: GetPropertyArgs, opts?: pulumi.InvokeOptions): Promise<GetPropertyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:properties/getProperty:getProperty", {
        "name": args.name,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getProperty.
 */
export interface GetPropertyArgs {
    readonly name: string;
    readonly version?: number;
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
