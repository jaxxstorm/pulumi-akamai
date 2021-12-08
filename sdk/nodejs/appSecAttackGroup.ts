// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Attack group
 *
 * Modify an attack group's action, conditions, and exceptions. Attack groups are collections of Kona Rule Set rules used to streamline the management of website protections.
 *
 * **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/attack-groups/{attackGroupId}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroup) *and* [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/attack-groups/{attackGroupId}/condition-exception](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception)
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
 * const attackGroup = new akamai.AppSecAttackGroup("attackGroup", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: "gms1_134637",
 *     attackGroup: "SQL",
 *     attackGroupAction: "deny",
 *     conditionException: fs.readFileSync(`${path.module}/condition_exception.json`),
 * });
 * ```
 */
export class AppSecAttackGroup extends pulumi.CustomResource {
    /**
     * Get an existing AppSecAttackGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecAttackGroupState, opts?: pulumi.CustomResourceOptions): AppSecAttackGroup {
        return new AppSecAttackGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecAttackGroup:AppSecAttackGroup';

    /**
     * Returns true if the given object is an instance of AppSecAttackGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecAttackGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecAttackGroup.__pulumiType;
    }

    /**
     * . Unique name of the attack group being modified.
     */
    public readonly attackGroup!: pulumi.Output<string>;
    /**
     * . Action taken any time the attack group is triggered. Allowed values are:
     * - **alert**. Record information about the request.
     * - **deny**. Block the request,
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    public readonly attackGroupAction!: pulumi.Output<string>;
    /**
     * . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group. You can view a sample JSON file in the [Modify the exceptions of an attack group](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception) section of the Application Security API documentation.
     */
    public readonly conditionException!: pulumi.Output<string | undefined>;
    /**
     * . Unique identifier of the security configuration associated with the attack group being modified.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * . Unique identifier of the security policy associated with the attack group being modified.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;

    /**
     * Create a AppSecAttackGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecAttackGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecAttackGroupArgs | AppSecAttackGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecAttackGroupState | undefined;
            inputs["attackGroup"] = state ? state.attackGroup : undefined;
            inputs["attackGroupAction"] = state ? state.attackGroupAction : undefined;
            inputs["conditionException"] = state ? state.conditionException : undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
        } else {
            const args = argsOrState as AppSecAttackGroupArgs | undefined;
            if ((!args || args.attackGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attackGroup'");
            }
            if ((!args || args.attackGroupAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attackGroupAction'");
            }
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            inputs["attackGroup"] = args ? args.attackGroup : undefined;
            inputs["attackGroupAction"] = args ? args.attackGroupAction : undefined;
            inputs["conditionException"] = args ? args.conditionException : undefined;
            inputs["configId"] = args ? args.configId : undefined;
            inputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecAttackGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecAttackGroup resources.
 */
export interface AppSecAttackGroupState {
    /**
     * . Unique name of the attack group being modified.
     */
    attackGroup?: pulumi.Input<string>;
    /**
     * . Action taken any time the attack group is triggered. Allowed values are:
     * - **alert**. Record information about the request.
     * - **deny**. Block the request,
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    attackGroupAction?: pulumi.Input<string>;
    /**
     * . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group. You can view a sample JSON file in the [Modify the exceptions of an attack group](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception) section of the Application Security API documentation.
     */
    conditionException?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration associated with the attack group being modified.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the attack group being modified.
     */
    securityPolicyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecAttackGroup resource.
 */
export interface AppSecAttackGroupArgs {
    /**
     * . Unique name of the attack group being modified.
     */
    attackGroup: pulumi.Input<string>;
    /**
     * . Action taken any time the attack group is triggered. Allowed values are:
     * - **alert**. Record information about the request.
     * - **deny**. Block the request,
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    attackGroupAction: pulumi.Input<string>;
    /**
     * . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group. You can view a sample JSON file in the [Modify the exceptions of an attack group](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception) section of the Application Security API documentation.
     */
    conditionException?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration associated with the attack group being modified.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the attack group being modified.
     */
    securityPolicyId: pulumi.Input<string>;
}
