// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication {
    /**
     * @return Whether you want to enable the Online Certificate Status Protocol (OCSP) stapling for client certificates.
     * 
     */
    private @Nullable Boolean ocspEnabled;
    /**
     * @return Whether you want to enable the server to send the certificate authority (CA) list to the client.
     * 
     */
    private @Nullable Boolean sendCaListToClient;
    /**
     * @return The identifier of the set of trust chains, created in [Trust Chain Manager](https://techdocs.akamai.com/trust-chain-mgr/docs/welcome-trust-chain-manager).
     * 
     */
    private @Nullable String setId;

    private CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication() {}
    /**
     * @return Whether you want to enable the Online Certificate Status Protocol (OCSP) stapling for client certificates.
     * 
     */
    public Optional<Boolean> ocspEnabled() {
        return Optional.ofNullable(this.ocspEnabled);
    }
    /**
     * @return Whether you want to enable the server to send the certificate authority (CA) list to the client.
     * 
     */
    public Optional<Boolean> sendCaListToClient() {
        return Optional.ofNullable(this.sendCaListToClient);
    }
    /**
     * @return The identifier of the set of trust chains, created in [Trust Chain Manager](https://techdocs.akamai.com/trust-chain-mgr/docs/welcome-trust-chain-manager).
     * 
     */
    public Optional<String> setId() {
        return Optional.ofNullable(this.setId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean ocspEnabled;
        private @Nullable Boolean sendCaListToClient;
        private @Nullable String setId;
        public Builder() {}
        public Builder(CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ocspEnabled = defaults.ocspEnabled;
    	      this.sendCaListToClient = defaults.sendCaListToClient;
    	      this.setId = defaults.setId;
        }

        @CustomType.Setter
        public Builder ocspEnabled(@Nullable Boolean ocspEnabled) {
            this.ocspEnabled = ocspEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder sendCaListToClient(@Nullable Boolean sendCaListToClient) {
            this.sendCaListToClient = sendCaListToClient;
            return this;
        }
        @CustomType.Setter
        public Builder setId(@Nullable String setId) {
            this.setId = setId;
            return this;
        }
        public CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication build() {
            final var o = new CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication();
            o.ocspEnabled = ocspEnabled;
            o.sendCaListToClient = sendCaListToClient;
            o.setId = setId;
            return o;
        }
    }
}
