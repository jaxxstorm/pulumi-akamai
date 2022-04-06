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
    public sealed class GtmResourceResourceInstance
    {
        /// <summary>
        /// A unique identifier for an existing data center in the domain.
        /// </summary>
        public readonly int DatacenterId;
        /// <summary>
        /// Identifies the load object file used to report real-time information about the current load, maximum allowable load, and target load on each resource.
        /// </summary>
        public readonly string? LoadObject;
        /// <summary>
        /// Specifies the TCP port of the `load_object`.
        /// </summary>
        public readonly int? LoadObjectPort;
        /// <summary>
        /// (List) Specifies a list of servers from which to request the load object.
        /// </summary>
        public readonly ImmutableArray<string> LoadServers;
        /// <summary>
        /// A boolean that indicates whether a default `load_object` is used for the resources.
        /// </summary>
        public readonly bool? UseDefaultLoadObject;

        [OutputConstructor]
        private GtmResourceResourceInstance(
            int datacenterId,

            string? loadObject,

            int? loadObjectPort,

            ImmutableArray<string> loadServers,

            bool? useDefaultLoadObject)
        {
            DatacenterId = datacenterId;
            LoadObject = loadObject;
            LoadObjectPort = loadObjectPort;
            LoadServers = loadServers;
            UseDefaultLoadObject = useDefaultLoadObject;
        }
    }
}
