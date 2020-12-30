// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getDnsRecordSet(args: GetDnsRecordSetArgs, opts?: pulumi.InvokeOptions): Promise<GetDnsRecordSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getDnsRecordSet:getDnsRecordSet", {
        "host": args.host,
        "recordType": args.recordType,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getDnsRecordSet.
 */
export interface GetDnsRecordSetArgs {
    readonly host: string;
    readonly recordType: string;
    readonly zone: string;
}

/**
 * A collection of values returned by getDnsRecordSet.
 */
export interface GetDnsRecordSetResult {
    readonly host: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly rdatas: string[];
    readonly recordType: string;
    readonly zone: string;
}
