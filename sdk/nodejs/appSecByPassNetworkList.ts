// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security configuration
 *
 * Specifies the networks that appear on the bypass network list. Networks on this list are allowed to bypass the Web Application Firewall.
 *
 * Note that this resource is only applicable to WAP (Web Application Protector) configurations.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/bypass-network-lists](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putbypassnetworklistsforawapconfigversion)
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
 * const bypassNetworkLists = new akamai.AppSecByPassNetworkList("bypassNetworkLists", {
 *     configId: configuration.then(configuration => configuration.configId),
 *     bypassNetworkLists: [
 *         "DocumentationNetworkList",
 *         "TrainingNetworkList",
 *     ],
 * });
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `outputText`. Tabular report showing the updated list of bypass network IDs.
 */
export class AppSecByPassNetworkList extends pulumi.CustomResource {
    /**
     * Get an existing AppSecByPassNetworkList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppSecByPassNetworkListState, opts?: pulumi.CustomResourceOptions): AppSecByPassNetworkList {
        return new AppSecByPassNetworkList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/appSecByPassNetworkList:AppSecByPassNetworkList';

    /**
     * Returns true if the given object is an instance of AppSecByPassNetworkList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppSecByPassNetworkList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppSecByPassNetworkList.__pulumiType;
    }

    /**
     * . JSON array of network IDs that comprise the bypass list.
     */
    public readonly bypassNetworkLists!: pulumi.Output<string[]>;
    /**
     * . Unique identifier of the security configuration associated with the network bypass lists being modified.
     */
    public readonly configId!: pulumi.Output<number>;

    /**
     * Create a AppSecByPassNetworkList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppSecByPassNetworkListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppSecByPassNetworkListArgs | AppSecByPassNetworkListState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppSecByPassNetworkListState | undefined;
            inputs["bypassNetworkLists"] = state ? state.bypassNetworkLists : undefined;
            inputs["configId"] = state ? state.configId : undefined;
        } else {
            const args = argsOrState as AppSecByPassNetworkListArgs | undefined;
            if ((!args || args.bypassNetworkLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bypassNetworkLists'");
            }
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            inputs["bypassNetworkLists"] = args ? args.bypassNetworkLists : undefined;
            inputs["configId"] = args ? args.configId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppSecByPassNetworkList.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppSecByPassNetworkList resources.
 */
export interface AppSecByPassNetworkListState {
    /**
     * . JSON array of network IDs that comprise the bypass list.
     */
    bypassNetworkLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * . Unique identifier of the security configuration associated with the network bypass lists being modified.
     */
    configId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AppSecByPassNetworkList resource.
 */
export interface AppSecByPassNetworkListArgs {
    /**
     * . JSON array of network IDs that comprise the bypass list.
     */
    bypassNetworkLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * . Unique identifier of the security configuration associated with the network bypass lists being modified.
     */
    configId: pulumi.Input<number>;
}
