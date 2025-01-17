// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class CpsDvEnrollmentNetworkConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration for client mutual authentication. Specifies the trust chain that is used to verify client certificates and some configuration options.
        /// </summary>
        [Input("clientMutualAuthentication")]
        public Input<Inputs.CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationGetArgs>? ClientMutualAuthentication { get; set; }

        /// <summary>
        /// Whether CPS should direct traffic using all the SANs you listed in the SANs parameter when you created your enrollment.
        /// </summary>
        [Input("cloneDnsNames")]
        public Input<bool>? CloneDnsNames { get; set; }

        [Input("disallowedTlsVersions")]
        private InputList<string>? _disallowedTlsVersions;

        /// <summary>
        /// The TLS protocol version to disallow. CPS uses the TLS protocols that Akamai currently supports as a best practice.
        /// </summary>
        public InputList<string> DisallowedTlsVersions
        {
            get => _disallowedTlsVersions ?? (_disallowedTlsVersions = new InputList<string>());
            set => _disallowedTlsVersions = value;
        }

        /// <summary>
        /// Lists where you can deploy the certificate. Either `core` to specify worldwide deployment (including China and Russia), `china+core` to specify worldwide deployment and China, or `russia+core` to specify worldwide deployment and Russia. You can only use the setting to include China and Russia if your Akamai contract specifies your ability to do so and you have approval from the Chinese and Russian government.
        /// </summary>
        [Input("geography", required: true)]
        public Input<string> Geography { get; set; } = null!;

        /// <summary>
        /// The ciphers to include for the enrollment while deploying it on the network. Defaults to `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
        /// </summary>
        [Input("mustHaveCiphers")]
        public Input<string>? MustHaveCiphers { get; set; }

        /// <summary>
        /// Whether to use OCSP stapling for the enrollment, either `on`, `off` or `not-set`. OCSP Stapling improves performance by including a valid OCSP response in every TLS handshake. This option allows the visitors on your site to query the Online Certificate Status Protocol (OCSP) server at regular intervals to obtain a signed time-stamped OCSP response. This response must be signed by the CA, not the server, therefore ensuring security. Disable OSCP Stapling if you want visitors to your site to contact the CA directly for an OSCP response. OCSP allows you to obtain the revocation status of a certificate.
        /// </summary>
        [Input("ocspStapling")]
        public Input<string>? OcspStapling { get; set; }

        /// <summary>
        /// Ciphers that you preferably want to include for the enrollment while deploying it on the network. Defaults to `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
        /// </summary>
        [Input("preferredCiphers")]
        public Input<string>? PreferredCiphers { get; set; }

        /// <summary>
        /// Whether to use the QUIC transport layer network protocol.
        /// </summary>
        [Input("quicEnabled")]
        public Input<bool>? QuicEnabled { get; set; }

        public CpsDvEnrollmentNetworkConfigurationGetArgs()
        {
        }
    }
}
