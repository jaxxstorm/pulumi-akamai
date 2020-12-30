// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class PropertyVariables extends pulumi.CustomResource {
    /**
     * Get an existing PropertyVariables resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertyVariablesState, opts?: pulumi.CustomResourceOptions): PropertyVariables {
        return new PropertyVariables(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/propertyVariables:PropertyVariables';

    /**
     * Returns true if the given object is an instance of PropertyVariables.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PropertyVariables {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PropertyVariables.__pulumiType;
    }

    /**
     * JSON variables representation
     */
    public /*out*/ readonly json!: pulumi.Output<string>;
    /**
     * @deprecated resource "akamai_property_variables" is no longer supported - See Akamai Terraform Upgrade Guide
     */
    public readonly variables!: pulumi.Output<outputs.PropertyVariablesVariable[] | undefined>;

    /**
     * Create a PropertyVariables resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PropertyVariablesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PropertyVariablesArgs | PropertyVariablesState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PropertyVariablesState | undefined;
            inputs["json"] = state ? state.json : undefined;
            inputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as PropertyVariablesArgs | undefined;
            inputs["variables"] = args ? args.variables : undefined;
            inputs["json"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "akamai:properties/propertyVariables:PropertyVariables" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PropertyVariables.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PropertyVariables resources.
 */
export interface PropertyVariablesState {
    /**
     * JSON variables representation
     */
    readonly json?: pulumi.Input<string>;
    /**
     * @deprecated resource "akamai_property_variables" is no longer supported - See Akamai Terraform Upgrade Guide
     */
    readonly variables?: pulumi.Input<pulumi.Input<inputs.PropertyVariablesVariable>[]>;
}

/**
 * The set of arguments for constructing a PropertyVariables resource.
 */
export interface PropertyVariablesArgs {
    /**
     * @deprecated resource "akamai_property_variables" is no longer supported - See Akamai Terraform Upgrade Guide
     */
    readonly variables?: pulumi.Input<pulumi.Input<inputs.PropertyVariablesVariable>[]>;
}
