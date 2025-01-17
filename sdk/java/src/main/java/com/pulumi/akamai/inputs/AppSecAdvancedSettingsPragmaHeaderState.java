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


public final class AppSecAdvancedSettingsPragmaHeaderState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecAdvancedSettingsPragmaHeaderState Empty = new AppSecAdvancedSettingsPragmaHeaderState();

    /**
     * . Unique identifier of the security configuration associated with the pragma header settings being modified.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the pragma header settings being modified.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation&#39;s response except in cases where you set `excludeCondition`.
     * 
     */
    @Import(name="pragmaHeader")
    private @Nullable Output<String> pragmaHeader;

    /**
     * @return . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation&#39;s response except in cases where you set `excludeCondition`.
     * 
     */
    public Optional<Output<String>> pragmaHeader() {
        return Optional.ofNullable(this.pragmaHeader);
    }

    /**
     * . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private AppSecAdvancedSettingsPragmaHeaderState() {}

    private AppSecAdvancedSettingsPragmaHeaderState(AppSecAdvancedSettingsPragmaHeaderState $) {
        this.configId = $.configId;
        this.pragmaHeader = $.pragmaHeader;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecAdvancedSettingsPragmaHeaderState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecAdvancedSettingsPragmaHeaderState $;

        public Builder() {
            $ = new AppSecAdvancedSettingsPragmaHeaderState();
        }

        public Builder(AppSecAdvancedSettingsPragmaHeaderState defaults) {
            $ = new AppSecAdvancedSettingsPragmaHeaderState(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the pragma header settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the pragma header settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param pragmaHeader . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation&#39;s response except in cases where you set `excludeCondition`.
         * 
         * @return builder
         * 
         */
        public Builder pragmaHeader(@Nullable Output<String> pragmaHeader) {
            $.pragmaHeader = pragmaHeader;
            return this;
        }

        /**
         * @param pragmaHeader . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation&#39;s response except in cases where you set `excludeCondition`.
         * 
         * @return builder
         * 
         */
        public Builder pragmaHeader(String pragmaHeader) {
            return pragmaHeader(Output.of(pragmaHeader));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecAdvancedSettingsPragmaHeaderState build() {
            return $;
        }
    }

}
