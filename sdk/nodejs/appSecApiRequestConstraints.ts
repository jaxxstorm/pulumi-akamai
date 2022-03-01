// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
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
 * const apiEndpoint = configuration.then(configuration => akamai.getAppSecApiEndpoints({
 *     configId: configuration.configId,
 *     securityPolicyId: "gms1_134637",
 *     apiName: "Contracts",
 * }));
 * const apiRequestConstraints = new akamai.AppSecApiRequestConstraints("apiRequestConstraints", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: "gms1_134637",
 *     apiEndpointId: apiEndpoint.then(apiEndpoint => apiEndpoint.id),
 *     action: "alert",
 * });
 * ```
 */
export class AppSecApiRequestConstraints extends pulumi.CustomResource {
    /**
     * Get an existing AppSecApiRequestConstraints resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecApiRequestConstraintsState, opts?: pulumi.CustomResourceOptions): AppSecApiRequestConstraints {
        return new AppSecApiRequestConstraints(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecApiRequestConstraints:AppSecApiRequestConstraints';

    /**
     * Returns true if the given object is an instance of AppSecApiRequestConstraints.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecApiRequestConstraints {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecApiRequestConstraints.__pulumiType;
    }

    /**
     * . Action to assign to the API request constraint. Allowed values are:
     * - **alert**, Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * . ID of the API endpoint the constraint will be assigned to.
     */
    public readonly apiEndpointId!: pulumi.Output<number | undefined>;
    /**
     * . Unique identifier of the security configuration associated with the API request constraint settings being modified.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * . Unique identifier of the security policy associated with the API request constraint settings being modified.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;

    /**
     * Create a AppSecApiRequestConstraints resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecApiRequestConstraintsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecApiRequestConstraintsArgs | AppSecApiRequestConstraintsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecApiRequestConstraintsState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["apiEndpointId"] = state ? state.apiEndpointId : undefined;
            resourceInputs["configId"] = state ? state.configId : undefined;
            resourceInputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
        } else {
            const args = argsOrState as AppSecApiRequestConstraintsArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["apiEndpointId"] = args ? args.apiEndpointId : undefined;
            resourceInputs["configId"] = args ? args.configId : undefined;
            resourceInputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppSecApiRequestConstraints.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecApiRequestConstraints resources.
 */
export interface AppSecApiRequestConstraintsState {
    /**
     * . Action to assign to the API request constraint. Allowed values are:
     * - **alert**, Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    action?: pulumi.Input<string>;
    /**
     * . ID of the API endpoint the constraint will be assigned to.
     */
    apiEndpointId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security configuration associated with the API request constraint settings being modified.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the API request constraint settings being modified.
     */
    securityPolicyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecApiRequestConstraints resource.
 */
export interface AppSecApiRequestConstraintsArgs {
    /**
     * . Action to assign to the API request constraint. Allowed values are:
     * - **alert**, Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     */
    action: pulumi.Input<string>;
    /**
     * . ID of the API endpoint the constraint will be assigned to.
     */
    apiEndpointId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security configuration associated with the API request constraint settings being modified.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the API request constraint settings being modified.
     */
    securityPolicyId: pulumi.Input<string>;
}
