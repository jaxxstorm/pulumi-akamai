// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getGroup:getGroup", {
        "contract": args.contract,
        "contractId": args.contractId,
        "groupName": args.groupName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * @deprecated The setting "contract" has been deprecated.
     */
    readonly contract?: string;
    readonly contractId?: string;
    readonly groupName?: string;
    /**
     * @deprecated The setting "name" has been deprecated.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * @deprecated The setting "contract" has been deprecated.
     */
    readonly contract: string;
    readonly contractId: string;
    readonly groupName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * @deprecated The setting "name" has been deprecated.
     */
    readonly name: string;
}
