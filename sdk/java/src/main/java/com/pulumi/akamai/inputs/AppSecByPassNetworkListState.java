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


public final class AppSecByPassNetworkListState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecByPassNetworkListState Empty = new AppSecByPassNetworkListState();

    /**
     * . JSON array of network IDs that comprise the bypass list.
     * 
     */
    @Import(name="bypassNetworkLists")
    private @Nullable Output<List<String>> bypassNetworkLists;

    /**
     * @return . JSON array of network IDs that comprise the bypass list.
     * 
     */
    public Optional<Output<List<String>>> bypassNetworkLists() {
        return Optional.ofNullable(this.bypassNetworkLists);
    }

    /**
     * . Unique identifier of the security configuration associated with the bypass network lists being modified.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the bypass network lists being modified.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Unique identifier of the security policy associated with the bypass network lists.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the bypass network lists.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private AppSecByPassNetworkListState() {}

    private AppSecByPassNetworkListState(AppSecByPassNetworkListState $) {
        this.bypassNetworkLists = $.bypassNetworkLists;
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecByPassNetworkListState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecByPassNetworkListState $;

        public Builder() {
            $ = new AppSecByPassNetworkListState();
        }

        public Builder(AppSecByPassNetworkListState defaults) {
            $ = new AppSecByPassNetworkListState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bypassNetworkLists . JSON array of network IDs that comprise the bypass list.
         * 
         * @return builder
         * 
         */
        public Builder bypassNetworkLists(@Nullable Output<List<String>> bypassNetworkLists) {
            $.bypassNetworkLists = bypassNetworkLists;
            return this;
        }

        /**
         * @param bypassNetworkLists . JSON array of network IDs that comprise the bypass list.
         * 
         * @return builder
         * 
         */
        public Builder bypassNetworkLists(List<String> bypassNetworkLists) {
            return bypassNetworkLists(Output.of(bypassNetworkLists));
        }

        /**
         * @param bypassNetworkLists . JSON array of network IDs that comprise the bypass list.
         * 
         * @return builder
         * 
         */
        public Builder bypassNetworkLists(String... bypassNetworkLists) {
            return bypassNetworkLists(List.of(bypassNetworkLists));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the bypass network lists being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the bypass network lists being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the bypass network lists.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the bypass network lists.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecByPassNetworkListState build() {
            return $;
        }
    }

}
