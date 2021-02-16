// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

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
        return new CpCode(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/cpCode:CpCode';

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
     * @deprecated use "contract_id" attribute instead
     */
    public readonly contract!: pulumi.Output<string>;
    public readonly contractId!: pulumi.Output<string>;
    /**
     * @deprecated use "group_id" attribute instead
     */
    public readonly group!: pulumi.Output<string>;
    public readonly groupId!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly product!: pulumi.Output<string>;

    /**
     * Create a CpCode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CpCodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CpCodeArgs | CpCodeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CpCodeState | undefined;
            inputs["contract"] = state ? state.contract : undefined;
            inputs["contractId"] = state ? state.contractId : undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["product"] = state ? state.product : undefined;
        } else {
            const args = argsOrState as CpCodeArgs | undefined;
            if ((!args || args.product === undefined) && !opts.urn) {
                throw new Error("Missing required property 'product'");
            }
            inputs["contract"] = args ? args.contract : undefined;
            inputs["contractId"] = args ? args.contractId : undefined;
            inputs["group"] = args ? args.group : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["product"] = args ? args.product : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "akamai:properties/cpCode:CpCode" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(CpCode.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CpCode resources.
 */
export interface CpCodeState {
    /**
     * @deprecated use "contract_id" attribute instead
     */
    readonly contract?: pulumi.Input<string>;
    readonly contractId?: pulumi.Input<string>;
    /**
     * @deprecated use "group_id" attribute instead
     */
    readonly group?: pulumi.Input<string>;
    readonly groupId?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly product?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CpCode resource.
 */
export interface CpCodeArgs {
    /**
     * @deprecated use "contract_id" attribute instead
     */
    readonly contract?: pulumi.Input<string>;
    readonly contractId?: pulumi.Input<string>;
    /**
     * @deprecated use "group_id" attribute instead
     */
    readonly group?: pulumi.Input<string>;
    readonly groupId?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly product: pulumi.Input<string>;
}
