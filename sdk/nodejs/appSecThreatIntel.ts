// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy
 *
 * Enables or disables threat intelligence for a security policy. This resource is only available to organizations running the Adaptive Security Engine (ASE) beta Please contact your Akamai representative for more information.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/threat-intel](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putthreatintelligence)
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
 * const threatIntel = new akamai.AppSecThreatIntel("threatIntel", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: "gms1_134637",
 *     threatIntel: "on",
 * });
 * ```
 */
export class AppSecThreatIntel extends pulumi.CustomResource {
    /**
     * Get an existing AppSecThreatIntel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecThreatIntelState, opts?: pulumi.CustomResourceOptions): AppSecThreatIntel {
        return new AppSecThreatIntel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecThreatIntel:AppSecThreatIntel';

    /**
     * Returns true if the given object is an instance of AppSecThreatIntel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecThreatIntel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecThreatIntel.__pulumiType;
    }

    /**
     * . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;
    /**
     * . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
     */
    public readonly threatIntel!: pulumi.Output<string>;

    /**
     * Create a AppSecThreatIntel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecThreatIntelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecThreatIntelArgs | AppSecThreatIntelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecThreatIntelState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
            inputs["threatIntel"] = state ? state.threatIntel : undefined;
        } else {
            const args = argsOrState as AppSecThreatIntelArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            if ((!args || args.threatIntel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'threatIntel'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
            inputs["threatIntel"] = args ? args.threatIntel : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecThreatIntel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecThreatIntel resources.
 */
export interface AppSecThreatIntelState {
    /**
     * . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
     */
    threatIntel?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecThreatIntel resource.
 */
export interface AppSecThreatIntelArgs {
    /**
     * . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
     */
    securityPolicyId: pulumi.Input<string>;
    /**
     * . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
     */
    threatIntel: pulumi.Input<string>;
}
