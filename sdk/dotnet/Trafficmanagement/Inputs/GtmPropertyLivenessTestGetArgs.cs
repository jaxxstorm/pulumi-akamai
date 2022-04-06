// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Trafficmanagement.Inputs
{

    public sealed class GtmPropertyLivenessTestGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If `test_object_protocol` is DNS, enter a boolean value if an answer is needed for the DNS query to be successful.
        /// </summary>
        [Input("answersRequired")]
        public Input<bool>? AnswersRequired { get; set; }

        /// <summary>
        /// A boolean that if set to `true`, disables warnings when non-standard ports are used.
        /// </summary>
        [Input("disableNonstandardPortWarning")]
        public Input<bool>? DisableNonstandardPortWarning { get; set; }

        /// <summary>
        /// A boolean indicating whether the liveness test is disabled. When disabled, GTM stops running the test, effectively treating it as if it no longer exists.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Specifies the score that's reported if the liveness test encounters an error other than timeout, such as connection refused, and 404.
        /// </summary>
        [Input("errorPenalty")]
        public Input<double>? ErrorPenalty { get; set; }

        /// <summary>
        /// A boolean that if set to `true`, treats a 3xx HTTP response as a failure if the `test_object_protocol` is `http`, `https`, or `ftp`.
        /// </summary>
        [Input("httpError3xx")]
        public Input<bool>? HttpError3xx { get; set; }

        /// <summary>
        /// A boolean that if set to `true`, treats a 4xx HTTP response as a failure if the `test_object_protocol` is `http`, `https`, or `ftp`.
        /// </summary>
        [Input("httpError4xx")]
        public Input<bool>? HttpError4xx { get; set; }

        /// <summary>
        /// A boolean that if set to `true`, treats a 5xx HTTP response as a failure if the `test_object_protocol` is `http`, `https`, or `ftp`.
        /// </summary>
        [Input("httpError5xx")]
        public Input<bool>? HttpError5xx { get; set; }

        [Input("httpHeaders")]
        private InputList<Inputs.GtmPropertyLivenessTestHttpHeaderGetArgs>? _httpHeaders;

        /// <summary>
        /// Contains HTTP headers to send if the `test_object_protocol` is `http` or `https`. You can have multiple `http_header` entries. Requires these arguments:
        /// </summary>
        public InputList<Inputs.GtmPropertyLivenessTestHttpHeaderGetArgs> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputList<Inputs.GtmPropertyLivenessTestHttpHeaderGetArgs>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// Name of HTTP header.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// A boolean that if set to `true`, validates the origin certificate. Applies only to tests with `test_object_protocol` of https.
        /// </summary>
        [Input("peerCertificateVerification")]
        public Input<bool>? PeerCertificateVerification { get; set; }

        /// <summary>
        /// A boolean indicating whether the `test_object_protocol` is DNS. The DNS query is recursive.
        /// </summary>
        [Input("recursionRequested")]
        public Input<bool>? RecursionRequested { get; set; }

        /// <summary>
        /// Specifies a request string.
        /// </summary>
        [Input("requestString")]
        public Input<string>? RequestString { get; set; }

        /// <summary>
        /// Specifies the query type, if `test_object_protocol` is DNS.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// Specifies a response string.
        /// </summary>
        [Input("responseString")]
        public Input<string>? ResponseString { get; set; }

        /// <summary>
        /// Indicates a Base64-encoded certificate. SSL client certificates are available for livenessTests that use secure protocols.
        /// </summary>
        [Input("sslClientCertificate")]
        public Input<string>? SslClientCertificate { get; set; }

        /// <summary>
        /// Indicates a Base64-encoded private key. The private key used to generate or request a certificate for livenessTests can't have a passphrase nor be used for any other purpose.
        /// </summary>
        [Input("sslClientPrivateKey")]
        public Input<string>? SslClientPrivateKey { get; set; }

        /// <summary>
        /// Indicates the interval at which the liveness test is run, in seconds. Requires a minimum of 10 seconds.
        /// </summary>
        [Input("testInterval", required: true)]
        public Input<int> TestInterval { get; set; } = null!;

        /// <summary>
        /// Specifies the static text that acts as a stand-in for the data that you're sending on the network.
        /// </summary>
        [Input("testObject", required: true)]
        public Input<string> TestObject { get; set; } = null!;

        /// <summary>
        /// Specifies the test object's password. It is required if testObjectProtocol is ftp.
        /// </summary>
        [Input("testObjectPassword")]
        public Input<string>? TestObjectPassword { get; set; }

        /// <summary>
        /// Specifies the port number for the testObject.
        /// </summary>
        [Input("testObjectPort")]
        public Input<int>? TestObjectPort { get; set; }

        /// <summary>
        /// Specifies the test protocol. Possible values include `DNS`, `HTTP`, `HTTPS`, `FTP`, `POP`, `POPS`, `SMTP`, `SMTPS`, `TCP`, or `TCPS`.
        /// </summary>
        [Input("testObjectProtocol", required: true)]
        public Input<string> TestObjectProtocol { get; set; } = null!;

        /// <summary>
        /// A descriptive name for the testObject.
        /// </summary>
        [Input("testObjectUsername")]
        public Input<string>? TestObjectUsername { get; set; }

        /// <summary>
        /// Specifies the duration of the liveness test before it fails. The range is from 0.001 to 60 seconds.
        /// </summary>
        [Input("testTimeout", required: true)]
        public Input<double> TestTimeout { get; set; } = null!;

        /// <summary>
        /// Specifies the score to be reported if the liveness test times out.
        /// </summary>
        [Input("timeoutPenalty")]
        public Input<double>? TimeoutPenalty { get; set; }

        public GtmPropertyLivenessTestGetArgs()
        {
        }
    }
}
