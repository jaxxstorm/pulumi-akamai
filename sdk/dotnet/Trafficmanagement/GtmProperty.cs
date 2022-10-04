// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Trafficmanagement
{
    /// <summary>
    /// Use the `akamai.GtmProperty` resource to create, configure and import a GTM property, a set of IP addresses or CNAMEs that GTM provides in response to DNS queries based on a set of rules.
    /// 
    /// &gt; **Note** Import requires an ID with this format: `existing_domain_name`:`existing_property_name`.
    /// 
    /// ## Example Usage
    /// 
    /// Basic usage:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Akamai = Pulumi.Akamai;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var demoProperty = new Akamai.GtmProperty("demoProperty", new Akamai.GtmPropertyArgs
    ///         {
    ///             Domain = "demo_domain.akadns.net",
    ///             HandoutLimit = 5,
    ///             HandoutMode = "normal",
    ///             ScoreAggregationType = "median",
    ///             TrafficTargets = 
    ///             {
    ///                 new Akamai.Inputs.GtmPropertyTrafficTargetArgs
    ///                 {
    ///                     DatacenterId = 3131,
    ///                 },
    ///             },
    ///             Type = "weighted-round-robin",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [Obsolete(@"akamai.trafficmanagement.GtmProperty has been deprecated in favor of akamai.GtmProperty")]
    [AkamaiResourceType("akamai:trafficmanagement/gtmProperty:GtmProperty")]
    public partial class GtmProperty : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies a backup CNAME. If GTM declares that all of the servers configured for your property are down, the backup CNAME is handed out. If a backup CNAME is set, do not set a backup IP.
        /// </summary>
        [Output("backupCname")]
        public Output<string?> BackupCname { get; private set; } = null!;

        /// <summary>
        /// Specifies a backup IP. When GTM declares that all of the targets are down, the backup IP is handed out. If a backup IP is set, do not set a backup CNAME.
        /// </summary>
        [Output("backupIp")]
        public Output<string?> BackupIp { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates whether download score based load balancing is enabled.
        /// </summary>
        [Output("balanceByDownloadScore")]
        public Output<bool?> BalanceByDownloadScore { get; private set; } = null!;

        /// <summary>
        /// Indicates the fully qualified name aliased to a particular property.
        /// </summary>
        [Output("cname")]
        public Output<string?> Cname { get; private set; } = null!;

        /// <summary>
        /// A descriptive note about changes to the domain. The maximum is 4000 characters.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// DNS name for the GTM Domain set that includes this Property.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Indicates the TTL in seconds for records that might change dynamically based on liveness and load balancing such as A and AAAA records, and CNAMEs.
        /// </summary>
        [Output("dynamicTtl")]
        public Output<int?> DynamicTtl { get; private set; } = null!;

        /// <summary>
        /// Specifies the failback delay in seconds.
        /// </summary>
        [Output("failbackDelay")]
        public Output<int?> FailbackDelay { get; private set; } = null!;

        /// <summary>
        /// Specifies the failover delay in seconds.
        /// </summary>
        [Output("failoverDelay")]
        public Output<int?> FailoverDelay { get; private set; } = null!;

        /// <summary>
        /// Use load estimates from Akamai Ghost utilization messages.
        /// </summary>
        [Output("ghostDemandReporting")]
        public Output<bool?> GhostDemandReporting { get; private set; } = null!;

        /// <summary>
        /// Indicates the limit for the number of live IPs handed out to a DNS request.
        /// </summary>
        [Output("handoutLimit")]
        public Output<int> HandoutLimit { get; private set; } = null!;

        /// <summary>
        /// Specifies how IPs are returned when more than one IP is alive and available.
        /// </summary>
        [Output("handoutMode")]
        public Output<string> HandoutMode { get; private set; } = null!;

        /// <summary>
        /// Defines the absolute limit beyond which IPs are declared unhealthy.
        /// </summary>
        [Output("healthMax")]
        public Output<double?> HealthMax { get; private set; } = null!;

        /// <summary>
        /// Configures a cutoff value that is computed from the median scores.
        /// </summary>
        [Output("healthMultiplier")]
        public Output<double?> HealthMultiplier { get; private set; } = null!;

        /// <summary>
        /// Configures a cutoff value that is computed from the median scores.
        /// </summary>
        [Output("healthThreshold")]
        public Output<double?> HealthThreshold { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates the type of IP address handed out by a GTM property.
        /// </summary>
        [Output("ipv6")]
        public Output<bool?> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// Contains information about the liveness tests, which are run periodically to determine whether your servers respond to requests. You can have multiple `liveness_test` arguments. If used, requires these arguments:
        /// </summary>
        [Output("livenessTests")]
        public Output<ImmutableArray<Outputs.GtmPropertyLivenessTest>> LivenessTests { get; private set; } = null!;

        /// <summary>
        /// Indicates the percent of load imbalance factor (LIF) for the property.
        /// </summary>
        [Output("loadImbalancePercentage")]
        public Output<double?> LoadImbalancePercentage { get; private set; } = null!;

        /// <summary>
        /// A descriptive label for a GeographicMap or a CidrMap that's required if the property is either geographic or cidrmapping, in which case mapName needs to reference either an existing GeographicMap or CidrMap in the same domain.
        /// </summary>
        [Output("mapName")]
        public Output<string?> MapName { get; private set; } = null!;

        /// <summary>
        /// For performance domains, this specifies a penalty value that's added to liveness test scores when data centers show an aggregated loss fraction higher than the penalty value.
        /// </summary>
        [Output("maxUnreachablePenalty")]
        public Output<int?> MaxUnreachablePenalty { get; private set; } = null!;

        /// <summary>
        /// Specifies what fraction of the servers need to respond to requests so GTM considers the data center up and able to receive traffic.
        /// </summary>
        [Output("minLiveFraction")]
        public Output<double?> MinLiveFraction { get; private set; } = null!;

        /// <summary>
        /// Name of HTTP header.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies how GTM aggregates liveness test scores across different tests, when multiple tests are configured.
        /// </summary>
        [Output("scoreAggregationType")]
        public Output<string> ScoreAggregationType { get; private set; } = null!;

        /// <summary>
        /// Contains static record sets. You can have multiple `static_rr_set` entries. Requires these arguments:
        /// </summary>
        [Output("staticRrSets")]
        public Output<ImmutableArray<Outputs.GtmPropertyStaticRrSet>> StaticRrSets { get; private set; } = null!;

        [Output("staticTtl")]
        public Output<int?> StaticTtl { get; private set; } = null!;

        /// <summary>
        /// Specifies a constant used to configure data center affinity.
        /// </summary>
        [Output("stickinessBonusConstant")]
        public Output<int?> StickinessBonusConstant { get; private set; } = null!;

        /// <summary>
        /// Specifies a percentage used to configure data center affinity.
        /// </summary>
        [Output("stickinessBonusPercentage")]
        public Output<int?> StickinessBonusPercentage { get; private set; } = null!;

        /// <summary>
        /// Contains information about where to direct data center traffic. You can have multiple `traffic_target` arguments. If used, includes these arguments:
        /// </summary>
        [Output("trafficTargets")]
        public Output<ImmutableArray<Outputs.GtmPropertyTrafficTarget>> TrafficTargets { get; private set; } = null!;

        /// <summary>
        /// The record type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// For performance domains, this specifies a penalty value that's added to liveness test scores when data centers have an aggregated loss fraction higher than this value.
        /// </summary>
        [Output("unreachableThreshold")]
        public Output<double?> UnreachableThreshold { get; private set; } = null!;

        /// <summary>
        /// For load-feedback domains only, a boolean that indicates whether you want GTM to automatically compute target load.
        /// </summary>
        [Output("useComputedTargets")]
        public Output<bool?> UseComputedTargets { get; private set; } = null!;

        /// <summary>
        /// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
        /// </summary>
        [Output("waitOnComplete")]
        public Output<bool?> WaitOnComplete { get; private set; } = null!;

        [Output("weightedHashBitsForIpv4")]
        public Output<int> WeightedHashBitsForIpv4 { get; private set; } = null!;

        [Output("weightedHashBitsForIpv6")]
        public Output<int> WeightedHashBitsForIpv6 { get; private set; } = null!;


        /// <summary>
        /// Create a GtmProperty resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GtmProperty(string name, GtmPropertyArgs args, CustomResourceOptions? options = null)
            : base("akamai:trafficmanagement/gtmProperty:GtmProperty", name, args ?? new GtmPropertyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GtmProperty(string name, Input<string> id, GtmPropertyState? state = null, CustomResourceOptions? options = null)
            : base("akamai:trafficmanagement/gtmProperty:GtmProperty", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GtmProperty resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GtmProperty Get(string name, Input<string> id, GtmPropertyState? state = null, CustomResourceOptions? options = null)
        {
            return new GtmProperty(name, id, state, options);
        }
    }

    public sealed class GtmPropertyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a backup CNAME. If GTM declares that all of the servers configured for your property are down, the backup CNAME is handed out. If a backup CNAME is set, do not set a backup IP.
        /// </summary>
        [Input("backupCname")]
        public Input<string>? BackupCname { get; set; }

        /// <summary>
        /// Specifies a backup IP. When GTM declares that all of the targets are down, the backup IP is handed out. If a backup IP is set, do not set a backup CNAME.
        /// </summary>
        [Input("backupIp")]
        public Input<string>? BackupIp { get; set; }

        /// <summary>
        /// A boolean that indicates whether download score based load balancing is enabled.
        /// </summary>
        [Input("balanceByDownloadScore")]
        public Input<bool>? BalanceByDownloadScore { get; set; }

        /// <summary>
        /// Indicates the fully qualified name aliased to a particular property.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// A descriptive note about changes to the domain. The maximum is 4000 characters.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// DNS name for the GTM Domain set that includes this Property.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Indicates the TTL in seconds for records that might change dynamically based on liveness and load balancing such as A and AAAA records, and CNAMEs.
        /// </summary>
        [Input("dynamicTtl")]
        public Input<int>? DynamicTtl { get; set; }

        /// <summary>
        /// Specifies the failback delay in seconds.
        /// </summary>
        [Input("failbackDelay")]
        public Input<int>? FailbackDelay { get; set; }

        /// <summary>
        /// Specifies the failover delay in seconds.
        /// </summary>
        [Input("failoverDelay")]
        public Input<int>? FailoverDelay { get; set; }

        /// <summary>
        /// Use load estimates from Akamai Ghost utilization messages.
        /// </summary>
        [Input("ghostDemandReporting")]
        public Input<bool>? GhostDemandReporting { get; set; }

        /// <summary>
        /// Indicates the limit for the number of live IPs handed out to a DNS request.
        /// </summary>
        [Input("handoutLimit", required: true)]
        public Input<int> HandoutLimit { get; set; } = null!;

        /// <summary>
        /// Specifies how IPs are returned when more than one IP is alive and available.
        /// </summary>
        [Input("handoutMode", required: true)]
        public Input<string> HandoutMode { get; set; } = null!;

        /// <summary>
        /// Defines the absolute limit beyond which IPs are declared unhealthy.
        /// </summary>
        [Input("healthMax")]
        public Input<double>? HealthMax { get; set; }

        /// <summary>
        /// Configures a cutoff value that is computed from the median scores.
        /// </summary>
        [Input("healthMultiplier")]
        public Input<double>? HealthMultiplier { get; set; }

        /// <summary>
        /// Configures a cutoff value that is computed from the median scores.
        /// </summary>
        [Input("healthThreshold")]
        public Input<double>? HealthThreshold { get; set; }

        /// <summary>
        /// A boolean that indicates the type of IP address handed out by a GTM property.
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        [Input("livenessTests")]
        private InputList<Inputs.GtmPropertyLivenessTestArgs>? _livenessTests;

        /// <summary>
        /// Contains information about the liveness tests, which are run periodically to determine whether your servers respond to requests. You can have multiple `liveness_test` arguments. If used, requires these arguments:
        /// </summary>
        public InputList<Inputs.GtmPropertyLivenessTestArgs> LivenessTests
        {
            get => _livenessTests ?? (_livenessTests = new InputList<Inputs.GtmPropertyLivenessTestArgs>());
            set => _livenessTests = value;
        }

        /// <summary>
        /// Indicates the percent of load imbalance factor (LIF) for the property.
        /// </summary>
        [Input("loadImbalancePercentage")]
        public Input<double>? LoadImbalancePercentage { get; set; }

        /// <summary>
        /// A descriptive label for a GeographicMap or a CidrMap that's required if the property is either geographic or cidrmapping, in which case mapName needs to reference either an existing GeographicMap or CidrMap in the same domain.
        /// </summary>
        [Input("mapName")]
        public Input<string>? MapName { get; set; }

        /// <summary>
        /// For performance domains, this specifies a penalty value that's added to liveness test scores when data centers show an aggregated loss fraction higher than the penalty value.
        /// </summary>
        [Input("maxUnreachablePenalty")]
        public Input<int>? MaxUnreachablePenalty { get; set; }

        /// <summary>
        /// Specifies what fraction of the servers need to respond to requests so GTM considers the data center up and able to receive traffic.
        /// </summary>
        [Input("minLiveFraction")]
        public Input<double>? MinLiveFraction { get; set; }

        /// <summary>
        /// Name of HTTP header.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies how GTM aggregates liveness test scores across different tests, when multiple tests are configured.
        /// </summary>
        [Input("scoreAggregationType", required: true)]
        public Input<string> ScoreAggregationType { get; set; } = null!;

        [Input("staticRrSets")]
        private InputList<Inputs.GtmPropertyStaticRrSetArgs>? _staticRrSets;

        /// <summary>
        /// Contains static record sets. You can have multiple `static_rr_set` entries. Requires these arguments:
        /// </summary>
        public InputList<Inputs.GtmPropertyStaticRrSetArgs> StaticRrSets
        {
            get => _staticRrSets ?? (_staticRrSets = new InputList<Inputs.GtmPropertyStaticRrSetArgs>());
            set => _staticRrSets = value;
        }

        [Input("staticTtl")]
        public Input<int>? StaticTtl { get; set; }

        /// <summary>
        /// Specifies a constant used to configure data center affinity.
        /// </summary>
        [Input("stickinessBonusConstant")]
        public Input<int>? StickinessBonusConstant { get; set; }

        /// <summary>
        /// Specifies a percentage used to configure data center affinity.
        /// </summary>
        [Input("stickinessBonusPercentage")]
        public Input<int>? StickinessBonusPercentage { get; set; }

        [Input("trafficTargets")]
        private InputList<Inputs.GtmPropertyTrafficTargetArgs>? _trafficTargets;

        /// <summary>
        /// Contains information about where to direct data center traffic. You can have multiple `traffic_target` arguments. If used, includes these arguments:
        /// </summary>
        public InputList<Inputs.GtmPropertyTrafficTargetArgs> TrafficTargets
        {
            get => _trafficTargets ?? (_trafficTargets = new InputList<Inputs.GtmPropertyTrafficTargetArgs>());
            set => _trafficTargets = value;
        }

        /// <summary>
        /// The record type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// For performance domains, this specifies a penalty value that's added to liveness test scores when data centers have an aggregated loss fraction higher than this value.
        /// </summary>
        [Input("unreachableThreshold")]
        public Input<double>? UnreachableThreshold { get; set; }

        /// <summary>
        /// For load-feedback domains only, a boolean that indicates whether you want GTM to automatically compute target load.
        /// </summary>
        [Input("useComputedTargets")]
        public Input<bool>? UseComputedTargets { get; set; }

        /// <summary>
        /// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
        /// </summary>
        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        public GtmPropertyArgs()
        {
        }
    }

    public sealed class GtmPropertyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a backup CNAME. If GTM declares that all of the servers configured for your property are down, the backup CNAME is handed out. If a backup CNAME is set, do not set a backup IP.
        /// </summary>
        [Input("backupCname")]
        public Input<string>? BackupCname { get; set; }

        /// <summary>
        /// Specifies a backup IP. When GTM declares that all of the targets are down, the backup IP is handed out. If a backup IP is set, do not set a backup CNAME.
        /// </summary>
        [Input("backupIp")]
        public Input<string>? BackupIp { get; set; }

        /// <summary>
        /// A boolean that indicates whether download score based load balancing is enabled.
        /// </summary>
        [Input("balanceByDownloadScore")]
        public Input<bool>? BalanceByDownloadScore { get; set; }

        /// <summary>
        /// Indicates the fully qualified name aliased to a particular property.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// A descriptive note about changes to the domain. The maximum is 4000 characters.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// DNS name for the GTM Domain set that includes this Property.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Indicates the TTL in seconds for records that might change dynamically based on liveness and load balancing such as A and AAAA records, and CNAMEs.
        /// </summary>
        [Input("dynamicTtl")]
        public Input<int>? DynamicTtl { get; set; }

        /// <summary>
        /// Specifies the failback delay in seconds.
        /// </summary>
        [Input("failbackDelay")]
        public Input<int>? FailbackDelay { get; set; }

        /// <summary>
        /// Specifies the failover delay in seconds.
        /// </summary>
        [Input("failoverDelay")]
        public Input<int>? FailoverDelay { get; set; }

        /// <summary>
        /// Use load estimates from Akamai Ghost utilization messages.
        /// </summary>
        [Input("ghostDemandReporting")]
        public Input<bool>? GhostDemandReporting { get; set; }

        /// <summary>
        /// Indicates the limit for the number of live IPs handed out to a DNS request.
        /// </summary>
        [Input("handoutLimit")]
        public Input<int>? HandoutLimit { get; set; }

        /// <summary>
        /// Specifies how IPs are returned when more than one IP is alive and available.
        /// </summary>
        [Input("handoutMode")]
        public Input<string>? HandoutMode { get; set; }

        /// <summary>
        /// Defines the absolute limit beyond which IPs are declared unhealthy.
        /// </summary>
        [Input("healthMax")]
        public Input<double>? HealthMax { get; set; }

        /// <summary>
        /// Configures a cutoff value that is computed from the median scores.
        /// </summary>
        [Input("healthMultiplier")]
        public Input<double>? HealthMultiplier { get; set; }

        /// <summary>
        /// Configures a cutoff value that is computed from the median scores.
        /// </summary>
        [Input("healthThreshold")]
        public Input<double>? HealthThreshold { get; set; }

        /// <summary>
        /// A boolean that indicates the type of IP address handed out by a GTM property.
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        [Input("livenessTests")]
        private InputList<Inputs.GtmPropertyLivenessTestGetArgs>? _livenessTests;

        /// <summary>
        /// Contains information about the liveness tests, which are run periodically to determine whether your servers respond to requests. You can have multiple `liveness_test` arguments. If used, requires these arguments:
        /// </summary>
        public InputList<Inputs.GtmPropertyLivenessTestGetArgs> LivenessTests
        {
            get => _livenessTests ?? (_livenessTests = new InputList<Inputs.GtmPropertyLivenessTestGetArgs>());
            set => _livenessTests = value;
        }

        /// <summary>
        /// Indicates the percent of load imbalance factor (LIF) for the property.
        /// </summary>
        [Input("loadImbalancePercentage")]
        public Input<double>? LoadImbalancePercentage { get; set; }

        /// <summary>
        /// A descriptive label for a GeographicMap or a CidrMap that's required if the property is either geographic or cidrmapping, in which case mapName needs to reference either an existing GeographicMap or CidrMap in the same domain.
        /// </summary>
        [Input("mapName")]
        public Input<string>? MapName { get; set; }

        /// <summary>
        /// For performance domains, this specifies a penalty value that's added to liveness test scores when data centers show an aggregated loss fraction higher than the penalty value.
        /// </summary>
        [Input("maxUnreachablePenalty")]
        public Input<int>? MaxUnreachablePenalty { get; set; }

        /// <summary>
        /// Specifies what fraction of the servers need to respond to requests so GTM considers the data center up and able to receive traffic.
        /// </summary>
        [Input("minLiveFraction")]
        public Input<double>? MinLiveFraction { get; set; }

        /// <summary>
        /// Name of HTTP header.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies how GTM aggregates liveness test scores across different tests, when multiple tests are configured.
        /// </summary>
        [Input("scoreAggregationType")]
        public Input<string>? ScoreAggregationType { get; set; }

        [Input("staticRrSets")]
        private InputList<Inputs.GtmPropertyStaticRrSetGetArgs>? _staticRrSets;

        /// <summary>
        /// Contains static record sets. You can have multiple `static_rr_set` entries. Requires these arguments:
        /// </summary>
        public InputList<Inputs.GtmPropertyStaticRrSetGetArgs> StaticRrSets
        {
            get => _staticRrSets ?? (_staticRrSets = new InputList<Inputs.GtmPropertyStaticRrSetGetArgs>());
            set => _staticRrSets = value;
        }

        [Input("staticTtl")]
        public Input<int>? StaticTtl { get; set; }

        /// <summary>
        /// Specifies a constant used to configure data center affinity.
        /// </summary>
        [Input("stickinessBonusConstant")]
        public Input<int>? StickinessBonusConstant { get; set; }

        /// <summary>
        /// Specifies a percentage used to configure data center affinity.
        /// </summary>
        [Input("stickinessBonusPercentage")]
        public Input<int>? StickinessBonusPercentage { get; set; }

        [Input("trafficTargets")]
        private InputList<Inputs.GtmPropertyTrafficTargetGetArgs>? _trafficTargets;

        /// <summary>
        /// Contains information about where to direct data center traffic. You can have multiple `traffic_target` arguments. If used, includes these arguments:
        /// </summary>
        public InputList<Inputs.GtmPropertyTrafficTargetGetArgs> TrafficTargets
        {
            get => _trafficTargets ?? (_trafficTargets = new InputList<Inputs.GtmPropertyTrafficTargetGetArgs>());
            set => _trafficTargets = value;
        }

        /// <summary>
        /// The record type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// For performance domains, this specifies a penalty value that's added to liveness test scores when data centers have an aggregated loss fraction higher than this value.
        /// </summary>
        [Input("unreachableThreshold")]
        public Input<double>? UnreachableThreshold { get; set; }

        /// <summary>
        /// For load-feedback domains only, a boolean that indicates whether you want GTM to automatically compute target load.
        /// </summary>
        [Input("useComputedTargets")]
        public Input<bool>? UseComputedTargets { get; set; }

        /// <summary>
        /// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
        /// </summary>
        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        [Input("weightedHashBitsForIpv4")]
        public Input<int>? WeightedHashBitsForIpv4 { get; set; }

        [Input("weightedHashBitsForIpv6")]
        public Input<int>? WeightedHashBitsForIpv6 { get; set; }

        public GtmPropertyState()
        {
        }
    }
}
