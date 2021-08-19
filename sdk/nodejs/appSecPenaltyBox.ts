// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecPenaltyBox` resource to update the penalty box settings for a given security policy.
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
 * const penaltyBox = new akamai.AppSecPenaltyBox("penaltyBox", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: _var.security_policy_id,
 *     penaltyBoxProtection: true,
 *     penaltyBoxAction: _var.action,
 * });
 * ```
 */
export class AppSecPenaltyBox extends pulumi.CustomResource {
    /**
     * Get an existing AppSecPenaltyBox resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecPenaltyBoxState, opts?: pulumi.CustomResourceOptions): AppSecPenaltyBox {
        return new AppSecPenaltyBox(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecPenaltyBox:AppSecPenaltyBox';

    /**
     * Returns true if the given object is an instance of AppSecPenaltyBox.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecPenaltyBox {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecPenaltyBox.__pulumiType;
    }

    /**
     * The ID of the security configuration to use.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penaltyBoxProtection` is set to `false`.
     */
    public readonly penaltyBoxAction!: pulumi.Output<string>;
    /**
     * A boolean value indicating whether to enable penalty box protection.
     */
    public readonly penaltyBoxProtection!: pulumi.Output<boolean>;
    /**
     * The ID of the security policy to use.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;

    /**
     * Create a AppSecPenaltyBox resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecPenaltyBoxArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecPenaltyBoxArgs | AppSecPenaltyBoxState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecPenaltyBoxState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["penaltyBoxAction"] = state ? state.penaltyBoxAction : undefined;
            inputs["penaltyBoxProtection"] = state ? state.penaltyBoxProtection : undefined;
            inputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
        } else {
            const args = argsOrState as AppSecPenaltyBoxArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.penaltyBoxAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'penaltyBoxAction'");
            }
            if ((!args || args.penaltyBoxProtection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'penaltyBoxProtection'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["penaltyBoxAction"] = args ? args.penaltyBoxAction : undefined;
            inputs["penaltyBoxProtection"] = args ? args.penaltyBoxProtection : undefined;
            inputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecPenaltyBox.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecPenaltyBox resources.
 */
export interface AppSecPenaltyBoxState {
    /**
     * The ID of the security configuration to use.
     */
    configId?: pulumi.Input<number>;
    /**
     * The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penaltyBoxProtection` is set to `false`.
     */
    penaltyBoxAction?: pulumi.Input<string>;
    /**
     * A boolean value indicating whether to enable penalty box protection.
     */
    penaltyBoxProtection?: pulumi.Input<boolean>;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecPenaltyBox resource.
 */
export interface AppSecPenaltyBoxArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: pulumi.Input<number>;
    /**
     * The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penaltyBoxProtection` is set to `false`.
     */
    penaltyBoxAction: pulumi.Input<string>;
    /**
     * A boolean value indicating whether to enable penalty box protection.
     */
    penaltyBoxProtection: pulumi.Input<boolean>;
    /**
     * The ID of the security policy to use.
     */
    securityPolicyId: pulumi.Input<string>;
}
