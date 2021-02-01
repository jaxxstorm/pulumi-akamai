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
    /// `akamai.GtmAsmap` provides the resource for creating, configuring and importing a gtm AS Map to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existing_domain_name`:`existing_map_name`
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
    ///         var demoAsmap = new Akamai.GtmAsmap("demoAsmap", new Akamai.GtmAsmapArgs
    ///         {
    ///             DefaultDatacenter = new Akamai.Inputs.GtmAsmapDefaultDatacenterArgs
    ///             {
    ///                 DatacenterId = 5400,
    ///                 Nickname = "All Other AS numbers",
    ///             },
    ///             Domain = "demo_domain.akadns.net",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/gtmAsmap:GtmAsmap")]
    public partial class GtmAsmap : Pulumi.CustomResource
    {
        /// <summary>
        /// * `datacenter_id`
        /// * `nickname`
        /// </summary>
        [Output("assignments")]
        public Output<ImmutableArray<Outputs.GtmAsmapAssignment>> Assignments { get; private set; } = null!;

        [Output("defaultDatacenter")]
        public Output<Outputs.GtmAsmapDefaultDatacenter> DefaultDatacenter { get; private set; } = null!;

        /// <summary>
        /// Domain name
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// * `default_datacenter`
        /// * `datacenter_id`
        /// * `nickname`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Wait for transaction to complete
        /// </summary>
        [Output("waitOnComplete")]
        public Output<bool?> WaitOnComplete { get; private set; } = null!;


        /// <summary>
        /// Create a GtmAsmap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GtmAsmap(string name, GtmAsmapArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/gtmAsmap:GtmAsmap", name, args ?? new GtmAsmapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GtmAsmap(string name, Input<string> id, GtmAsmapState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/gtmAsmap:GtmAsmap", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "akamai:trafficmanagement/gtmASmap:GtmASmap"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GtmAsmap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GtmAsmap Get(string name, Input<string> id, GtmAsmapState? state = null, CustomResourceOptions? options = null)
        {
            return new GtmAsmap(name, id, state, options);
        }
    }

    public sealed class GtmAsmapArgs : Pulumi.ResourceArgs
    {
        [Input("assignments")]
        private InputList<Inputs.GtmAsmapAssignmentArgs>? _assignments;

        /// <summary>
        /// * `datacenter_id`
        /// * `nickname`
        /// </summary>
        public InputList<Inputs.GtmAsmapAssignmentArgs> Assignments
        {
            get => _assignments ?? (_assignments = new InputList<Inputs.GtmAsmapAssignmentArgs>());
            set => _assignments = value;
        }

        [Input("defaultDatacenter", required: true)]
        public Input<Inputs.GtmAsmapDefaultDatacenterArgs> DefaultDatacenter { get; set; } = null!;

        /// <summary>
        /// Domain name
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Resource name
        /// * `default_datacenter`
        /// * `datacenter_id`
        /// * `nickname`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Wait for transaction to complete
        /// </summary>
        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        public GtmAsmapArgs()
        {
        }
    }

    public sealed class GtmAsmapState : Pulumi.ResourceArgs
    {
        [Input("assignments")]
        private InputList<Inputs.GtmAsmapAssignmentGetArgs>? _assignments;

        /// <summary>
        /// * `datacenter_id`
        /// * `nickname`
        /// </summary>
        public InputList<Inputs.GtmAsmapAssignmentGetArgs> Assignments
        {
            get => _assignments ?? (_assignments = new InputList<Inputs.GtmAsmapAssignmentGetArgs>());
            set => _assignments = value;
        }

        [Input("defaultDatacenter")]
        public Input<Inputs.GtmAsmapDefaultDatacenterGetArgs>? DefaultDatacenter { get; set; }

        /// <summary>
        /// Domain name
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Resource name
        /// * `default_datacenter`
        /// * `datacenter_id`
        /// * `nickname`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Wait for transaction to complete
        /// </summary>
        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        public GtmAsmapState()
        {
        }
    }
}
