// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class DatastreamConfigFrequencyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// - (Required) The time in seconds after which the system bundles log lines into a file and sends it to a destination. `30` or `60` are the possible values.
        /// </summary>
        [Input("timeInSec", required: true)]
        public Input<int> TimeInSec { get; set; } = null!;

        public DatastreamConfigFrequencyGetArgs()
        {
        }
    }
}
