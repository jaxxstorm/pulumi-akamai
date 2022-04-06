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
    public sealed class GtmPropertyTrafficTarget
    {
        /// <summary>
        /// A unique identifier for an existing data center in the domain.
        /// </summary>
        public readonly int? DatacenterId;
        /// <summary>
        /// A boolean indicating whether the traffic target is used. You can also omit the traffic target, which has the same result as the false value.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Specifies an optional data center for the property. Used when there are no servers configured for the property.
        /// </summary>
        public readonly string? HandoutCname;
        /// <summary>
        /// Name of HTTP header.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (List) Identifies the IP address or the hostnames of the servers.
        /// </summary>
        public readonly ImmutableArray<string> Servers;
        /// <summary>
        /// Specifies the traffic weight for the target.
        /// </summary>
        public readonly double? Weight;

        [OutputConstructor]
        private GtmPropertyTrafficTarget(
            int? datacenterId,

            bool? enabled,

            string? handoutCname,

            string? name,

            ImmutableArray<string> servers,

            double? weight)
        {
            DatacenterId = datacenterId;
            Enabled = enabled;
            HandoutCname = handoutCname;
            Name = name;
            Servers = servers;
            Weight = weight;
        }
    }
}
