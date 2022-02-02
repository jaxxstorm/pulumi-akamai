// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.CpCode` data source to retrieve the ID for a content provider (CP) code.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const example = pulumi.output(akamai.getCpCode({
 *     contractId: "ctr_1-AB123",
 *     groupId: "grp_123",
 *     name: "my cpcode name",
 * }));
 * ```
 *
 * Here's a real-world example that includes other data sources as dependencies:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const groupName = "example group name";
 * const cpcodeName = "My CP code Name";
 * const exampleContract = akamai.getContract({
 *     groupName: groupName,
 * });
 * const exampleGroup = exampleContract.then(exampleContract => akamai.getGroup({
 *     groupName: groupName,
 *     contractId: exampleContract.id,
 * }));
 * const exampleCpCode = Promise.all([exampleGroup, exampleContract]).then(([exampleGroup, exampleContract]) => akamai.getCpCode({
 *     name: cpcodeName,
 *     groupId: exampleGroup.id,
 *     contractId: exampleContract.id,
 * }));
 * ```
 * ## Attributes reference
 *
 * This data source returns these attributes:
 *
 * * `id` - The ID of the CP code, including the `cpc_` prefix.
 * * `productIds` - An array of product IDs associated with this CP code. Each ID returned includes the `prd_` prefix.
 */
export function getCpCode(args: GetCpCodeArgs, opts?: pulumi.InvokeOptions): Promise<GetCpCodeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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
     * Replaced by `contractId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "contract" has been deprecated.
     */
    contract?: string;
    /**
     * - (Required) A contract's unique ID, including the `ctr_` prefix.
     */
    contractId?: string;
    /**
     * Replaced by `groupId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "group" has been deprecated.
     */
    group?: string;
    /**
     * The group's unique ID, including the `grp_` prefix.
     */
    groupId?: string;
    /**
     * The name of the CP code.
     */
    name: string;
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

export function getCpCodeOutput(args: GetCpCodeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCpCodeResult> {
    return pulumi.output(args).apply(a => getCpCode(a, opts))
}

/**
 * A collection of arguments for invoking getCpCode.
 */
export interface GetCpCodeOutputArgs {
    /**
     * Replaced by `contractId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "contract" has been deprecated.
     */
    contract?: pulumi.Input<string>;
    /**
     * - (Required) A contract's unique ID, including the `ctr_` prefix.
     */
    contractId?: pulumi.Input<string>;
    /**
     * Replaced by `groupId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "group" has been deprecated.
     */
    group?: pulumi.Input<string>;
    /**
     * The group's unique ID, including the `grp_` prefix.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The name of the CP code.
     */
    name: pulumi.Input<string>;
}
