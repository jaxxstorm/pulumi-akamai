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
    public sealed class GtmASmapAssignment
    {
        /// <summary>
        /// Specifies an array of AS numbers.
        /// </summary>
        public readonly ImmutableArray<int> AsNumbers;
        /// <summary>
        /// A unique identifier for an existing data center in the domain.
        /// </summary>
        public readonly int DatacenterId;
        /// <summary>
        /// A descriptive label for the group.
        /// </summary>
        public readonly string Nickname;

        [OutputConstructor]
        private GtmASmapAssignment(
            ImmutableArray<int> asNumbers,

            int datacenterId,

            string nickname)
        {
            AsNumbers = asNumbers;
            DatacenterId = datacenterId;
            Nickname = nickname;
        }
    }
}
