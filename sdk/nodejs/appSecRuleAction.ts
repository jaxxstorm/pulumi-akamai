// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.AppSecRuleAction` resource to update what action a rule takes when it is triggered.
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
 * const ruleAction = new akamai.AppSecRuleAction("ruleAction", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     version: configuration.then(configuration => configuration.latestVersion),
 *     securityPolicyId: _var.security_policy_id,
 *     ruleId: _var.rule_id,
 *     ruleAction: _var.action,
 * });
 * ```
 */
export class AppSecRuleAction extends pulumi.CustomResource {
    /**
     * Get an existing AppSecRuleAction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecRuleActionState, opts?: pulumi.CustomResourceOptions): AppSecRuleAction {
        return new AppSecRuleAction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecRuleAction:AppSecRuleAction';

    /**
     * Returns true if the given object is an instance of AppSecRuleAction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecRuleAction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecRuleAction.__pulumiType;
    }

    /**
     * The ID of the security configuration to use.
     */
    public readonly configId!: pulumi.Output<number>;
    public readonly ruleAction!: pulumi.Output<string>;
    /**
     * The ID of the rule to use.
     */
    public readonly ruleId!: pulumi.Output<number>;
    /**
     * The ID of the security policy to use.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;
    /**
     * The version number of the security configuration to use.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a AppSecRuleAction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecRuleActionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecRuleActionArgs | AppSecRuleActionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecRuleActionState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["ruleAction"] = state ? state.ruleAction : undefined;
            inputs["ruleId"] = state ? state.ruleId : undefined;
            inputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AppSecRuleActionArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.ruleAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleAction'");
            }
            if ((!args || args.ruleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleId'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["ruleAction"] = args ? args.ruleAction : undefined;
            inputs["ruleId"] = args ? args.ruleId : undefined;
            inputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
            inputs["version"] = args ? args.version : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecRuleAction.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecRuleAction resources.
 */
export interface AppSecRuleActionState {
    /**
     * The ID of the security configuration to use.
     */
    readonly configId?: pulumi.Input<number>;
    readonly ruleAction?: pulumi.Input<string>;
    /**
     * The ID of the rule to use.
     */
    readonly ruleId?: pulumi.Input<number>;
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
 * The set of arguments for constructing a AppSecRuleAction resource.
 */
export interface AppSecRuleActionArgs {
    /**
     * The ID of the security configuration to use.
     */
    readonly configId: pulumi.Input<number>;
    readonly ruleAction: pulumi.Input<string>;
    /**
     * The ID of the rule to use.
     */
    readonly ruleId: pulumi.Input<number>;
    /**
     * The ID of the security policy to use.
     */
    readonly securityPolicyId: pulumi.Input<string>;
    /**
     * The version number of the security configuration to use.
     */
    readonly version: pulumi.Input<number>;
}
