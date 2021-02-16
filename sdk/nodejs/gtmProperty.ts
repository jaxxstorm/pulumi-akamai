// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `akamai.GtmProperty` provides the resource for creating, configuring and importing a gtm property to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existingDomainName`:`existingPropertyName`
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

    public readonly backupCname!: pulumi.Output<string | undefined>;
    public readonly backupIp!: pulumi.Output<string | undefined>;
    /**
     * * `staticTtl`
     * * `unreachableThreshold`
     * * `healthMultiplier`
     * * `dynamicTtl`
     * * `maxUnreachablePenalty`
     * * `mapName`
     * * `loadImbalancePercentage`
     * * `healthMax`
     * * `cname`
     * * `comments`
     * * `ghostDemandReporting`
     * * `minLiveFraction`
     */
    public readonly balanceByDownloadScore!: pulumi.Output<boolean | undefined>;
    public readonly cname!: pulumi.Output<string | undefined>;
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Domain name
     */
    public readonly domain!: pulumi.Output<string>;
    public readonly dynamicTtl!: pulumi.Output<number | undefined>;
    public readonly failbackDelay!: pulumi.Output<number | undefined>;
    public readonly failoverDelay!: pulumi.Output<number | undefined>;
    public readonly ghostDemandReporting!: pulumi.Output<boolean | undefined>;
    public readonly handoutLimit!: pulumi.Output<number>;
    public readonly handoutMode!: pulumi.Output<string>;
    public readonly healthMax!: pulumi.Output<number | undefined>;
    public readonly healthMultiplier!: pulumi.Output<number | undefined>;
    public readonly healthThreshold!: pulumi.Output<number | undefined>;
    /**
     * * `stickinessBonusPercentage`
     * * `stickinessBonusConstant`
     * * `healthThreshold`
     */
    public readonly ipv6!: pulumi.Output<boolean | undefined>;
    public readonly livenessTests!: pulumi.Output<outputs.GtmPropertyLivenessTest[] | undefined>;
    public readonly loadImbalancePercentage!: pulumi.Output<number | undefined>;
    public readonly mapName!: pulumi.Output<string | undefined>;
    public readonly maxUnreachablePenalty!: pulumi.Output<number | undefined>;
    public readonly minLiveFraction!: pulumi.Output<number | undefined>;
    /**
     * Liveness test name
     * * `testInterval`
     * * `testObjectProtocol`
     * * `testTimeout`
     */
    public readonly name!: pulumi.Output<string>;
    public readonly scoreAggregationType!: pulumi.Output<string>;
    /**
     * * `type`
     * * `ttl`
     */
    public readonly staticRrSets!: pulumi.Output<outputs.GtmPropertyStaticRrSet[] | undefined>;
    public readonly staticTtl!: pulumi.Output<number | undefined>;
    public readonly stickinessBonusConstant!: pulumi.Output<number | undefined>;
    public readonly stickinessBonusPercentage!: pulumi.Output<number | undefined>;
    /**
     * * `datacenterId`
     */
    public readonly trafficTargets!: pulumi.Output<outputs.GtmPropertyTrafficTarget[]>;
    /**
     * Property type  
     * * `scoreAggregationType`
     */
    public readonly type!: pulumi.Output<string>;
    public readonly unreachableThreshold!: pulumi.Output<number | undefined>;
    /**
     * * `backupIp`
     */
    public readonly useComputedTargets!: pulumi.Output<boolean | undefined>;
    /**
     * Wait for transaction to complete
     * * `failoverDelay`
     * * `failbackDelay`
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmPropertyState | undefined;
            inputs["backupCname"] = state ? state.backupCname : undefined;
            inputs["backupIp"] = state ? state.backupIp : undefined;
            inputs["balanceByDownloadScore"] = state ? state.balanceByDownloadScore : undefined;
            inputs["cname"] = state ? state.cname : undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["dynamicTtl"] = state ? state.dynamicTtl : undefined;
            inputs["failbackDelay"] = state ? state.failbackDelay : undefined;
            inputs["failoverDelay"] = state ? state.failoverDelay : undefined;
            inputs["ghostDemandReporting"] = state ? state.ghostDemandReporting : undefined;
            inputs["handoutLimit"] = state ? state.handoutLimit : undefined;
            inputs["handoutMode"] = state ? state.handoutMode : undefined;
            inputs["healthMax"] = state ? state.healthMax : undefined;
            inputs["healthMultiplier"] = state ? state.healthMultiplier : undefined;
            inputs["healthThreshold"] = state ? state.healthThreshold : undefined;
            inputs["ipv6"] = state ? state.ipv6 : undefined;
            inputs["livenessTests"] = state ? state.livenessTests : undefined;
            inputs["loadImbalancePercentage"] = state ? state.loadImbalancePercentage : undefined;
            inputs["mapName"] = state ? state.mapName : undefined;
            inputs["maxUnreachablePenalty"] = state ? state.maxUnreachablePenalty : undefined;
            inputs["minLiveFraction"] = state ? state.minLiveFraction : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["scoreAggregationType"] = state ? state.scoreAggregationType : undefined;
            inputs["staticRrSets"] = state ? state.staticRrSets : undefined;
            inputs["staticTtl"] = state ? state.staticTtl : undefined;
            inputs["stickinessBonusConstant"] = state ? state.stickinessBonusConstant : undefined;
            inputs["stickinessBonusPercentage"] = state ? state.stickinessBonusPercentage : undefined;
            inputs["trafficTargets"] = state ? state.trafficTargets : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["unreachableThreshold"] = state ? state.unreachableThreshold : undefined;
            inputs["useComputedTargets"] = state ? state.useComputedTargets : undefined;
            inputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
            inputs["weightedHashBitsForIpv4"] = state ? state.weightedHashBitsForIpv4 : undefined;
            inputs["weightedHashBitsForIpv6"] = state ? state.weightedHashBitsForIpv6 : undefined;
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
            if ((!args || args.trafficTargets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficTargets'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["backupCname"] = args ? args.backupCname : undefined;
            inputs["backupIp"] = args ? args.backupIp : undefined;
            inputs["balanceByDownloadScore"] = args ? args.balanceByDownloadScore : undefined;
            inputs["cname"] = args ? args.cname : undefined;
            inputs["comments"] = args ? args.comments : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["dynamicTtl"] = args ? args.dynamicTtl : undefined;
            inputs["failbackDelay"] = args ? args.failbackDelay : undefined;
            inputs["failoverDelay"] = args ? args.failoverDelay : undefined;
            inputs["ghostDemandReporting"] = args ? args.ghostDemandReporting : undefined;
            inputs["handoutLimit"] = args ? args.handoutLimit : undefined;
            inputs["handoutMode"] = args ? args.handoutMode : undefined;
            inputs["healthMax"] = args ? args.healthMax : undefined;
            inputs["healthMultiplier"] = args ? args.healthMultiplier : undefined;
            inputs["healthThreshold"] = args ? args.healthThreshold : undefined;
            inputs["ipv6"] = args ? args.ipv6 : undefined;
            inputs["livenessTests"] = args ? args.livenessTests : undefined;
            inputs["loadImbalancePercentage"] = args ? args.loadImbalancePercentage : undefined;
            inputs["mapName"] = args ? args.mapName : undefined;
            inputs["maxUnreachablePenalty"] = args ? args.maxUnreachablePenalty : undefined;
            inputs["minLiveFraction"] = args ? args.minLiveFraction : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scoreAggregationType"] = args ? args.scoreAggregationType : undefined;
            inputs["staticRrSets"] = args ? args.staticRrSets : undefined;
            inputs["staticTtl"] = args ? args.staticTtl : undefined;
            inputs["stickinessBonusConstant"] = args ? args.stickinessBonusConstant : undefined;
            inputs["stickinessBonusPercentage"] = args ? args.stickinessBonusPercentage : undefined;
            inputs["trafficTargets"] = args ? args.trafficTargets : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["unreachableThreshold"] = args ? args.unreachableThreshold : undefined;
            inputs["useComputedTargets"] = args ? args.useComputedTargets : undefined;
            inputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
            inputs["weightedHashBitsForIpv4"] = undefined /*out*/;
            inputs["weightedHashBitsForIpv6"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "akamai:trafficmanagement/gtmProperty:GtmProperty" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GtmProperty.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmProperty resources.
 */
export interface GtmPropertyState {
    readonly backupCname?: pulumi.Input<string>;
    readonly backupIp?: pulumi.Input<string>;
    /**
     * * `staticTtl`
     * * `unreachableThreshold`
     * * `healthMultiplier`
     * * `dynamicTtl`
     * * `maxUnreachablePenalty`
     * * `mapName`
     * * `loadImbalancePercentage`
     * * `healthMax`
     * * `cname`
     * * `comments`
     * * `ghostDemandReporting`
     * * `minLiveFraction`
     */
    readonly balanceByDownloadScore?: pulumi.Input<boolean>;
    readonly cname?: pulumi.Input<string>;
    readonly comments?: pulumi.Input<string>;
    /**
     * Domain name
     */
    readonly domain?: pulumi.Input<string>;
    readonly dynamicTtl?: pulumi.Input<number>;
    readonly failbackDelay?: pulumi.Input<number>;
    readonly failoverDelay?: pulumi.Input<number>;
    readonly ghostDemandReporting?: pulumi.Input<boolean>;
    readonly handoutLimit?: pulumi.Input<number>;
    readonly handoutMode?: pulumi.Input<string>;
    readonly healthMax?: pulumi.Input<number>;
    readonly healthMultiplier?: pulumi.Input<number>;
    readonly healthThreshold?: pulumi.Input<number>;
    /**
     * * `stickinessBonusPercentage`
     * * `stickinessBonusConstant`
     * * `healthThreshold`
     */
    readonly ipv6?: pulumi.Input<boolean>;
    readonly livenessTests?: pulumi.Input<pulumi.Input<inputs.GtmPropertyLivenessTest>[]>;
    readonly loadImbalancePercentage?: pulumi.Input<number>;
    readonly mapName?: pulumi.Input<string>;
    readonly maxUnreachablePenalty?: pulumi.Input<number>;
    readonly minLiveFraction?: pulumi.Input<number>;
    /**
     * Liveness test name
     * * `testInterval`
     * * `testObjectProtocol`
     * * `testTimeout`
     */
    readonly name?: pulumi.Input<string>;
    readonly scoreAggregationType?: pulumi.Input<string>;
    /**
     * * `type`
     * * `ttl`
     */
    readonly staticRrSets?: pulumi.Input<pulumi.Input<inputs.GtmPropertyStaticRrSet>[]>;
    readonly staticTtl?: pulumi.Input<number>;
    readonly stickinessBonusConstant?: pulumi.Input<number>;
    readonly stickinessBonusPercentage?: pulumi.Input<number>;
    /**
     * * `datacenterId`
     */
    readonly trafficTargets?: pulumi.Input<pulumi.Input<inputs.GtmPropertyTrafficTarget>[]>;
    /**
     * Property type  
     * * `scoreAggregationType`
     */
    readonly type?: pulumi.Input<string>;
    readonly unreachableThreshold?: pulumi.Input<number>;
    /**
     * * `backupIp`
     */
    readonly useComputedTargets?: pulumi.Input<boolean>;
    /**
     * Wait for transaction to complete
     * * `failoverDelay`
     * * `failbackDelay`
     */
    readonly waitOnComplete?: pulumi.Input<boolean>;
    readonly weightedHashBitsForIpv4?: pulumi.Input<number>;
    readonly weightedHashBitsForIpv6?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GtmProperty resource.
 */
export interface GtmPropertyArgs {
    readonly backupCname?: pulumi.Input<string>;
    readonly backupIp?: pulumi.Input<string>;
    /**
     * * `staticTtl`
     * * `unreachableThreshold`
     * * `healthMultiplier`
     * * `dynamicTtl`
     * * `maxUnreachablePenalty`
     * * `mapName`
     * * `loadImbalancePercentage`
     * * `healthMax`
     * * `cname`
     * * `comments`
     * * `ghostDemandReporting`
     * * `minLiveFraction`
     */
    readonly balanceByDownloadScore?: pulumi.Input<boolean>;
    readonly cname?: pulumi.Input<string>;
    readonly comments?: pulumi.Input<string>;
    /**
     * Domain name
     */
    readonly domain: pulumi.Input<string>;
    readonly dynamicTtl?: pulumi.Input<number>;
    readonly failbackDelay?: pulumi.Input<number>;
    readonly failoverDelay?: pulumi.Input<number>;
    readonly ghostDemandReporting?: pulumi.Input<boolean>;
    readonly handoutLimit: pulumi.Input<number>;
    readonly handoutMode: pulumi.Input<string>;
    readonly healthMax?: pulumi.Input<number>;
    readonly healthMultiplier?: pulumi.Input<number>;
    readonly healthThreshold?: pulumi.Input<number>;
    /**
     * * `stickinessBonusPercentage`
     * * `stickinessBonusConstant`
     * * `healthThreshold`
     */
    readonly ipv6?: pulumi.Input<boolean>;
    readonly livenessTests?: pulumi.Input<pulumi.Input<inputs.GtmPropertyLivenessTest>[]>;
    readonly loadImbalancePercentage?: pulumi.Input<number>;
    readonly mapName?: pulumi.Input<string>;
    readonly maxUnreachablePenalty?: pulumi.Input<number>;
    readonly minLiveFraction?: pulumi.Input<number>;
    /**
     * Liveness test name
     * * `testInterval`
     * * `testObjectProtocol`
     * * `testTimeout`
     */
    readonly name?: pulumi.Input<string>;
    readonly scoreAggregationType: pulumi.Input<string>;
    /**
     * * `type`
     * * `ttl`
     */
    readonly staticRrSets?: pulumi.Input<pulumi.Input<inputs.GtmPropertyStaticRrSet>[]>;
    readonly staticTtl?: pulumi.Input<number>;
    readonly stickinessBonusConstant?: pulumi.Input<number>;
    readonly stickinessBonusPercentage?: pulumi.Input<number>;
    /**
     * * `datacenterId`
     */
    readonly trafficTargets: pulumi.Input<pulumi.Input<inputs.GtmPropertyTrafficTarget>[]>;
    /**
     * Property type  
     * * `scoreAggregationType`
     */
    readonly type: pulumi.Input<string>;
    readonly unreachableThreshold?: pulumi.Input<number>;
    /**
     * * `backupIp`
     */
    readonly useComputedTargets?: pulumi.Input<boolean>;
    /**
     * Wait for transaction to complete
     * * `failoverDelay`
     * * `failbackDelay`
     */
    readonly waitOnComplete?: pulumi.Input<boolean>;
}
