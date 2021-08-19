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
    /// Use the `akamai.NetworkList` resource to create a network list, or to modify an existing list.
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
    ///         var networkList = new Akamai.NetworkList("networkList", new Akamai.NetworkListArgs
    ///         {
    ///             Type = "IP",
    ///             Description = "network list description",
    ///             Lists = @var.List,
    ///             Mode = "APPEND",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/networkList:NetworkList")]
    public partial class NetworkList : Pulumi.CustomResource
    {
        /// <summary>
        /// The description to be assigned to the network list.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
        /// removed from an existing list.
        /// </summary>
        [Output("lists")]
        public Output<ImmutableArray<string>> Lists { get; private set; } = null!;

        /// <summary>
        /// A string specifying the interpretation of the `list` parameter. Must be one of the following:
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// The name to be assigned to the network list.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the network list.
        /// </summary>
        [Output("networkListId")]
        public Output<string> NetworkListId { get; private set; } = null!;

        /// <summary>
        /// An integer that identifies the current version of the network list; this value is incremented each time
        /// the list is modified.
        /// </summary>
        [Output("syncPoint")]
        public Output<int> SyncPoint { get; private set; } = null!;

        /// <summary>
        /// The type of the network list; must be either "IP" or "GEO".
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// uniqueId
        /// </summary>
        [Output("uniqueid")]
        public Output<string> Uniqueid { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkList(string name, NetworkListArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/networkList:NetworkList", name, args ?? new NetworkListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkList(string name, Input<string> id, NetworkListState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/networkList:NetworkList", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkList Get(string name, Input<string> id, NetworkListState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkList(name, id, state, options);
        }
    }

    public sealed class NetworkListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description to be assigned to the network list.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("lists")]
        private InputList<string>? _lists;

        /// <summary>
        /// : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
        /// removed from an existing list.
        /// </summary>
        public InputList<string> Lists
        {
            get => _lists ?? (_lists = new InputList<string>());
            set => _lists = value;
        }

        /// <summary>
        /// A string specifying the interpretation of the `list` parameter. Must be one of the following:
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// The name to be assigned to the network list.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the network list; must be either "IP" or "GEO".
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public NetworkListArgs()
        {
        }
    }

    public sealed class NetworkListState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description to be assigned to the network list.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("lists")]
        private InputList<string>? _lists;

        /// <summary>
        /// : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
        /// removed from an existing list.
        /// </summary>
        public InputList<string> Lists
        {
            get => _lists ?? (_lists = new InputList<string>());
            set => _lists = value;
        }

        /// <summary>
        /// A string specifying the interpretation of the `list` parameter. Must be one of the following:
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name to be assigned to the network list.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the network list.
        /// </summary>
        [Input("networkListId")]
        public Input<string>? NetworkListId { get; set; }

        /// <summary>
        /// An integer that identifies the current version of the network list; this value is incremented each time
        /// the list is modified.
        /// </summary>
        [Input("syncPoint")]
        public Input<int>? SyncPoint { get; set; }

        /// <summary>
        /// The type of the network list; must be either "IP" or "GEO".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// uniqueId
        /// </summary>
        [Input("uniqueid")]
        public Input<string>? Uniqueid { get; set; }

        public NetworkListState()
        {
        }
    }
}