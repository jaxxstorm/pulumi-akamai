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
    /// Use the `akamai.EdgeWorkersActivation` resource to activate a specific EdgeWorker version. An activation deploys the version to either the Akamai staging or production network.
    /// 
    /// Before activating on production, activate on staging first. This way you can detect any problems in staging before your changes progress to production.
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
    ///         var test = new Akamai.EdgeWorkersActivation("test", new Akamai.EdgeWorkersActivationArgs
    ///         {
    ///             EdgeworkerId = 1234,
    ///             Network = "STAGING",
    ///             Version = "test1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/edgeWorkersActivation:EdgeWorkersActivation")]
    public partial class EdgeWorkersActivation : Pulumi.CustomResource
    {
        /// <summary>
        /// (Required) Unique identifier of the activation.
        /// </summary>
        [Output("activationId")]
        public Output<int> ActivationId { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the EdgeWorker ID you want to activate.
        /// </summary>
        [Output("edgeworkerId")]
        public Output<int> EdgeworkerId { get; private set; } = null!;

        /// <summary>
        /// The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The EdgeWorker version you want to activate.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeWorkersActivation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeWorkersActivation(string name, EdgeWorkersActivationArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/edgeWorkersActivation:EdgeWorkersActivation", name, args ?? new EdgeWorkersActivationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeWorkersActivation(string name, Input<string> id, EdgeWorkersActivationState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/edgeWorkersActivation:EdgeWorkersActivation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EdgeWorkersActivation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeWorkersActivation Get(string name, Input<string> id, EdgeWorkersActivationState? state = null, CustomResourceOptions? options = null)
        {
            return new EdgeWorkersActivation(name, id, state, options);
        }
    }

    public sealed class EdgeWorkersActivationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for the EdgeWorker ID you want to activate.
        /// </summary>
        [Input("edgeworkerId", required: true)]
        public Input<int> EdgeworkerId { get; set; } = null!;

        /// <summary>
        /// The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The EdgeWorker version you want to activate.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public EdgeWorkersActivationArgs()
        {
        }
    }

    public sealed class EdgeWorkersActivationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Required) Unique identifier of the activation.
        /// </summary>
        [Input("activationId")]
        public Input<int>? ActivationId { get; set; }

        /// <summary>
        /// A unique identifier for the EdgeWorker ID you want to activate.
        /// </summary>
        [Input("edgeworkerId")]
        public Input<int>? EdgeworkerId { get; set; }

        /// <summary>
        /// The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The EdgeWorker version you want to activate.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public EdgeWorkersActivationState()
        {
        }
    }
}
