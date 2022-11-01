// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSecIPGeoState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecIPGeoState Empty = new AppSecIPGeoState();

    /**
     * . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
     * 
     */
    @Import(name="exceptionIpNetworkLists")
    private @Nullable Output<List<String>> exceptionIpNetworkLists;

    /**
     * @return . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
     * 
     */
    public Optional<Output<List<String>>> exceptionIpNetworkLists() {
        return Optional.ofNullable(this.exceptionIpNetworkLists);
    }

    /**
     * . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
     * 
     */
    @Import(name="geoNetworkLists")
    private @Nullable Output<List<String>> geoNetworkLists;

    /**
     * @return . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
     * 
     */
    public Optional<Output<List<String>>> geoNetworkLists() {
        return Optional.ofNullable(this.geoNetworkLists);
    }

    /**
     * . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
     * 
     */
    @Import(name="ipNetworkLists")
    private @Nullable Output<List<String>> ipNetworkLists;

    /**
     * @return . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
     * 
     */
    public Optional<Output<List<String>>> ipNetworkLists() {
        return Optional.ofNullable(this.ipNetworkLists);
    }

    /**
     * . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * . Unique identifier of the security policy associated with the IP/Geo lists being modified.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the IP/Geo lists being modified.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private AppSecIPGeoState() {}

    private AppSecIPGeoState(AppSecIPGeoState $) {
        this.configId = $.configId;
        this.exceptionIpNetworkLists = $.exceptionIpNetworkLists;
        this.geoNetworkLists = $.geoNetworkLists;
        this.ipNetworkLists = $.ipNetworkLists;
        this.mode = $.mode;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecIPGeoState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecIPGeoState $;

        public Builder() {
            $ = new AppSecIPGeoState();
        }

        public Builder(AppSecIPGeoState defaults) {
            $ = new AppSecIPGeoState(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param exceptionIpNetworkLists . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
         * 
         * @return builder
         * 
         */
        public Builder exceptionIpNetworkLists(@Nullable Output<List<String>> exceptionIpNetworkLists) {
            $.exceptionIpNetworkLists = exceptionIpNetworkLists;
            return this;
        }

        /**
         * @param exceptionIpNetworkLists . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
         * 
         * @return builder
         * 
         */
        public Builder exceptionIpNetworkLists(List<String> exceptionIpNetworkLists) {
            return exceptionIpNetworkLists(Output.of(exceptionIpNetworkLists));
        }

        /**
         * @param exceptionIpNetworkLists . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
         * 
         * @return builder
         * 
         */
        public Builder exceptionIpNetworkLists(String... exceptionIpNetworkLists) {
            return exceptionIpNetworkLists(List.of(exceptionIpNetworkLists));
        }

        /**
         * @param geoNetworkLists . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
         * 
         * @return builder
         * 
         */
        public Builder geoNetworkLists(@Nullable Output<List<String>> geoNetworkLists) {
            $.geoNetworkLists = geoNetworkLists;
            return this;
        }

        /**
         * @param geoNetworkLists . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
         * 
         * @return builder
         * 
         */
        public Builder geoNetworkLists(List<String> geoNetworkLists) {
            return geoNetworkLists(Output.of(geoNetworkLists));
        }

        /**
         * @param geoNetworkLists . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
         * 
         * @return builder
         * 
         */
        public Builder geoNetworkLists(String... geoNetworkLists) {
            return geoNetworkLists(List.of(geoNetworkLists));
        }

        /**
         * @param ipNetworkLists . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
         * 
         * @return builder
         * 
         */
        public Builder ipNetworkLists(@Nullable Output<List<String>> ipNetworkLists) {
            $.ipNetworkLists = ipNetworkLists;
            return this;
        }

        /**
         * @param ipNetworkLists . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
         * 
         * @return builder
         * 
         */
        public Builder ipNetworkLists(List<String> ipNetworkLists) {
            return ipNetworkLists(Output.of(ipNetworkLists));
        }

        /**
         * @param ipNetworkLists . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
         * 
         * @return builder
         * 
         */
        public Builder ipNetworkLists(String... ipNetworkLists) {
            return ipNetworkLists(List.of(ipNetworkLists));
        }

        /**
         * @param mode . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the IP/Geo lists being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the IP/Geo lists being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecIPGeoState build() {
            return $;
        }
    }

}