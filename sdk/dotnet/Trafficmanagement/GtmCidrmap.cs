// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Trafficmanagement
{
    [Obsolete(@"akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap")]
    [AkamaiResourceType("akamai:trafficmanagement/gtmCidrmap:GtmCidrmap")]
    public partial class GtmCidrmap : Pulumi.CustomResource
    {
        [Output("assignments")]
        public Output<ImmutableArray<Outputs.GtmCidrmapAssignment>> Assignments { get; private set; } = null!;

        [Output("defaultDatacenter")]
        public Output<Outputs.GtmCidrmapDefaultDatacenter> DefaultDatacenter { get; private set; } = null!;

        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("waitOnComplete")]
        public Output<bool?> WaitOnComplete { get; private set; } = null!;


        /// <summary>
        /// Create a GtmCidrmap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GtmCidrmap(string name, GtmCidrmapArgs args, CustomResourceOptions? options = null)
            : base("akamai:trafficmanagement/gtmCidrmap:GtmCidrmap", name, args ?? new GtmCidrmapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GtmCidrmap(string name, Input<string> id, GtmCidrmapState? state = null, CustomResourceOptions? options = null)
            : base("akamai:trafficmanagement/gtmCidrmap:GtmCidrmap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GtmCidrmap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GtmCidrmap Get(string name, Input<string> id, GtmCidrmapState? state = null, CustomResourceOptions? options = null)
        {
            return new GtmCidrmap(name, id, state, options);
        }
    }

    public sealed class GtmCidrmapArgs : Pulumi.ResourceArgs
    {
        [Input("assignments")]
        private InputList<Inputs.GtmCidrmapAssignmentArgs>? _assignments;
        public InputList<Inputs.GtmCidrmapAssignmentArgs> Assignments
        {
            get => _assignments ?? (_assignments = new InputList<Inputs.GtmCidrmapAssignmentArgs>());
            set => _assignments = value;
        }

        [Input("defaultDatacenter", required: true)]
        public Input<Inputs.GtmCidrmapDefaultDatacenterArgs> DefaultDatacenter { get; set; } = null!;

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        public GtmCidrmapArgs()
        {
        }
    }

    public sealed class GtmCidrmapState : Pulumi.ResourceArgs
    {
        [Input("assignments")]
        private InputList<Inputs.GtmCidrmapAssignmentGetArgs>? _assignments;
        public InputList<Inputs.GtmCidrmapAssignmentGetArgs> Assignments
        {
            get => _assignments ?? (_assignments = new InputList<Inputs.GtmCidrmapAssignmentGetArgs>());
            set => _assignments = value;
        }

        [Input("defaultDatacenter")]
        public Input<Inputs.GtmCidrmapDefaultDatacenterGetArgs>? DefaultDatacenter { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("waitOnComplete")]
        public Input<bool>? WaitOnComplete { get; set; }

        public GtmCidrmapState()
        {
        }
    }
}
