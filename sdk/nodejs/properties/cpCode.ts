// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The `akamai.CpCode` resource lets you create or reuse content provider (CP) codes.  CP codes track web traffic handled by Akamai servers. Akamai gives you a CP code when you purchase a product. You need this code when you activate associated properties.
 *
 * You can create additional CP codes to support more detailed billing and reporting functions.
 *
 * By default, the Akamai Provider uses your existing CP code instead of creating a new one.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const cpCode = new akamai.CpCode("cpCode", {
 *     contractId: akamai_contract.contract.id,
 *     groupId: akamai_group.group.id,
 *     productId: "prd_Object_Delivery",
 * });
 * ```
 *
 * Here's a real-life example that includes other data sources as dependencies:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const groupName = "example group name";
 * const cpcodeName = "My CP Code";
 * const exampleContract = akamai.getContract({
 *     groupName: groupName,
 * });
 * const exampleGroup = exampleContract.then(exampleContract => akamai.getGroup({
 *     name: groupName,
 *     contractId: exampleContract.id,
 * }));
 * const exampleCp = new akamai.CpCode("exampleCp", {
 *     groupId: exampleGroup.then(exampleGroup => exampleGroup.id),
 *     contractId: exampleContract.then(exampleContract => exampleContract.id),
 *     productId: "prd_Object_Delivery",
 * });
 * ```
 * ## Attributes reference
 *
 * * `id` - The ID of the CP code.
 *
 * ## Import
 *
 * Basic Usagehcl resource "akamai_cp_code" "example" {
 *
 * # (resource arguments)
 *
 *  } You can import your Akamai CP codes using a comma-delimited string of the CP code, contract, and group IDs. You have to enter the IDs in this order`cpcode_id,contract_id,group_id` For example
 *
 * ```sh
 *  $ pulumi import akamai:properties/cpCode:CpCode example cpc_123,ctr_1-AB123,grp_123
 * ```
 *
 * @deprecated akamai.properties.CpCode has been deprecated in favor of akamai.CpCode
 */
export class CpCode extends pulumi.CustomResource {
    /**
     * Get an existing CpCode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CpCodeState, opts?: pulumi.CustomResourceOptions): CpCode {
        pulumi.log.warn("CpCode is deprecated: akamai.properties.CpCode has been deprecated in favor of akamai.CpCode")
        return new CpCode(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:properties/cpCode:CpCode';

    /**
     * Returns true if the given object is an instance of CpCode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CpCode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CpCode.__pulumiType;
    }

    /**
     * Replaced by `contractId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "contract" has been deprecated.
     */
    public readonly contract!: pulumi.Output<string>;
    /**
     * - (Required) A contract's unique ID, including the `ctr_` prefix.
     */
    public readonly contractId!: pulumi.Output<string>;
    /**
     * Replaced by `groupId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "group" has been deprecated.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * - (Required) A group's unique ID, including the `grp_` prefix.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Replaced by `productId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "product" has been deprecated.
     */
    public readonly product!: pulumi.Output<string>;
    public readonly productId!: pulumi.Output<string>;

    /**
     * Create a CpCode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.properties.CpCode has been deprecated in favor of akamai.CpCode */
    constructor(name: string, args?: CpCodeArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.properties.CpCode has been deprecated in favor of akamai.CpCode */
    constructor(name: string, argsOrState?: CpCodeArgs | CpCodeState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("CpCode is deprecated: akamai.properties.CpCode has been deprecated in favor of akamai.CpCode")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CpCodeState | undefined;
            resourceInputs["contract"] = state ? state.contract : undefined;
            resourceInputs["contractId"] = state ? state.contractId : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["product"] = state ? state.product : undefined;
            resourceInputs["productId"] = state ? state.productId : undefined;
        } else {
            const args = argsOrState as CpCodeArgs | undefined;
            resourceInputs["contract"] = args ? args.contract : undefined;
            resourceInputs["contractId"] = args ? args.contractId : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["product"] = args ? args.product : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CpCode.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CpCode resources.
 */
export interface CpCodeState {
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
     * - (Required) A group's unique ID, including the `grp_` prefix.
     */
    groupId?: pulumi.Input<string>;
    /**
     * - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
     */
    name?: pulumi.Input<string>;
    /**
     * Replaced by `productId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "product" has been deprecated.
     */
    product?: pulumi.Input<string>;
    productId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CpCode resource.
 */
export interface CpCodeArgs {
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
     * - (Required) A group's unique ID, including the `grp_` prefix.
     */
    groupId?: pulumi.Input<string>;
    /**
     * - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
     */
    name?: pulumi.Input<string>;
    /**
     * Replaced by `productId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "product" has been deprecated.
     */
    product?: pulumi.Input<string>;
    productId?: pulumi.Input<string>;
}
