// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * The `akamai.PropertyActivation` resource lets you activate a property version. An activation deploys the version to either the Akamai staging or production network. You can activate a specific version multiple times if you need to.
 *
 * Before activating on production, activate on staging first. This way you can detect any problems in staging before your changes progress to production.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 * import * as fs from "fs";
 *
 * const email = "user@example.org";
 * const ruleFormat = "v2020-03-04";
 * const example = new akamai.Property("example", {
 *     productId: "prd_SPM",
 *     contractId: _var.contractid,
 *     groupId: _var.groupid,
 *     hostnames: {
 *         "example.org": "example.org.edgesuite.net",
 *         "www.example.org": "example.org.edgesuite.net",
 *         "sub.example.org": "sub.example.org.edgesuite.net",
 *     },
 *     ruleFormat: ruleFormat,
 *     rules: fs.readFileSync(`${path.module}/main.json`),
 * });
 * const exampleStaging = new akamai.PropertyActivation("exampleStaging", {
 *     propertyId: example.id,
 *     contacts: [email],
 *     version: example.latestVersion,
 *     note: "Sample activation",
 * });
 * const exampleProd = new akamai.PropertyActivation("exampleProd", {
 *     propertyId: example.id,
 *     network: "PRODUCTION",
 *     version: 3,
 *     contacts: [email],
 * }, {
 *     dependsOn: [exampleStaging],
 * });
 * ```
 *
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

    /**
     * The ID given to the activation event while it's in progress.
     */
    public readonly activationId!: pulumi.Output<string>;
    /**
     * Whether the activation should proceed despite any warnings. By default set to `true`.
     */
    public readonly autoAcknowledgeRuleWarnings!: pulumi.Output<boolean | undefined>;
    /**
     * One or more email addresses to send activation status changes to.
     */
    public readonly contacts!: pulumi.Output<string[]>;
    /**
     * The contents of `errors` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
     */
    public /*out*/ readonly errors!: pulumi.Output<string>;
    /**
     * Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
     */
    public readonly network!: pulumi.Output<string | undefined>;
    /**
     * A log message you can assign to the activation request.
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * - (Deprecated) Replaced by `propertyId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "property" has been deprecated.
     */
    public readonly property!: pulumi.Output<string>;
    /**
     * - (Required) The property's unique identifier, including the `prp_` prefix.
     */
    public readonly propertyId!: pulumi.Output<string>;
    public readonly ruleErrors!: pulumi.Output<outputs.properties.PropertyActivationRuleError[]>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    public readonly ruleWarnings!: pulumi.Output<outputs.properties.PropertyActivationRuleWarning[]>;
    /**
     * The property version's activation status on the selected network.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The property version to activate. Previously this field was optional. It now depends on the `akamai.Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latestVersion` attribute in the `akamai.Property` resource labeled `example`.
     */
    public readonly version!: pulumi.Output<number>;
    /**
     * The contents of `warnings` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
     */
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PropertyActivationState | undefined;
            resourceInputs["activationId"] = state ? state.activationId : undefined;
            resourceInputs["autoAcknowledgeRuleWarnings"] = state ? state.autoAcknowledgeRuleWarnings : undefined;
            resourceInputs["contacts"] = state ? state.contacts : undefined;
            resourceInputs["errors"] = state ? state.errors : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["note"] = state ? state.note : undefined;
            resourceInputs["property"] = state ? state.property : undefined;
            resourceInputs["propertyId"] = state ? state.propertyId : undefined;
            resourceInputs["ruleErrors"] = state ? state.ruleErrors : undefined;
            resourceInputs["ruleWarnings"] = state ? state.ruleWarnings : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["warnings"] = state ? state.warnings : undefined;
        } else {
            const args = argsOrState as PropertyActivationArgs | undefined;
            if ((!args || args.contacts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contacts'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["activationId"] = args ? args.activationId : undefined;
            resourceInputs["autoAcknowledgeRuleWarnings"] = args ? args.autoAcknowledgeRuleWarnings : undefined;
            resourceInputs["contacts"] = args ? args.contacts : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["note"] = args ? args.note : undefined;
            resourceInputs["property"] = args ? args.property : undefined;
            resourceInputs["propertyId"] = args ? args.propertyId : undefined;
            resourceInputs["ruleErrors"] = args ? args.ruleErrors : undefined;
            resourceInputs["ruleWarnings"] = args ? args.ruleWarnings : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["errors"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["warnings"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PropertyActivation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PropertyActivation resources.
 */
export interface PropertyActivationState {
    /**
     * The ID given to the activation event while it's in progress.
     */
    activationId?: pulumi.Input<string>;
    /**
     * Whether the activation should proceed despite any warnings. By default set to `true`.
     */
    autoAcknowledgeRuleWarnings?: pulumi.Input<boolean>;
    /**
     * One or more email addresses to send activation status changes to.
     */
    contacts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The contents of `errors` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
     */
    errors?: pulumi.Input<string>;
    /**
     * Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
     */
    network?: pulumi.Input<string>;
    /**
     * A log message you can assign to the activation request.
     */
    note?: pulumi.Input<string>;
    /**
     * - (Deprecated) Replaced by `propertyId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "property" has been deprecated.
     */
    property?: pulumi.Input<string>;
    /**
     * - (Required) The property's unique identifier, including the `prp_` prefix.
     */
    propertyId?: pulumi.Input<string>;
    ruleErrors?: pulumi.Input<pulumi.Input<inputs.properties.PropertyActivationRuleError>[]>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    ruleWarnings?: pulumi.Input<pulumi.Input<inputs.properties.PropertyActivationRuleWarning>[]>;
    /**
     * The property version's activation status on the selected network.
     */
    status?: pulumi.Input<string>;
    /**
     * The property version to activate. Previously this field was optional. It now depends on the `akamai.Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latestVersion` attribute in the `akamai.Property` resource labeled `example`.
     */
    version?: pulumi.Input<number>;
    /**
     * The contents of `warnings` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
     */
    warnings?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PropertyActivation resource.
 */
export interface PropertyActivationArgs {
    /**
     * The ID given to the activation event while it's in progress.
     */
    activationId?: pulumi.Input<string>;
    /**
     * Whether the activation should proceed despite any warnings. By default set to `true`.
     */
    autoAcknowledgeRuleWarnings?: pulumi.Input<boolean>;
    /**
     * One or more email addresses to send activation status changes to.
     */
    contacts: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
     */
    network?: pulumi.Input<string>;
    /**
     * A log message you can assign to the activation request.
     */
    note?: pulumi.Input<string>;
    /**
     * - (Deprecated) Replaced by `propertyId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "property" has been deprecated.
     */
    property?: pulumi.Input<string>;
    /**
     * - (Required) The property's unique identifier, including the `prp_` prefix.
     */
    propertyId?: pulumi.Input<string>;
    ruleErrors?: pulumi.Input<pulumi.Input<inputs.properties.PropertyActivationRuleError>[]>;
    /**
     * @deprecated Rule warnings will not be set in state anymore
     */
    ruleWarnings?: pulumi.Input<pulumi.Input<inputs.properties.PropertyActivationRuleWarning>[]>;
    /**
     * The property version to activate. Previously this field was optional. It now depends on the `akamai.Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latestVersion` attribute in the `akamai.Property` resource labeled `example`.
     */
    version: pulumi.Input<number>;
}
