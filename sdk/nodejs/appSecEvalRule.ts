// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Evaluation rule
 *
 * Creates or modifies an evaluation rule's action, conditions, and exceptions.
 * Evaluation rules are Kona Rule Set rules used when running a security configuration in evaluation mode.
 * Changes to these rules do not affect the rules used on your production network.
 *
 * **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval-rules/{ruleId}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevalrule) *and* [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval-rules/{ruleId}/condition-exception](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevalconditionsexceptions)
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 * import * from "fs";
 *
 * const configuration = akamai.getAppSecConfiguration({
 *     name: "Documentation",
 * });
 * const evalRule = new akamai.AppSecEvalRule("evalRule", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: "gms1_134637",
 *     ruleId: 60029316,
 *     ruleAction: "deny",
 *     conditionException: fs.readFileSync(`${path.module}/condition_exception.json`),
 * });
 * ```
 */
export class AppSecEvalRule extends pulumi.CustomResource {
    /**
     * Get an existing AppSecEvalRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecEvalRuleState, opts?: pulumi.CustomResourceOptions): AppSecEvalRule {
        return new AppSecEvalRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecEvalRule:AppSecEvalRule';

    /**
     * Returns true if the given object is an instance of AppSecEvalRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecEvalRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecEvalRule.__pulumiType;
    }

    /**
     * . Path to a JSON file containing the conditions and exceptions to be applied to the evaluation rule. To view a sample JSON file, see the [Modify the conditions and exceptions for an evaluation rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevalconditionsexceptions) section of the Application Security API documentation.
     */
    public readonly conditionException!: pulumi.Output<string | undefined>;
    /**
     * . Unique identifier of the security configuration in evaluation mode.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * . Action to be taken any time the evaluation rule is triggered, Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    public readonly ruleAction!: pulumi.Output<string>;
    /**
     * . Unique identifier of the evaluation rule being modified.
     */
    public readonly ruleId!: pulumi.Output<number>;
    /**
     * . Unique identifier of the security policy associated with the evaluation process.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;

    /**
     * Create a AppSecEvalRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecEvalRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecEvalRuleArgs | AppSecEvalRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecEvalRuleState | undefined;
            inputs["conditionException"] = state ? state.conditionException : undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["ruleAction"] = state ? state.ruleAction : undefined;
            inputs["ruleId"] = state ? state.ruleId : undefined;
            inputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
        } else {
            const args = argsOrState as AppSecEvalRuleArgs | undefined;
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
            inputs["conditionException"] = args ? args.conditionException : undefined;
            inputs["configId"] = args ? args.configId : undefined;
            inputs["ruleAction"] = args ? args.ruleAction : undefined;
            inputs["ruleId"] = args ? args.ruleId : undefined;
            inputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecEvalRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecEvalRule resources.
 */
export interface AppSecEvalRuleState {
    /**
     * . Path to a JSON file containing the conditions and exceptions to be applied to the evaluation rule. To view a sample JSON file, see the [Modify the conditions and exceptions for an evaluation rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevalconditionsexceptions) section of the Application Security API documentation.
     */
    conditionException?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration in evaluation mode.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Action to be taken any time the evaluation rule is triggered, Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    ruleAction?: pulumi.Input<string>;
    /**
     * . Unique identifier of the evaluation rule being modified.
     */
    ruleId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the evaluation process.
     */
    securityPolicyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecEvalRule resource.
 */
export interface AppSecEvalRuleArgs {
    /**
     * . Path to a JSON file containing the conditions and exceptions to be applied to the evaluation rule. To view a sample JSON file, see the [Modify the conditions and exceptions for an evaluation rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevalconditionsexceptions) section of the Application Security API documentation.
     */
    conditionException?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration in evaluation mode.
     */
    configId: pulumi.Input<number>;
    /**
     * . Action to be taken any time the evaluation rule is triggered, Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    ruleAction: pulumi.Input<string>;
    /**
     * . Unique identifier of the evaluation rule being modified.
     */
    ruleId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the evaluation process.
     */
    securityPolicyId: pulumi.Input<string>;
}
