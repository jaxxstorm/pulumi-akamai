// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy
 *
 * Enables or disables rate protection for a security policy.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/protections](https://techdocs.akamai.com/application-security/reference/put-policy-protections)
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
 *     name: "Documentation",
 * });
 * const protection = new akamai.AppSecRateProtection("protection", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: "gms1_134637",
 *     enabled: true,
 * });
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `outputText`. Tabular report showing the current protection settings.
 */
export class AppSecRateProtection extends pulumi.CustomResource {
    /**
     * Get an existing AppSecRateProtection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecRateProtectionState, opts?: pulumi.CustomResourceOptions): AppSecRateProtection {
        return new AppSecRateProtection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecRateProtection:AppSecRateProtection';

    /**
     * Returns true if the given object is an instance of AppSecRateProtection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecRateProtection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecRateProtection.__pulumiType;
    }

    /**
     * . Unique identifier of the security configuration associated with the rate protection settings being modified.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * . Set to **true** to enable rate protection; set to **false** to disable rate protection.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Text representation
     */
    public /*out*/ readonly outputText!: pulumi.Output<string>;
    /**
     * . Unique identifier of the security policy associated with the rate protection settings being modified.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;

    /**
     * Create a AppSecRateProtection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecRateProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecRateProtectionArgs | AppSecRateProtectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecRateProtectionState | undefined;
            resourceInputs["configId"] = state ? state.configId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["outputText"] = state ? state.outputText : undefined;
            resourceInputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
        } else {
            const args = argsOrState as AppSecRateProtectionArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            resourceInputs["configId"] = args ? args.configId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
            resourceInputs["outputText"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppSecRateProtection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecRateProtection resources.
 */
export interface AppSecRateProtectionState {
    /**
     * . Unique identifier of the security configuration associated with the rate protection settings being modified.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Set to **true** to enable rate protection; set to **false** to disable rate protection.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Text representation
     */
    outputText?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security policy associated with the rate protection settings being modified.
     */
    securityPolicyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecRateProtection resource.
 */
export interface AppSecRateProtectionArgs {
    /**
     * . Unique identifier of the security configuration associated with the rate protection settings being modified.
     */
    configId: pulumi.Input<number>;
    /**
     * . Set to **true** to enable rate protection; set to **false** to disable rate protection.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * . Unique identifier of the security policy associated with the rate protection settings being modified.
     */
    securityPolicyId: pulumi.Input<string>;
}
