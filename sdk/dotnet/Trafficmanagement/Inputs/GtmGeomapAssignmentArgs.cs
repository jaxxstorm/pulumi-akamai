// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Trafficmanagement.Inputs
{

    public sealed class GtmGeomapAssignmentArgs : Pulumi.ResourceArgs
    {
        [Input("countries")]
        private InputList<string>? _countries;
        public InputList<string> Countries
        {
            get => _countries ?? (_countries = new InputList<string>());
            set => _countries = value;
        }

        [Input("datacenterId", required: true)]
        public Input<int> DatacenterId { get; set; } = null!;

        [Input("nickname", required: true)]
        public Input<string> Nickname { get; set; } = null!;

        public GtmGeomapAssignmentArgs()
        {
        }
    }
}
