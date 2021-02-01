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
    /// `akamai.GtmDatacenter` provides the resource for creating, configuring and importing a gtm datacenter to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existing_domain_name`:`existing_datacenter_id`
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
    /// </summary>
    [AkamaiResourceType("akamai:index/gtmDatacenter:GtmDatacenter")]
    public partial class GtmDatacenter : Pulumi.CustomResource
    {
        [Output("city")]
        public Output<string?> City { get; private set; } = null!;

        [Output("cloneOf")]
        public Output<int?> CloneOf { get; private set; } = null!;

        /// <summary>
        /// * `continent`
        /// * `country`
        /// * `latitude`
        /// * `longitude`
        /// * `state_or_province`
        /// </summary>
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

        /// <summary>
        /// Domain name
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        [Output("latitude")]
        public Output<double?> Latitude { get; private set; } = null!;

        [Output("longitude")]
        public Output<double?> Longitude { get; private set; } = null!;

        /// <summary>
        /// datacenter nickname
        /// * `default_load_object`
        /// * `load_object`
        /// * `load_object_port`
        /// </summary>
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

        /// <summary>
        /// Wait for transaction to complete
        /// </summary>
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

        /// <summary>
        /// * `continent`
        /// * `country`
        /// * `latitude`
        /// * `longitude`
        /// * `state_or_province`
        /// </summary>
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

        /// <summary>
        /// Domain name
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("latitude")]
        public Input<double>? Latitude { get; set; }

        [Input("longitude")]
        public Input<double>? Longitude { get; set; }

        /// <summary>
        /// datacenter nickname
        /// * `default_load_object`
        /// * `load_object`
        /// * `load_object_port`
        /// </summary>
        [Input("nickname")]
        public Input<string>? Nickname { get; set; }

        [Input("stateOrProvince")]
        public Input<string>? StateOrProvince { get; set; }

        /// <summary>
        /// Wait for transaction to complete
        /// </summary>
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

        /// <summary>
        /// * `continent`
        /// * `country`
        /// * `latitude`
        /// * `longitude`
        /// * `state_or_province`
        /// </summary>
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

        /// <summary>
        /// Domain name
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("latitude")]
        public Input<double>? Latitude { get; set; }

        [Input("longitude")]
        public Input<double>? Longitude { get; set; }

        /// <summary>
        /// datacenter nickname
        /// * `default_load_object`
        /// * `load_object`
        /// * `load_object_port`
        /// </summary>
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

        /// <summary>
        /// Wait for transaction to complete
        /// </summary>
        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        public GtmDatacenterState()
        {
        }
    }
}
