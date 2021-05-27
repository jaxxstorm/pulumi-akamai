// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The `akamai.GtmResource` lets you create, configure, and import a GTM resource. In GTM, a resource is anything you can measure whose scarcity affects load balancing. Examples of resources include bandwidth, CPU load average, database queries per second, or disk operations per second.
 *
 * > **Note** Import requires an ID with this format: `existingDomainName`:
 * `existingResourceName`.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const demoResource = new akamai.GtmResource("demo_resource", {
 *     aggregationType: "latest",
 *     domain: "demo_domain.akadns.net",
 *     type: "XML load object via HTTP",
 * });
 * ```
 * ## Argument reference
 *
 * This resource supports these arguments:
 *
 * * `domain` - (Required) DNS name for the GTM Domain set that includes this property.
 * * `name` - (Required) A descriptive label for the GTM resource.
 * * `aggregationType` - (Required) Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
 * * `type` - (Required) Indicates the kind of `loadObject` format used to determine the load on the resource.
 * * `waitOnComplete` - (Optional) A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
 * * `resourceInstance`  - (Optional) (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
 *   * `datacenterId` - (Optional) A unique identifier for an existing data center in the domain.
 *   * `loadObject` - (Optional) Identifies the load object file used to report real-time information about the current load, maximum allowable load, and target load on each resource.
 *   * `loadObjectPort` - (Optional) Specifies the TCP port of the `loadObject`.
 *   * `loadServers` - (Optional) (List) Specifies a list of servers from which to request the load object.
 *   * `useDefaultLoadObject` - (Optional) A boolean that indicates whether a default `loadObject` is used for the resources.
 * * `hostHeader` - (Optional) Optionally specifies the host header used when fetching the load object.
 * * `leastSquaresDecay` - (Optional) For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
 * * `upperBound` - (Optional) An optional sanity check that specifies the maximum allowed value for any component of the load object.
 * * `description` - (Optional) A descriptive note to help you track what the resource constrains.
 * * `leaderString` - (Optional) Specifies the text that comes before the `loadObject`.
 * * `constrainedProperty` - (Optional) Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
 * * `loadImbalancePercent` - (Optional) Indicates the percent of load imbalance factor (LIF) for the property.
 * * `maxUMultiplicativeIncrement` - (Optional) For Akamai internal use only. You can omit the value or set it to `null`.
 * * `decayRate` - (Optional) For Akamai internal use only. You can omit the value or set it to `null`.
 *
 * ## Schema reference
 *
 * You can download the GTM Resource backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#resource) page.
 */
export class GtmResource extends pulumi.CustomResource {
    /**
     * Get an existing GtmResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GtmResourceState, opts?: pulumi.CustomResourceOptions): GtmResource {
        return new GtmResource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/gtmResource:GtmResource';

    /**
     * Returns true if the given object is an instance of GtmResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GtmResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GtmResource.__pulumiType;
    }

    public readonly aggregationType!: pulumi.Output<string>;
    public readonly constrainedProperty!: pulumi.Output<string | undefined>;
    public readonly decayRate!: pulumi.Output<number | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly domain!: pulumi.Output<string>;
    public readonly hostHeader!: pulumi.Output<string | undefined>;
    public readonly leaderString!: pulumi.Output<string | undefined>;
    public readonly leastSquaresDecay!: pulumi.Output<number | undefined>;
    public readonly loadImbalancePercentage!: pulumi.Output<number | undefined>;
    public readonly maxUMultiplicativeIncrement!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly resourceInstances!: pulumi.Output<outputs.GtmResourceResourceInstance[] | undefined>;
    public readonly type!: pulumi.Output<string>;
    public readonly upperBound!: pulumi.Output<number | undefined>;
    public readonly waitOnComplete!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GtmResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GtmResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GtmResourceArgs | GtmResourceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmResourceState | undefined;
            inputs["aggregationType"] = state ? state.aggregationType : undefined;
            inputs["constrainedProperty"] = state ? state.constrainedProperty : undefined;
            inputs["decayRate"] = state ? state.decayRate : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["hostHeader"] = state ? state.hostHeader : undefined;
            inputs["leaderString"] = state ? state.leaderString : undefined;
            inputs["leastSquaresDecay"] = state ? state.leastSquaresDecay : undefined;
            inputs["loadImbalancePercentage"] = state ? state.loadImbalancePercentage : undefined;
            inputs["maxUMultiplicativeIncrement"] = state ? state.maxUMultiplicativeIncrement : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceInstances"] = state ? state.resourceInstances : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["upperBound"] = state ? state.upperBound : undefined;
            inputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
        } else {
            const args = argsOrState as GtmResourceArgs | undefined;
            if ((!args || args.aggregationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aggregationType'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["aggregationType"] = args ? args.aggregationType : undefined;
            inputs["constrainedProperty"] = args ? args.constrainedProperty : undefined;
            inputs["decayRate"] = args ? args.decayRate : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["hostHeader"] = args ? args.hostHeader : undefined;
            inputs["leaderString"] = args ? args.leaderString : undefined;
            inputs["leastSquaresDecay"] = args ? args.leastSquaresDecay : undefined;
            inputs["loadImbalancePercentage"] = args ? args.loadImbalancePercentage : undefined;
            inputs["maxUMultiplicativeIncrement"] = args ? args.maxUMultiplicativeIncrement : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceInstances"] = args ? args.resourceInstances : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["upperBound"] = args ? args.upperBound : undefined;
            inputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "akamai:trafficmanagement/gtmResource:GtmResource" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GtmResource.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmResource resources.
 */
export interface GtmResourceState {
    aggregationType?: pulumi.Input<string>;
    constrainedProperty?: pulumi.Input<string>;
    decayRate?: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    domain?: pulumi.Input<string>;
    hostHeader?: pulumi.Input<string>;
    leaderString?: pulumi.Input<string>;
    leastSquaresDecay?: pulumi.Input<number>;
    loadImbalancePercentage?: pulumi.Input<number>;
    maxUMultiplicativeIncrement?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    resourceInstances?: pulumi.Input<pulumi.Input<inputs.GtmResourceResourceInstance>[]>;
    type?: pulumi.Input<string>;
    upperBound?: pulumi.Input<number>;
    waitOnComplete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GtmResource resource.
 */
export interface GtmResourceArgs {
    aggregationType: pulumi.Input<string>;
    constrainedProperty?: pulumi.Input<string>;
    decayRate?: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    domain: pulumi.Input<string>;
    hostHeader?: pulumi.Input<string>;
    leaderString?: pulumi.Input<string>;
    leastSquaresDecay?: pulumi.Input<number>;
    loadImbalancePercentage?: pulumi.Input<number>;
    maxUMultiplicativeIncrement?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    resourceInstances?: pulumi.Input<pulumi.Input<inputs.GtmResourceResourceInstance>[]>;
    type: pulumi.Input<string>;
    upperBound?: pulumi.Input<number>;
    waitOnComplete?: pulumi.Input<boolean>;
}
