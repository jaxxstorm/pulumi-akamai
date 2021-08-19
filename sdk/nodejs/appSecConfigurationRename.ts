// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `akamai.AppSecConfigurationRename` resource allows you to rename an existing security configuration.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const configurationAppSecConfiguration = akamai.getAppSecConfiguration({
 *     name: _var.security_configuration,
 * });
 * const configurationAppSecConfigurationRename = new akamai.AppSecConfigurationRename("configurationAppSecConfigurationRename", {
 *     configId: configurationAppSecConfiguration.then(configurationAppSecConfiguration => configurationAppSecConfiguration.configId),
 *     description: _var.description,
 * });
 * ```
 */
export class AppSecConfigurationRename extends pulumi.CustomResource {
    /**
     * Get an existing AppSecConfigurationRename resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecConfigurationRenameState, opts?: pulumi.CustomResourceOptions): AppSecConfigurationRename {
        return new AppSecConfigurationRename(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecConfigurationRename:AppSecConfigurationRename';

    /**
     * Returns true if the given object is an instance of AppSecConfigurationRename.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecConfigurationRename {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecConfigurationRename.__pulumiType;
    }

    /**
     * The ID of the security configuration to be renamed.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * The description to be applied to the configuration.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The new name to be given to the configuration.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a AppSecConfigurationRename resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecConfigurationRenameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecConfigurationRenameArgs | AppSecConfigurationRenameState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecConfigurationRenameState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as AppSecConfigurationRenameArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecConfigurationRename.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecConfigurationRename resources.
 */
export interface AppSecConfigurationRenameState {
    /**
     * The ID of the security configuration to be renamed.
     */
    configId?: pulumi.Input<number>;
    /**
     * The description to be applied to the configuration.
     */
    description?: pulumi.Input<string>;
    /**
     * The new name to be given to the configuration.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecConfigurationRename resource.
 */
export interface AppSecConfigurationRenameArgs {
    /**
     * The ID of the security configuration to be renamed.
     */
    configId: pulumi.Input<number>;
    /**
     * The description to be applied to the configuration.
     */
    description: pulumi.Input<string>;
    /**
     * The new name to be given to the configuration.
     */
    name?: pulumi.Input<string>;
}