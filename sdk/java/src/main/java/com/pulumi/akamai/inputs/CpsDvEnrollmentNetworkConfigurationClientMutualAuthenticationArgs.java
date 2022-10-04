// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs extends com.pulumi.resources.ResourceArgs {

    public static final CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs Empty = new CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs();

    /**
     * Whether you want to enable the Online Certificate Status Protocol (OCSP) stapling for client certificates.
     * 
     */
    @Import(name="ocspEnabled")
    private @Nullable Output<Boolean> ocspEnabled;

    /**
     * @return Whether you want to enable the Online Certificate Status Protocol (OCSP) stapling for client certificates.
     * 
     */
    public Optional<Output<Boolean>> ocspEnabled() {
        return Optional.ofNullable(this.ocspEnabled);
    }

    /**
     * Whether you want to enable the server to send the certificate authority (CA) list to the client.
     * 
     */
    @Import(name="sendCaListToClient")
    private @Nullable Output<Boolean> sendCaListToClient;

    /**
     * @return Whether you want to enable the server to send the certificate authority (CA) list to the client.
     * 
     */
    public Optional<Output<Boolean>> sendCaListToClient() {
        return Optional.ofNullable(this.sendCaListToClient);
    }

    /**
     * The identifier of the set of trust chains, created in [Trust Chain Manager](https://techdocs.akamai.com/trust-chain-mgr/docs/welcome-trust-chain-manager).
     * 
     */
    @Import(name="setId")
    private @Nullable Output<String> setId;

    /**
     * @return The identifier of the set of trust chains, created in [Trust Chain Manager](https://techdocs.akamai.com/trust-chain-mgr/docs/welcome-trust-chain-manager).
     * 
     */
    public Optional<Output<String>> setId() {
        return Optional.ofNullable(this.setId);
    }

    private CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs() {}

    private CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs(CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs $) {
        this.ocspEnabled = $.ocspEnabled;
        this.sendCaListToClient = $.sendCaListToClient;
        this.setId = $.setId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs $;

        public Builder() {
            $ = new CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs();
        }

        public Builder(CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs defaults) {
            $ = new CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ocspEnabled Whether you want to enable the Online Certificate Status Protocol (OCSP) stapling for client certificates.
         * 
         * @return builder
         * 
         */
        public Builder ocspEnabled(@Nullable Output<Boolean> ocspEnabled) {
            $.ocspEnabled = ocspEnabled;
            return this;
        }

        /**
         * @param ocspEnabled Whether you want to enable the Online Certificate Status Protocol (OCSP) stapling for client certificates.
         * 
         * @return builder
         * 
         */
        public Builder ocspEnabled(Boolean ocspEnabled) {
            return ocspEnabled(Output.of(ocspEnabled));
        }

        /**
         * @param sendCaListToClient Whether you want to enable the server to send the certificate authority (CA) list to the client.
         * 
         * @return builder
         * 
         */
        public Builder sendCaListToClient(@Nullable Output<Boolean> sendCaListToClient) {
            $.sendCaListToClient = sendCaListToClient;
            return this;
        }

        /**
         * @param sendCaListToClient Whether you want to enable the server to send the certificate authority (CA) list to the client.
         * 
         * @return builder
         * 
         */
        public Builder sendCaListToClient(Boolean sendCaListToClient) {
            return sendCaListToClient(Output.of(sendCaListToClient));
        }

        /**
         * @param setId The identifier of the set of trust chains, created in [Trust Chain Manager](https://techdocs.akamai.com/trust-chain-mgr/docs/welcome-trust-chain-manager).
         * 
         * @return builder
         * 
         */
        public Builder setId(@Nullable Output<String> setId) {
            $.setId = setId;
            return this;
        }

        /**
         * @param setId The identifier of the set of trust chains, created in [Trust Chain Manager](https://techdocs.akamai.com/trust-chain-mgr/docs/welcome-trust-chain-manager).
         * 
         * @return builder
         * 
         */
        public Builder setId(String setId) {
            return setId(Output.of(setId));
        }

        public CpsDvEnrollmentNetworkConfigurationClientMutualAuthenticationArgs build() {
            return $;
        }
    }

}
