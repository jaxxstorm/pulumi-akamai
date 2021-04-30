// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecWafProtection` resource to enable or disable WAF protection for a given configuration version and security policy.
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
 * const protection = new akamai.AppSecWafProtection("protection", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     version: configuration.then(configuration => configuration.latestVersion),
 *     securityPolicyId: _var.security_policy_id,
 *     enabled: _var.enabled,
 * });
 * ```
 */
export class AppSecWafProtection extends pulumi.CustomResource {
    /**
     * Get an existing AppSecWafProtection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecWafProtectionState, opts?: pulumi.CustomResourceOptions): AppSecWafProtection {
        return new AppSecWafProtection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecWafProtection:AppSecWafProtection';

    /**
     * Returns true if the given object is an instance of AppSecWafProtection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecWafProtection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecWafProtection.__pulumiType;
    }

    /**
     * The ID of the security configuration to use.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * Whether to enable WAF controls: either `true` or `false`.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * A tabular display showing the current protection settings.
     */
    public /*out*/ readonly outputText!: pulumi.Output<string>;
    /**
     * The ID of the security policy to use.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;
    /**
     * The version number of the security configuration to use.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a AppSecWafProtection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecWafProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecWafProtectionArgs | AppSecWafProtectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecWafProtectionState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["outputText"] = state ? state.outputText : undefined;
            inputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AppSecWafProtectionArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["outputText"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecWafProtection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecWafProtection resources.
 */
export interface AppSecWafProtectionState {
    /**
     * The ID of the security configuration to use.
     */
    readonly configId?: pulumi.Input<number>;
    /**
     * Whether to enable WAF controls: either `true` or `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * A tabular display showing the current protection settings.
     */
    readonly outputText?: pulumi.Input<string>;
    /**
     * The ID of the security policy to use.
     */
    readonly securityPolicyId?: pulumi.Input<string>;
    /**
     * The version number of the security configuration to use.
     */
    readonly version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AppSecWafProtection resource.
 */
export interface AppSecWafProtectionArgs {
    /**
     * The ID of the security configuration to use.
     */
    readonly configId: pulumi.Input<number>;
    /**
     * Whether to enable WAF controls: either `true` or `false`.
     */
    readonly enabled: pulumi.Input<boolean>;
    /**
     * The ID of the security policy to use.
     */
    readonly securityPolicyId: pulumi.Input<string>;
    /**
     * The version number of the security configuration to use.
     */
    readonly version: pulumi.Input<number>;
}
