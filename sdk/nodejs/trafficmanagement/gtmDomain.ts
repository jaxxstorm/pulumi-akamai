// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use the `akamai.GtmDomain` resource to create, configure, and import a GTM Domain, which is a basic building block of a traffic management configuration.
 *
 * > **Note** Import requires an ID with this format: `existingDomainName`.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const demodomain = new akamai.GtmDomain("demodomain", {
 *     comment: "some comment",
 *     contract: "XXX",
 *     group: "100",
 *     type: "basic",
 * });
 * ```
 * ## Schema reference
 *
 * You can download the GTM Domain backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#domain) page.
 *
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

    /**
     * A boolean that if set to `true`, GTM collapses CNAME redirections in DNS answers when it knows the target of the CNAME.
     */
    public readonly cnameCoalescingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A descriptive note about changes to the domain. The maximum is 4000 characters.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * If creating a domain, the contract ID.
     */
    public readonly contract!: pulumi.Output<string | undefined>;
    /**
     * Specifies the download penalty score. The default is `75`. If the download encounters an error, the web agent computes a score that is either the download time in seconds or a penalty score.
     */
    public readonly defaultErrorPenalty!: pulumi.Output<number | undefined>;
    public /*out*/ readonly defaultHealthMax!: pulumi.Output<number>;
    public /*out*/ readonly defaultHealthMultiplier!: pulumi.Output<number>;
    public /*out*/ readonly defaultHealthThreshold!: pulumi.Output<number>;
    public /*out*/ readonly defaultMaxUnreachablePenalty!: pulumi.Output<number>;
    /**
     * Specifies an optional Base64-encoded certificate that corresponds with the private key for TLS-based liveness tests (HTTPS, SMTPS, POPS, and TCPS).
     */
    public readonly defaultSslClientCertificate!: pulumi.Output<string | undefined>;
    /**
     * Specifies a Base64-encoded private key that corresponds with the TLS certificate for HTTPS, SMTPS, POPS, and TCPS liveness tests.
     */
    public readonly defaultSslClientPrivateKey!: pulumi.Output<string | undefined>;
    /**
     * Specifies the timeout penalty score. Default is `25`.
     */
    public readonly defaultTimeoutPenalty!: pulumi.Output<number | undefined>;
    public /*out*/ readonly defaultUnreachableThreshold!: pulumi.Output<number>;
    /**
     * A list of email addresses to notify when a change is made to the domain.
     */
    public readonly emailNotificationLists!: pulumi.Output<string[] | undefined>;
    /**
     * A boolean indicating whether whether the GTM Domain is using end user client subnet mapping.
     */
    public readonly endUserMappingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * If creating a domain, the currently selected group ID.
     */
    public readonly group!: pulumi.Output<string | undefined>;
    /**
     * A boolean indicating whether one or more measurements of load (resources) are defined by you and supplied by each data center in real time to balance load.
     */
    public readonly loadFeedback!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates the percentage of load imbalance factor (LIF) for the domain.
     */
    public readonly loadImbalancePercentage!: pulumi.Output<number | undefined>;
    public /*out*/ readonly mapUpdateInterval!: pulumi.Output<number>;
    public /*out*/ readonly maxProperties!: pulumi.Output<number>;
    public /*out*/ readonly maxResources!: pulumi.Output<number>;
    public /*out*/ readonly maxTestTimeout!: pulumi.Output<number>;
    public /*out*/ readonly maxTtl!: pulumi.Output<number>;
    public /*out*/ readonly minPingableRegionFraction!: pulumi.Output<number>;
    public /*out*/ readonly minTestInterval!: pulumi.Output<number>;
    public /*out*/ readonly minTtl!: pulumi.Output<number>;
    /**
     * The DNS name for a collection of GTM Properties.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly pingInterval!: pulumi.Output<number>;
    public /*out*/ readonly pingPacketSize!: pulumi.Output<number>;
    public /*out*/ readonly roundRobinPrefix!: pulumi.Output<string>;
    public /*out*/ readonly servermonitorLivenessCount!: pulumi.Output<number>;
    public /*out*/ readonly servermonitorLoadCount!: pulumi.Output<number>;
    public /*out*/ readonly servermonitorPool!: pulumi.Output<string>;
    /**
     * Th type of GTM domain. Options include `failover-only`, `static`, `weighted`, `basic`, or `full`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * A boolean that, if set to `true`, waits for transaction to complete.
     */
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmDomainState | undefined;
            resourceInputs["cnameCoalescingEnabled"] = state ? state.cnameCoalescingEnabled : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["contract"] = state ? state.contract : undefined;
            resourceInputs["defaultErrorPenalty"] = state ? state.defaultErrorPenalty : undefined;
            resourceInputs["defaultHealthMax"] = state ? state.defaultHealthMax : undefined;
            resourceInputs["defaultHealthMultiplier"] = state ? state.defaultHealthMultiplier : undefined;
            resourceInputs["defaultHealthThreshold"] = state ? state.defaultHealthThreshold : undefined;
            resourceInputs["defaultMaxUnreachablePenalty"] = state ? state.defaultMaxUnreachablePenalty : undefined;
            resourceInputs["defaultSslClientCertificate"] = state ? state.defaultSslClientCertificate : undefined;
            resourceInputs["defaultSslClientPrivateKey"] = state ? state.defaultSslClientPrivateKey : undefined;
            resourceInputs["defaultTimeoutPenalty"] = state ? state.defaultTimeoutPenalty : undefined;
            resourceInputs["defaultUnreachableThreshold"] = state ? state.defaultUnreachableThreshold : undefined;
            resourceInputs["emailNotificationLists"] = state ? state.emailNotificationLists : undefined;
            resourceInputs["endUserMappingEnabled"] = state ? state.endUserMappingEnabled : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["loadFeedback"] = state ? state.loadFeedback : undefined;
            resourceInputs["loadImbalancePercentage"] = state ? state.loadImbalancePercentage : undefined;
            resourceInputs["mapUpdateInterval"] = state ? state.mapUpdateInterval : undefined;
            resourceInputs["maxProperties"] = state ? state.maxProperties : undefined;
            resourceInputs["maxResources"] = state ? state.maxResources : undefined;
            resourceInputs["maxTestTimeout"] = state ? state.maxTestTimeout : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["minPingableRegionFraction"] = state ? state.minPingableRegionFraction : undefined;
            resourceInputs["minTestInterval"] = state ? state.minTestInterval : undefined;
            resourceInputs["minTtl"] = state ? state.minTtl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pingInterval"] = state ? state.pingInterval : undefined;
            resourceInputs["pingPacketSize"] = state ? state.pingPacketSize : undefined;
            resourceInputs["roundRobinPrefix"] = state ? state.roundRobinPrefix : undefined;
            resourceInputs["servermonitorLivenessCount"] = state ? state.servermonitorLivenessCount : undefined;
            resourceInputs["servermonitorLoadCount"] = state ? state.servermonitorLoadCount : undefined;
            resourceInputs["servermonitorPool"] = state ? state.servermonitorPool : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
        } else {
            const args = argsOrState as GtmDomainArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["cnameCoalescingEnabled"] = args ? args.cnameCoalescingEnabled : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["contract"] = args ? args.contract : undefined;
            resourceInputs["defaultErrorPenalty"] = args ? args.defaultErrorPenalty : undefined;
            resourceInputs["defaultSslClientCertificate"] = args ? args.defaultSslClientCertificate : undefined;
            resourceInputs["defaultSslClientPrivateKey"] = args ? args.defaultSslClientPrivateKey : undefined;
            resourceInputs["defaultTimeoutPenalty"] = args ? args.defaultTimeoutPenalty : undefined;
            resourceInputs["emailNotificationLists"] = args ? args.emailNotificationLists : undefined;
            resourceInputs["endUserMappingEnabled"] = args ? args.endUserMappingEnabled : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["loadFeedback"] = args ? args.loadFeedback : undefined;
            resourceInputs["loadImbalancePercentage"] = args ? args.loadImbalancePercentage : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
            resourceInputs["defaultHealthMax"] = undefined /*out*/;
            resourceInputs["defaultHealthMultiplier"] = undefined /*out*/;
            resourceInputs["defaultHealthThreshold"] = undefined /*out*/;
            resourceInputs["defaultMaxUnreachablePenalty"] = undefined /*out*/;
            resourceInputs["defaultUnreachableThreshold"] = undefined /*out*/;
            resourceInputs["mapUpdateInterval"] = undefined /*out*/;
            resourceInputs["maxProperties"] = undefined /*out*/;
            resourceInputs["maxResources"] = undefined /*out*/;
            resourceInputs["maxTestTimeout"] = undefined /*out*/;
            resourceInputs["maxTtl"] = undefined /*out*/;
            resourceInputs["minPingableRegionFraction"] = undefined /*out*/;
            resourceInputs["minTestInterval"] = undefined /*out*/;
            resourceInputs["minTtl"] = undefined /*out*/;
            resourceInputs["pingInterval"] = undefined /*out*/;
            resourceInputs["pingPacketSize"] = undefined /*out*/;
            resourceInputs["roundRobinPrefix"] = undefined /*out*/;
            resourceInputs["servermonitorLivenessCount"] = undefined /*out*/;
            resourceInputs["servermonitorLoadCount"] = undefined /*out*/;
            resourceInputs["servermonitorPool"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GtmDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmDomain resources.
 */
export interface GtmDomainState {
    /**
     * A boolean that if set to `true`, GTM collapses CNAME redirections in DNS answers when it knows the target of the CNAME.
     */
    cnameCoalescingEnabled?: pulumi.Input<boolean>;
    /**
     * A descriptive note about changes to the domain. The maximum is 4000 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * If creating a domain, the contract ID.
     */
    contract?: pulumi.Input<string>;
    /**
     * Specifies the download penalty score. The default is `75`. If the download encounters an error, the web agent computes a score that is either the download time in seconds or a penalty score.
     */
    defaultErrorPenalty?: pulumi.Input<number>;
    defaultHealthMax?: pulumi.Input<number>;
    defaultHealthMultiplier?: pulumi.Input<number>;
    defaultHealthThreshold?: pulumi.Input<number>;
    defaultMaxUnreachablePenalty?: pulumi.Input<number>;
    /**
     * Specifies an optional Base64-encoded certificate that corresponds with the private key for TLS-based liveness tests (HTTPS, SMTPS, POPS, and TCPS).
     */
    defaultSslClientCertificate?: pulumi.Input<string>;
    /**
     * Specifies a Base64-encoded private key that corresponds with the TLS certificate for HTTPS, SMTPS, POPS, and TCPS liveness tests.
     */
    defaultSslClientPrivateKey?: pulumi.Input<string>;
    /**
     * Specifies the timeout penalty score. Default is `25`.
     */
    defaultTimeoutPenalty?: pulumi.Input<number>;
    defaultUnreachableThreshold?: pulumi.Input<number>;
    /**
     * A list of email addresses to notify when a change is made to the domain.
     */
    emailNotificationLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A boolean indicating whether whether the GTM Domain is using end user client subnet mapping.
     */
    endUserMappingEnabled?: pulumi.Input<boolean>;
    /**
     * If creating a domain, the currently selected group ID.
     */
    group?: pulumi.Input<string>;
    /**
     * A boolean indicating whether one or more measurements of load (resources) are defined by you and supplied by each data center in real time to balance load.
     */
    loadFeedback?: pulumi.Input<boolean>;
    /**
     * Indicates the percentage of load imbalance factor (LIF) for the domain.
     */
    loadImbalancePercentage?: pulumi.Input<number>;
    mapUpdateInterval?: pulumi.Input<number>;
    maxProperties?: pulumi.Input<number>;
    maxResources?: pulumi.Input<number>;
    maxTestTimeout?: pulumi.Input<number>;
    maxTtl?: pulumi.Input<number>;
    minPingableRegionFraction?: pulumi.Input<number>;
    minTestInterval?: pulumi.Input<number>;
    minTtl?: pulumi.Input<number>;
    /**
     * The DNS name for a collection of GTM Properties.
     */
    name?: pulumi.Input<string>;
    pingInterval?: pulumi.Input<number>;
    pingPacketSize?: pulumi.Input<number>;
    roundRobinPrefix?: pulumi.Input<string>;
    servermonitorLivenessCount?: pulumi.Input<number>;
    servermonitorLoadCount?: pulumi.Input<number>;
    servermonitorPool?: pulumi.Input<string>;
    /**
     * Th type of GTM domain. Options include `failover-only`, `static`, `weighted`, `basic`, or `full`.
     */
    type?: pulumi.Input<string>;
    /**
     * A boolean that, if set to `true`, waits for transaction to complete.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GtmDomain resource.
 */
export interface GtmDomainArgs {
    /**
     * A boolean that if set to `true`, GTM collapses CNAME redirections in DNS answers when it knows the target of the CNAME.
     */
    cnameCoalescingEnabled?: pulumi.Input<boolean>;
    /**
     * A descriptive note about changes to the domain. The maximum is 4000 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * If creating a domain, the contract ID.
     */
    contract?: pulumi.Input<string>;
    /**
     * Specifies the download penalty score. The default is `75`. If the download encounters an error, the web agent computes a score that is either the download time in seconds or a penalty score.
     */
    defaultErrorPenalty?: pulumi.Input<number>;
    /**
     * Specifies an optional Base64-encoded certificate that corresponds with the private key for TLS-based liveness tests (HTTPS, SMTPS, POPS, and TCPS).
     */
    defaultSslClientCertificate?: pulumi.Input<string>;
    /**
     * Specifies a Base64-encoded private key that corresponds with the TLS certificate for HTTPS, SMTPS, POPS, and TCPS liveness tests.
     */
    defaultSslClientPrivateKey?: pulumi.Input<string>;
    /**
     * Specifies the timeout penalty score. Default is `25`.
     */
    defaultTimeoutPenalty?: pulumi.Input<number>;
    /**
     * A list of email addresses to notify when a change is made to the domain.
     */
    emailNotificationLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A boolean indicating whether whether the GTM Domain is using end user client subnet mapping.
     */
    endUserMappingEnabled?: pulumi.Input<boolean>;
    /**
     * If creating a domain, the currently selected group ID.
     */
    group?: pulumi.Input<string>;
    /**
     * A boolean indicating whether one or more measurements of load (resources) are defined by you and supplied by each data center in real time to balance load.
     */
    loadFeedback?: pulumi.Input<boolean>;
    /**
     * Indicates the percentage of load imbalance factor (LIF) for the domain.
     */
    loadImbalancePercentage?: pulumi.Input<number>;
    /**
     * The DNS name for a collection of GTM Properties.
     */
    name?: pulumi.Input<string>;
    /**
     * Th type of GTM domain. Options include `failover-only`, `static`, `weighted`, `basic`, or `full`.
     */
    type: pulumi.Input<string>;
    /**
     * A boolean that, if set to `true`, waits for transaction to complete.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}
