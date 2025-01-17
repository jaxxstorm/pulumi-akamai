// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSecSiemSettingsState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecSiemSettingsState Empty = new AppSecSiemSettingsState();

    /**
     * . Unique identifier of the security configuration associated with the SIEM settings being modified.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the SIEM settings being modified.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
     * 
     */
    @Import(name="enableBotmanSiem")
    private @Nullable Output<Boolean> enableBotmanSiem;

    /**
     * @return . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
     * 
     */
    public Optional<Output<Boolean>> enableBotmanSiem() {
        return Optional.ofNullable(this.enableBotmanSiem);
    }

    /**
     * . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `security_policy_ids` argument.
     * 
     */
    @Import(name="enableForAllPolicies")
    private @Nullable Output<Boolean> enableForAllPolicies;

    /**
     * @return . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `security_policy_ids` argument.
     * 
     */
    public Optional<Output<Boolean>> enableForAllPolicies() {
        return Optional.ofNullable(this.enableForAllPolicies);
    }

    /**
     * . Set to **true** to enable SIEM; set to **false** to disable SIEM.
     * 
     */
    @Import(name="enableSiem")
    private @Nullable Output<Boolean> enableSiem;

    /**
     * @return . Set to **true** to enable SIEM; set to **false** to disable SIEM.
     * 
     */
    public Optional<Output<Boolean>> enableSiem() {
        return Optional.ofNullable(this.enableSiem);
    }

    /**
     * JSON array of IDs for the security policies where SIEM integration is to be enabled.
     * 
     */
    @Import(name="securityPolicyIds")
    private @Nullable Output<List<String>> securityPolicyIds;

    /**
     * @return JSON array of IDs for the security policies where SIEM integration is to be enabled.
     * 
     */
    public Optional<Output<List<String>>> securityPolicyIds() {
        return Optional.ofNullable(this.securityPolicyIds);
    }

    /**
     * . Unique identifier of the SIEM settings being modified.
     * 
     */
    @Import(name="siemId")
    private @Nullable Output<Integer> siemId;

    /**
     * @return . Unique identifier of the SIEM settings being modified.
     * 
     */
    public Optional<Output<Integer>> siemId() {
        return Optional.ofNullable(this.siemId);
    }

    private AppSecSiemSettingsState() {}

    private AppSecSiemSettingsState(AppSecSiemSettingsState $) {
        this.configId = $.configId;
        this.enableBotmanSiem = $.enableBotmanSiem;
        this.enableForAllPolicies = $.enableForAllPolicies;
        this.enableSiem = $.enableSiem;
        this.securityPolicyIds = $.securityPolicyIds;
        this.siemId = $.siemId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecSiemSettingsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecSiemSettingsState $;

        public Builder() {
            $ = new AppSecSiemSettingsState();
        }

        public Builder(AppSecSiemSettingsState defaults) {
            $ = new AppSecSiemSettingsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the SIEM settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the SIEM settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param enableBotmanSiem . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
         * 
         * @return builder
         * 
         */
        public Builder enableBotmanSiem(@Nullable Output<Boolean> enableBotmanSiem) {
            $.enableBotmanSiem = enableBotmanSiem;
            return this;
        }

        /**
         * @param enableBotmanSiem . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
         * 
         * @return builder
         * 
         */
        public Builder enableBotmanSiem(Boolean enableBotmanSiem) {
            return enableBotmanSiem(Output.of(enableBotmanSiem));
        }

        /**
         * @param enableForAllPolicies . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `security_policy_ids` argument.
         * 
         * @return builder
         * 
         */
        public Builder enableForAllPolicies(@Nullable Output<Boolean> enableForAllPolicies) {
            $.enableForAllPolicies = enableForAllPolicies;
            return this;
        }

        /**
         * @param enableForAllPolicies . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `security_policy_ids` argument.
         * 
         * @return builder
         * 
         */
        public Builder enableForAllPolicies(Boolean enableForAllPolicies) {
            return enableForAllPolicies(Output.of(enableForAllPolicies));
        }

        /**
         * @param enableSiem . Set to **true** to enable SIEM; set to **false** to disable SIEM.
         * 
         * @return builder
         * 
         */
        public Builder enableSiem(@Nullable Output<Boolean> enableSiem) {
            $.enableSiem = enableSiem;
            return this;
        }

        /**
         * @param enableSiem . Set to **true** to enable SIEM; set to **false** to disable SIEM.
         * 
         * @return builder
         * 
         */
        public Builder enableSiem(Boolean enableSiem) {
            return enableSiem(Output.of(enableSiem));
        }

        /**
         * @param securityPolicyIds JSON array of IDs for the security policies where SIEM integration is to be enabled.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyIds(@Nullable Output<List<String>> securityPolicyIds) {
            $.securityPolicyIds = securityPolicyIds;
            return this;
        }

        /**
         * @param securityPolicyIds JSON array of IDs for the security policies where SIEM integration is to be enabled.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyIds(List<String> securityPolicyIds) {
            return securityPolicyIds(Output.of(securityPolicyIds));
        }

        /**
         * @param securityPolicyIds JSON array of IDs for the security policies where SIEM integration is to be enabled.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyIds(String... securityPolicyIds) {
            return securityPolicyIds(List.of(securityPolicyIds));
        }

        /**
         * @param siemId . Unique identifier of the SIEM settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder siemId(@Nullable Output<Integer> siemId) {
            $.siemId = siemId;
            return this;
        }

        /**
         * @param siemId . Unique identifier of the SIEM settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder siemId(Integer siemId) {
            return siemId(Output.of(siemId));
        }

        public AppSecSiemSettingsState build() {
            return $;
        }
    }

}
