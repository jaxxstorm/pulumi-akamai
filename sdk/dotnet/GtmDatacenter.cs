// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    /// <summary>
    /// Use the `akamai.GtmDatacenter` resource to create, configure, and import a GTM data center. A GTM data center represents a customer data center and is also known as a traffic target, a location containing many servers GTM can direct traffic to.
    /// 
    /// GTM uses data centers to scale load balancing. For example, you might have data centers in both New York and Amsterdam and want to balance load between them. You can configure GTM to send US users to the New York data center and European users to the data center in Amsterdam.
    /// 
    /// &gt; **Note** Import requires an ID with this format: `existing_domain_name`:`existing_datacenter_id`.
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
    ///         var demoDatacenter = new Akamai.GtmDatacenter("demoDatacenter", new Akamai.GtmDatacenterArgs
    ///         {
    ///             Domain = "demo_domain.akadns.net",
    ///             Nickname = "demo_datacenter",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Argument reference
    /// 
    /// This resource supports these arguments:
    /// 
    /// * `domain` - (Required) The GTM domain name for the data center.
    /// * `wait_on_complete` - (Optional) A boolean, that if set to `true`, waits for transaction to complete.
    /// * `nickname` - (Optional) A descriptive label for the data center.
    /// * `default_load_object` - (Optional) Specifies the load reporting interface between you and the GTM system. If used, requires these additional arguments:
    ///   * `load_object` - A load object is a file that provides real-time information about the current load, maximum allowable load, and target load on each resource.
    ///   * `load_object_port` - Specifies the TCP port to connect to when requesting the load object.
    ///   * `load_servers` - Specifies a list of servers to request the load object from.
    /// * `city` - (Optional) The name of the city where the data center is located.
    /// * `clone_of` - (Optional) Identifies the data center’s `datacenter_id` of which this data center is a clone.
    /// * `cloud_server_targeting` - (Optional) A boolean indicating whether to balance load between two or more servers in a cloud environment.
    /// * `cloud_server_host_header_override` - (Optional) A boolean that, if set to `true`, Akamai's liveness test agents use the Host header configured in the liveness test.
    /// * `continent` - (Optional) A two-letter code that specifies the continent where the data center maps to.
    /// * `country` - (Optional) A two-letter ISO 3166 country code that specifies the country where the data center maps to.
    /// * `latitude` - (Optional) Specifies the geographical latitude of the data center’s position. See also longitude within this object.
    /// * `longitude` - (Optional) Specifies the geographic longitude of the data center’s position. See also latitude within this object.
    /// * `state_or_province` - (Optional) Specifies a two-letter ISO 3166 country code for the state or province where the data center is located.
    /// 
    /// ## Attribute reference
    /// 
    /// This resource returns these computed attributes in the state file:
    /// 
    /// * `datacenter_id` - A unique identifier for an existing data center in the domain.
    /// * `ping_interval`
    /// * `ping_packet_size`
    /// * `score_penalty`
    /// * `servermonitor_liveness_count`
    /// * `servermonitor_load_count`
    /// * `servermonitor_pool`
    /// * `virtual` - A boolean indicating whether the data center is virtual or physical, the latter meaning the data center has an Akamai Network Agent installed, and its physical location (`latitude`, `longitude`) is fixed. Either `true` if virtual or `false` if physical.
    /// 
    /// ## Schema reference
    /// 
    /// You can download the GTM Data Center backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#datacenter) page.
    /// </summary>
    [AkamaiResourceType("akamai:index/gtmDatacenter:GtmDatacenter")]
    public partial class GtmDatacenter : Pulumi.CustomResource
    {
        [Output("city")]
        public Output<string?> City { get; private set; } = null!;

        [Output("cloneOf")]
        public Output<int?> CloneOf { get; private set; } = null!;

        [Output("cloudServerHostHeaderOverride")]
        public Output<bool?> CloudServerHostHeaderOverride { get; private set; } = null!;

        [Output("cloudServerTargeting")]
        public Output<bool?> CloudServerTargeting { get; private set; } = null!;

        [Output("continent")]
        public Output<string?> Continent { get; private set; } = null!;

        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        [Output("datacenterId")]
        public Output<int> DatacenterId { get; private set; } = null!;

        [Output("defaultLoadObject")]
        public Output<Outputs.GtmDatacenterDefaultLoadObject?> DefaultLoadObject { get; private set; } = null!;

        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        [Output("latitude")]
        public Output<double?> Latitude { get; private set; } = null!;

        [Output("longitude")]
        public Output<double?> Longitude { get; private set; } = null!;

        [Output("nickname")]
        public Output<string?> Nickname { get; private set; } = null!;

        [Output("pingInterval")]
        public Output<int> PingInterval { get; private set; } = null!;

        [Output("pingPacketSize")]
        public Output<int> PingPacketSize { get; private set; } = null!;

        [Output("scorePenalty")]
        public Output<int> ScorePenalty { get; private set; } = null!;

        [Output("servermonitorLivenessCount")]
        public Output<int> ServermonitorLivenessCount { get; private set; } = null!;

        [Output("servermonitorLoadCount")]
        public Output<int> ServermonitorLoadCount { get; private set; } = null!;

        [Output("servermonitorPool")]
        public Output<string> ServermonitorPool { get; private set; } = null!;

        [Output("stateOrProvince")]
        public Output<string?> StateOrProvince { get; private set; } = null!;

        [Output("virtual")]
        public Output<bool> Virtual { get; private set; } = null!;

        [Output("waitOnComplete")]
        public Output<bool?> WaitOnComplete { get; private set; } = null!;


        /// <summary>
        /// Create a GtmDatacenter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GtmDatacenter(string name, GtmDatacenterArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/gtmDatacenter:GtmDatacenter", name, args ?? new GtmDatacenterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GtmDatacenter(string name, Input<string> id, GtmDatacenterState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/gtmDatacenter:GtmDatacenter", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "akamai:trafficmanagement/gtmDatacenter:GtmDatacenter"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GtmDatacenter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GtmDatacenter Get(string name, Input<string> id, GtmDatacenterState? state = null, CustomResourceOptions? options = null)
        {
            return new GtmDatacenter(name, id, state, options);
        }
    }

    public sealed class GtmDatacenterArgs : Pulumi.ResourceArgs
    {
        [Input("city")]
        public Input<string>? City { get; set; }

        [Input("cloneOf")]
        public Input<int>? CloneOf { get; set; }

        [Input("cloudServerHostHeaderOverride")]
        public Input<bool>? CloudServerHostHeaderOverride { get; set; }

        [Input("cloudServerTargeting")]
        public Input<bool>? CloudServerTargeting { get; set; }

        [Input("continent")]
        public Input<string>? Continent { get; set; }

        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("defaultLoadObject")]
        public Input<Inputs.GtmDatacenterDefaultLoadObjectArgs>? DefaultLoadObject { get; set; }

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("latitude")]
        public Input<double>? Latitude { get; set; }

        [Input("longitude")]
        public Input<double>? Longitude { get; set; }

        [Input("nickname")]
        public Input<string>? Nickname { get; set; }

        [Input("stateOrProvince")]
        public Input<string>? StateOrProvince { get; set; }

        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        public GtmDatacenterArgs()
        {
        }
    }

    public sealed class GtmDatacenterState : Pulumi.ResourceArgs
    {
        [Input("city")]
        public Input<string>? City { get; set; }

        [Input("cloneOf")]
        public Input<int>? CloneOf { get; set; }

        [Input("cloudServerHostHeaderOverride")]
        public Input<bool>? CloudServerHostHeaderOverride { get; set; }

        [Input("cloudServerTargeting")]
        public Input<bool>? CloudServerTargeting { get; set; }

        [Input("continent")]
        public Input<string>? Continent { get; set; }

        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("datacenterId")]
        public Input<int>? DatacenterId { get; set; }

        [Input("defaultLoadObject")]
        public Input<Inputs.GtmDatacenterDefaultLoadObjectGetArgs>? DefaultLoadObject { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("latitude")]
        public Input<double>? Latitude { get; set; }

        [Input("longitude")]
        public Input<double>? Longitude { get; set; }

        [Input("nickname")]
        public Input<string>? Nickname { get; set; }

        [Input("pingInterval")]
        public Input<int>? PingInterval { get; set; }

        [Input("pingPacketSize")]
        public Input<int>? PingPacketSize { get; set; }

        [Input("scorePenalty")]
        public Input<int>? ScorePenalty { get; set; }

        [Input("servermonitorLivenessCount")]
        public Input<int>? ServermonitorLivenessCount { get; set; }

        [Input("servermonitorLoadCount")]
        public Input<int>? ServermonitorLoadCount { get; set; }

        [Input("servermonitorPool")]
        public Input<string>? ServermonitorPool { get; set; }

        [Input("stateOrProvince")]
        public Input<string>? StateOrProvince { get; set; }

        [Input("virtual")]
        public Input<bool>? Virtual { get; set; }

        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        public GtmDatacenterState()
        {
        }
    }
}
