// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class GtmCidrmapAssignmentGetArgs : Pulumi.ResourceArgs
    {
        [Input("blocks")]
        private InputList<string>? _blocks;

        /// <summary>
        /// Specifies an array of CIDR blocks.
        /// </summary>
        public InputList<string> Blocks
        {
            get => _blocks ?? (_blocks = new InputList<string>());
            set => _blocks = value;
        }

        /// <summary>
        /// A unique identifier for an existing data center in the domain.
        /// </summary>
        [Input("datacenterId", required: true)]
        public Input<int> DatacenterId { get; set; } = null!;

        /// <summary>
        /// A descriptive label for the CIDR zone group, up to 256 characters.
        /// </summary>
        [Input("nickname", required: true)]
        public Input<string> Nickname { get; set; } = null!;

        public GtmCidrmapAssignmentGetArgs()
        {
        }
    }
}
