// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy
 *
 * Renames an existing security policy. Note that you can only change the name of the policy: once issued, the security policy ID can't be modified.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsecuritypolicy)
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
     * . Unique identifier of the security configuration associated with the security policy being renamed.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * . Unique identifier of the security policy being renamed.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;
    /**
     * . New name to be given to the security policy.
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
     * . Unique identifier of the security configuration associated with the security policy being renamed.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy being renamed.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * . New name to be given to the security policy.
     */
    securityPolicyName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecSecurityPolicyRename resource.
 */
export interface AppSecSecurityPolicyRenameArgs {
    /**
     * . Unique identifier of the security configuration associated with the security policy being renamed.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy being renamed.
     */
    securityPolicyId: pulumi.Input<string>;
    /**
     * . New name to be given to the security policy.
     */
    securityPolicyName: pulumi.Input<string>;
}
