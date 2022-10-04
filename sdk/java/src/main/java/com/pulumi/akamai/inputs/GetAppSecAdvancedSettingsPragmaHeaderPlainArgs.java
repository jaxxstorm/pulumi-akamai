// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecAdvancedSettingsPragmaHeaderPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecAdvancedSettingsPragmaHeaderPlainArgs Empty = new GetAppSecAdvancedSettingsPragmaHeaderPlainArgs();

    /**
     * . Unique identifier of the security configuration associated with the pragma header settings.
     * 
     */
    @Import(name="configId", required=true)
    private Integer configId;

    /**
     * @return . Unique identifier of the security configuration associated with the pragma header settings.
     * 
     */
    public Integer configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the security policy associated with the pragma header settings. If not included, information is returned for all your security policies.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable String securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the pragma header settings. If not included, information is returned for all your security policies.
     * 
     */
    public Optional<String> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private GetAppSecAdvancedSettingsPragmaHeaderPlainArgs() {}

    private GetAppSecAdvancedSettingsPragmaHeaderPlainArgs(GetAppSecAdvancedSettingsPragmaHeaderPlainArgs $) {
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecAdvancedSettingsPragmaHeaderPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecAdvancedSettingsPragmaHeaderPlainArgs $;

        public Builder() {
            $ = new GetAppSecAdvancedSettingsPragmaHeaderPlainArgs();
        }

        public Builder(GetAppSecAdvancedSettingsPragmaHeaderPlainArgs defaults) {
            $ = new GetAppSecAdvancedSettingsPragmaHeaderPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the pragma header settings.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the pragma header settings. If not included, information is returned for all your security policies.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable String securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        public GetAppSecAdvancedSettingsPragmaHeaderPlainArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            return $;
        }
    }

}
