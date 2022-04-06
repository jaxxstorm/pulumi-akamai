// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

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
 * ## Schema reference
 *
 * You can download the GTM Resource backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#resource) page.
 *
 * @deprecated akamai.trafficmanagement.GtmResource has been deprecated in favor of akamai.GtmResource
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
        pulumi.log.warn("GtmResource is deprecated: akamai.trafficmanagement.GtmResource has been deprecated in favor of akamai.GtmResource")
        return new GtmResource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:trafficmanagement/gtmResource:GtmResource';

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

    /**
     * Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
     */
    public readonly aggregationType!: pulumi.Output<string>;
    /**
     * Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
     */
    public readonly constrainedProperty!: pulumi.Output<string | undefined>;
    /**
     * For Akamai internal use only. You can omit the value or set it to `null`.
     */
    public readonly decayRate!: pulumi.Output<number | undefined>;
    /**
     * A descriptive note to help you track what the resource constrains.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * DNS name for the GTM Domain set that includes this property.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Optionally specifies the host header used when fetching the load object.
     */
    public readonly hostHeader!: pulumi.Output<string | undefined>;
    /**
     * Specifies the text that comes before the `loadObject`.
     */
    public readonly leaderString!: pulumi.Output<string | undefined>;
    /**
     * For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
     */
    public readonly leastSquaresDecay!: pulumi.Output<number | undefined>;
    public readonly loadImbalancePercentage!: pulumi.Output<number | undefined>;
    /**
     * For Akamai internal use only. You can omit the value or set it to `null`.
     */
    public readonly maxUMultiplicativeIncrement!: pulumi.Output<number | undefined>;
    /**
     * A descriptive label for the GTM resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
     */
    public readonly resourceInstances!: pulumi.Output<outputs.trafficmanagement.GtmResourceResourceInstance[] | undefined>;
    /**
     * Indicates the kind of `loadObject` format used to determine the load on the resource.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * An optional sanity check that specifies the maximum allowed value for any component of the load object.
     */
    public readonly upperBound!: pulumi.Output<number | undefined>;
    /**
     * A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     */
    public readonly waitOnComplete!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GtmResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.trafficmanagement.GtmResource has been deprecated in favor of akamai.GtmResource */
    constructor(name: string, args: GtmResourceArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.trafficmanagement.GtmResource has been deprecated in favor of akamai.GtmResource */
    constructor(name: string, argsOrState?: GtmResourceArgs | GtmResourceState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GtmResource is deprecated: akamai.trafficmanagement.GtmResource has been deprecated in favor of akamai.GtmResource")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmResourceState | undefined;
            resourceInputs["aggregationType"] = state ? state.aggregationType : undefined;
            resourceInputs["constrainedProperty"] = state ? state.constrainedProperty : undefined;
            resourceInputs["decayRate"] = state ? state.decayRate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["hostHeader"] = state ? state.hostHeader : undefined;
            resourceInputs["leaderString"] = state ? state.leaderString : undefined;
            resourceInputs["leastSquaresDecay"] = state ? state.leastSquaresDecay : undefined;
            resourceInputs["loadImbalancePercentage"] = state ? state.loadImbalancePercentage : undefined;
            resourceInputs["maxUMultiplicativeIncrement"] = state ? state.maxUMultiplicativeIncrement : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceInstances"] = state ? state.resourceInstances : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["upperBound"] = state ? state.upperBound : undefined;
            resourceInputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
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
            resourceInputs["aggregationType"] = args ? args.aggregationType : undefined;
            resourceInputs["constrainedProperty"] = args ? args.constrainedProperty : undefined;
            resourceInputs["decayRate"] = args ? args.decayRate : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["hostHeader"] = args ? args.hostHeader : undefined;
            resourceInputs["leaderString"] = args ? args.leaderString : undefined;
            resourceInputs["leastSquaresDecay"] = args ? args.leastSquaresDecay : undefined;
            resourceInputs["loadImbalancePercentage"] = args ? args.loadImbalancePercentage : undefined;
            resourceInputs["maxUMultiplicativeIncrement"] = args ? args.maxUMultiplicativeIncrement : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceInstances"] = args ? args.resourceInstances : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["upperBound"] = args ? args.upperBound : undefined;
            resourceInputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GtmResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmResource resources.
 */
export interface GtmResourceState {
    /**
     * Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
     */
    aggregationType?: pulumi.Input<string>;
    /**
     * Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
     */
    constrainedProperty?: pulumi.Input<string>;
    /**
     * For Akamai internal use only. You can omit the value or set it to `null`.
     */
    decayRate?: pulumi.Input<number>;
    /**
     * A descriptive note to help you track what the resource constrains.
     */
    description?: pulumi.Input<string>;
    /**
     * DNS name for the GTM Domain set that includes this property.
     */
    domain?: pulumi.Input<string>;
    /**
     * Optionally specifies the host header used when fetching the load object.
     */
    hostHeader?: pulumi.Input<string>;
    /**
     * Specifies the text that comes before the `loadObject`.
     */
    leaderString?: pulumi.Input<string>;
    /**
     * For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
     */
    leastSquaresDecay?: pulumi.Input<number>;
    loadImbalancePercentage?: pulumi.Input<number>;
    /**
     * For Akamai internal use only. You can omit the value or set it to `null`.
     */
    maxUMultiplicativeIncrement?: pulumi.Input<number>;
    /**
     * A descriptive label for the GTM resource.
     */
    name?: pulumi.Input<string>;
    /**
     * (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
     */
    resourceInstances?: pulumi.Input<pulumi.Input<inputs.trafficmanagement.GtmResourceResourceInstance>[]>;
    /**
     * Indicates the kind of `loadObject` format used to determine the load on the resource.
     */
    type?: pulumi.Input<string>;
    /**
     * An optional sanity check that specifies the maximum allowed value for any component of the load object.
     */
    upperBound?: pulumi.Input<number>;
    /**
     * A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GtmResource resource.
 */
export interface GtmResourceArgs {
    /**
     * Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
     */
    aggregationType: pulumi.Input<string>;
    /**
     * Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
     */
    constrainedProperty?: pulumi.Input<string>;
    /**
     * For Akamai internal use only. You can omit the value or set it to `null`.
     */
    decayRate?: pulumi.Input<number>;
    /**
     * A descriptive note to help you track what the resource constrains.
     */
    description?: pulumi.Input<string>;
    /**
     * DNS name for the GTM Domain set that includes this property.
     */
    domain: pulumi.Input<string>;
    /**
     * Optionally specifies the host header used when fetching the load object.
     */
    hostHeader?: pulumi.Input<string>;
    /**
     * Specifies the text that comes before the `loadObject`.
     */
    leaderString?: pulumi.Input<string>;
    /**
     * For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
     */
    leastSquaresDecay?: pulumi.Input<number>;
    loadImbalancePercentage?: pulumi.Input<number>;
    /**
     * For Akamai internal use only. You can omit the value or set it to `null`.
     */
    maxUMultiplicativeIncrement?: pulumi.Input<number>;
    /**
     * A descriptive label for the GTM resource.
     */
    name?: pulumi.Input<string>;
    /**
     * (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
     */
    resourceInstances?: pulumi.Input<pulumi.Input<inputs.trafficmanagement.GtmResourceResourceInstance>[]>;
    /**
     * Indicates the kind of `loadObject` format used to determine the load on the resource.
     */
    type: pulumi.Input<string>;
    /**
     * An optional sanity check that specifies the maximum allowed value for any component of the load object.
     */
    upperBound?: pulumi.Input<number>;
    /**
     * A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}
