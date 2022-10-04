// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class AppSecWafProtectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSecWafProtectionArgs Empty = new AppSecWafProtectionArgs();

    /**
     * . Unique identifier of the security configuration associated with the WAF protection settings being modified.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the WAF protection settings being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Set to **true** to enable WAF protection; set to **false** to disable WAF protection.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return . Set to **true** to enable WAF protection; set to **false** to disable WAF protection.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * . Unique identifier of the security policy associated with the WAF protection settings being modified.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the WAF protection settings being modified.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    private AppSecWafProtectionArgs() {}

    private AppSecWafProtectionArgs(AppSecWafProtectionArgs $) {
        this.configId = $.configId;
        this.enabled = $.enabled;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecWafProtectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecWafProtectionArgs $;

        public Builder() {
            $ = new AppSecWafProtectionArgs();
        }

        public Builder(AppSecWafProtectionArgs defaults) {
            $ = new AppSecWafProtectionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the WAF protection settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the WAF protection settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param enabled . Set to **true** to enable WAF protection; set to **false** to disable WAF protection.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled . Set to **true** to enable WAF protection; set to **false** to disable WAF protection.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the WAF protection settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the WAF protection settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecWafProtectionArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}
