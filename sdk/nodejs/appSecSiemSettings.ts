// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecSiemSettings` resource to mpdate the SIEM integration settings for a specific configuration.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const configuration = akamai.getAppSecConfiguration({
 *     name: _var.security_configuration,
 * });
 * const siemDefinition = akamai.getAppSecSiemDefinitions({
 *     siemDefinitionName: _var.siem_definition_name,
 * });
 * const securityPolicies = configuration.then(configuration => akamai.getAppSecSecurityPolicy({
 *     configId: configuration.configId,
 * }));
 * const siem = new akamai.AppSecSiemSettings("siem", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     enableSiem: true,
 *     enableForAllPolicies: false,
 *     enableBotmanSiem: true,
 *     siemId: siemDefinition.then(siemDefinition => siemDefinition.id),
 *     securityPolicyIds: securityPolicies.then(securityPolicies => securityPolicies.securityPolicyIdLists),
 * });
 * ```
 */
export class AppSecSiemSettings extends pulumi.CustomResource {
    /**
     * Get an existing AppSecSiemSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecSiemSettingsState, opts?: pulumi.CustomResourceOptions): AppSecSiemSettings {
        return new AppSecSiemSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecSiemSettings:AppSecSiemSettings';

    /**
     * Returns true if the given object is an instance of AppSecSiemSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecSiemSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecSiemSettings.__pulumiType;
    }

    /**
     * The configuration ID to use.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * Whether you enabled SIEM for the Bot Manager events.
     */
    public readonly enableBotmanSiem!: pulumi.Output<boolean>;
    /**
     * Whether you enabled SIEM for all the security policies in the configuration.
     */
    public readonly enableForAllPolicies!: pulumi.Output<boolean>;
    /**
     * Whether you enabled SIEM in a security configuration version.
     */
    public readonly enableSiem!: pulumi.Output<boolean>;
    /**
     * The list of security policy identifiers for which to enable the SIEM integration.
     */
    public readonly securityPolicyIds!: pulumi.Output<string[] | undefined>;
    /**
     * An integer that uniquely identifies the SIEM settings.
     */
    public readonly siemId!: pulumi.Output<number>;

    /**
     * Create a AppSecSiemSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecSiemSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecSiemSettingsArgs | AppSecSiemSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecSiemSettingsState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["enableBotmanSiem"] = state ? state.enableBotmanSiem : undefined;
            inputs["enableForAllPolicies"] = state ? state.enableForAllPolicies : undefined;
            inputs["enableSiem"] = state ? state.enableSiem : undefined;
            inputs["securityPolicyIds"] = state ? state.securityPolicyIds : undefined;
            inputs["siemId"] = state ? state.siemId : undefined;
        } else {
            const args = argsOrState as AppSecSiemSettingsArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.enableBotmanSiem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enableBotmanSiem'");
            }
            if ((!args || args.enableForAllPolicies === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enableForAllPolicies'");
            }
            if ((!args || args.enableSiem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enableSiem'");
            }
            if ((!args || args.siemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siemId'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["enableBotmanSiem"] = args ? args.enableBotmanSiem : undefined;
            inputs["enableForAllPolicies"] = args ? args.enableForAllPolicies : undefined;
            inputs["enableSiem"] = args ? args.enableSiem : undefined;
            inputs["securityPolicyIds"] = args ? args.securityPolicyIds : undefined;
            inputs["siemId"] = args ? args.siemId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecSiemSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecSiemSettings resources.
 */
export interface AppSecSiemSettingsState {
    /**
     * The configuration ID to use.
     */
    configId?: pulumi.Input<number>;
    /**
     * Whether you enabled SIEM for the Bot Manager events.
     */
    enableBotmanSiem?: pulumi.Input<boolean>;
    /**
     * Whether you enabled SIEM for all the security policies in the configuration.
     */
    enableForAllPolicies?: pulumi.Input<boolean>;
    /**
     * Whether you enabled SIEM in a security configuration version.
     */
    enableSiem?: pulumi.Input<boolean>;
    /**
     * The list of security policy identifiers for which to enable the SIEM integration.
     */
    securityPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An integer that uniquely identifies the SIEM settings.
     */
    siemId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AppSecSiemSettings resource.
 */
export interface AppSecSiemSettingsArgs {
    /**
     * The configuration ID to use.
     */
    configId: pulumi.Input<number>;
    /**
     * Whether you enabled SIEM for the Bot Manager events.
     */
    enableBotmanSiem: pulumi.Input<boolean>;
    /**
     * Whether you enabled SIEM for all the security policies in the configuration.
     */
    enableForAllPolicies: pulumi.Input<boolean>;
    /**
     * Whether you enabled SIEM in a security configuration version.
     */
    enableSiem: pulumi.Input<boolean>;
    /**
     * The list of security policy identifiers for which to enable the SIEM integration.
     */
    securityPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An integer that uniquely identifies the SIEM settings.
     */
    siemId: pulumi.Input<number>;
}
