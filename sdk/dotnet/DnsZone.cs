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
    /// The `akamai.DnsZone` provides the resource for configuring a DNS zone to integrate easily with your existing DNS infrastructure to provide a secure, high performance, highly available and scalable solution for DNS hosting.
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
    ///         var demozone = new Akamai.DnsZone("demozone", new Akamai.DnsZoneArgs
    ///         {
    ///             Comment = "some comment",
    ///             Contract = "ctr_1-AB123",
    ///             Group = "100",
    ///             Masters = 
    ///             {
    ///                 "1.2.3.4",
    ///                 "1.2.3.5",
    ///             },
    ///             SignAndServe = false,
    ///             Type = "secondary",
    ///             Zone = "example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/dnsZone:DnsZone")]
    public partial class DnsZone : Pulumi.CustomResource
    {
        [Output("activationState")]
        public Output<string> ActivationState { get; private set; } = null!;

        [Output("aliasCount")]
        public Output<int> AliasCount { get; private set; } = null!;

        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// The contract ID.
        /// </summary>
        [Output("contract")]
        public Output<string> Contract { get; private set; } = null!;

        [Output("endCustomerId")]
        public Output<string?> EndCustomerId { get; private set; } = null!;

        /// <summary>
        /// The currently selected group ID.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The names or addresses of the customer’s nameservers from which the zone data should be retrieved.
        /// </summary>
        [Output("masters")]
        public Output<ImmutableArray<string>> Masters { get; private set; } = null!;

        /// <summary>
        /// Whether DNSSEC Sign&amp;Serve is enabled.
        /// </summary>
        [Output("signAndServe")]
        public Output<bool?> SignAndServe { get; private set; } = null!;

        /// <summary>
        /// Algorithm used by Sign&amp;Serve.
        /// </summary>
        [Output("signAndServeAlgorithm")]
        public Output<string?> SignAndServeAlgorithm { get; private set; } = null!;

        /// <summary>
        /// The name of the zone whose configuration this zone will copy.
        /// </summary>
        [Output("target")]
        public Output<string?> Target { get; private set; } = null!;

        /// <summary>
        /// TSIG Key used in secure zone transfers
        /// </summary>
        [Output("tsigKey")]
        public Output<Outputs.DnsZoneTsigKey?> TsigKey { get; private set; } = null!;

        /// <summary>
        /// Whether the zone is `primary` or `secondary`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;

        /// <summary>
        /// Domain zone, encapsulating any nested subdomains.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a DnsZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsZone(string name, DnsZoneArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/dnsZone:DnsZone", name, args ?? new DnsZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsZone(string name, Input<string> id, DnsZoneState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/dnsZone:DnsZone", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "akamai:edgedns/dnsZone:DnsZone"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DnsZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsZone Get(string name, Input<string> id, DnsZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsZone(name, id, state, options);
        }
    }

    public sealed class DnsZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The contract ID.
        /// </summary>
        [Input("contract", required: true)]
        public Input<string> Contract { get; set; } = null!;

        [Input("endCustomerId")]
        public Input<string>? EndCustomerId { get; set; }

        /// <summary>
        /// The currently selected group ID.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        [Input("masters")]
        private InputList<string>? _masters;

        /// <summary>
        /// The names or addresses of the customer’s nameservers from which the zone data should be retrieved.
        /// </summary>
        public InputList<string> Masters
        {
            get => _masters ?? (_masters = new InputList<string>());
            set => _masters = value;
        }

        /// <summary>
        /// Whether DNSSEC Sign&amp;Serve is enabled.
        /// </summary>
        [Input("signAndServe")]
        public Input<bool>? SignAndServe { get; set; }

        /// <summary>
        /// Algorithm used by Sign&amp;Serve.
        /// </summary>
        [Input("signAndServeAlgorithm")]
        public Input<string>? SignAndServeAlgorithm { get; set; }

        /// <summary>
        /// The name of the zone whose configuration this zone will copy.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// TSIG Key used in secure zone transfers
        /// </summary>
        [Input("tsigKey")]
        public Input<Inputs.DnsZoneTsigKeyArgs>? TsigKey { get; set; }

        /// <summary>
        /// Whether the zone is `primary` or `secondary`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Domain zone, encapsulating any nested subdomains.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public DnsZoneArgs()
        {
        }
    }

    public sealed class DnsZoneState : Pulumi.ResourceArgs
    {
        [Input("activationState")]
        public Input<string>? ActivationState { get; set; }

        [Input("aliasCount")]
        public Input<int>? AliasCount { get; set; }

        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The contract ID.
        /// </summary>
        [Input("contract")]
        public Input<string>? Contract { get; set; }

        [Input("endCustomerId")]
        public Input<string>? EndCustomerId { get; set; }

        /// <summary>
        /// The currently selected group ID.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("masters")]
        private InputList<string>? _masters;

        /// <summary>
        /// The names or addresses of the customer’s nameservers from which the zone data should be retrieved.
        /// </summary>
        public InputList<string> Masters
        {
            get => _masters ?? (_masters = new InputList<string>());
            set => _masters = value;
        }

        /// <summary>
        /// Whether DNSSEC Sign&amp;Serve is enabled.
        /// </summary>
        [Input("signAndServe")]
        public Input<bool>? SignAndServe { get; set; }

        /// <summary>
        /// Algorithm used by Sign&amp;Serve.
        /// </summary>
        [Input("signAndServeAlgorithm")]
        public Input<string>? SignAndServeAlgorithm { get; set; }

        /// <summary>
        /// The name of the zone whose configuration this zone will copy.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// TSIG Key used in secure zone transfers
        /// </summary>
        [Input("tsigKey")]
        public Input<Inputs.DnsZoneTsigKeyGetArgs>? TsigKey { get; set; }

        /// <summary>
        /// Whether the zone is `primary` or `secondary`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        /// <summary>
        /// Domain zone, encapsulating any nested subdomains.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DnsZoneState()
        {
        }
    }
}
