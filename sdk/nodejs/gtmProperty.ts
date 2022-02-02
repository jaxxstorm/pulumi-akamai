// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.GtmProperty` resource to create, configure and import a GTM property, a set of IP addresses or CNAMEs that GTM provides in response to DNS queries based on a set of rules.
 *
 * > **Note** Import requires an ID with this format: `existingDomainName`:`existingPropertyName`.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const demoProperty = new akamai.GtmProperty("demo_property", {
 *     domain: "demo_domain.akadns.net",
 *     handoutLimit: 5,
 *     handoutMode: "normal",
 *     scoreAggregationType: "median",
 *     trafficTargets: [{
 *         datacenterId: 3131,
 *     }],
 *     type: "weighted-round-robin",
 * });
 * ```
 * ## Schema reference
 *
 * You can download the GTM Property backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#property) page.
 */
export class GtmProperty extends pulumi.CustomResource {
    /**
     * Get an existing GtmProperty resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GtmPropertyState, opts?: pulumi.CustomResourceOptions): GtmProperty {
        return new GtmProperty(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/gtmProperty:GtmProperty';

    /**
     * Returns true if the given object is an instance of GtmProperty.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GtmProperty {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GtmProperty.__pulumiType;
    }

    /**
     * Specifies a backup CNAME. If GTM declares that all of the servers configured for your property are down, the backup CNAME is handed out. If a backup CNAME is set, do not set a backup IP.
     */
    public readonly backupCname!: pulumi.Output<string | undefined>;
    /**
     * Specifies a backup IP. When GTM declares that all of the targets are down, the backup IP is handed out. If a backup IP is set, do not set a backup CNAME.
     */
    public readonly backupIp!: pulumi.Output<string | undefined>;
    /**
     * A boolean that indicates whether download score based load balancing is enabled.
     */
    public readonly balanceByDownloadScore!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates the fully qualified name aliased to a particular property.
     */
    public readonly cname!: pulumi.Output<string | undefined>;
    /**
     * A descriptive note about changes to the domain. The maximum is 4000 characters.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * DNS name for the GTM Domain set that includes this Property.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Indicates the TTL in seconds for records that might change dynamically based on liveness and load balancing such as A and AAAA records, and CNAMEs.
     */
    public readonly dynamicTtl!: pulumi.Output<number | undefined>;
    /**
     * Specifies the failback delay in seconds.
     */
    public readonly failbackDelay!: pulumi.Output<number | undefined>;
    /**
     * Specifies the failover delay in seconds.
     */
    public readonly failoverDelay!: pulumi.Output<number | undefined>;
    /**
     * Use load estimates from Akamai Ghost utilization messages.
     */
    public readonly ghostDemandReporting!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates the limit for the number of live IPs handed out to a DNS request.
     */
    public readonly handoutLimit!: pulumi.Output<number>;
    /**
     * Specifies how IPs are returned when more than one IP is alive and available.
     */
    public readonly handoutMode!: pulumi.Output<string>;
    /**
     * Defines the absolute limit beyond which IPs are declared unhealthy.
     */
    public readonly healthMax!: pulumi.Output<number | undefined>;
    /**
     * Configures a cutoff value that is computed from the median scores.
     */
    public readonly healthMultiplier!: pulumi.Output<number | undefined>;
    /**
     * Configures a cutoff value that is computed from the median scores.
     */
    public readonly healthThreshold!: pulumi.Output<number | undefined>;
    /**
     * A boolean that indicates the type of IP address handed out by a GTM property.
     */
    public readonly ipv6!: pulumi.Output<boolean | undefined>;
    /**
     * Contains information about the liveness tests, which are run periodically to determine whether your servers respond to requests. You can have multiple `livenessTest` arguments. If used, requires these arguments:
     */
    public readonly livenessTests!: pulumi.Output<outputs.GtmPropertyLivenessTest[] | undefined>;
    /**
     * Indicates the percent of load imbalance factor (LIF) for the property.
     */
    public readonly loadImbalancePercentage!: pulumi.Output<number | undefined>;
    /**
     * A descriptive label for a GeographicMap or a CidrMap that's required if the property is either geographic or cidrmapping, in which case mapName needs to reference either an existing GeographicMap or CidrMap in the same domain.
     */
    public readonly mapName!: pulumi.Output<string | undefined>;
    /**
     * For performance domains, this specifies a penalty value that's added to liveness test scores when data centers show an aggregated loss fraction higher than the penalty value.
     */
    public readonly maxUnreachablePenalty!: pulumi.Output<number | undefined>;
    /**
     * Specifies what fraction of the servers need to respond to requests so GTM considers the data center up and able to receive traffic.
     */
    public readonly minLiveFraction!: pulumi.Output<number | undefined>;
    /**
     * Name of HTTP header.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies how GTM aggregates liveness test scores across different tests, when multiple tests are configured.
     */
    public readonly scoreAggregationType!: pulumi.Output<string>;
    /**
     * Contains static record sets. You can have multiple `staticRrSet` entries. Requires these arguments:
     */
    public readonly staticRrSets!: pulumi.Output<outputs.GtmPropertyStaticRrSet[] | undefined>;
    public readonly staticTtl!: pulumi.Output<number | undefined>;
    /**
     * Specifies a constant used to configure data center affinity.
     */
    public readonly stickinessBonusConstant!: pulumi.Output<number | undefined>;
    /**
     * Specifies a percentage used to configure data center affinity.
     */
    public readonly stickinessBonusPercentage!: pulumi.Output<number | undefined>;
    /**
     * Contains information about where to direct data center traffic. You can have multiple `trafficTarget` arguments. If used, includes these arguments:
     */
    public readonly trafficTargets!: pulumi.Output<outputs.GtmPropertyTrafficTarget[] | undefined>;
    /**
     * The record type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * For performance domains, this specifies a penalty value that's added to liveness test scores when data centers have an aggregated loss fraction higher than this value.
     */
    public readonly unreachableThreshold!: pulumi.Output<number | undefined>;
    /**
     * For load-feedback domains only, a boolean that indicates whether you want GTM to automatically compute target load.
     */
    public readonly useComputedTargets!: pulumi.Output<boolean | undefined>;
    /**
     * A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     */
    public readonly waitOnComplete!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly weightedHashBitsForIpv4!: pulumi.Output<number>;
    public /*out*/ readonly weightedHashBitsForIpv6!: pulumi.Output<number>;

    /**
     * Create a GtmProperty resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GtmPropertyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GtmPropertyArgs | GtmPropertyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmPropertyState | undefined;
            resourceInputs["backupCname"] = state ? state.backupCname : undefined;
            resourceInputs["backupIp"] = state ? state.backupIp : undefined;
            resourceInputs["balanceByDownloadScore"] = state ? state.balanceByDownloadScore : undefined;
            resourceInputs["cname"] = state ? state.cname : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["dynamicTtl"] = state ? state.dynamicTtl : undefined;
            resourceInputs["failbackDelay"] = state ? state.failbackDelay : undefined;
            resourceInputs["failoverDelay"] = state ? state.failoverDelay : undefined;
            resourceInputs["ghostDemandReporting"] = state ? state.ghostDemandReporting : undefined;
            resourceInputs["handoutLimit"] = state ? state.handoutLimit : undefined;
            resourceInputs["handoutMode"] = state ? state.handoutMode : undefined;
            resourceInputs["healthMax"] = state ? state.healthMax : undefined;
            resourceInputs["healthMultiplier"] = state ? state.healthMultiplier : undefined;
            resourceInputs["healthThreshold"] = state ? state.healthThreshold : undefined;
            resourceInputs["ipv6"] = state ? state.ipv6 : undefined;
            resourceInputs["livenessTests"] = state ? state.livenessTests : undefined;
            resourceInputs["loadImbalancePercentage"] = state ? state.loadImbalancePercentage : undefined;
            resourceInputs["mapName"] = state ? state.mapName : undefined;
            resourceInputs["maxUnreachablePenalty"] = state ? state.maxUnreachablePenalty : undefined;
            resourceInputs["minLiveFraction"] = state ? state.minLiveFraction : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scoreAggregationType"] = state ? state.scoreAggregationType : undefined;
            resourceInputs["staticRrSets"] = state ? state.staticRrSets : undefined;
            resourceInputs["staticTtl"] = state ? state.staticTtl : undefined;
            resourceInputs["stickinessBonusConstant"] = state ? state.stickinessBonusConstant : undefined;
            resourceInputs["stickinessBonusPercentage"] = state ? state.stickinessBonusPercentage : undefined;
            resourceInputs["trafficTargets"] = state ? state.trafficTargets : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["unreachableThreshold"] = state ? state.unreachableThreshold : undefined;
            resourceInputs["useComputedTargets"] = state ? state.useComputedTargets : undefined;
            resourceInputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
            resourceInputs["weightedHashBitsForIpv4"] = state ? state.weightedHashBitsForIpv4 : undefined;
            resourceInputs["weightedHashBitsForIpv6"] = state ? state.weightedHashBitsForIpv6 : undefined;
        } else {
            const args = argsOrState as GtmPropertyArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.handoutLimit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'handoutLimit'");
            }
            if ((!args || args.handoutMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'handoutMode'");
            }
            if ((!args || args.scoreAggregationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scoreAggregationType'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["backupCname"] = args ? args.backupCname : undefined;
            resourceInputs["backupIp"] = args ? args.backupIp : undefined;
            resourceInputs["balanceByDownloadScore"] = args ? args.balanceByDownloadScore : undefined;
            resourceInputs["cname"] = args ? args.cname : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["dynamicTtl"] = args ? args.dynamicTtl : undefined;
            resourceInputs["failbackDelay"] = args ? args.failbackDelay : undefined;
            resourceInputs["failoverDelay"] = args ? args.failoverDelay : undefined;
            resourceInputs["ghostDemandReporting"] = args ? args.ghostDemandReporting : undefined;
            resourceInputs["handoutLimit"] = args ? args.handoutLimit : undefined;
            resourceInputs["handoutMode"] = args ? args.handoutMode : undefined;
            resourceInputs["healthMax"] = args ? args.healthMax : undefined;
            resourceInputs["healthMultiplier"] = args ? args.healthMultiplier : undefined;
            resourceInputs["healthThreshold"] = args ? args.healthThreshold : undefined;
            resourceInputs["ipv6"] = args ? args.ipv6 : undefined;
            resourceInputs["livenessTests"] = args ? args.livenessTests : undefined;
            resourceInputs["loadImbalancePercentage"] = args ? args.loadImbalancePercentage : undefined;
            resourceInputs["mapName"] = args ? args.mapName : undefined;
            resourceInputs["maxUnreachablePenalty"] = args ? args.maxUnreachablePenalty : undefined;
            resourceInputs["minLiveFraction"] = args ? args.minLiveFraction : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scoreAggregationType"] = args ? args.scoreAggregationType : undefined;
            resourceInputs["staticRrSets"] = args ? args.staticRrSets : undefined;
            resourceInputs["staticTtl"] = args ? args.staticTtl : undefined;
            resourceInputs["stickinessBonusConstant"] = args ? args.stickinessBonusConstant : undefined;
            resourceInputs["stickinessBonusPercentage"] = args ? args.stickinessBonusPercentage : undefined;
            resourceInputs["trafficTargets"] = args ? args.trafficTargets : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["unreachableThreshold"] = args ? args.unreachableThreshold : undefined;
            resourceInputs["useComputedTargets"] = args ? args.useComputedTargets : undefined;
            resourceInputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
            resourceInputs["weightedHashBitsForIpv4"] = undefined /*out*/;
            resourceInputs["weightedHashBitsForIpv6"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "akamai:trafficmanagement/gtmProperty:GtmProperty" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GtmProperty.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmProperty resources.
 */
export interface GtmPropertyState {
    /**
     * Specifies a backup CNAME. If GTM declares that all of the servers configured for your property are down, the backup CNAME is handed out. If a backup CNAME is set, do not set a backup IP.
     */
    backupCname?: pulumi.Input<string>;
    /**
     * Specifies a backup IP. When GTM declares that all of the targets are down, the backup IP is handed out. If a backup IP is set, do not set a backup CNAME.
     */
    backupIp?: pulumi.Input<string>;
    /**
     * A boolean that indicates whether download score based load balancing is enabled.
     */
    balanceByDownloadScore?: pulumi.Input<boolean>;
    /**
     * Indicates the fully qualified name aliased to a particular property.
     */
    cname?: pulumi.Input<string>;
    /**
     * A descriptive note about changes to the domain. The maximum is 4000 characters.
     */
    comments?: pulumi.Input<string>;
    /**
     * DNS name for the GTM Domain set that includes this Property.
     */
    domain?: pulumi.Input<string>;
    /**
     * Indicates the TTL in seconds for records that might change dynamically based on liveness and load balancing such as A and AAAA records, and CNAMEs.
     */
    dynamicTtl?: pulumi.Input<number>;
    /**
     * Specifies the failback delay in seconds.
     */
    failbackDelay?: pulumi.Input<number>;
    /**
     * Specifies the failover delay in seconds.
     */
    failoverDelay?: pulumi.Input<number>;
    /**
     * Use load estimates from Akamai Ghost utilization messages.
     */
    ghostDemandReporting?: pulumi.Input<boolean>;
    /**
     * Indicates the limit for the number of live IPs handed out to a DNS request.
     */
    handoutLimit?: pulumi.Input<number>;
    /**
     * Specifies how IPs are returned when more than one IP is alive and available.
     */
    handoutMode?: pulumi.Input<string>;
    /**
     * Defines the absolute limit beyond which IPs are declared unhealthy.
     */
    healthMax?: pulumi.Input<number>;
    /**
     * Configures a cutoff value that is computed from the median scores.
     */
    healthMultiplier?: pulumi.Input<number>;
    /**
     * Configures a cutoff value that is computed from the median scores.
     */
    healthThreshold?: pulumi.Input<number>;
    /**
     * A boolean that indicates the type of IP address handed out by a GTM property.
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * Contains information about the liveness tests, which are run periodically to determine whether your servers respond to requests. You can have multiple `livenessTest` arguments. If used, requires these arguments:
     */
    livenessTests?: pulumi.Input<pulumi.Input<inputs.GtmPropertyLivenessTest>[]>;
    /**
     * Indicates the percent of load imbalance factor (LIF) for the property.
     */
    loadImbalancePercentage?: pulumi.Input<number>;
    /**
     * A descriptive label for a GeographicMap or a CidrMap that's required if the property is either geographic or cidrmapping, in which case mapName needs to reference either an existing GeographicMap or CidrMap in the same domain.
     */
    mapName?: pulumi.Input<string>;
    /**
     * For performance domains, this specifies a penalty value that's added to liveness test scores when data centers show an aggregated loss fraction higher than the penalty value.
     */
    maxUnreachablePenalty?: pulumi.Input<number>;
    /**
     * Specifies what fraction of the servers need to respond to requests so GTM considers the data center up and able to receive traffic.
     */
    minLiveFraction?: pulumi.Input<number>;
    /**
     * Name of HTTP header.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies how GTM aggregates liveness test scores across different tests, when multiple tests are configured.
     */
    scoreAggregationType?: pulumi.Input<string>;
    /**
     * Contains static record sets. You can have multiple `staticRrSet` entries. Requires these arguments:
     */
    staticRrSets?: pulumi.Input<pulumi.Input<inputs.GtmPropertyStaticRrSet>[]>;
    staticTtl?: pulumi.Input<number>;
    /**
     * Specifies a constant used to configure data center affinity.
     */
    stickinessBonusConstant?: pulumi.Input<number>;
    /**
     * Specifies a percentage used to configure data center affinity.
     */
    stickinessBonusPercentage?: pulumi.Input<number>;
    /**
     * Contains information about where to direct data center traffic. You can have multiple `trafficTarget` arguments. If used, includes these arguments:
     */
    trafficTargets?: pulumi.Input<pulumi.Input<inputs.GtmPropertyTrafficTarget>[]>;
    /**
     * The record type.
     */
    type?: pulumi.Input<string>;
    /**
     * For performance domains, this specifies a penalty value that's added to liveness test scores when data centers have an aggregated loss fraction higher than this value.
     */
    unreachableThreshold?: pulumi.Input<number>;
    /**
     * For load-feedback domains only, a boolean that indicates whether you want GTM to automatically compute target load.
     */
    useComputedTargets?: pulumi.Input<boolean>;
    /**
     * A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     */
    waitOnComplete?: pulumi.Input<boolean>;
    weightedHashBitsForIpv4?: pulumi.Input<number>;
    weightedHashBitsForIpv6?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GtmProperty resource.
 */
export interface GtmPropertyArgs {
    /**
     * Specifies a backup CNAME. If GTM declares that all of the servers configured for your property are down, the backup CNAME is handed out. If a backup CNAME is set, do not set a backup IP.
     */
    backupCname?: pulumi.Input<string>;
    /**
     * Specifies a backup IP. When GTM declares that all of the targets are down, the backup IP is handed out. If a backup IP is set, do not set a backup CNAME.
     */
    backupIp?: pulumi.Input<string>;
    /**
     * A boolean that indicates whether download score based load balancing is enabled.
     */
    balanceByDownloadScore?: pulumi.Input<boolean>;
    /**
     * Indicates the fully qualified name aliased to a particular property.
     */
    cname?: pulumi.Input<string>;
    /**
     * A descriptive note about changes to the domain. The maximum is 4000 characters.
     */
    comments?: pulumi.Input<string>;
    /**
     * DNS name for the GTM Domain set that includes this Property.
     */
    domain: pulumi.Input<string>;
    /**
     * Indicates the TTL in seconds for records that might change dynamically based on liveness and load balancing such as A and AAAA records, and CNAMEs.
     */
    dynamicTtl?: pulumi.Input<number>;
    /**
     * Specifies the failback delay in seconds.
     */
    failbackDelay?: pulumi.Input<number>;
    /**
     * Specifies the failover delay in seconds.
     */
    failoverDelay?: pulumi.Input<number>;
    /**
     * Use load estimates from Akamai Ghost utilization messages.
     */
    ghostDemandReporting?: pulumi.Input<boolean>;
    /**
     * Indicates the limit for the number of live IPs handed out to a DNS request.
     */
    handoutLimit: pulumi.Input<number>;
    /**
     * Specifies how IPs are returned when more than one IP is alive and available.
     */
    handoutMode: pulumi.Input<string>;
    /**
     * Defines the absolute limit beyond which IPs are declared unhealthy.
     */
    healthMax?: pulumi.Input<number>;
    /**
     * Configures a cutoff value that is computed from the median scores.
     */
    healthMultiplier?: pulumi.Input<number>;
    /**
     * Configures a cutoff value that is computed from the median scores.
     */
    healthThreshold?: pulumi.Input<number>;
    /**
     * A boolean that indicates the type of IP address handed out by a GTM property.
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * Contains information about the liveness tests, which are run periodically to determine whether your servers respond to requests. You can have multiple `livenessTest` arguments. If used, requires these arguments:
     */
    livenessTests?: pulumi.Input<pulumi.Input<inputs.GtmPropertyLivenessTest>[]>;
    /**
     * Indicates the percent of load imbalance factor (LIF) for the property.
     */
    loadImbalancePercentage?: pulumi.Input<number>;
    /**
     * A descriptive label for a GeographicMap or a CidrMap that's required if the property is either geographic or cidrmapping, in which case mapName needs to reference either an existing GeographicMap or CidrMap in the same domain.
     */
    mapName?: pulumi.Input<string>;
    /**
     * For performance domains, this specifies a penalty value that's added to liveness test scores when data centers show an aggregated loss fraction higher than the penalty value.
     */
    maxUnreachablePenalty?: pulumi.Input<number>;
    /**
     * Specifies what fraction of the servers need to respond to requests so GTM considers the data center up and able to receive traffic.
     */
    minLiveFraction?: pulumi.Input<number>;
    /**
     * Name of HTTP header.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies how GTM aggregates liveness test scores across different tests, when multiple tests are configured.
     */
    scoreAggregationType: pulumi.Input<string>;
    /**
     * Contains static record sets. You can have multiple `staticRrSet` entries. Requires these arguments:
     */
    staticRrSets?: pulumi.Input<pulumi.Input<inputs.GtmPropertyStaticRrSet>[]>;
    staticTtl?: pulumi.Input<number>;
    /**
     * Specifies a constant used to configure data center affinity.
     */
    stickinessBonusConstant?: pulumi.Input<number>;
    /**
     * Specifies a percentage used to configure data center affinity.
     */
    stickinessBonusPercentage?: pulumi.Input<number>;
    /**
     * Contains information about where to direct data center traffic. You can have multiple `trafficTarget` arguments. If used, includes these arguments:
     */
    trafficTargets?: pulumi.Input<pulumi.Input<inputs.GtmPropertyTrafficTarget>[]>;
    /**
     * The record type.
     */
    type: pulumi.Input<string>;
    /**
     * For performance domains, this specifies a penalty value that's added to liveness test scores when data centers have an aggregated loss fraction higher than this value.
     */
    unreachableThreshold?: pulumi.Input<number>;
    /**
     * For load-feedback domains only, a boolean that indicates whether you want GTM to automatically compute target load.
     */
    useComputedTargets?: pulumi.Input<boolean>;
    /**
     * A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}
