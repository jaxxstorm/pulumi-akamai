// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class PropertyOriginArgs : Pulumi.ResourceArgs
    {
        [Input("cacheKeyHostname")]
        public Input<string>? CacheKeyHostname { get; set; }

        [Input("compress")]
        public Input<bool>? Compress { get; set; }

        [Input("enableTrueClientIp")]
        public Input<bool>? EnableTrueClientIp { get; set; }

        [Input("forwardHostname")]
        public Input<string>? ForwardHostname { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        public PropertyOriginArgs()
        {
        }
    }
}
