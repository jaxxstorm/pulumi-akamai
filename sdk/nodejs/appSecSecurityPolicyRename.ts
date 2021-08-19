// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `akamai.AppSecSecurityPolicyRename` resource allows you to rename an existing security policy.
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
 * const securityPolicyRename = new akamai.AppSecSecurityPolicy("securityPolicyRename", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: _var.security_policy_id,
 *     securityPolicyName: _var.name,
 * });
 * ```
 */
export class AppSecSecurityPolicyRename extends pulumi.CustomResource {
    /**
     * Get an existing AppSecSecurityPolicyRename resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecSecurityPolicyRenameState, opts?: pulumi.CustomResourceOptions): AppSecSecurityPolicyRename {
        return new AppSecSecurityPolicyRename(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecSecurityPolicyRename:AppSecSecurityPolicyRename';

    /**
     * Returns true if the given object is an instance of AppSecSecurityPolicyRename.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecSecurityPolicyRename {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecSecurityPolicyRename.__pulumiType;
    }

    /**
     * The ID of the security configuration to use.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * The ID of the security policy to be renamed.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;
    /**
     * The new name to be given to the security policy.
     */
    public readonly securityPolicyName!: pulumi.Output<string>;

    /**
     * Create a AppSecSecurityPolicyRename resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecSecurityPolicyRenameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecSecurityPolicyRenameArgs | AppSecSecurityPolicyRenameState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecSecurityPolicyRenameState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
            inputs["securityPolicyName"] = state ? state.securityPolicyName : undefined;
        } else {
            const args = argsOrState as AppSecSecurityPolicyRenameArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            if ((!args || args.securityPolicyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyName'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
            inputs["securityPolicyName"] = args ? args.securityPolicyName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecSecurityPolicyRename.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecSecurityPolicyRename resources.
 */
export interface AppSecSecurityPolicyRenameState {
    /**
     * The ID of the security configuration to use.
     */
    configId?: pulumi.Input<number>;
    /**
     * The ID of the security policy to be renamed.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * The new name to be given to the security policy.
     */
    securityPolicyName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecSecurityPolicyRename resource.
 */
export interface AppSecSecurityPolicyRenameArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: pulumi.Input<number>;
    /**
     * The ID of the security policy to be renamed.
     */
    securityPolicyId: pulumi.Input<string>;
    /**
     * The new name to be given to the security policy.
     */
    securityPolicyName: pulumi.Input<string>;
}