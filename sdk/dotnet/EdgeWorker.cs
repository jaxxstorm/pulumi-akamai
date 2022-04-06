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
    /// The `akamai.EdgeWorker` resource lets you deploy custom code on thousands of edge servers and apply logic that creates powerful web experiences.
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
    ///         var ew = new Akamai.EdgeWorker("ew", new Akamai.EdgeWorkerArgs
    ///         {
    ///             GroupId = 72297,
    ///             ResourceTierId = 100,
    ///             LocalBundle = @var.Bundle_path,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Attributes reference
    /// 
    /// * `edgeworker_id` - Unique identifier for an EdgeWorker ID.
    /// * `local_bundle_hash` - A SHA-256 hash digest of the EdgeWorkers code bundle.
    /// * `version` - Unique identifier for a specific EdgeWorker version.
    /// * `warnings` - List of validation warnings.
    /// </summary>
    [AkamaiResourceType("akamai:index/edgeWorker:EdgeWorker")]
    public partial class EdgeWorker : Pulumi.CustomResource
    {
        /// <summary>
        /// The unique identifier of the EdgeWorker
        /// </summary>
        [Output("edgeworkerId")]
        public Output<int> EdgeworkerId { get; private set; } = null!;

        /// <summary>
        /// - (Required) Identifies a group to assign to the EdgeWorker ID.
        /// </summary>
        [Output("groupId")]
        public Output<int> GroupId { get; private set; } = null!;

        /// <summary>
        /// - (Optional) The path to the EdgeWorkers code bundle.
        /// </summary>
        [Output("localBundle")]
        public Output<string?> LocalBundle { get; private set; } = null!;

        /// <summary>
        /// The local bundle hash for the EdgeWorker
        /// </summary>
        [Output("localBundleHash")]
        public Output<string> LocalBundleHash { get; private set; } = null!;

        /// <summary>
        /// - (Required) The name of the EdgeWorker ID.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// - (Required) Unique identifier of the resource tier.
        /// </summary>
        [Output("resourceTierId")]
        public Output<int> ResourceTierId { get; private set; } = null!;

        /// <summary>
        /// The bundle version
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The list of warnings returned by EdgeWorker validation
        /// </summary>
        [Output("warnings")]
        public Output<ImmutableArray<string>> Warnings { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeWorker resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeWorker(string name, EdgeWorkerArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/edgeWorker:EdgeWorker", name, args ?? new EdgeWorkerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeWorker(string name, Input<string> id, EdgeWorkerState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/edgeWorker:EdgeWorker", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EdgeWorker resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeWorker Get(string name, Input<string> id, EdgeWorkerState? state = null, CustomResourceOptions? options = null)
        {
            return new EdgeWorker(name, id, state, options);
        }
    }

    public sealed class EdgeWorkerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// - (Required) Identifies a group to assign to the EdgeWorker ID.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        /// <summary>
        /// - (Optional) The path to the EdgeWorkers code bundle.
        /// </summary>
        [Input("localBundle")]
        public Input<string>? LocalBundle { get; set; }

        /// <summary>
        /// - (Required) The name of the EdgeWorker ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// - (Required) Unique identifier of the resource tier.
        /// </summary>
        [Input("resourceTierId", required: true)]
        public Input<int> ResourceTierId { get; set; } = null!;

        public EdgeWorkerArgs()
        {
        }
    }

    public sealed class EdgeWorkerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the EdgeWorker
        /// </summary>
        [Input("edgeworkerId")]
        public Input<int>? EdgeworkerId { get; set; }

        /// <summary>
        /// - (Required) Identifies a group to assign to the EdgeWorker ID.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// - (Optional) The path to the EdgeWorkers code bundle.
        /// </summary>
        [Input("localBundle")]
        public Input<string>? LocalBundle { get; set; }

        /// <summary>
        /// The local bundle hash for the EdgeWorker
        /// </summary>
        [Input("localBundleHash")]
        public Input<string>? LocalBundleHash { get; set; }

        /// <summary>
        /// - (Required) The name of the EdgeWorker ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// - (Required) Unique identifier of the resource tier.
        /// </summary>
        [Input("resourceTierId")]
        public Input<int>? ResourceTierId { get; set; }

        /// <summary>
        /// The bundle version
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("warnings")]
        private InputList<string>? _warnings;

        /// <summary>
        /// The list of warnings returned by EdgeWorker validation
        /// </summary>
        public InputList<string> Warnings
        {
            get => _warnings ?? (_warnings = new InputList<string>());
            set => _warnings = value;
        }

        public EdgeWorkerState()
        {
        }
    }
}
