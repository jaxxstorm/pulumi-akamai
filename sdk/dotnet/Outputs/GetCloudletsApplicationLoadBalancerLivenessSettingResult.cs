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
    public sealed class GetCloudletsApplicationLoadBalancerLivenessSettingResult
    {
        public readonly ImmutableDictionary<string, string> AdditionalHeaders;
        public readonly string HostHeader;
        public readonly int Interval;
        public readonly string Path;
        public readonly bool PeerCertificateVerification;
        public readonly int Port;
        public readonly string Protocol;
        public readonly string RequestString;
        public readonly string ResponseString;
        public readonly bool Status3xxFailure;
        public readonly bool Status4xxFailure;
        public readonly bool Status5xxFailure;
        public readonly double Timeout;

        [OutputConstructor]
        private GetCloudletsApplicationLoadBalancerLivenessSettingResult(
            ImmutableDictionary<string, string> additionalHeaders,

            string hostHeader,

            int interval,

            string path,

            bool peerCertificateVerification,

            int port,

            string protocol,

            string requestString,

            string responseString,

            bool status3xxFailure,

            bool status4xxFailure,

            bool status5xxFailure,

            double timeout)
        {
            AdditionalHeaders = additionalHeaders;
            HostHeader = hostHeader;
            Interval = interval;
            Path = path;
            PeerCertificateVerification = peerCertificateVerification;
            Port = port;
            Protocol = protocol;
            RequestString = requestString;
            ResponseString = responseString;
            Status3xxFailure = status3xxFailure;
            Status4xxFailure = status4xxFailure;
            Status5xxFailure = status5xxFailure;
            Timeout = timeout;
        }
    }
}
