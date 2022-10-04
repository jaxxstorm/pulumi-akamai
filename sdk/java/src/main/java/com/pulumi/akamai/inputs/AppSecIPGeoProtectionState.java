// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSecIPGeoProtectionState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecIPGeoProtectionState Empty = new AppSecIPGeoProtectionState();

    /**
     * . Unique identifier of the security configuration associated with the IP/Geo protection settings being modified.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the IP/Geo protection settings being modified.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Set to **true** to enable IP/Geo protection; set to **false** to disable IP/Geo protection.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return . Set to **true** to enable IP/Geo protection; set to **false** to disable IP/Geo protection.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Text representation
     * 
     */
    @Import(name="outputText")
    private @Nullable Output<String> outputText;

    /**
     * @return Text representation
     * 
     */
    public Optional<Output<String>> outputText() {
        return Optional.ofNullable(this.outputText);
    }

    /**
     * . Unique identifier of the security policy associated with the IP/Geo protection settings being modified.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the IP/Geo protection settings being modified.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private AppSecIPGeoProtectionState() {}

    private AppSecIPGeoProtectionState(AppSecIPGeoProtectionState $) {
        this.configId = $.configId;
        this.enabled = $.enabled;
        this.outputText = $.outputText;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecIPGeoProtectionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecIPGeoProtectionState $;

        public Builder() {
            $ = new AppSecIPGeoProtectionState();
        }

        public Builder(AppSecIPGeoProtectionState defaults) {
            $ = new AppSecIPGeoProtectionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the IP/Geo protection settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the IP/Geo protection settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param enabled . Set to **true** to enable IP/Geo protection; set to **false** to disable IP/Geo protection.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled . Set to **true** to enable IP/Geo protection; set to **false** to disable IP/Geo protection.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param outputText Text representation
         * 
         * @return builder
         * 
         */
        public Builder outputText(@Nullable Output<String> outputText) {
            $.outputText = outputText;
            return this;
        }

        /**
         * @param outputText Text representation
         * 
         * @return builder
         * 
         */
        public Builder outputText(String outputText) {
            return outputText(Output.of(outputText));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the IP/Geo protection settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the IP/Geo protection settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecIPGeoProtectionState build() {
            return $;
        }
    }

}
