// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Trafficmanagement.Inputs
{

    public sealed class GtmResourceResourceInstanceGetArgs : Pulumi.ResourceArgs
    {
        [Input("datacenterId", required: true)]
        public Input<int> DatacenterId { get; set; } = null!;

        [Input("loadObject")]
        public Input<string>? LoadObject { get; set; }

        [Input("loadObjectPort")]
        public Input<int>? LoadObjectPort { get; set; }

        [Input("loadServers")]
        private InputList<string>? _loadServers;

        /// <summary>
        /// — (List)
        /// </summary>
        public InputList<string> LoadServers
        {
            get => _loadServers ?? (_loadServers = new InputList<string>());
            set => _loadServers = value;
        }

        /// <summary>
        /// — (Boolean)
        /// * `host_header`
        /// * `least_squares_decay`
        /// * `upper_bound`
        /// * `description`
        /// * `leader_string`
        /// * `constrained_property`
        /// * `load_imbalance_percent`
        /// * `max_u_multiplicative_increment`
        /// * `decay_rate`
        /// </summary>
        [Input("useDefaultLoadObject")]
        public Input<bool>? UseDefaultLoadObject { get; set; }

        public GtmResourceResourceInstanceGetArgs()
        {
        }
    }
}