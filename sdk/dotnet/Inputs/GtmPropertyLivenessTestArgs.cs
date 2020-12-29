// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class GtmPropertyLivenessTestArgs : Pulumi.ResourceArgs
    {
        [Input("answersRequired")]
        public Input<bool>? AnswersRequired { get; set; }

        /// <summary>
        /// * `error_penalty`
        /// </summary>
        [Input("disableNonstandardPortWarning")]
        public Input<bool>? DisableNonstandardPortWarning { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("errorPenalty")]
        public Input<double>? ErrorPenalty { get; set; }

        [Input("httpError3xx")]
        public Input<bool>? HttpError3xx { get; set; }

        [Input("httpError4xx")]
        public Input<bool>? HttpError4xx { get; set; }

        [Input("httpError5xx")]
        public Input<bool>? HttpError5xx { get; set; }

        [Input("httpHeaders")]
        private InputList<Inputs.GtmPropertyLivenessTestHttpHeaderArgs>? _httpHeaders;

        /// <summary>
        /// `name`
        /// `value`
        /// </summary>
        public InputList<Inputs.GtmPropertyLivenessTestHttpHeaderArgs> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputList<Inputs.GtmPropertyLivenessTestHttpHeaderArgs>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// Liveness test name
        /// * `test_interval`
        /// * `test_object_protocol`
        /// * `test_timeout`
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("peerCertificateVerification")]
        public Input<bool>? PeerCertificateVerification { get; set; }

        /// <summary>
        /// * `request_string`
        /// * `resource_type`
        /// * `response_string`
        /// * `ssl_client_certificate`
        /// * `ssl_client_private_key`
        /// * `test_object`
        /// * `test_object_password`
        /// * `test_object_port`
        /// * `test_object_username`
        /// * `timeout_penalty`
        /// </summary>
        [Input("recursionRequested")]
        public Input<bool>? RecursionRequested { get; set; }

        [Input("requestString")]
        public Input<string>? RequestString { get; set; }

        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        [Input("responseString")]
        public Input<string>? ResponseString { get; set; }

        [Input("sslClientCertificate")]
        public Input<string>? SslClientCertificate { get; set; }

        [Input("sslClientPrivateKey")]
        public Input<string>? SslClientPrivateKey { get; set; }

        [Input("testInterval", required: true)]
        public Input<int> TestInterval { get; set; } = null!;

        [Input("testObject", required: true)]
        public Input<string> TestObject { get; set; } = null!;

        [Input("testObjectPassword")]
        public Input<string>? TestObjectPassword { get; set; }

        [Input("testObjectPort")]
        public Input<int>? TestObjectPort { get; set; }

        [Input("testObjectProtocol", required: true)]
        public Input<string> TestObjectProtocol { get; set; } = null!;

        [Input("testObjectUsername")]
        public Input<string>? TestObjectUsername { get; set; }

        [Input("testTimeout", required: true)]
        public Input<double> TestTimeout { get; set; } = null!;

        [Input("timeoutPenalty")]
        public Input<double>? TimeoutPenalty { get; set; }

        public GtmPropertyLivenessTestArgs()
        {
        }
    }
}
