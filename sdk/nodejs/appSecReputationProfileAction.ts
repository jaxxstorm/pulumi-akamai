// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Reputation profile
 *
 * Modifies the action taken when a reputation profile is triggered.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/reputation-profiles/{reputationProfileId}](https://techdocs.akamai.com/application-security/reference/put-reputation-profile-action)
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
 * const appsecReputationProfileAction = new akamai.AppSecReputationProfileAction("appsecReputationProfileAction", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: "gms1_134637",
 *     reputationProfileId: 130713,
 *     action: "alert",
 * });
 * export const reputationProfileId = appsecReputationProfileAction.reputationProfileId;
 * export const reputationProfileAction = appsecReputationProfileAction.action;
 * ```
 */
export class AppSecReputationProfileAction extends pulumi.CustomResource {
    /**
     * Get an existing AppSecReputationProfileAction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecReputationProfileActionState, opts?: pulumi.CustomResourceOptions): AppSecReputationProfileAction {
        return new AppSecReputationProfileAction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecReputationProfileAction:AppSecReputationProfileAction';

    /**
     * Returns true if the given object is an instance of AppSecReputationProfileAction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecReputationProfileAction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecReputationProfileAction.__pulumiType;
    }

    /**
     * . Action taken any time the reputation profile is triggered. Allows values are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * . Unique identifier of the security configuration associated with the reputation profile action being modified.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * . Unique identifier of the reputation profile whose action is being modified.
     */
    public readonly reputationProfileId!: pulumi.Output<number>;
    /**
     * . Unique identifier of the security policy associated with the reputation profile action being modified.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;

    /**
     * Create a AppSecReputationProfileAction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecReputationProfileActionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecReputationProfileActionArgs | AppSecReputationProfileActionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecReputationProfileActionState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["configId"] = state ? state.configId : undefined;
            resourceInputs["reputationProfileId"] = state ? state.reputationProfileId : undefined;
            resourceInputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
        } else {
            const args = argsOrState as AppSecReputationProfileActionArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.reputationProfileId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reputationProfileId'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["configId"] = args ? args.configId : undefined;
            resourceInputs["reputationProfileId"] = args ? args.reputationProfileId : undefined;
            resourceInputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppSecReputationProfileAction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecReputationProfileAction resources.
 */
export interface AppSecReputationProfileActionState {
    /**
     * . Action taken any time the reputation profile is triggered. Allows values are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    action?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration associated with the reputation profile action being modified.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the reputation profile whose action is being modified.
     */
    reputationProfileId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the reputation profile action being modified.
     */
    securityPolicyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecReputationProfileAction resource.
 */
export interface AppSecReputationProfileActionArgs {
    /**
     * . Action taken any time the reputation profile is triggered. Allows values are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    action: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration associated with the reputation profile action being modified.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the reputation profile whose action is being modified.
     */
    reputationProfileId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the reputation profile action being modified.
     */
    securityPolicyId: pulumi.Input<string>;
}
