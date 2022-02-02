// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * @deprecated akamai.properties.Property has been deprecated in favor of akamai.Property
 */
export class Property extends pulumi.CustomResource {
    /**
     * Get an existing Property resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertyState, opts?: pulumi.CustomResourceOptions): Property {
        pulumi.log.warn("Property is deprecated: akamai.properties.Property has been deprecated in favor of akamai.Property")
        return new Property(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:properties/property:Property';

    /**
     * Returns true if the given object is an instance of Property.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Property {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Property.__pulumiType;
    }

    /**
     * @deprecated The setting "contact" has been deprecated.
     */
    public readonly contacts!: pulumi.Output<string[] | undefined>;
    /**
     * @deprecated The setting "contract" has been deprecated.
     */
    public readonly contract!: pulumi.Output<string>;
    /**
     * Contract ID to be assigned to the Property
     */
    public readonly contractId!: pulumi.Output<string>;
    /**
     * @deprecated The setting "cp_code" has been deprecated.
     */
    public readonly cpCode!: pulumi.Output<string | undefined>;
    /**
     * @deprecated The setting "group" has been deprecated.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * Group ID to be assigned to the Property
     */
    public readonly groupId!: pulumi.Output<string>;
    public readonly hostnames!: pulumi.Output<outputs.properties.PropertyHostname[] | undefined>;
    /**
     * @deprecated The setting "is_secure" has been deprecated.
     */
    public readonly isSecure!: pulumi.Output<boolean | undefined>;
    /**
     * Property's current latest version number
     */
    public /*out*/ readonly latestVersion!: pulumi.Output<number>;
    /**
     * Name to give to the Property (must be unique)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * @deprecated The setting "origin" has been deprecated.
     */
    public readonly origins!: pulumi.Output<outputs.properties.PropertyOrigin[] | undefined>;
    /**
     * @deprecated The setting "product" has been deprecated.
     */
    public readonly product!: pulumi.Output<string>;
    /**
     * Product ID to be assigned to the Property
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * Property's version currently activated in production (zero when not active in production)
     */
    public /*out*/ readonly productionVersion!: pulumi.Output<number>;
    /**
     * Required property's version to be read
     */
    public /*out*/ readonly readVersion!: pulumi.Output<number>;
    public /*out*/ readonly ruleErrors!: pulumi.Output<outputs.properties.PropertyRuleError[]>;
    /**
     * Specify the rule format version (defaults to latest version available when created)
     */
    public readonly ruleFormat!: pulumi.Output<string>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    public readonly ruleWarnings!: pulumi.Output<outputs.properties.PropertyRuleWarning[]>;
    /**
     * Property Rules as JSON
     */
    public readonly rules!: pulumi.Output<string>;
    /**
     * Property's version currently activated in staging (zero when not active in staging)
     */
    public /*out*/ readonly stagingVersion!: pulumi.Output<number>;
    /**
     * @deprecated The setting "variables" has been deprecated.
     */
    public readonly variables!: pulumi.Output<string | undefined>;

    /**
     * Create a Property resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.properties.Property has been deprecated in favor of akamai.Property */
    constructor(name: string, args?: PropertyArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.properties.Property has been deprecated in favor of akamai.Property */
    constructor(name: string, argsOrState?: PropertyArgs | PropertyState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Property is deprecated: akamai.properties.Property has been deprecated in favor of akamai.Property")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PropertyState | undefined;
            resourceInputs["contacts"] = state ? state.contacts : undefined;
            resourceInputs["contract"] = state ? state.contract : undefined;
            resourceInputs["contractId"] = state ? state.contractId : undefined;
            resourceInputs["cpCode"] = state ? state.cpCode : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["hostnames"] = state ? state.hostnames : undefined;
            resourceInputs["isSecure"] = state ? state.isSecure : undefined;
            resourceInputs["latestVersion"] = state ? state.latestVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["origins"] = state ? state.origins : undefined;
            resourceInputs["product"] = state ? state.product : undefined;
            resourceInputs["productId"] = state ? state.productId : undefined;
            resourceInputs["productionVersion"] = state ? state.productionVersion : undefined;
            resourceInputs["readVersion"] = state ? state.readVersion : undefined;
            resourceInputs["ruleErrors"] = state ? state.ruleErrors : undefined;
            resourceInputs["ruleFormat"] = state ? state.ruleFormat : undefined;
            resourceInputs["ruleWarnings"] = state ? state.ruleWarnings : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["stagingVersion"] = state ? state.stagingVersion : undefined;
            resourceInputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as PropertyArgs | undefined;
            resourceInputs["contacts"] = args ? args.contacts : undefined;
            resourceInputs["contract"] = args ? args.contract : undefined;
            resourceInputs["contractId"] = args ? args.contractId : undefined;
            resourceInputs["cpCode"] = args ? args.cpCode : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["hostnames"] = args ? args.hostnames : undefined;
            resourceInputs["isSecure"] = args ? args.isSecure : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["origins"] = args ? args.origins : undefined;
            resourceInputs["product"] = args ? args.product : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["ruleFormat"] = args ? args.ruleFormat : undefined;
            resourceInputs["ruleWarnings"] = args ? args.ruleWarnings : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
            resourceInputs["latestVersion"] = undefined /*out*/;
            resourceInputs["productionVersion"] = undefined /*out*/;
            resourceInputs["readVersion"] = undefined /*out*/;
            resourceInputs["ruleErrors"] = undefined /*out*/;
            resourceInputs["stagingVersion"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Property.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Property resources.
 */
export interface PropertyState {
    /**
     * @deprecated The setting "contact" has been deprecated.
     */
    contacts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated The setting "contract" has been deprecated.
     */
    contract?: pulumi.Input<string>;
    /**
     * Contract ID to be assigned to the Property
     */
    contractId?: pulumi.Input<string>;
    /**
     * @deprecated The setting "cp_code" has been deprecated.
     */
    cpCode?: pulumi.Input<string>;
    /**
     * @deprecated The setting "group" has been deprecated.
     */
    group?: pulumi.Input<string>;
    /**
     * Group ID to be assigned to the Property
     */
    groupId?: pulumi.Input<string>;
    hostnames?: pulumi.Input<pulumi.Input<inputs.properties.PropertyHostname>[]>;
    /**
     * @deprecated The setting "is_secure" has been deprecated.
     */
    isSecure?: pulumi.Input<boolean>;
    /**
     * Property's current latest version number
     */
    latestVersion?: pulumi.Input<number>;
    /**
     * Name to give to the Property (must be unique)
     */
    name?: pulumi.Input<string>;
    /**
     * @deprecated The setting "origin" has been deprecated.
     */
    origins?: pulumi.Input<pulumi.Input<inputs.properties.PropertyOrigin>[]>;
    /**
     * @deprecated The setting "product" has been deprecated.
     */
    product?: pulumi.Input<string>;
    /**
     * Product ID to be assigned to the Property
     */
    productId?: pulumi.Input<string>;
    /**
     * Property's version currently activated in production (zero when not active in production)
     */
    productionVersion?: pulumi.Input<number>;
    /**
     * Required property's version to be read
     */
    readVersion?: pulumi.Input<number>;
    ruleErrors?: pulumi.Input<pulumi.Input<inputs.properties.PropertyRuleError>[]>;
    /**
     * Specify the rule format version (defaults to latest version available when created)
     */
    ruleFormat?: pulumi.Input<string>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    ruleWarnings?: pulumi.Input<pulumi.Input<inputs.properties.PropertyRuleWarning>[]>;
    /**
     * Property Rules as JSON
     */
    rules?: pulumi.Input<string>;
    /**
     * Property's version currently activated in staging (zero when not active in staging)
     */
    stagingVersion?: pulumi.Input<number>;
    /**
     * @deprecated The setting "variables" has been deprecated.
     */
    variables?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Property resource.
 */
export interface PropertyArgs {
    /**
     * @deprecated The setting "contact" has been deprecated.
     */
    contacts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated The setting "contract" has been deprecated.
     */
    contract?: pulumi.Input<string>;
    /**
     * Contract ID to be assigned to the Property
     */
    contractId?: pulumi.Input<string>;
    /**
     * @deprecated The setting "cp_code" has been deprecated.
     */
    cpCode?: pulumi.Input<string>;
    /**
     * @deprecated The setting "group" has been deprecated.
     */
    group?: pulumi.Input<string>;
    /**
     * Group ID to be assigned to the Property
     */
    groupId?: pulumi.Input<string>;
    hostnames?: pulumi.Input<pulumi.Input<inputs.properties.PropertyHostname>[]>;
    /**
     * @deprecated The setting "is_secure" has been deprecated.
     */
    isSecure?: pulumi.Input<boolean>;
    /**
     * Name to give to the Property (must be unique)
     */
    name?: pulumi.Input<string>;
    /**
     * @deprecated The setting "origin" has been deprecated.
     */
    origins?: pulumi.Input<pulumi.Input<inputs.properties.PropertyOrigin>[]>;
    /**
     * @deprecated The setting "product" has been deprecated.
     */
    product?: pulumi.Input<string>;
    /**
     * Product ID to be assigned to the Property
     */
    productId?: pulumi.Input<string>;
    /**
     * Specify the rule format version (defaults to latest version available when created)
     */
    ruleFormat?: pulumi.Input<string>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    ruleWarnings?: pulumi.Input<pulumi.Input<inputs.properties.PropertyRuleWarning>[]>;
    /**
     * Property Rules as JSON
     */
    rules?: pulumi.Input<string>;
    /**
     * @deprecated The setting "variables" has been deprecated.
     */
    variables?: pulumi.Input<string>;
}
