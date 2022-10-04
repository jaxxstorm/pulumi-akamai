// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSecAdvancedSettingsLoggingArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSecAdvancedSettingsLoggingArgs Empty = new AppSecAdvancedSettingsLoggingArgs();

    /**
     * . Unique identifier of the security configuration containing the logging settings being modified.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration containing the logging settings being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Path to a JSON file containing the logging settings to be configured.
     * 
     */
    @Import(name="logging", required=true)
    private Output<String> logging;

    /**
     * @return . Path to a JSON file containing the logging settings to be configured.
     * 
     */
    public Output<String> logging() {
        return this.logging;
    }

    /**
     * . Unique identifier of the security policies whose settings are being modified. If not included, the logging settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policies whose settings are being modified. If not included, the logging settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private AppSecAdvancedSettingsLoggingArgs() {}

    private AppSecAdvancedSettingsLoggingArgs(AppSecAdvancedSettingsLoggingArgs $) {
        this.configId = $.configId;
        this.logging = $.logging;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecAdvancedSettingsLoggingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecAdvancedSettingsLoggingArgs $;

        public Builder() {
            $ = new AppSecAdvancedSettingsLoggingArgs();
        }

        public Builder(AppSecAdvancedSettingsLoggingArgs defaults) {
            $ = new AppSecAdvancedSettingsLoggingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration containing the logging settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration containing the logging settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param logging . Path to a JSON file containing the logging settings to be configured.
         * 
         * @return builder
         * 
         */
        public Builder logging(Output<String> logging) {
            $.logging = logging;
            return this;
        }

        /**
         * @param logging . Path to a JSON file containing the logging settings to be configured.
         * 
         * @return builder
         * 
         */
        public Builder logging(String logging) {
            return logging(Output.of(logging));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policies whose settings are being modified. If not included, the logging settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policies whose settings are being modified. If not included, the logging settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecAdvancedSettingsLoggingArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.logging = Objects.requireNonNull($.logging, "expected parameter 'logging' to be non-null");
            return $;
        }
    }

}
