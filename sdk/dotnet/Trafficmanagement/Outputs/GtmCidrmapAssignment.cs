// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Trafficmanagement.Outputs
{

    [OutputType]
    public sealed class GtmCidrmapAssignment
    {
        public readonly ImmutableArray<string> Blocks;
        public readonly int DatacenterId;
        public readonly string Nickname;

        [OutputConstructor]
        private GtmCidrmapAssignment(
            ImmutableArray<string> blocks,

            int datacenterId,

            string nickname)
        {
            Blocks = blocks;
            DatacenterId = datacenterId;
            Nickname = nickname;
        }
    }
}
