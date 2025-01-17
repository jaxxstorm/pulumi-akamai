// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecApiEndpointsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecApiEndpointsArgs Empty = new GetAppSecApiEndpointsArgs();

    /**
     * . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
     * 
     */
    @Import(name="apiName")
    private @Nullable Output<String> apiName;

    /**
     * @return . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
     * 
     */
    public Optional<Output<String>> apiName() {
        return Optional.ofNullable(this.apiName);
    }

    /**
     * . Unique identifier of the security configuration associated with the API endpoints.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the API endpoints.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private GetAppSecApiEndpointsArgs() {}

    private GetAppSecApiEndpointsArgs(GetAppSecApiEndpointsArgs $) {
        this.apiName = $.apiName;
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecApiEndpointsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecApiEndpointsArgs $;

        public Builder() {
            $ = new GetAppSecApiEndpointsArgs();
        }

        public Builder(GetAppSecApiEndpointsArgs defaults) {
            $ = new GetAppSecApiEndpointsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiName . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
         * 
         * @return builder
         * 
         */
        public Builder apiName(@Nullable Output<String> apiName) {
            $.apiName = apiName;
            return this;
        }

        /**
         * @param apiName . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
         * 
         * @return builder
         * 
         */
        public Builder apiName(String apiName) {
            return apiName(Output.of(apiName));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the API endpoints.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the API endpoints.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public GetAppSecApiEndpointsArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            return $;
        }
    }

}
