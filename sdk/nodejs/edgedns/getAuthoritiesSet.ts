// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/** @deprecated akamai.edgedns.getAuthoritiesSet has been deprecated in favor of akamai.getAuthoritiesSet */
export function getAuthoritiesSet(args: GetAuthoritiesSetArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthoritiesSetResult> {
    pulumi.log.warn("getAuthoritiesSet is deprecated: akamai.edgedns.getAuthoritiesSet has been deprecated in favor of akamai.getAuthoritiesSet")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:edgedns/getAuthoritiesSet:getAuthoritiesSet", {
        "contract": args.contract,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthoritiesSet.
 */
export interface GetAuthoritiesSetArgs {
    contract: string;
}

/**
 * A collection of values returned by getAuthoritiesSet.
 */
export interface GetAuthoritiesSetResult {
    readonly authorities: string[];
    readonly contract: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
