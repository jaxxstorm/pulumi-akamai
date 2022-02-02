// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.GtmDatacenter` resource to create, configure, and import a GTM data center. A GTM data center represents a customer data center and is also known as a traffic target, a location containing many servers GTM can direct traffic to.
 *
 * GTM uses data centers to scale load balancing. For example, you might have data centers in both New York and Amsterdam and want to balance load between them. You can configure GTM to send US users to the New York data center and European users to the data center in Amsterdam.
 *
 * > **Note** Import requires an ID with this format: `existingDomainName`:`existingDatacenterId`.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const demoDatacenter = new akamai.GtmDatacenter("demo_datacenter", {
 *     domain: "demo_domain.akadns.net",
 *     nickname: "demo_datacenter",
 * });
 * ```
 * ## Schema reference
 *
 * You can download the GTM Data Center backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#datacenter) page.
 */
export class GtmDatacenter extends pulumi.CustomResource {
    /**
     * Get an existing GtmDatacenter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GtmDatacenterState, opts?: pulumi.CustomResourceOptions): GtmDatacenter {
        return new GtmDatacenter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/gtmDatacenter:GtmDatacenter';

    /**
     * Returns true if the given object is an instance of GtmDatacenter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GtmDatacenter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GtmDatacenter.__pulumiType;
    }

    /**
     * The name of the city where the data center is located.
     */
    public readonly city!: pulumi.Output<string | undefined>;
    /**
     * Identifies the data center's `datacenterId` of which this data center is a clone.
     */
    public readonly cloneOf!: pulumi.Output<number | undefined>;
    /**
     * A boolean that, if set to `true`, Akamai's liveness test agents use the Host header configured in the liveness test.
     */
    public readonly cloudServerHostHeaderOverride!: pulumi.Output<boolean | undefined>;
    /**
     * A boolean indicating whether to balance load between two or more servers in a cloud environment.
     */
    public readonly cloudServerTargeting!: pulumi.Output<boolean | undefined>;
    /**
     * A two-letter code that specifies the continent where the data center maps to.
     */
    public readonly continent!: pulumi.Output<string | undefined>;
    /**
     * A two-letter ISO 3166 country code that specifies the country where the data center maps to.
     */
    public readonly country!: pulumi.Output<string | undefined>;
    /**
     * A unique identifier for an existing data center in the domain.
     * * `pingInterval`
     * * `pingPacketSize`
     * * `scorePenalty`
     * * `servermonitorLivenessCount`
     * * `servermonitorLoadCount`
     * * `servermonitorPool`
     */
    public /*out*/ readonly datacenterId!: pulumi.Output<number>;
    /**
     * Specifies the load reporting interface between you and the GTM system. If used, requires these additional arguments:
     */
    public readonly defaultLoadObject!: pulumi.Output<outputs.GtmDatacenterDefaultLoadObject | undefined>;
    /**
     * The GTM domain name for the data center.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Specifies the geographical latitude of the data center's position. See also longitude within this object.
     */
    public readonly latitude!: pulumi.Output<number | undefined>;
    /**
     * Specifies the geographic longitude of the data center's position. See also latitude within this object.
     */
    public readonly longitude!: pulumi.Output<number | undefined>;
    /**
     * A descriptive label for the data center.
     */
    public readonly nickname!: pulumi.Output<string | undefined>;
    public /*out*/ readonly pingInterval!: pulumi.Output<number>;
    public /*out*/ readonly pingPacketSize!: pulumi.Output<number>;
    public /*out*/ readonly scorePenalty!: pulumi.Output<number>;
    public /*out*/ readonly servermonitorLivenessCount!: pulumi.Output<number>;
    public /*out*/ readonly servermonitorLoadCount!: pulumi.Output<number>;
    public /*out*/ readonly servermonitorPool!: pulumi.Output<string>;
    /**
     * Specifies a two-letter ISO 3166 country code for the state or province where the data center is located.
     */
    public readonly stateOrProvince!: pulumi.Output<string | undefined>;
    /**
     * A boolean indicating whether the data center is virtual or physical, the latter meaning the data center has an Akamai Network Agent installed, and its physical location (`latitude`, `longitude`) is fixed. Either `true` if virtual or `false` if physical.
     */
    public /*out*/ readonly virtual!: pulumi.Output<boolean>;
    /**
     * A boolean, that if set to `true`, waits for transaction to complete.
     */
    public readonly waitOnComplete!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GtmDatacenter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GtmDatacenterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GtmDatacenterArgs | GtmDatacenterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmDatacenterState | undefined;
            resourceInputs["city"] = state ? state.city : undefined;
            resourceInputs["cloneOf"] = state ? state.cloneOf : undefined;
            resourceInputs["cloudServerHostHeaderOverride"] = state ? state.cloudServerHostHeaderOverride : undefined;
            resourceInputs["cloudServerTargeting"] = state ? state.cloudServerTargeting : undefined;
            resourceInputs["continent"] = state ? state.continent : undefined;
            resourceInputs["country"] = state ? state.country : undefined;
            resourceInputs["datacenterId"] = state ? state.datacenterId : undefined;
            resourceInputs["defaultLoadObject"] = state ? state.defaultLoadObject : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["latitude"] = state ? state.latitude : undefined;
            resourceInputs["longitude"] = state ? state.longitude : undefined;
            resourceInputs["nickname"] = state ? state.nickname : undefined;
            resourceInputs["pingInterval"] = state ? state.pingInterval : undefined;
            resourceInputs["pingPacketSize"] = state ? state.pingPacketSize : undefined;
            resourceInputs["scorePenalty"] = state ? state.scorePenalty : undefined;
            resourceInputs["servermonitorLivenessCount"] = state ? state.servermonitorLivenessCount : undefined;
            resourceInputs["servermonitorLoadCount"] = state ? state.servermonitorLoadCount : undefined;
            resourceInputs["servermonitorPool"] = state ? state.servermonitorPool : undefined;
            resourceInputs["stateOrProvince"] = state ? state.stateOrProvince : undefined;
            resourceInputs["virtual"] = state ? state.virtual : undefined;
            resourceInputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
        } else {
            const args = argsOrState as GtmDatacenterArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            resourceInputs["city"] = args ? args.city : undefined;
            resourceInputs["cloneOf"] = args ? args.cloneOf : undefined;
            resourceInputs["cloudServerHostHeaderOverride"] = args ? args.cloudServerHostHeaderOverride : undefined;
            resourceInputs["cloudServerTargeting"] = args ? args.cloudServerTargeting : undefined;
            resourceInputs["continent"] = args ? args.continent : undefined;
            resourceInputs["country"] = args ? args.country : undefined;
            resourceInputs["defaultLoadObject"] = args ? args.defaultLoadObject : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["latitude"] = args ? args.latitude : undefined;
            resourceInputs["longitude"] = args ? args.longitude : undefined;
            resourceInputs["nickname"] = args ? args.nickname : undefined;
            resourceInputs["stateOrProvince"] = args ? args.stateOrProvince : undefined;
            resourceInputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
            resourceInputs["datacenterId"] = undefined /*out*/;
            resourceInputs["pingInterval"] = undefined /*out*/;
            resourceInputs["pingPacketSize"] = undefined /*out*/;
            resourceInputs["scorePenalty"] = undefined /*out*/;
            resourceInputs["servermonitorLivenessCount"] = undefined /*out*/;
            resourceInputs["servermonitorLoadCount"] = undefined /*out*/;
            resourceInputs["servermonitorPool"] = undefined /*out*/;
            resourceInputs["virtual"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "akamai:trafficmanagement/gtmDatacenter:GtmDatacenter" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GtmDatacenter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmDatacenter resources.
 */
export interface GtmDatacenterState {
    /**
     * The name of the city where the data center is located.
     */
    city?: pulumi.Input<string>;
    /**
     * Identifies the data center's `datacenterId` of which this data center is a clone.
     */
    cloneOf?: pulumi.Input<number>;
    /**
     * A boolean that, if set to `true`, Akamai's liveness test agents use the Host header configured in the liveness test.
     */
    cloudServerHostHeaderOverride?: pulumi.Input<boolean>;
    /**
     * A boolean indicating whether to balance load between two or more servers in a cloud environment.
     */
    cloudServerTargeting?: pulumi.Input<boolean>;
    /**
     * A two-letter code that specifies the continent where the data center maps to.
     */
    continent?: pulumi.Input<string>;
    /**
     * A two-letter ISO 3166 country code that specifies the country where the data center maps to.
     */
    country?: pulumi.Input<string>;
    /**
     * A unique identifier for an existing data center in the domain.
     * * `pingInterval`
     * * `pingPacketSize`
     * * `scorePenalty`
     * * `servermonitorLivenessCount`
     * * `servermonitorLoadCount`
     * * `servermonitorPool`
     */
    datacenterId?: pulumi.Input<number>;
    /**
     * Specifies the load reporting interface between you and the GTM system. If used, requires these additional arguments:
     */
    defaultLoadObject?: pulumi.Input<inputs.GtmDatacenterDefaultLoadObject>;
    /**
     * The GTM domain name for the data center.
     */
    domain?: pulumi.Input<string>;
    /**
     * Specifies the geographical latitude of the data center's position. See also longitude within this object.
     */
    latitude?: pulumi.Input<number>;
    /**
     * Specifies the geographic longitude of the data center's position. See also latitude within this object.
     */
    longitude?: pulumi.Input<number>;
    /**
     * A descriptive label for the data center.
     */
    nickname?: pulumi.Input<string>;
    pingInterval?: pulumi.Input<number>;
    pingPacketSize?: pulumi.Input<number>;
    scorePenalty?: pulumi.Input<number>;
    servermonitorLivenessCount?: pulumi.Input<number>;
    servermonitorLoadCount?: pulumi.Input<number>;
    servermonitorPool?: pulumi.Input<string>;
    /**
     * Specifies a two-letter ISO 3166 country code for the state or province where the data center is located.
     */
    stateOrProvince?: pulumi.Input<string>;
    /**
     * A boolean indicating whether the data center is virtual or physical, the latter meaning the data center has an Akamai Network Agent installed, and its physical location (`latitude`, `longitude`) is fixed. Either `true` if virtual or `false` if physical.
     */
    virtual?: pulumi.Input<boolean>;
    /**
     * A boolean, that if set to `true`, waits for transaction to complete.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GtmDatacenter resource.
 */
export interface GtmDatacenterArgs {
    /**
     * The name of the city where the data center is located.
     */
    city?: pulumi.Input<string>;
    /**
     * Identifies the data center's `datacenterId` of which this data center is a clone.
     */
    cloneOf?: pulumi.Input<number>;
    /**
     * A boolean that, if set to `true`, Akamai's liveness test agents use the Host header configured in the liveness test.
     */
    cloudServerHostHeaderOverride?: pulumi.Input<boolean>;
    /**
     * A boolean indicating whether to balance load between two or more servers in a cloud environment.
     */
    cloudServerTargeting?: pulumi.Input<boolean>;
    /**
     * A two-letter code that specifies the continent where the data center maps to.
     */
    continent?: pulumi.Input<string>;
    /**
     * A two-letter ISO 3166 country code that specifies the country where the data center maps to.
     */
    country?: pulumi.Input<string>;
    /**
     * Specifies the load reporting interface between you and the GTM system. If used, requires these additional arguments:
     */
    defaultLoadObject?: pulumi.Input<inputs.GtmDatacenterDefaultLoadObject>;
    /**
     * The GTM domain name for the data center.
     */
    domain: pulumi.Input<string>;
    /**
     * Specifies the geographical latitude of the data center's position. See also longitude within this object.
     */
    latitude?: pulumi.Input<number>;
    /**
     * Specifies the geographic longitude of the data center's position. See also latitude within this object.
     */
    longitude?: pulumi.Input<number>;
    /**
     * A descriptive label for the data center.
     */
    nickname?: pulumi.Input<string>;
    /**
     * Specifies a two-letter ISO 3166 country code for the state or province where the data center is located.
     */
    stateOrProvince?: pulumi.Input<string>;
    /**
     * A boolean, that if set to `true`, waits for transaction to complete.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}
