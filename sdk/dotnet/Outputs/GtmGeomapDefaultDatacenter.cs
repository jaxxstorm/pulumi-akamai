// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Outputs
{

    [OutputType]
    public sealed class GtmGeomapDefaultDatacenter
    {
        /// <summary>
        /// A unique identifier for an existing data center in the domain.
        /// </summary>
        public readonly int DatacenterId;
        /// <summary>
        /// A descriptive label for the group.
        /// </summary>
        public readonly string? Nickname;

        [OutputConstructor]
        private GtmGeomapDefaultDatacenter(
            int datacenterId,

            string? nickname)
        {
            DatacenterId = datacenterId;
            Nickname = nickname;
        }
    }
}
