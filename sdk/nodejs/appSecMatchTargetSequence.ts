// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `akamai.AppSecMatchTargetSequence` resource allows you to specify the order in which match targets are applied within a given security configuration.
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
 * const matchTargetSequence = new akamai.AppSecMatchTargetSequence("matchTargetSequence", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     matchTargetSequence: fs.readFileSync(`${path.module}/match_targets.json`),
 * });
 * ```
 */
export class AppSecMatchTargetSequence extends pulumi.CustomResource {
    /**
     * Get an existing AppSecMatchTargetSequence resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecMatchTargetSequenceState, opts?: pulumi.CustomResourceOptions): AppSecMatchTargetSequence {
        return new AppSecMatchTargetSequence(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence';

    /**
     * Returns true if the given object is an instance of AppSecMatchTargetSequence.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecMatchTargetSequence {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecMatchTargetSequence.__pulumiType;
    }

    /**
     * The ID of the security configuration to use.
     */
    public readonly configId!: pulumi.Output<number>;
    /**
     * The name of a JSON file containing the sequence of all match targets defined for the specified security configuration ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
     */
    public readonly matchTargetSequence!: pulumi.Output<string | undefined>;

    /**
     * Create a AppSecMatchTargetSequence resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecMatchTargetSequenceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecMatchTargetSequenceArgs | AppSecMatchTargetSequenceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecMatchTargetSequenceState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["matchTargetSequence"] = state ? state.matchTargetSequence : undefined;
        } else {
            const args = argsOrState as AppSecMatchTargetSequenceArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["matchTargetSequence"] = args ? args.matchTargetSequence : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecMatchTargetSequence.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecMatchTargetSequence resources.
 */
export interface AppSecMatchTargetSequenceState {
    /**
     * The ID of the security configuration to use.
     */
    configId?: pulumi.Input<number>;
    /**
     * The name of a JSON file containing the sequence of all match targets defined for the specified security configuration ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
     */
    matchTargetSequence?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecMatchTargetSequence resource.
 */
export interface AppSecMatchTargetSequenceArgs {
    /**
     * The ID of the security configuration to use.
     */
    configId: pulumi.Input<number>;
    /**
     * The name of a JSON file containing the sequence of all match targets defined for the specified security configuration ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
     */
    matchTargetSequence?: pulumi.Input<string>;
}
