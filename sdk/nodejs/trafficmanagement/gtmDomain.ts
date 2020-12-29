// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * @deprecated akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain
 */
export class GtmDomain extends pulumi.CustomResource {
    /**
     * Get an existing GtmDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GtmDomainState, opts?: pulumi.CustomResourceOptions): GtmDomain {
        pulumi.log.warn("GtmDomain is deprecated: akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain")
        return new GtmDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:trafficmanagement/gtmDomain:GtmDomain';

    /**
     * Returns true if the given object is an instance of GtmDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GtmDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GtmDomain.__pulumiType;
    }

    public readonly cnameCoalescingEnabled!: pulumi.Output<boolean | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly contract!: pulumi.Output<string | undefined>;
    public readonly defaultErrorPenalty!: pulumi.Output<number | undefined>;
    public /*out*/ readonly defaultHealthMax!: pulumi.Output<number>;
    public /*out*/ readonly defaultHealthMultiplier!: pulumi.Output<number>;
    public /*out*/ readonly defaultHealthThreshold!: pulumi.Output<number>;
    public /*out*/ readonly defaultMaxUnreachablePenalty!: pulumi.Output<number>;
    public readonly defaultSslClientCertificate!: pulumi.Output<string | undefined>;
    public readonly defaultSslClientPrivateKey!: pulumi.Output<string | undefined>;
    public readonly defaultTimeoutPenalty!: pulumi.Output<number | undefined>;
    public /*out*/ readonly defaultUnreachableThreshold!: pulumi.Output<number>;
    public readonly emailNotificationLists!: pulumi.Output<string[] | undefined>;
    public readonly endUserMappingEnabled!: pulumi.Output<boolean | undefined>;
    public readonly group!: pulumi.Output<string | undefined>;
    public readonly loadFeedback!: pulumi.Output<boolean | undefined>;
    public readonly loadImbalancePercentage!: pulumi.Output<number | undefined>;
    public /*out*/ readonly mapUpdateInterval!: pulumi.Output<number>;
    public /*out*/ readonly maxProperties!: pulumi.Output<number>;
    public /*out*/ readonly maxResources!: pulumi.Output<number>;
    public /*out*/ readonly maxTestTimeout!: pulumi.Output<number>;
    public /*out*/ readonly maxTtl!: pulumi.Output<number>;
    public /*out*/ readonly minPingableRegionFraction!: pulumi.Output<number>;
    public /*out*/ readonly minTestInterval!: pulumi.Output<number>;
    public /*out*/ readonly minTtl!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly pingInterval!: pulumi.Output<number>;
    public /*out*/ readonly pingPacketSize!: pulumi.Output<number>;
    public /*out*/ readonly roundRobinPrefix!: pulumi.Output<string>;
    public /*out*/ readonly servermonitorLivenessCount!: pulumi.Output<number>;
    public /*out*/ readonly servermonitorLoadCount!: pulumi.Output<number>;
    public /*out*/ readonly servermonitorPool!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string>;
    public readonly waitOnComplete!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GtmDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain */
    constructor(name: string, args: GtmDomainArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain */
    constructor(name: string, argsOrState?: GtmDomainArgs | GtmDomainState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GtmDomain is deprecated: akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GtmDomainState | undefined;
            inputs["cnameCoalescingEnabled"] = state ? state.cnameCoalescingEnabled : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["contract"] = state ? state.contract : undefined;
            inputs["defaultErrorPenalty"] = state ? state.defaultErrorPenalty : undefined;
            inputs["defaultHealthMax"] = state ? state.defaultHealthMax : undefined;
            inputs["defaultHealthMultiplier"] = state ? state.defaultHealthMultiplier : undefined;
            inputs["defaultHealthThreshold"] = state ? state.defaultHealthThreshold : undefined;
            inputs["defaultMaxUnreachablePenalty"] = state ? state.defaultMaxUnreachablePenalty : undefined;
            inputs["defaultSslClientCertificate"] = state ? state.defaultSslClientCertificate : undefined;
            inputs["defaultSslClientPrivateKey"] = state ? state.defaultSslClientPrivateKey : undefined;
            inputs["defaultTimeoutPenalty"] = state ? state.defaultTimeoutPenalty : undefined;
            inputs["defaultUnreachableThreshold"] = state ? state.defaultUnreachableThreshold : undefined;
            inputs["emailNotificationLists"] = state ? state.emailNotificationLists : undefined;
            inputs["endUserMappingEnabled"] = state ? state.endUserMappingEnabled : undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["loadFeedback"] = state ? state.loadFeedback : undefined;
            inputs["loadImbalancePercentage"] = state ? state.loadImbalancePercentage : undefined;
            inputs["mapUpdateInterval"] = state ? state.mapUpdateInterval : undefined;
            inputs["maxProperties"] = state ? state.maxProperties : undefined;
            inputs["maxResources"] = state ? state.maxResources : undefined;
            inputs["maxTestTimeout"] = state ? state.maxTestTimeout : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["minPingableRegionFraction"] = state ? state.minPingableRegionFraction : undefined;
            inputs["minTestInterval"] = state ? state.minTestInterval : undefined;
            inputs["minTtl"] = state ? state.minTtl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pingInterval"] = state ? state.pingInterval : undefined;
            inputs["pingPacketSize"] = state ? state.pingPacketSize : undefined;
            inputs["roundRobinPrefix"] = state ? state.roundRobinPrefix : undefined;
            inputs["servermonitorLivenessCount"] = state ? state.servermonitorLivenessCount : undefined;
            inputs["servermonitorLoadCount"] = state ? state.servermonitorLoadCount : undefined;
            inputs["servermonitorPool"] = state ? state.servermonitorPool : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
        } else {
            const args = argsOrState as GtmDomainArgs | undefined;
            if ((!args || args.type === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'type'");
            }
            inputs["cnameCoalescingEnabled"] = args ? args.cnameCoalescingEnabled : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["contract"] = args ? args.contract : undefined;
            inputs["defaultErrorPenalty"] = args ? args.defaultErrorPenalty : undefined;
            inputs["defaultSslClientCertificate"] = args ? args.defaultSslClientCertificate : undefined;
            inputs["defaultSslClientPrivateKey"] = args ? args.defaultSslClientPrivateKey : undefined;
            inputs["defaultTimeoutPenalty"] = args ? args.defaultTimeoutPenalty : undefined;
            inputs["emailNotificationLists"] = args ? args.emailNotificationLists : undefined;
            inputs["endUserMappingEnabled"] = args ? args.endUserMappingEnabled : undefined;
            inputs["group"] = args ? args.group : undefined;
            inputs["loadFeedback"] = args ? args.loadFeedback : undefined;
            inputs["loadImbalancePercentage"] = args ? args.loadImbalancePercentage : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
            inputs["defaultHealthMax"] = undefined /*out*/;
            inputs["defaultHealthMultiplier"] = undefined /*out*/;
            inputs["defaultHealthThreshold"] = undefined /*out*/;
            inputs["defaultMaxUnreachablePenalty"] = undefined /*out*/;
            inputs["defaultUnreachableThreshold"] = undefined /*out*/;
            inputs["mapUpdateInterval"] = undefined /*out*/;
            inputs["maxProperties"] = undefined /*out*/;
            inputs["maxResources"] = undefined /*out*/;
            inputs["maxTestTimeout"] = undefined /*out*/;
            inputs["maxTtl"] = undefined /*out*/;
            inputs["minPingableRegionFraction"] = undefined /*out*/;
            inputs["minTestInterval"] = undefined /*out*/;
            inputs["minTtl"] = undefined /*out*/;
            inputs["pingInterval"] = undefined /*out*/;
            inputs["pingPacketSize"] = undefined /*out*/;
            inputs["roundRobinPrefix"] = undefined /*out*/;
            inputs["servermonitorLivenessCount"] = undefined /*out*/;
            inputs["servermonitorLoadCount"] = undefined /*out*/;
            inputs["servermonitorPool"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GtmDomain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmDomain resources.
 */
export interface GtmDomainState {
    readonly cnameCoalescingEnabled?: pulumi.Input<boolean>;
    readonly comment?: pulumi.Input<string>;
    readonly contract?: pulumi.Input<string>;
    readonly defaultErrorPenalty?: pulumi.Input<number>;
    readonly defaultHealthMax?: pulumi.Input<number>;
    readonly defaultHealthMultiplier?: pulumi.Input<number>;
    readonly defaultHealthThreshold?: pulumi.Input<number>;
    readonly defaultMaxUnreachablePenalty?: pulumi.Input<number>;
    readonly defaultSslClientCertificate?: pulumi.Input<string>;
    readonly defaultSslClientPrivateKey?: pulumi.Input<string>;
    readonly defaultTimeoutPenalty?: pulumi.Input<number>;
    readonly defaultUnreachableThreshold?: pulumi.Input<number>;
    readonly emailNotificationLists?: pulumi.Input<pulumi.Input<string>[]>;
    readonly endUserMappingEnabled?: pulumi.Input<boolean>;
    readonly group?: pulumi.Input<string>;
    readonly loadFeedback?: pulumi.Input<boolean>;
    readonly loadImbalancePercentage?: pulumi.Input<number>;
    readonly mapUpdateInterval?: pulumi.Input<number>;
    readonly maxProperties?: pulumi.Input<number>;
    readonly maxResources?: pulumi.Input<number>;
    readonly maxTestTimeout?: pulumi.Input<number>;
    readonly maxTtl?: pulumi.Input<number>;
    readonly minPingableRegionFraction?: pulumi.Input<number>;
    readonly minTestInterval?: pulumi.Input<number>;
    readonly minTtl?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly pingInterval?: pulumi.Input<number>;
    readonly pingPacketSize?: pulumi.Input<number>;
    readonly roundRobinPrefix?: pulumi.Input<string>;
    readonly servermonitorLivenessCount?: pulumi.Input<number>;
    readonly servermonitorLoadCount?: pulumi.Input<number>;
    readonly servermonitorPool?: pulumi.Input<string>;
    readonly type?: pulumi.Input<string>;
    readonly waitOnComplete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GtmDomain resource.
 */
export interface GtmDomainArgs {
    readonly cnameCoalescingEnabled?: pulumi.Input<boolean>;
    readonly comment?: pulumi.Input<string>;
    readonly contract?: pulumi.Input<string>;
    readonly defaultErrorPenalty?: pulumi.Input<number>;
    readonly defaultSslClientCertificate?: pulumi.Input<string>;
    readonly defaultSslClientPrivateKey?: pulumi.Input<string>;
    readonly defaultTimeoutPenalty?: pulumi.Input<number>;
    readonly emailNotificationLists?: pulumi.Input<pulumi.Input<string>[]>;
    readonly endUserMappingEnabled?: pulumi.Input<boolean>;
    readonly group?: pulumi.Input<string>;
    readonly loadFeedback?: pulumi.Input<boolean>;
    readonly loadImbalancePercentage?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly type: pulumi.Input<string>;
    readonly waitOnComplete?: pulumi.Input<boolean>;
}
