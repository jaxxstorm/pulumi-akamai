// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AppSecWapSelectedHostnames extends pulumi.CustomResource {
    /**
     * Get an existing AppSecWapSelectedHostnames resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecWapSelectedHostnamesState, opts?: pulumi.CustomResourceOptions): AppSecWapSelectedHostnames {
        return new AppSecWapSelectedHostnames(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecWapSelectedHostnames:AppSecWapSelectedHostnames';

    /**
     * Returns true if the given object is an instance of AppSecWapSelectedHostnames.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecWapSelectedHostnames {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecWapSelectedHostnames.__pulumiType;
    }

    /**
     * . Unique identifier of the security configuration associated with the hostnames being protected or evaluated.
     */
    public readonly configId!: pulumi.Output<number>;
    public readonly evaluatedHosts!: pulumi.Output<string[] | undefined>;
    public readonly protectedHosts!: pulumi.Output<string[] | undefined>;
    /**
     * . Unique identifier of the security policy responsible for protecting or evaluating the specified hosts.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;

    /**
     * Create a AppSecWapSelectedHostnames resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecWapSelectedHostnamesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecWapSelectedHostnamesArgs | AppSecWapSelectedHostnamesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecWapSelectedHostnamesState | undefined;
            resourceInputs["configId"] = state ? state.configId : undefined;
            resourceInputs["evaluatedHosts"] = state ? state.evaluatedHosts : undefined;
            resourceInputs["protectedHosts"] = state ? state.protectedHosts : undefined;
            resourceInputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
        } else {
            const args = argsOrState as AppSecWapSelectedHostnamesArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.securityPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityPolicyId'");
            }
            resourceInputs["configId"] = args ? args.configId : undefined;
            resourceInputs["evaluatedHosts"] = args ? args.evaluatedHosts : undefined;
            resourceInputs["protectedHosts"] = args ? args.protectedHosts : undefined;
            resourceInputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppSecWapSelectedHostnames.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecWapSelectedHostnames resources.
 */
export interface AppSecWapSelectedHostnamesState {
    /**
     * . Unique identifier of the security configuration associated with the hostnames being protected or evaluated.
     */
    configId?: pulumi.Input<number>;
    evaluatedHosts?: pulumi.Input<pulumi.Input<string>[]>;
    protectedHosts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * . Unique identifier of the security policy responsible for protecting or evaluating the specified hosts.
     */
    securityPolicyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppSecWapSelectedHostnames resource.
 */
export interface AppSecWapSelectedHostnamesArgs {
    /**
     * . Unique identifier of the security configuration associated with the hostnames being protected or evaluated.
     */
    configId: pulumi.Input<number>;
    evaluatedHosts?: pulumi.Input<pulumi.Input<string>[]>;
    protectedHosts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * . Unique identifier of the security policy responsible for protecting or evaluating the specified hosts.
     */
    securityPolicyId: pulumi.Input<string>;
}
