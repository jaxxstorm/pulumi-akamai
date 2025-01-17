// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetAppSecIPGeoPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecIPGeoPlainArgs Empty = new GetAppSecIPGeoPlainArgs();

    /**
     * . Unique identifier of the security configuration associated with the IP/Geo lists.
     * 
     */
    @Import(name="configId", required=true)
    private Integer configId;

    /**
     * @return . Unique identifier of the security configuration associated with the IP/Geo lists.
     * 
     */
    public Integer configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the security policy associated with the IP/Geo lists. If not included, information is returned for all your security policies.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private String securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the IP/Geo lists. If not included, information is returned for all your security policies.
     * 
     */
    public String securityPolicyId() {
        return this.securityPolicyId;
    }

    private GetAppSecIPGeoPlainArgs() {}

    private GetAppSecIPGeoPlainArgs(GetAppSecIPGeoPlainArgs $) {
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecIPGeoPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecIPGeoPlainArgs $;

        public Builder() {
            $ = new GetAppSecIPGeoPlainArgs();
        }

        public Builder(GetAppSecIPGeoPlainArgs defaults) {
            $ = new GetAppSecIPGeoPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the IP/Geo lists.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the IP/Geo lists. If not included, information is returned for all your security policies.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        public GetAppSecIPGeoPlainArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}
