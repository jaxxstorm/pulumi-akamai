// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security configuration; security policy
 *
 * Specifies the headers you can exclude from inspection when you are working with a Pragma debug header, a header that provides information about such things as: the edge routers used in a transaction; the Akamai IP addresses involved; whether a request was cached or not; etc. By default, pragma headers are removed from all responses.
 *
 * This operation can be applied at the security configuration level (in which case it applies to all the security policies in the configuration), or can be customized for an individual security policy.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/pragma-header](https://techdocs.akamai.com/application-security/reference/put-policies-pragma-header)
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 * import * as fs from "fs";
 *
 * const configuration = akamai.getAppSecConfiguration({
 *     name: "Documentation",
 * });
 * const pragmaHeader = new akamai.AppSecAdvancedSettingsPragmaHeader("pragmaHeader", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     securityPolicyId: "gms1_134637",
 *     pragmaHeader: fs.readFileSync(`${path.module}/pragma_header.json`),
 * });
 * ```
 */
export class AppSecAdvancedSettingsPragmaHeader extends pulumi.CustomResource {
    /**
     * Get an existing AppSecAdvancedSettingsPragmaHeader resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecAdvancedSettingsPragmaHeaderState, opts?: pulumi.CustomResourceOptions): AppSecAdvancedSettingsPragmaHeader {
        return new AppSecAdvancedSettingsPragmaHeader(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecAdvancedSettingsPragmaHeader:AppSecAdvancedSettingsPragmaHeader';

    /**
     * Returns true if the given object is an instance of AppSecAdvancedSettingsPragmaHeader.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecAdvancedSettingsPragmaHeader {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecAdvancedSettingsPragmaHeader.__pulumiType;
    }

    /**
     * . Unique identifier of the security configuration associated with the pragma header settings being modified.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
     */
    public readonly pragmaHeader!: pulumi.Output<string>;
    /**
     * . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     */
    public readonly securityPolicyId!: pulumi.Output<string | undefined>;

    /**
     * Create a AppSecAdvancedSettingsPragmaHeader resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecAdvancedSettingsPragmaHeaderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecAdvancedSettingsPragmaHeaderArgs | AppSecAdvancedSettingsPragmaHeaderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecAdvancedSettingsPragmaHeaderState | undefined;
            resourceInputs["configId"] = state ? state.configId : undefined;
            resourceInputs["pragmaHeader"] = state ? state.pragmaHeader : undefined;
            resourceInputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
        } else {
            const args = argsOrState as AppSecAdvancedSettingsPragmaHeaderArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.pragmaHeader === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pragmaHeader'");
            }
            resourceInputs["configId"] = args ? args.configId : undefined;
            resourceInputs["pragmaHeader"] = args ? args.pragmaHeader : undefined;
            resourceInputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppSecAdvancedSettingsPragmaHeader.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecAdvancedSettingsPragmaHeader resources.
 */
export interface AppSecAdvancedSettingsPragmaHeaderState {
    /**
     * . Unique identifier of the security configuration associated with the pragma header settings being modified.
     */
    configId?: pulumi.Input<number>;
    /**
     * . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
     */
    pragmaHeader?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     */
    securityPolicyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecAdvancedSettingsPragmaHeader resource.
 */
export interface AppSecAdvancedSettingsPragmaHeaderArgs {
    /**
     * . Unique identifier of the security configuration associated with the pragma header settings being modified.
     */
    configId: pulumi.Input<number>;
    /**
     * . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
     */
    pragmaHeader: pulumi.Input<string>;
    /**
     * . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     */
    securityPolicyId?: pulumi.Input<string>;
}
