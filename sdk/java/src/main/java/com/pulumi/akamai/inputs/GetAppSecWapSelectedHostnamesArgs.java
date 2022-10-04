// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetAppSecWapSelectedHostnamesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecWapSelectedHostnamesArgs Empty = new GetAppSecWapSelectedHostnamesArgs();

    /**
     * . Unique identifier of the security configuration associated with the hostnames.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the hostnames.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the security policy associated with the hostnames.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the hostnames.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    private GetAppSecWapSelectedHostnamesArgs() {}

    private GetAppSecWapSelectedHostnamesArgs(GetAppSecWapSelectedHostnamesArgs $) {
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecWapSelectedHostnamesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecWapSelectedHostnamesArgs $;

        public Builder() {
            $ = new GetAppSecWapSelectedHostnamesArgs();
        }

        public Builder(GetAppSecWapSelectedHostnamesArgs defaults) {
            $ = new GetAppSecWapSelectedHostnamesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the hostnames.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the hostnames.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the hostnames.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the hostnames.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public GetAppSecWapSelectedHostnamesArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}
