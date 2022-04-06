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
    public sealed class GtmCidrmapDefaultDatacenter
    {
        /// <summary>
        /// A unique identifier for an existing data center in the domain.
        /// </summary>
        public readonly int DatacenterId;
        /// <summary>
        /// A descriptive label for the CIDR zone group, up to 256 characters.
        /// </summary>
        public readonly string? Nickname;

        [OutputConstructor]
        private GtmCidrmapDefaultDatacenter(
            int datacenterId,

            string? nickname)
        {
            DatacenterId = datacenterId;
            Nickname = nickname;
        }
    }
}
