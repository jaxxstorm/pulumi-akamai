// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `akamai.AppSecMatchTarget` resource allows you to create or modify a match target associated with a given security configuration version.
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
 *     name: "Akamai Tools",
 * });
 * const matchTarget = new akamai.AppSecMatchTarget("matchTarget", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     version: configuration.then(configuration => configuration.latestVersion),
 *     json: fs.readFileSync(`${path.module}/match_targets.json`),
 * });
 * ```
 */
export class AppSecMatchTarget extends pulumi.CustomResource {
    /**
     * Get an existing AppSecMatchTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecMatchTargetState, opts?: pulumi.CustomResourceOptions): AppSecMatchTarget {
        return new AppSecMatchTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecMatchTarget:AppSecMatchTarget';

    /**
     * Returns true if the given object is an instance of AppSecMatchTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecMatchTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecMatchTarget.__pulumiType;
    }

    /**
     * The ID of the security configuration to use.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * The name of a JSON file containing one or more match target definitions ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postmatchtargets)).
     */
    public readonly json!: pulumi.Output<string>;
    /**
     * The ID of the match target.
     */
    public /*out*/ readonly matchTargetId!: pulumi.Output<number>;
    /**
     * The version number of the security configuration to use.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a AppSecMatchTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecMatchTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecMatchTargetArgs | AppSecMatchTargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecMatchTargetState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["json"] = state ? state.json : undefined;
            inputs["matchTargetId"] = state ? state.matchTargetId : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AppSecMatchTargetArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.json === undefined) && !opts.urn) {
                throw new Error("Missing required property 'json'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["json"] = args ? args.json : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["matchTargetId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecMatchTarget.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecMatchTarget resources.
 */
export interface AppSecMatchTargetState {
    /**
     * The ID of the security configuration to use.
     */
    readonly configId?: pulumi.Input<number>;
    /**
     * The name of a JSON file containing one or more match target definitions ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postmatchtargets)).
     */
    readonly json?: pulumi.Input<string>;
    /**
     * The ID of the match target.
     */
    readonly matchTargetId?: pulumi.Input<number>;
    /**
     * The version number of the security configuration to use.
     */
    readonly version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AppSecMatchTarget resource.
 */
export interface AppSecMatchTargetArgs {
    /**
     * The ID of the security configuration to use.
     */
    readonly configId: pulumi.Input<number>;
    /**
     * The name of a JSON file containing one or more match target definitions ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postmatchtargets)).
     */
    readonly json: pulumi.Input<string>;
    /**
     * The version number of the security configuration to use.
     */
    readonly version: pulumi.Input<number>;
}
