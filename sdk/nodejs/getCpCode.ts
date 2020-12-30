// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getCpCode(args: GetCpCodeArgs, opts?: pulumi.InvokeOptions): Promise<GetCpCodeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getCpCode:getCpCode", {
        "contract": args.contract,
        "contractId": args.contractId,
        "group": args.group,
        "groupId": args.groupId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCpCode.
 */
export interface GetCpCodeArgs {
    /**
     * @deprecated The setting "contract" has been deprecated.
     */
    readonly contract?: string;
    readonly contractId?: string;
    /**
     * @deprecated The setting "group" has been deprecated.
     */
    readonly group?: string;
    readonly groupId?: string;
    readonly name: string;
}

/**
 * A collection of values returned by getCpCode.
 */
export interface GetCpCodeResult {
    /**
     * @deprecated The setting "contract" has been deprecated.
     */
    readonly contract: string;
    readonly contractId: string;
    /**
     * @deprecated The setting "group" has been deprecated.
     */
    readonly group: string;
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly productIds: string[];
}