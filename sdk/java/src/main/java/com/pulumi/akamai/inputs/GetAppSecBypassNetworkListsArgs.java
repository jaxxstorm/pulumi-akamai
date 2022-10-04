// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetAppSecBypassNetworkListsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecBypassNetworkListsArgs Empty = new GetAppSecBypassNetworkListsArgs();

    /**
     * . Unique identifier of the security configuration associated with the bypass network lists.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the bypass network lists.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the security policy associated with the bypass network lists.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the bypass network lists.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    private GetAppSecBypassNetworkListsArgs() {}

    private GetAppSecBypassNetworkListsArgs(GetAppSecBypassNetworkListsArgs $) {
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecBypassNetworkListsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecBypassNetworkListsArgs $;

        public Builder() {
            $ = new GetAppSecBypassNetworkListsArgs();
        }

        public Builder(GetAppSecBypassNetworkListsArgs defaults) {
            $ = new GetAppSecBypassNetworkListsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the bypass network lists.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the bypass network lists.
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
        public Builder securityPolicyId(Output<String> securityPolicyId) {
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

        public GetAppSecBypassNetworkListsArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}
