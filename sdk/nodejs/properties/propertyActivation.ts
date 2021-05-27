// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * @deprecated akamai.properties.PropertyActivation has been deprecated in favor of akamai.PropertyActivation
 */
export class PropertyActivation extends pulumi.CustomResource {
    /**
     * Get an existing PropertyActivation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertyActivationState, opts?: pulumi.CustomResourceOptions): PropertyActivation {
        pulumi.log.warn("PropertyActivation is deprecated: akamai.properties.PropertyActivation has been deprecated in favor of akamai.PropertyActivation")
        return new PropertyActivation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:properties/propertyActivation:PropertyActivation';

    /**
     * Returns true if the given object is an instance of PropertyActivation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PropertyActivation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PropertyActivation.__pulumiType;
    }

    public readonly activationId!: pulumi.Output<string>;
    /**
     * automatically acknowledge all rule warnings for activation to continue. default is true
     */
    public readonly autoAcknowledgeRuleWarnings!: pulumi.Output<boolean | undefined>;
    public readonly contacts!: pulumi.Output<string[]>;
    public /*out*/ readonly errors!: pulumi.Output<string>;
    public readonly network!: pulumi.Output<string | undefined>;
    /**
     * @deprecated The setting "property" has been deprecated.
     */
    public readonly property!: pulumi.Output<string>;
    public readonly propertyId!: pulumi.Output<string>;
    public readonly ruleErrors!: pulumi.Output<outputs.properties.PropertyActivationRuleError[]>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    public readonly ruleWarnings!: pulumi.Output<outputs.properties.PropertyActivationRuleWarning[]>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly version!: pulumi.Output<number>;
    public /*out*/ readonly warnings!: pulumi.Output<string>;

    /**
     * Create a PropertyActivation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.properties.PropertyActivation has been deprecated in favor of akamai.PropertyActivation */
    constructor(name: string, args: PropertyActivationArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.properties.PropertyActivation has been deprecated in favor of akamai.PropertyActivation */
    constructor(name: string, argsOrState?: PropertyActivationArgs | PropertyActivationState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("PropertyActivation is deprecated: akamai.properties.PropertyActivation has been deprecated in favor of akamai.PropertyActivation")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PropertyActivationState | undefined;
            inputs["activationId"] = state ? state.activationId : undefined;
            inputs["autoAcknowledgeRuleWarnings"] = state ? state.autoAcknowledgeRuleWarnings : undefined;
            inputs["contacts"] = state ? state.contacts : undefined;
            inputs["errors"] = state ? state.errors : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["property"] = state ? state.property : undefined;
            inputs["propertyId"] = state ? state.propertyId : undefined;
            inputs["ruleErrors"] = state ? state.ruleErrors : undefined;
            inputs["ruleWarnings"] = state ? state.ruleWarnings : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["warnings"] = state ? state.warnings : undefined;
        } else {
            const args = argsOrState as PropertyActivationArgs | undefined;
            if ((!args || args.contacts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contacts'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["activationId"] = args ? args.activationId : undefined;
            inputs["autoAcknowledgeRuleWarnings"] = args ? args.autoAcknowledgeRuleWarnings : undefined;
            inputs["contacts"] = args ? args.contacts : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["property"] = args ? args.property : undefined;
            inputs["propertyId"] = args ? args.propertyId : undefined;
            inputs["ruleErrors"] = args ? args.ruleErrors : undefined;
            inputs["ruleWarnings"] = args ? args.ruleWarnings : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["errors"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["warnings"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PropertyActivation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PropertyActivation resources.
 */
export interface PropertyActivationState {
    activationId?: pulumi.Input<string>;
    /**
     * automatically acknowledge all rule warnings for activation to continue. default is true
     */
    autoAcknowledgeRuleWarnings?: pulumi.Input<boolean>;
    contacts?: pulumi.Input<pulumi.Input<string>[]>;
    errors?: pulumi.Input<string>;
    network?: pulumi.Input<string>;
    /**
     * @deprecated The setting "property" has been deprecated.
     */
    property?: pulumi.Input<string>;
    propertyId?: pulumi.Input<string>;
    ruleErrors?: pulumi.Input<pulumi.Input<inputs.properties.PropertyActivationRuleError>[]>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    ruleWarnings?: pulumi.Input<pulumi.Input<inputs.properties.PropertyActivationRuleWarning>[]>;
    status?: pulumi.Input<string>;
    version?: pulumi.Input<number>;
    warnings?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PropertyActivation resource.
 */
export interface PropertyActivationArgs {
    activationId?: pulumi.Input<string>;
    /**
     * automatically acknowledge all rule warnings for activation to continue. default is true
     */
    autoAcknowledgeRuleWarnings?: pulumi.Input<boolean>;
    contacts: pulumi.Input<pulumi.Input<string>[]>;
    network?: pulumi.Input<string>;
    /**
     * @deprecated The setting "property" has been deprecated.
     */
    property?: pulumi.Input<string>;
    propertyId?: pulumi.Input<string>;
    ruleErrors?: pulumi.Input<pulumi.Input<inputs.properties.PropertyActivationRuleError>[]>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    ruleWarnings?: pulumi.Input<pulumi.Input<inputs.properties.PropertyActivationRuleWarning>[]>;
    version: pulumi.Input<number>;
}
