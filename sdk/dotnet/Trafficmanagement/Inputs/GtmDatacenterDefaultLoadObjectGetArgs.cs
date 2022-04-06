// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Trafficmanagement.Inputs
{

    public sealed class GtmDatacenterDefaultLoadObjectGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A load object is a file that provides real-time information about the current load, maximum allowable load, and target load on each resource.
        /// </summary>
        [Input("loadObject")]
        public Input<string>? LoadObject { get; set; }

        /// <summary>
        /// Specifies the TCP port to connect to when requesting the load object.
        /// </summary>
        [Input("loadObjectPort")]
        public Input<int>? LoadObjectPort { get; set; }

        [Input("loadServers")]
        private InputList<string>? _loadServers;

        /// <summary>
        /// Specifies a list of servers to request the load object from.
        /// </summary>
        public InputList<string> LoadServers
        {
            get => _loadServers ?? (_loadServers = new InputList<string>());
            set => _loadServers = value;
        }

        public GtmDatacenterDefaultLoadObjectGetArgs()
        {
        }
    }
}
