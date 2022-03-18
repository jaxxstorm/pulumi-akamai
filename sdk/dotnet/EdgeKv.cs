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
    /// The `akamai.EdgeKv` resource lets you control EdgeKV database functions outside EdgeWorkers JavaScript code. Refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/welcome-to-edgekv) for more information.
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
    ///         var testStaging = new Akamai.EdgeKv("testStaging", new Akamai.EdgeKvArgs
    ///         {
    ///             GeoLocation = "US",
    ///             GroupId = 4284,
    ///             InitialDatas = 
    ///             {
    ///                 new Akamai.Inputs.EdgeKvInitialDataArgs
    ///                 {
    ///                     Group = "translations",
    ///                     Key = "lang",
    ///                     Value = "English",
    ///                 },
    ///             },
    ///             NamespaceName = "Marketing",
    ///             Network = "staging",
    ///             RetentionInSeconds = 15724800,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Attributes reference
    /// 
    /// There are no supported arguments for this resource.
    /// </summary>
    [AkamaiResourceType("akamai:index/edgeKv:EdgeKv")]
    public partial class EdgeKv : Pulumi.CustomResource
    {
        /// <summary>
        /// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
        /// </summary>
        [Output("geoLocation")]
        public Output<string?> GeoLocation { get; private set; } = null!;

        /// <summary>
        /// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
        /// </summary>
        [Output("groupId")]
        public Output<int> GroupId { get; private set; } = null!;

        /// <summary>
        /// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
        /// </summary>
        [Output("initialDatas")]
        public Output<ImmutableArray<Outputs.EdgeKvInitialData>> InitialDatas { get; private set; } = null!;

        /// <summary>
        /// - (Required) The name of the namespace.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
        /// </summary>
        [Output("retentionInSeconds")]
        public Output<int> RetentionInSeconds { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeKv resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeKv(string name, EdgeKvArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/edgeKv:EdgeKv", name, args ?? new EdgeKvArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeKv(string name, Input<string> id, EdgeKvState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/edgeKv:EdgeKv", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EdgeKv resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeKv Get(string name, Input<string> id, EdgeKvState? state = null, CustomResourceOptions? options = null)
        {
            return new EdgeKv(name, id, state, options);
        }
    }

    public sealed class EdgeKvArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
        /// </summary>
        [Input("geoLocation")]
        public Input<string>? GeoLocation { get; set; }

        /// <summary>
        /// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        [Input("initialDatas")]
        private InputList<Inputs.EdgeKvInitialDataArgs>? _initialDatas;

        /// <summary>
        /// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
        /// </summary>
        public InputList<Inputs.EdgeKvInitialDataArgs> InitialDatas
        {
            get => _initialDatas ?? (_initialDatas = new InputList<Inputs.EdgeKvInitialDataArgs>());
            set => _initialDatas = value;
        }

        /// <summary>
        /// - (Required) The name of the namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
        /// </summary>
        [Input("retentionInSeconds", required: true)]
        public Input<int> RetentionInSeconds { get; set; } = null!;

        public EdgeKvArgs()
        {
        }
    }

    public sealed class EdgeKvState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
        /// </summary>
        [Input("geoLocation")]
        public Input<string>? GeoLocation { get; set; }

        /// <summary>
        /// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        [Input("initialDatas")]
        private InputList<Inputs.EdgeKvInitialDataGetArgs>? _initialDatas;

        /// <summary>
        /// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
        /// </summary>
        public InputList<Inputs.EdgeKvInitialDataGetArgs> InitialDatas
        {
            get => _initialDatas ?? (_initialDatas = new InputList<Inputs.EdgeKvInitialDataGetArgs>());
            set => _initialDatas = value;
        }

        /// <summary>
        /// - (Required) The name of the namespace.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
        /// </summary>
        [Input("retentionInSeconds")]
        public Input<int>? RetentionInSeconds { get; set; }

        public EdgeKvState()
        {
        }
    }
}
