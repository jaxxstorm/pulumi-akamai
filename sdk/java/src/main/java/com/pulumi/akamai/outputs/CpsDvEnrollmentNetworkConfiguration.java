// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.akamai.outputs.CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CpsDvEnrollmentNetworkConfiguration {
    /**
     * @return The configuration for client mutual authentication. Specifies the trust chain that is used to verify client certificates and some configuration options.
     * 
     */
    private @Nullable CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication clientMutualAuthentication;
    /**
     * @return Whether CPS should direct traffic using all the SANs you listed in the SANs parameter when you created your enrollment.
     * 
     */
    private @Nullable Boolean cloneDnsNames;
    /**
     * @return The TLS protocol version to disallow. CPS uses the TLS protocols that Akamai currently supports as a best practice.
     * 
     */
    private @Nullable List<String> disallowedTlsVersions;
    /**
     * @return Lists where you can deploy the certificate. Either `core` to specify worldwide deployment (including China and Russia), `china+core` to specify worldwide deployment and China, or `russia+core` to specify worldwide deployment and Russia. You can only use the setting to include China and Russia if your Akamai contract specifies your ability to do so and you have approval from the Chinese and Russian government.
     * 
     */
    private String geography;
    /**
     * @return The ciphers to include for the enrollment while deploying it on the network. Defaults to `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
     * 
     */
    private @Nullable String mustHaveCiphers;
    /**
     * @return Whether to use OCSP stapling for the enrollment, either `on`, `off` or `not-set`. OCSP Stapling improves performance by including a valid OCSP response in every TLS handshake. This option allows the visitors on your site to query the Online Certificate Status Protocol (OCSP) server at regular intervals to obtain a signed time-stamped OCSP response. This response must be signed by the CA, not the server, therefore ensuring security. Disable OSCP Stapling if you want visitors to your site to contact the CA directly for an OSCP response. OCSP allows you to obtain the revocation status of a certificate.
     * 
     */
    private @Nullable String ocspStapling;
    /**
     * @return Ciphers that you preferably want to include for the enrollment while deploying it on the network. Defaults to `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
     * 
     */
    private @Nullable String preferredCiphers;
    /**
     * @return Whether to use the QUIC transport layer network protocol.
     * 
     */
    private @Nullable Boolean quicEnabled;

    private CpsDvEnrollmentNetworkConfiguration() {}
    /**
     * @return The configuration for client mutual authentication. Specifies the trust chain that is used to verify client certificates and some configuration options.
     * 
     */
    public Optional<CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication> clientMutualAuthentication() {
        return Optional.ofNullable(this.clientMutualAuthentication);
    }
    /**
     * @return Whether CPS should direct traffic using all the SANs you listed in the SANs parameter when you created your enrollment.
     * 
     */
    public Optional<Boolean> cloneDnsNames() {
        return Optional.ofNullable(this.cloneDnsNames);
    }
    /**
     * @return The TLS protocol version to disallow. CPS uses the TLS protocols that Akamai currently supports as a best practice.
     * 
     */
    public List<String> disallowedTlsVersions() {
        return this.disallowedTlsVersions == null ? List.of() : this.disallowedTlsVersions;
    }
    /**
     * @return Lists where you can deploy the certificate. Either `core` to specify worldwide deployment (including China and Russia), `china+core` to specify worldwide deployment and China, or `russia+core` to specify worldwide deployment and Russia. You can only use the setting to include China and Russia if your Akamai contract specifies your ability to do so and you have approval from the Chinese and Russian government.
     * 
     */
    public String geography() {
        return this.geography;
    }
    /**
     * @return The ciphers to include for the enrollment while deploying it on the network. Defaults to `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
     * 
     */
    public Optional<String> mustHaveCiphers() {
        return Optional.ofNullable(this.mustHaveCiphers);
    }
    /**
     * @return Whether to use OCSP stapling for the enrollment, either `on`, `off` or `not-set`. OCSP Stapling improves performance by including a valid OCSP response in every TLS handshake. This option allows the visitors on your site to query the Online Certificate Status Protocol (OCSP) server at regular intervals to obtain a signed time-stamped OCSP response. This response must be signed by the CA, not the server, therefore ensuring security. Disable OSCP Stapling if you want visitors to your site to contact the CA directly for an OSCP response. OCSP allows you to obtain the revocation status of a certificate.
     * 
     */
    public Optional<String> ocspStapling() {
        return Optional.ofNullable(this.ocspStapling);
    }
    /**
     * @return Ciphers that you preferably want to include for the enrollment while deploying it on the network. Defaults to `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
     * 
     */
    public Optional<String> preferredCiphers() {
        return Optional.ofNullable(this.preferredCiphers);
    }
    /**
     * @return Whether to use the QUIC transport layer network protocol.
     * 
     */
    public Optional<Boolean> quicEnabled() {
        return Optional.ofNullable(this.quicEnabled);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CpsDvEnrollmentNetworkConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication clientMutualAuthentication;
        private @Nullable Boolean cloneDnsNames;
        private @Nullable List<String> disallowedTlsVersions;
        private String geography;
        private @Nullable String mustHaveCiphers;
        private @Nullable String ocspStapling;
        private @Nullable String preferredCiphers;
        private @Nullable Boolean quicEnabled;
        public Builder() {}
        public Builder(CpsDvEnrollmentNetworkConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientMutualAuthentication = defaults.clientMutualAuthentication;
    	      this.cloneDnsNames = defaults.cloneDnsNames;
    	      this.disallowedTlsVersions = defaults.disallowedTlsVersions;
    	      this.geography = defaults.geography;
    	      this.mustHaveCiphers = defaults.mustHaveCiphers;
    	      this.ocspStapling = defaults.ocspStapling;
    	      this.preferredCiphers = defaults.preferredCiphers;
    	      this.quicEnabled = defaults.quicEnabled;
        }

        @CustomType.Setter
        public Builder clientMutualAuthentication(@Nullable CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication clientMutualAuthentication) {
            this.clientMutualAuthentication = clientMutualAuthentication;
            return this;
        }
        @CustomType.Setter
        public Builder cloneDnsNames(@Nullable Boolean cloneDnsNames) {
            this.cloneDnsNames = cloneDnsNames;
            return this;
        }
        @CustomType.Setter
        public Builder disallowedTlsVersions(@Nullable List<String> disallowedTlsVersions) {
            this.disallowedTlsVersions = disallowedTlsVersions;
            return this;
        }
        public Builder disallowedTlsVersions(String... disallowedTlsVersions) {
            return disallowedTlsVersions(List.of(disallowedTlsVersions));
        }
        @CustomType.Setter
        public Builder geography(String geography) {
            this.geography = Objects.requireNonNull(geography);
            return this;
        }
        @CustomType.Setter
        public Builder mustHaveCiphers(@Nullable String mustHaveCiphers) {
            this.mustHaveCiphers = mustHaveCiphers;
            return this;
        }
        @CustomType.Setter
        public Builder ocspStapling(@Nullable String ocspStapling) {
            this.ocspStapling = ocspStapling;
            return this;
        }
        @CustomType.Setter
        public Builder preferredCiphers(@Nullable String preferredCiphers) {
            this.preferredCiphers = preferredCiphers;
            return this;
        }
        @CustomType.Setter
        public Builder quicEnabled(@Nullable Boolean quicEnabled) {
            this.quicEnabled = quicEnabled;
            return this;
        }
        public CpsDvEnrollmentNetworkConfiguration build() {
            final var o = new CpsDvEnrollmentNetworkConfiguration();
            o.clientMutualAuthentication = clientMutualAuthentication;
            o.cloneDnsNames = cloneDnsNames;
            o.disallowedTlsVersions = disallowedTlsVersions;
            o.geography = geography;
            o.mustHaveCiphers = mustHaveCiphers;
            o.ocspStapling = ocspStapling;
            o.preferredCiphers = preferredCiphers;
            o.quicEnabled = quicEnabled;
            return o;
        }
    }
}
