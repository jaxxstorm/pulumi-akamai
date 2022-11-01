// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetAppSecSlowPostPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecSlowPostPlainArgs Empty = new GetAppSecSlowPostPlainArgs();

    /**
     * . Unique identifier of the security configuration associated with the slow POST settings.
     * 
     */
    @Import(name="configId", required=true)
    private Integer configId;

    /**
     * @return . Unique identifier of the security configuration associated with the slow POST settings.
     * 
     */
    public Integer configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the security policy associated with the slow POST settings.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private String securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the slow POST settings.
     * 
     */
    public String securityPolicyId() {
        return this.securityPolicyId;
    }

    private GetAppSecSlowPostPlainArgs() {}

    private GetAppSecSlowPostPlainArgs(GetAppSecSlowPostPlainArgs $) {
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecSlowPostPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecSlowPostPlainArgs $;

        public Builder() {
            $ = new GetAppSecSlowPostPlainArgs();
        }

        public Builder(GetAppSecSlowPostPlainArgs defaults) {
            $ = new GetAppSecSlowPostPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the slow POST settings.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the slow POST settings.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        public GetAppSecSlowPostPlainArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}